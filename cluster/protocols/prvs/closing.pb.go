// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.5.1
// source: closing.proto

package prvs

import (
	network "github.com/yamakiller/velcro-go/network"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求网关关闭客户端
type RequestGatewayCloseClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *network.ClientID `protobuf:"bytes,1,opt,name=Target,proto3" json:"Target,omitempty"`
}

func (x *RequestGatewayCloseClient) Reset() {
	*x = RequestGatewayCloseClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_closing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGatewayCloseClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGatewayCloseClient) ProtoMessage() {}

func (x *RequestGatewayCloseClient) ProtoReflect() protoreflect.Message {
	mi := &file_closing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGatewayCloseClient.ProtoReflect.Descriptor instead.
func (*RequestGatewayCloseClient) Descriptor() ([]byte, []int) {
	return file_closing_proto_rawDescGZIP(), []int{0}
}

func (x *RequestGatewayCloseClient) GetTarget() *network.ClientID {
	if x != nil {
		return x.Target
	}
	return nil
}

// 通知某服务客户端连接已关闭
type ClientClosed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID *network.ClientID `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
}

func (x *ClientClosed) Reset() {
	*x = ClientClosed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_closing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientClosed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientClosed) ProtoMessage() {}

func (x *ClientClosed) ProtoReflect() protoreflect.Message {
	mi := &file_closing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientClosed.ProtoReflect.Descriptor instead.
func (*ClientClosed) Descriptor() ([]byte, []int) {
	return file_closing_proto_rawDescGZIP(), []int{1}
}

func (x *ClientClosed) GetClientID() *network.ClientID {
	if x != nil {
		return x.ClientID
	}
	return nil
}

var File_closing_proto protoreflect.FileDescriptor

var file_closing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x70, 0x72, 0x76, 0x73, 0x1a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x19, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x3d,
	0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x2d,
	0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x3b, 0x70, 0x72, 0x76, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_closing_proto_rawDescOnce sync.Once
	file_closing_proto_rawDescData = file_closing_proto_rawDesc
)

func file_closing_proto_rawDescGZIP() []byte {
	file_closing_proto_rawDescOnce.Do(func() {
		file_closing_proto_rawDescData = protoimpl.X.CompressGZIP(file_closing_proto_rawDescData)
	})
	return file_closing_proto_rawDescData
}

var file_closing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_closing_proto_goTypes = []interface{}{
	(*RequestGatewayCloseClient)(nil), // 0: prvs.RequestGatewayCloseClient
	(*ClientClosed)(nil),              // 1: prvs.ClientClosed
	(*network.ClientID)(nil),          // 2: network.ClientID
}
var file_closing_proto_depIdxs = []int32{
	2, // 0: prvs.RequestGatewayCloseClient.Target:type_name -> network.ClientID
	2, // 1: prvs.ClientClosed.ClientID:type_name -> network.ClientID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_closing_proto_init() }
func file_closing_proto_init() {
	if File_closing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_closing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGatewayCloseClient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_closing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientClosed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_closing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_closing_proto_goTypes,
		DependencyIndexes: file_closing_proto_depIdxs,
		MessageInfos:      file_closing_proto_msgTypes,
	}.Build()
	File_closing_proto = out.File
	file_closing_proto_rawDesc = nil
	file_closing_proto_goTypes = nil
	file_closing_proto_depIdxs = nil
}
