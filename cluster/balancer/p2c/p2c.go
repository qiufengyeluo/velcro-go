package p2c

import (
	"container/list"
	"context"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/yamakiller/velcro-go/cluster/balancer"
	"github.com/yamakiller/velcro-go/cluster/naming"
	"github.com/yamakiller/velcro-go/metric"
)

const (
	// "cost"的平均寿命,在 Tau*ln(2) 之后达到半衰期.
	tau = int64(time.Millisecond * 600)
	// 如果没有收集统计数据,我们会给端点添加一个很大的滞后惩罚.
	penalty = uint64(time.Second * 20)

	forceGap = int64(time.Second * 3)

	Name = "p2c"
)

var lags = [10]time.Duration{
	time.Millisecond * 2, time.Millisecond * 5, time.Millisecond * 20, time.Millisecond * 50,
	time.Millisecond * 100, time.Millisecond * 200, time.Millisecond * 400,
	time.Millisecond * 800, time.Millisecond * 1600, time.Millisecond * 3000,
}

type subConn struct {
	node *naming.Instance
	// 客户统计数据
	lag       int64
	success   uint64
	inflight  int64
	inflights *list.List

	// 最后收集的时间戳
	stamp int64
	//last pick timestamp
	pick int64
	// request number in a period time
	reqs int64

	predictTs int64
	predict   int64
	lk        sync.RWMutex
	//2ms,5ms,20ms,50ms,100ms,200ms,400ms,800ms,1600ms,>1600ms
	lags [10]metric.RollingCounter
}

func newSubConn(node *naming.Instance) *subConn {
	s := &subConn{
		node:      node,
		lag:       0,
		success:   1000,
		inflight:  1,
		inflights: list.New(),
	}
	for i := range s.lags {
		s.lags[i] = metric.NewRollingCounter(metric.RollingCounterOpts{Size: 20, BucketDuration: time.Millisecond * 50})
	}
	return s
}

func (sc *subConn) valid() bool {
	return sc.health() >= 500
}

func (sc *subConn) health() uint64 {
	return atomic.LoadUint64(&sc.success)
}

func (sc *subConn) load(now int64) uint64 {
	avgLag := atomic.LoadInt64(&sc.lag) + 1

	lastPredictTs := atomic.LoadInt64(&sc.predictTs)

	if now-lastPredictTs > int64(time.Millisecond*10) {
		if atomic.CompareAndSwapInt64(&sc.predictTs, lastPredictTs, now) {
			var total int64
			var count int
			var predict int64
			sc.lk.RLock()
			first := sc.inflights.Front()
			for {
				if first == nil {
					break
				}
				lag := now - first.Value.(int64)
				if lag > avgLag {
					count++
					total += lag
				}
				first = first.Next()
			}
			sc.lk.RUnlock()
			if count > sc.inflights.Len()/2 {
				predict = total / int64(count)
			}
			atomic.StoreInt64(&sc.predict, predict)

		}
	}

	predict := atomic.LoadInt64(&sc.predict)
	if predict > avgLag {
		avgLag = predict
	}
	load := uint64(avgLag) * uint64(atomic.LoadInt64(&sc.inflight))
	if load == 0 {
		// penalty是初始化没有数据时的惩罚值，默认为1e9 * 20
		load = penalty * uint64(atomic.LoadInt64(&sc.inflight))
	}
	return load
}

// statistics is info for log
type statistic struct {
	addr     string
	score    float64
	cs       uint64
	lantency time.Duration
	load     uint64
	inflight int64
	reqs     int64
	predict  time.Duration
}

// New p2c
func New(errHandler func(error) bool,
	printDebugHandler func(fmts string, args ...interface{})) balancer.Balancer {
	p := &P2cPicker{
		r:                 rand.New(rand.NewSource(time.Now().UnixNano())),
		subConns:          make(map[string]*subConn),
		errHandler:        errHandler,
		printDebugHandler: printDebugHandler,
	}
	return p
}

type P2cPicker struct {
	// subConns is the snapshot of the weighted-roundrobin balancer when this picker was
	// created. The slice is immutable. Each Get() will do a round robin
	// selection from it and return the selected SubConn.
	subConns          map[string]*subConn
	logTs             int64
	r                 *rand.Rand
	lk                sync.Mutex
	errHandler        func(err error) (isErr bool)
	printDebugHandler func(fmts string, args ...interface{})
}

// choose two distinct nodes
func (p *P2cPicker) prePick(nodes []naming.Instance) (nodeA *subConn, nodeB *subConn) {
	for i := 0; i < 2; i++ {
		p.lk.Lock()
		a := p.r.Intn(len(nodes))
		b := p.r.Intn(len(nodes) - 1)
		if b >= a {
			b = b + 1
		}
		nodeA, nodeB = p.subConns[nodes[a].Addr()], p.subConns[nodes[b].Addr()]
		if nodeA == nil {
			nodeA = newSubConn(&nodes[a])
			p.subConns[nodeA.node.Addr()] = nodeA
		}
		if nodeB == nil {
			nodeB = newSubConn(&nodes[b])
			p.subConns[nodeB.node.Addr()] = nodeB
		}
		p.lk.Unlock()

		if nodeA.valid() || nodeB.valid() {
			break
		}
	}
	return
}

func (p *P2cPicker) Pick(ctx context.Context, nodes []naming.Instance) (*naming.Instance, func(di balancer.DoneInfo), error) {
	var pc, upc *subConn
	start := time.Now().UnixNano()

	if len(nodes) == 0 {
		return nil, func(di balancer.DoneInfo) {}, nil
	} else if len(nodes) == 1 {
		p.lk.Lock()
		pc = p.subConns[nodes[0].Addr()]
		if pc == nil {
			pc = newSubConn(&nodes[0])
			p.subConns[nodes[0].Addr()] = pc
		}
		p.lk.Unlock()
	} else {
		nodeA, nodeB := p.prePick(nodes)
		// meta.Weight为服务发布者在disocvery中设置的权重
		if nodeA.load(start)*nodeB.health() > nodeB.load(start)*nodeA.health() {
			pc, upc = nodeB, nodeA
		} else {
			pc, upc = nodeA, nodeB
		}
		// 如果选中的节点，在forceGap期间内没有被选中一次，那么强制一次
		// 利用强制的机会，来触发成功率、延迟的衰减
		// 原子锁conn.pick保证并发安全，放行一次
		pick := atomic.LoadInt64(&upc.pick)
		if start-pick > forceGap && atomic.CompareAndSwapInt64(&upc.pick, pick, start) {
			pc = upc
		}
	}

	// 节点未发生切换才更新pick时间
	if pc != upc {
		atomic.StoreInt64(&pc.pick, start)
	}
	atomic.AddInt64(&pc.inflight, 1)
	atomic.AddInt64(&pc.reqs, 1)
	pc.lk.Lock()
	e := pc.inflights.PushBack(start)
	pc.lk.Unlock()

	return pc.node, func(di balancer.DoneInfo) {
		pc.lk.Lock()
		pc.inflights.Remove(e)
		pc.lk.Unlock()
		atomic.AddInt64(&pc.inflight, -1)

		now := time.Now().UnixNano()
		// get moving average ratio w
		stamp := atomic.SwapInt64(&pc.stamp, now)
		td := now - stamp
		if td < 0 {
			td = 0
		}
		w := math.Exp(float64(-td) / float64(tau))

		lag := now - start
		if lag < 0 {
			lag = 0
		}
		oldLag := atomic.LoadInt64(&pc.lag)
		if oldLag == 0 {
			w = 0.0
		}
		lag = int64(float64(oldLag)*w + float64(lag)*(1.0-w))
		atomic.StoreInt64(&pc.lag, int64(lag))

		success := uint64(1000) // error value ,if error set 1
		if di.Err != nil {
			if p.errHandler != nil {
				if p.errHandler(di.Err) {
					success = 0
				}
			} else if errors.Is(context.DeadlineExceeded, di.Err) || errors.Is(context.Canceled, di.Err) || errors.FromError(di.Err).Code >= 500 {
				success = 0
			}
		}
		oldSuc := atomic.LoadUint64(&pc.success)
		success = uint64(float64(oldSuc)*w + float64(success)*(1.0-w))
		atomic.StoreUint64(&pc.success, success)

		logTs := atomic.LoadInt64(&p.logTs)
		if now-logTs > int64(time.Second*3) {
			if atomic.CompareAndSwapInt64(&p.logTs, logTs, now) {
				p.PrintStats()
			}
		}
	}, nil
}

func (p *P2cPicker) PrintStats() {
	if len(p.subConns) == 0 {
		return
	}
	stats := make([]statistic, 0, len(p.subConns))
	var serverName string
	var reqs int64
	var now = time.Now().UnixNano()
	for _, conn := range p.subConns {
		var stat statistic
		stat.addr = conn.node.Addr()
		stat.cs = atomic.LoadUint64(&conn.success)
		stat.inflight = atomic.LoadInt64(&conn.inflight)
		stat.lantency = time.Duration(atomic.LoadInt64(&conn.lag))
		stat.reqs = atomic.SwapInt64(&conn.reqs, 0)
		stat.load = conn.load(now)
		stat.predict = time.Duration(atomic.LoadInt64(&conn.predict))
		stats = append(stats, stat)
		if serverName == "" {
			serverName = conn.node.Service.Name
		}
		reqs += stat.reqs
	}
	if reqs > 10 && p.printDebugHandler != nil {

		p.printDebugHandler("p2c %s : %+v", serverName, stats)
	}
}

func (p *P2cPicker) Schema() string {
	return Name
}
