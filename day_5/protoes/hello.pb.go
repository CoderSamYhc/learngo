// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: protoes/hello.proto

package protoes

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// required：表示该字段有且只有1个，在3.0中该修饰符被移除
// optional：表示该字段可以是0或1个，后面可加default默认值，如果不加，使用默认值
// repeated：表示该字段可以是0到多个，packed=true 代表使用高效编码格式
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoes_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protoes_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_protoes_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReplay) Reset() {
	*x = HelloReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoes_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReplay) ProtoMessage() {}

func (x *HelloReplay) ProtoReflect() protoreflect.Message {
	mi := &file_protoes_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReplay.ProtoReflect.Descriptor instead.
func (*HelloReplay) Descriptor() ([]byte, []int) {
	return file_protoes_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReplay) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloMessage) Reset() {
	*x = HelloMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoes_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloMessage) ProtoMessage() {}

func (x *HelloMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protoes_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloMessage.ProtoReflect.Descriptor instead.
func (*HelloMessage) Descriptor() ([]byte, []int) {
	return file_protoes_hello_proto_rawDescGZIP(), []int{2}
}

func (x *HelloMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_protoes_hello_proto protoreflect.FileDescriptor

var file_protoes_hello_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x73, 0x22, 0x22,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x7f, 0x0a,
	0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x05,
	0x53, 0x61, 0x79, 0x48, 0x69, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x73, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x73, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protoes_hello_proto_rawDescOnce sync.Once
	file_protoes_hello_proto_rawDescData = file_protoes_hello_proto_rawDesc
)

func file_protoes_hello_proto_rawDescGZIP() []byte {
	file_protoes_hello_proto_rawDescOnce.Do(func() {
		file_protoes_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoes_hello_proto_rawDescData)
	})
	return file_protoes_hello_proto_rawDescData
}

var file_protoes_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protoes_hello_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: protoes.HelloRequest
	(*HelloReplay)(nil),  // 1: protoes.HelloReplay
	(*HelloMessage)(nil), // 2: protoes.HelloMessage
}
var file_protoes_hello_proto_depIdxs = []int32{
	0, // 0: protoes.HelloServer.SayHi:input_type -> protoes.HelloRequest
	0, // 1: protoes.HelloServer.GetMsg:input_type -> protoes.HelloRequest
	1, // 2: protoes.HelloServer.SayHi:output_type -> protoes.HelloReplay
	2, // 3: protoes.HelloServer.GetMsg:output_type -> protoes.HelloMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoes_hello_proto_init() }
func file_protoes_hello_proto_init() {
	if File_protoes_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoes_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_protoes_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReplay); i {
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
		file_protoes_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloMessage); i {
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
			RawDescriptor: file_protoes_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoes_hello_proto_goTypes,
		DependencyIndexes: file_protoes_hello_proto_depIdxs,
		MessageInfos:      file_protoes_hello_proto_msgTypes,
	}.Build()
	File_protoes_hello_proto = out.File
	file_protoes_hello_proto_rawDesc = nil
	file_protoes_hello_proto_goTypes = nil
	file_protoes_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloServerClient is the client API for HelloServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServerClient interface {
	SayHi(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReplay, error)
	GetMsg(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloMessage, error)
}

type helloServerClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServerClient(cc grpc.ClientConnInterface) HelloServerClient {
	return &helloServerClient{cc}
}

func (c *helloServerClient) SayHi(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReplay, error) {
	out := new(HelloReplay)
	err := c.cc.Invoke(ctx, "/protoes.HelloServer/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServerClient) GetMsg(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloMessage, error) {
	out := new(HelloMessage)
	err := c.cc.Invoke(ctx, "/protoes.HelloServer/GetMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServerServer is the server API for HelloServer service.
type HelloServerServer interface {
	SayHi(context.Context, *HelloRequest) (*HelloReplay, error)
	GetMsg(context.Context, *HelloRequest) (*HelloMessage, error)
}

// UnimplementedHelloServerServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServerServer struct {
}

func (*UnimplementedHelloServerServer) SayHi(context.Context, *HelloRequest) (*HelloReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (*UnimplementedHelloServerServer) GetMsg(context.Context, *HelloRequest) (*HelloMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsg not implemented")
}

func RegisterHelloServerServer(s *grpc.Server, srv HelloServerServer) {
	s.RegisterService(&_HelloServer_serviceDesc, srv)
}

func _HelloServer_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoes.HelloServer/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).SayHi(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloServer_GetMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).GetMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoes.HelloServer/GetMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).GetMsg(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoes.HelloServer",
	HandlerType: (*HelloServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _HelloServer_SayHi_Handler,
		},
		{
			MethodName: "GetMsg",
			Handler:    _HelloServer_GetMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoes/hello.proto",
}
