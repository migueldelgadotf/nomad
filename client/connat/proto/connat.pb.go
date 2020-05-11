// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/connat/proto/connat.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StartRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2abbff11c8f23ebf, []int{0}
}

func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (m *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(m, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2abbff11c8f23ebf, []int{1}
}

func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (m *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(m, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

type StopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2abbff11c8f23ebf, []int{2}
}

func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (m *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(m, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

type StopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2abbff11c8f23ebf, []int{3}
}

func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (m *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(m, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StartRequest)(nil), "hashicorp.nomad.client.connat.proto.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "hashicorp.nomad.client.connat.proto.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "hashicorp.nomad.client.connat.proto.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "hashicorp.nomad.client.connat.proto.StopResponse")
}

func init() {
	proto.RegisterFile("client/connat/proto/connat.proto", fileDescriptor_2abbff11c8f23ebf)
}

var fileDescriptor_2abbff11c8f23ebf = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x4f, 0xce, 0xcf, 0xcb, 0x4b, 0x2c, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87,
	0x72, 0xf4, 0xc0, 0x1c, 0x21, 0xe5, 0x8c, 0xc4, 0xe2, 0x8c, 0xcc, 0xe4, 0xfc, 0xa2, 0x02, 0xbd,
	0xbc, 0xfc, 0xdc, 0xc4, 0x14, 0x3d, 0x88, 0x0e, 0x3d, 0x64, 0x45, 0x4a, 0x7c, 0x5c, 0x3c, 0xc1,
	0x25, 0x89, 0x45, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0xfc, 0x5c, 0xbc, 0x50,
	0x7e, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x12, 0x2f, 0x17, 0x77, 0x70, 0x49, 0x7e, 0x01, 0x4c,
	0x1e, 0xac, 0x1e, 0xc4, 0x85, 0x48, 0x1b, 0xbd, 0x64, 0xe4, 0x62, 0x73, 0xce, 0xcf, 0xf3, 0x4b,
	0x2c, 0x11, 0x2a, 0xe0, 0x62, 0x05, 0x6b, 0x15, 0x32, 0xd4, 0x23, 0xc2, 0x66, 0x3d, 0x64, 0x6b,
	0xa5, 0x8c, 0x48, 0xd1, 0x02, 0x75, 0x19, 0x83, 0x50, 0x2e, 0x17, 0x0b, 0xc8, 0x31, 0x42, 0x06,
	0x44, 0xea, 0x86, 0x7b, 0x43, 0xca, 0x90, 0x04, 0x1d, 0x30, 0xeb, 0x9c, 0xd8, 0xa3, 0x58, 0xc1,
	0xe2, 0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x57, 0x9f, 0x87, 0x6e, 0x84,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConNatClient is the client API for ConNat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConNatClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type conNatClient struct {
	cc grpc.ClientConnInterface
}

func NewConNatClient(cc grpc.ClientConnInterface) ConNatClient {
	return &conNatClient{cc}
}

func (c *conNatClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.client.connat.proto.ConNat/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conNatClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.client.connat.proto.ConNat/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConNatServer is the server API for ConNat service.
type ConNatServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

// UnimplementedConNatServer can be embedded to have forward compatible implementations.
type UnimplementedConNatServer struct {
}

func (*UnimplementedConNatServer) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedConNatServer) Stop(ctx context.Context, req *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterConNatServer(s *grpc.Server, srv ConNatServer) {
	s.RegisterService(&_ConNat_serviceDesc, srv)
}

func _ConNat_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConNatServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.client.connat.proto.ConNat/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConNatServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConNat_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConNatServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.client.connat.proto.ConNat/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConNatServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConNat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.client.connat.proto.ConNat",
	HandlerType: (*ConNatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _ConNat_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ConNat_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/connat/proto/connat.proto",
}
