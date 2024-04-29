// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package prvs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"strings"
	"regexp"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal
// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

// Attributes:
//  - BattleSpaceID
//  - VerifiyCode
//  - NatAddr
type ReportNat struct {
  BattleSpaceID string `thrift:"BattleSpaceID,1" db:"BattleSpaceID" json:"BattleSpaceID"`
  VerifiyCode string `thrift:"VerifiyCode,2" db:"VerifiyCode" json:"VerifiyCode"`
  NatAddr string `thrift:"NatAddr,3" db:"NatAddr" json:"NatAddr"`
}

func NewReportNat() *ReportNat {
  return &ReportNat{}
}


func (p *ReportNat) GetBattleSpaceID() string {
  return p.BattleSpaceID
}

func (p *ReportNat) GetVerifiyCode() string {
  return p.VerifiyCode
}

func (p *ReportNat) GetNatAddr() string {
  return p.NatAddr
}
func (p *ReportNat) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ReportNat)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.BattleSpaceID = v
}
  return nil
}

func (p *ReportNat)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.VerifiyCode = v
}
  return nil
}

func (p *ReportNat)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.NatAddr = v
}
  return nil
}

func (p *ReportNat) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "ReportNat"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ReportNat) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "BattleSpaceID", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:BattleSpaceID: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.BattleSpaceID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.BattleSpaceID (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:BattleSpaceID: ", p), err) }
  return err
}

func (p *ReportNat) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "VerifiyCode", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:VerifiyCode: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.VerifiyCode)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.VerifiyCode (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:VerifiyCode: ", p), err) }
  return err
}

func (p *ReportNat) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "NatAddr", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:NatAddr: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.NatAddr)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.NatAddr (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:NatAddr: ", p), err) }
  return err
}

func (p *ReportNat) Equals(other *ReportNat) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.BattleSpaceID != other.BattleSpaceID { return false }
  if p.VerifiyCode != other.VerifiyCode { return false }
  if p.NatAddr != other.NatAddr { return false }
  return true
}

func (p *ReportNat) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ReportNat(%+v)", *p)
}

func (p *ReportNat) Validate() error {
  return nil
}
type ReportNatService interface {
  // Parameters:
  //  - Req
  OnReportNat(ctx context.Context, req *ReportNat) (_r *ReportNat, _err error)
}

type ReportNatServiceClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewReportNatServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ReportNatServiceClient {
  return &ReportNatServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewReportNatServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ReportNatServiceClient {
  return &ReportNatServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewReportNatServiceClient(c thrift.TClient) *ReportNatServiceClient {
  return &ReportNatServiceClient{
    c: c,
  }
}

func (p *ReportNatServiceClient) Client_() thrift.TClient {
  return p.c
}

func (p *ReportNatServiceClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *ReportNatServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - Req
func (p *ReportNatServiceClient) OnReportNat(ctx context.Context, req *ReportNat) (_r *ReportNat, _err error) {
  var _args0 ReportNatServiceOnReportNatArgs
  _args0.Req = req
  var _result2 ReportNatServiceOnReportNatResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "OnReportNat", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  if _ret3 := _result2.GetSuccess(); _ret3 != nil {
    return _ret3, nil
  }
  return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "OnReportNat failed: unknown result")
}

type ReportNatServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler ReportNatService
}

func (p *ReportNatServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *ReportNatServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *ReportNatServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewReportNatServiceProcessor(handler ReportNatService) *ReportNatServiceProcessor {

  self4 := &ReportNatServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["OnReportNat"] = &reportNatServiceProcessorOnReportNat{handler:handler}
return self4
}

func (p *ReportNatServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x5.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x5

}

type reportNatServiceProcessorOnReportNat struct {
  handler ReportNatService
}

func (p *reportNatServiceProcessorOnReportNat) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  var _write_err6 error
  args := ReportNatServiceOnReportNatArgs{}
  if err2 := args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "OnReportNat", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelCauseFunc
    ctx, cancel = context.WithCancelCause(ctx)
    defer cancel(nil)
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelCauseFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel(thrift.ErrAbandonRequest)
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := ReportNatServiceOnReportNatResult{}
  if retval, err2 := p.handler.OnReportNat(ctx, args.Req); err2 != nil {
    tickerCancel()
    err = thrift.WrapTException(err2)
    if errors.Is(err2, thrift.ErrAbandonRequest) {
      return false, thrift.WrapTException(err2)
    }
    if errors.Is(err2, context.Canceled) {
      if err := context.Cause(ctx); errors.Is(err, thrift.ErrAbandonRequest) {
        return false, thrift.WrapTException(err)
      }
    }
    _exc7 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing OnReportNat: " + err2.Error())
    if err2 := oprot.WriteMessageBegin(ctx, "OnReportNat", thrift.EXCEPTION, seqId); err2 != nil {
      _write_err6 = thrift.WrapTException(err2)
    }
    if err2 := _exc7.Write(ctx, oprot); _write_err6 == nil && err2 != nil {
      _write_err6 = thrift.WrapTException(err2)
    }
    if err2 := oprot.WriteMessageEnd(ctx); _write_err6 == nil && err2 != nil {
      _write_err6 = thrift.WrapTException(err2)
    }
    if err2 := oprot.Flush(ctx); _write_err6 == nil && err2 != nil {
      _write_err6 = thrift.WrapTException(err2)
    }
    if _write_err6 != nil {
      return false, thrift.WrapTException(_write_err6)
    }
    return true, err
  } else {
    result.Success = retval
  }
  tickerCancel()
  if err2 := oprot.WriteMessageBegin(ctx, "OnReportNat", thrift.REPLY, seqId); err2 != nil {
    _write_err6 = thrift.WrapTException(err2)
  }
  if err2 := result.Write(ctx, oprot); _write_err6 == nil && err2 != nil {
    _write_err6 = thrift.WrapTException(err2)
  }
  if err2 := oprot.WriteMessageEnd(ctx); _write_err6 == nil && err2 != nil {
    _write_err6 = thrift.WrapTException(err2)
  }
  if err2 := oprot.Flush(ctx); _write_err6 == nil && err2 != nil {
    _write_err6 = thrift.WrapTException(err2)
  }
  if _write_err6 != nil {
    return false, thrift.WrapTException(_write_err6)
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Req
type ReportNatServiceOnReportNatArgs struct {
  Req *ReportNat `thrift:"req,1" db:"req" json:"req"`
}

func NewReportNatServiceOnReportNatArgs() *ReportNatServiceOnReportNatArgs {
  return &ReportNatServiceOnReportNatArgs{}
}

var ReportNatServiceOnReportNatArgs_Req_DEFAULT *ReportNat
func (p *ReportNatServiceOnReportNatArgs) GetReq() *ReportNat {
  if !p.IsSetReq() {
    return ReportNatServiceOnReportNatArgs_Req_DEFAULT
  }
return p.Req
}
func (p *ReportNatServiceOnReportNatArgs) IsSetReq() bool {
  return p.Req != nil
}

func (p *ReportNatServiceOnReportNatArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ReportNatServiceOnReportNatArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.Req = &ReportNat{}
  if err := p.Req.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
  }
  return nil
}

func (p *ReportNatServiceOnReportNatArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "OnReportNat_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ReportNatServiceOnReportNatArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "req", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err) }
  if err := p.Req.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err) }
  return err
}

func (p *ReportNatServiceOnReportNatArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ReportNatServiceOnReportNatArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ReportNatServiceOnReportNatResult struct {
  Success *ReportNat `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewReportNatServiceOnReportNatResult() *ReportNatServiceOnReportNatResult {
  return &ReportNatServiceOnReportNatResult{}
}

var ReportNatServiceOnReportNatResult_Success_DEFAULT *ReportNat
func (p *ReportNatServiceOnReportNatResult) GetSuccess() *ReportNat {
  if !p.IsSetSuccess() {
    return ReportNatServiceOnReportNatResult_Success_DEFAULT
  }
return p.Success
}
func (p *ReportNatServiceOnReportNatResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *ReportNatServiceOnReportNatResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ReportNatServiceOnReportNatResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  p.Success = &ReportNat{}
  if err := p.Success.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *ReportNatServiceOnReportNatResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "OnReportNat_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ReportNatServiceOnReportNatResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *ReportNatServiceOnReportNatResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ReportNatServiceOnReportNatResult(%+v)", *p)
}

