// Code generated by protoc-gen-go. DO NOT EDIT.
// source: halo.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	halo.proto

It has these top-level messages:
	HaloRequest
	HaloResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HaloRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HaloRequest) Reset()                    { *m = HaloRequest{} }
func (m *HaloRequest) String() string            { return proto.CompactTextString(m) }
func (*HaloRequest) ProtoMessage()               {}
func (*HaloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HaloRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HaloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HaloResponse) Reset()                    { *m = HaloResponse{} }
func (m *HaloResponse) String() string            { return proto.CompactTextString(m) }
func (*HaloResponse) ProtoMessage()               {}
func (*HaloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HaloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HaloRequest)(nil), "pb.haloRequest")
	proto.RegisterType((*HaloResponse)(nil), "pb.haloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Halo service

type HaloClient interface {
	Halo(ctx context.Context, in *HaloRequest, opts ...grpc.CallOption) (*HaloResponse, error)
}

type haloClient struct {
	cc *grpc.ClientConn
}

func NewHaloClient(cc *grpc.ClientConn) HaloClient {
	return &haloClient{cc}
}

func (c *haloClient) Halo(ctx context.Context, in *HaloRequest, opts ...grpc.CallOption) (*HaloResponse, error) {
	out := new(HaloResponse)
	err := grpc.Invoke(ctx, "/pb.halo/halo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Halo service

type HaloServer interface {
	Halo(context.Context, *HaloRequest) (*HaloResponse, error)
}

func RegisterHaloServer(s *grpc.Server, srv HaloServer) {
	s.RegisterService(&_Halo_serviceDesc, srv)
}

func _Halo_Halo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HaloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HaloServer).Halo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.halo/Halo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HaloServer).Halo(ctx, req.(*HaloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Halo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.halo",
	HandlerType: (*HaloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "halo",
			Handler:    _Halo_Halo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "halo.proto",
}

func init() { proto.RegisterFile("halo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x48, 0xcc, 0xc9,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe7, 0xe2, 0x06, 0x89,
	0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x4a, 0x1a, 0x5c, 0x3c, 0x10,
	0x85, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0xb8, 0x55, 0x1a, 0x19, 0x73, 0xb1, 0x80, 0x54, 0x0a,
	0x69, 0x43, 0x69, 0x7e, 0xbd, 0x82, 0x24, 0x3d, 0x24, 0x4b, 0xa4, 0x04, 0x10, 0x02, 0x10, 0xc3,
	0x94, 0x18, 0x92, 0xd8, 0xc0, 0x4e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x95, 0x49,
	0x3c, 0xa0, 0x00, 0x00, 0x00,
}
