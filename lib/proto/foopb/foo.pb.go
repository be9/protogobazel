// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.3
// source: foo/foo.proto

package foopb

import (
	context "context"
	barpb "github.com/be9/protogobazel/lib/proto/barpb"
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

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B *barpb.Bar `protobuf:"bytes,1,opt,name=b,proto3" json:"b,omitempty"`
	F FooEnum    `protobuf:"varint,2,opt,name=f,proto3,enum=foo.FooEnum" json:"f,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_foo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_foo_foo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_foo_foo_proto_rawDescGZIP(), []int{0}
}

func (x *Foo) GetB() *barpb.Bar {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *Foo) GetF() FooEnum {
	if x != nil {
		return x.F
	}
	return FooEnum_FOO_ENUM_UNKNOWN
}

type MethRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MethRequest) Reset() {
	*x = MethRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_foo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethRequest) ProtoMessage() {}

func (x *MethRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foo_foo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethRequest.ProtoReflect.Descriptor instead.
func (*MethRequest) Descriptor() ([]byte, []int) {
	return file_foo_foo_proto_rawDescGZIP(), []int{1}
}

type MethResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MethResponse) Reset() {
	*x = MethResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_foo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethResponse) ProtoMessage() {}

func (x *MethResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foo_foo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethResponse.ProtoReflect.Descriptor instead.
func (*MethResponse) Descriptor() ([]byte, []int) {
	return file_foo_foo_proto_rawDescGZIP(), []int{2}
}

var File_foo_foo_proto protoreflect.FileDescriptor

var file_foo_foo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x6f, 0x6f, 0x2f, 0x66, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x66, 0x6f, 0x6f, 0x1a, 0x0d, 0x62, 0x61, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x66, 0x6f, 0x6f, 0x2f, 0x66, 0x6f, 0x6f, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x16,
	0x0a, 0x01, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x62, 0x61, 0x72, 0x2e,
	0x42, 0x61, 0x72, 0x52, 0x01, 0x62, 0x12, 0x1a, 0x0a, 0x01, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x46, 0x6f, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x01, 0x66, 0x22, 0x0d, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x2f, 0x0a, 0x01, 0x53, 0x12, 0x2a, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x68, 0x12, 0x10,
	0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x65, 0x39, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x6f, 0x62, 0x61, 0x7a, 0x65,
	0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x6f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foo_foo_proto_rawDescOnce sync.Once
	file_foo_foo_proto_rawDescData = file_foo_foo_proto_rawDesc
)

func file_foo_foo_proto_rawDescGZIP() []byte {
	file_foo_foo_proto_rawDescOnce.Do(func() {
		file_foo_foo_proto_rawDescData = protoimpl.X.CompressGZIP(file_foo_foo_proto_rawDescData)
	})
	return file_foo_foo_proto_rawDescData
}

var file_foo_foo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foo_foo_proto_goTypes = []interface{}{
	(*Foo)(nil),          // 0: foo.Foo
	(*MethRequest)(nil),  // 1: foo.MethRequest
	(*MethResponse)(nil), // 2: foo.MethResponse
	(*barpb.Bar)(nil),    // 3: bar.Bar
	(FooEnum)(0),         // 4: foo.FooEnum
}
var file_foo_foo_proto_depIdxs = []int32{
	3, // 0: foo.Foo.b:type_name -> bar.Bar
	4, // 1: foo.Foo.f:type_name -> foo.FooEnum
	1, // 2: foo.S.Meth:input_type -> foo.MethRequest
	1, // 3: foo.S.Meth:output_type -> foo.MethRequest
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foo_foo_proto_init() }
func file_foo_foo_proto_init() {
	if File_foo_foo_proto != nil {
		return
	}
	file_foo_foo_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_foo_foo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
		file_foo_foo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethRequest); i {
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
		file_foo_foo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethResponse); i {
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
			RawDescriptor: file_foo_foo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foo_foo_proto_goTypes,
		DependencyIndexes: file_foo_foo_proto_depIdxs,
		MessageInfos:      file_foo_foo_proto_msgTypes,
	}.Build()
	File_foo_foo_proto = out.File
	file_foo_foo_proto_rawDesc = nil
	file_foo_foo_proto_goTypes = nil
	file_foo_foo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SClient is the client API for S service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SClient interface {
	Meth(ctx context.Context, in *MethRequest, opts ...grpc.CallOption) (*MethRequest, error)
}

type sClient struct {
	cc grpc.ClientConnInterface
}

func NewSClient(cc grpc.ClientConnInterface) SClient {
	return &sClient{cc}
}

func (c *sClient) Meth(ctx context.Context, in *MethRequest, opts ...grpc.CallOption) (*MethRequest, error) {
	out := new(MethRequest)
	err := c.cc.Invoke(ctx, "/foo.S/Meth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SServer is the server API for S service.
type SServer interface {
	Meth(context.Context, *MethRequest) (*MethRequest, error)
}

// UnimplementedSServer can be embedded to have forward compatible implementations.
type UnimplementedSServer struct {
}

func (*UnimplementedSServer) Meth(context.Context, *MethRequest) (*MethRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Meth not implemented")
}

func RegisterSServer(s *grpc.Server, srv SServer) {
	s.RegisterService(&_S_serviceDesc, srv)
}

func _S_Meth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MethRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SServer).Meth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.S/Meth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SServer).Meth(ctx, req.(*MethRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _S_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foo.S",
	HandlerType: (*SServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Meth",
			Handler:    _S_Meth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foo/foo.proto",
}
