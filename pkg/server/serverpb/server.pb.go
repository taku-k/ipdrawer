// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/serverpb/server.proto

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	server/serverpb/server.proto

It has these top-level messages:
	DrawIPRequest
	DrawIPResponse
	GetPrefixIncludingIPRequest
	GetPrefixIncludingIPResponse
	ActivateIPRequest
	ActivateIPResponse
	GetPrefixRequest
	GetPrefixResponse
	CreatePrefixRequest
	CreatePrefixResponse
	CreatePoolRequest
	CreatePoolResponse
	Tag
	Pool
*/
package serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/mwitkow/go-proto-validators"

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

type DrawIPRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask int32  `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
	Tags []*Tag `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
}

func (m *DrawIPRequest) Reset()                    { *m = DrawIPRequest{} }
func (m *DrawIPRequest) String() string            { return proto.CompactTextString(m) }
func (*DrawIPRequest) ProtoMessage()               {}
func (*DrawIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DrawIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DrawIPRequest) GetMask() int32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *DrawIPRequest) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type DrawIPResponse struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *DrawIPResponse) Reset()                    { *m = DrawIPResponse{} }
func (m *DrawIPResponse) String() string            { return proto.CompactTextString(m) }
func (*DrawIPResponse) ProtoMessage()               {}
func (*DrawIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DrawIPResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type GetPrefixIncludingIPRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *GetPrefixIncludingIPRequest) Reset()                    { *m = GetPrefixIncludingIPRequest{} }
func (m *GetPrefixIncludingIPRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPrefixIncludingIPRequest) ProtoMessage()               {}
func (*GetPrefixIncludingIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetPrefixIncludingIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type GetPrefixIncludingIPResponse struct {
}

func (m *GetPrefixIncludingIPResponse) Reset()                    { *m = GetPrefixIncludingIPResponse{} }
func (m *GetPrefixIncludingIPResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPrefixIncludingIPResponse) ProtoMessage()               {}
func (*GetPrefixIncludingIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ActivateIPRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *ActivateIPRequest) Reset()                    { *m = ActivateIPRequest{} }
func (m *ActivateIPRequest) String() string            { return proto.CompactTextString(m) }
func (*ActivateIPRequest) ProtoMessage()               {}
func (*ActivateIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ActivateIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ActivateIPResponse struct {
}

func (m *ActivateIPResponse) Reset()                    { *m = ActivateIPResponse{} }
func (m *ActivateIPResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivateIPResponse) ProtoMessage()               {}
func (*ActivateIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetPrefixRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask int32  `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
}

func (m *GetPrefixRequest) Reset()                    { *m = GetPrefixRequest{} }
func (m *GetPrefixRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPrefixRequest) ProtoMessage()               {}
func (*GetPrefixRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetPrefixRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *GetPrefixRequest) GetMask() int32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

type GetPrefixResponse struct {
	Ipnet           string   `protobuf:"bytes,1,opt,name=ipnet" json:"ipnet,omitempty"`
	DefaultGateways []string `protobuf:"bytes,2,rep,name=default_gateways,json=defaultGateways" json:"default_gateways,omitempty"`
	Broadcast       string   `protobuf:"bytes,3,opt,name=broadcast" json:"broadcast,omitempty"`
	Netmask         string   `protobuf:"bytes,4,opt,name=netmask" json:"netmask,omitempty"`
	Tags            []*Tag   `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
}

func (m *GetPrefixResponse) Reset()                    { *m = GetPrefixResponse{} }
func (m *GetPrefixResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPrefixResponse) ProtoMessage()               {}
func (*GetPrefixResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetPrefixResponse) GetIpnet() string {
	if m != nil {
		return m.Ipnet
	}
	return ""
}

func (m *GetPrefixResponse) GetDefaultGateways() []string {
	if m != nil {
		return m.DefaultGateways
	}
	return nil
}

func (m *GetPrefixResponse) GetBroadcast() string {
	if m != nil {
		return m.Broadcast
	}
	return ""
}

func (m *GetPrefixResponse) GetNetmask() string {
	if m != nil {
		return m.Netmask
	}
	return ""
}

func (m *GetPrefixResponse) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreatePrefixRequest struct {
	Ip              string   `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask            int32    `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
	DefaultGateways []string `protobuf:"bytes,3,rep,name=default_gateways,json=defaultGateways" json:"default_gateways,omitempty"`
	Broadcast       string   `protobuf:"bytes,4,opt,name=broadcast" json:"broadcast,omitempty"`
	Netmask         string   `protobuf:"bytes,5,opt,name=netmask" json:"netmask,omitempty"`
	Tags            []*Tag   `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
}

func (m *CreatePrefixRequest) Reset()                    { *m = CreatePrefixRequest{} }
func (m *CreatePrefixRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePrefixRequest) ProtoMessage()               {}
func (*CreatePrefixRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreatePrefixRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreatePrefixRequest) GetMask() int32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *CreatePrefixRequest) GetDefaultGateways() []string {
	if m != nil {
		return m.DefaultGateways
	}
	return nil
}

func (m *CreatePrefixRequest) GetBroadcast() string {
	if m != nil {
		return m.Broadcast
	}
	return ""
}

func (m *CreatePrefixRequest) GetNetmask() string {
	if m != nil {
		return m.Netmask
	}
	return ""
}

func (m *CreatePrefixRequest) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreatePrefixResponse struct {
}

func (m *CreatePrefixResponse) Reset()                    { *m = CreatePrefixResponse{} }
func (m *CreatePrefixResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePrefixResponse) ProtoMessage()               {}
func (*CreatePrefixResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type CreatePoolRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask int32  `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
	Pool *Pool  `protobuf:"bytes,3,opt,name=pool" json:"pool,omitempty"`
}

func (m *CreatePoolRequest) Reset()                    { *m = CreatePoolRequest{} }
func (m *CreatePoolRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePoolRequest) ProtoMessage()               {}
func (*CreatePoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreatePoolRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreatePoolRequest) GetMask() int32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *CreatePoolRequest) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

type CreatePoolResponse struct {
}

func (m *CreatePoolResponse) Reset()                    { *m = CreatePoolResponse{} }
func (m *CreatePoolResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePoolResponse) ProtoMessage()               {}
func (*CreatePoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type Tag struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Pool struct {
	Start  string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End    string `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	Status int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	Tags   []*Tag `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
}

func (m *Pool) Reset()                    { *m = Pool{} }
func (m *Pool) String() string            { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()               {}
func (*Pool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Pool) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Pool) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *Pool) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Pool) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*DrawIPRequest)(nil), "serverpb.DrawIPRequest")
	proto.RegisterType((*DrawIPResponse)(nil), "serverpb.DrawIPResponse")
	proto.RegisterType((*GetPrefixIncludingIPRequest)(nil), "serverpb.GetPrefixIncludingIPRequest")
	proto.RegisterType((*GetPrefixIncludingIPResponse)(nil), "serverpb.GetPrefixIncludingIPResponse")
	proto.RegisterType((*ActivateIPRequest)(nil), "serverpb.ActivateIPRequest")
	proto.RegisterType((*ActivateIPResponse)(nil), "serverpb.ActivateIPResponse")
	proto.RegisterType((*GetPrefixRequest)(nil), "serverpb.GetPrefixRequest")
	proto.RegisterType((*GetPrefixResponse)(nil), "serverpb.GetPrefixResponse")
	proto.RegisterType((*CreatePrefixRequest)(nil), "serverpb.CreatePrefixRequest")
	proto.RegisterType((*CreatePrefixResponse)(nil), "serverpb.CreatePrefixResponse")
	proto.RegisterType((*CreatePoolRequest)(nil), "serverpb.CreatePoolRequest")
	proto.RegisterType((*CreatePoolResponse)(nil), "serverpb.CreatePoolResponse")
	proto.RegisterType((*Tag)(nil), "serverpb.Tag")
	proto.RegisterType((*Pool)(nil), "serverpb.Pool")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PrefixService service

type PrefixServiceClient interface {
	DrawIP(ctx context.Context, in *DrawIPRequest, opts ...grpc.CallOption) (*DrawIPResponse, error)
	GetPrefix(ctx context.Context, in *GetPrefixRequest, opts ...grpc.CallOption) (*GetPrefixResponse, error)
	CreatePrefix(ctx context.Context, in *CreatePrefixRequest, opts ...grpc.CallOption) (*CreatePrefixResponse, error)
	CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error)
}

type prefixServiceClient struct {
	cc *grpc.ClientConn
}

func NewPrefixServiceClient(cc *grpc.ClientConn) PrefixServiceClient {
	return &prefixServiceClient{cc}
}

func (c *prefixServiceClient) DrawIP(ctx context.Context, in *DrawIPRequest, opts ...grpc.CallOption) (*DrawIPResponse, error) {
	out := new(DrawIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.PrefixService/DrawIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prefixServiceClient) GetPrefix(ctx context.Context, in *GetPrefixRequest, opts ...grpc.CallOption) (*GetPrefixResponse, error) {
	out := new(GetPrefixResponse)
	err := grpc.Invoke(ctx, "/serverpb.PrefixService/GetPrefix", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prefixServiceClient) CreatePrefix(ctx context.Context, in *CreatePrefixRequest, opts ...grpc.CallOption) (*CreatePrefixResponse, error) {
	out := new(CreatePrefixResponse)
	err := grpc.Invoke(ctx, "/serverpb.PrefixService/CreatePrefix", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prefixServiceClient) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error) {
	out := new(CreatePoolResponse)
	err := grpc.Invoke(ctx, "/serverpb.PrefixService/CreatePool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PrefixService service

type PrefixServiceServer interface {
	DrawIP(context.Context, *DrawIPRequest) (*DrawIPResponse, error)
	GetPrefix(context.Context, *GetPrefixRequest) (*GetPrefixResponse, error)
	CreatePrefix(context.Context, *CreatePrefixRequest) (*CreatePrefixResponse, error)
	CreatePool(context.Context, *CreatePoolRequest) (*CreatePoolResponse, error)
}

func RegisterPrefixServiceServer(s *grpc.Server, srv PrefixServiceServer) {
	s.RegisterService(&_PrefixService_serviceDesc, srv)
}

func _PrefixService_DrawIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixServiceServer).DrawIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.PrefixService/DrawIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixServiceServer).DrawIP(ctx, req.(*DrawIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrefixService_GetPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixServiceServer).GetPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.PrefixService/GetPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixServiceServer).GetPrefix(ctx, req.(*GetPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrefixService_CreatePrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixServiceServer).CreatePrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.PrefixService/CreatePrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixServiceServer).CreatePrefix(ctx, req.(*CreatePrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrefixService_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixServiceServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.PrefixService/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixServiceServer).CreatePool(ctx, req.(*CreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrefixService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.PrefixService",
	HandlerType: (*PrefixServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DrawIP",
			Handler:    _PrefixService_DrawIP_Handler,
		},
		{
			MethodName: "GetPrefix",
			Handler:    _PrefixService_GetPrefix_Handler,
		},
		{
			MethodName: "CreatePrefix",
			Handler:    _PrefixService_CreatePrefix_Handler,
		},
		{
			MethodName: "CreatePool",
			Handler:    _PrefixService_CreatePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/server.proto",
}

// Client API for IPService service

type IPServiceClient interface {
	GetPrefixIncludingIP(ctx context.Context, in *GetPrefixIncludingIPRequest, opts ...grpc.CallOption) (*GetPrefixIncludingIPResponse, error)
	ActivateIP(ctx context.Context, in *ActivateIPRequest, opts ...grpc.CallOption) (*ActivateIPResponse, error)
}

type iPServiceClient struct {
	cc *grpc.ClientConn
}

func NewIPServiceClient(cc *grpc.ClientConn) IPServiceClient {
	return &iPServiceClient{cc}
}

func (c *iPServiceClient) GetPrefixIncludingIP(ctx context.Context, in *GetPrefixIncludingIPRequest, opts ...grpc.CallOption) (*GetPrefixIncludingIPResponse, error) {
	out := new(GetPrefixIncludingIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.IPService/GetPrefixIncludingIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPServiceClient) ActivateIP(ctx context.Context, in *ActivateIPRequest, opts ...grpc.CallOption) (*ActivateIPResponse, error) {
	out := new(ActivateIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.IPService/ActivateIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IPService service

type IPServiceServer interface {
	GetPrefixIncludingIP(context.Context, *GetPrefixIncludingIPRequest) (*GetPrefixIncludingIPResponse, error)
	ActivateIP(context.Context, *ActivateIPRequest) (*ActivateIPResponse, error)
}

func RegisterIPServiceServer(s *grpc.Server, srv IPServiceServer) {
	s.RegisterService(&_IPService_serviceDesc, srv)
}

func _IPService_GetPrefixIncludingIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrefixIncludingIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPServiceServer).GetPrefixIncludingIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.IPService/GetPrefixIncludingIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPServiceServer).GetPrefixIncludingIP(ctx, req.(*GetPrefixIncludingIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPService_ActivateIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPServiceServer).ActivateIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.IPService/ActivateIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPServiceServer).ActivateIP(ctx, req.(*ActivateIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.IPService",
	HandlerType: (*IPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrefixIncludingIP",
			Handler:    _IPService_GetPrefixIncludingIP_Handler,
		},
		{
			MethodName: "ActivateIP",
			Handler:    _IPService_ActivateIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/server.proto",
}

func init() { proto.RegisterFile("server/serverpb/server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 844 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x4f, 0x8f, 0xdb, 0x44,
	0x14, 0xd7, 0xd8, 0x49, 0x20, 0xaf, 0xec, 0xb2, 0x3b, 0xdd, 0x76, 0x8d, 0x37, 0xdd, 0xa4, 0x86,
	0x56, 0x5e, 0xa4, 0xc4, 0xa9, 0x97, 0x22, 0x2d, 0x12, 0x07, 0x0a, 0x52, 0xb5, 0xb7, 0x55, 0xe8,
	0x01, 0x29, 0x04, 0x34, 0x49, 0x66, 0xcd, 0x28, 0x59, 0x8f, 0xb1, 0x27, 0x09, 0x55, 0x36, 0x42,
	0x54, 0x02, 0x89, 0x0b, 0x17, 0x38, 0x73, 0xe3, 0x4b, 0x70, 0x43, 0x1c, 0x11, 0x27, 0x6e, 0x5c,
	0x22, 0x45, 0x7c, 0x08, 0x6e, 0x20, 0x8f, 0x27, 0x71, 0x42, 0x1c, 0xfe, 0x5c, 0x4a, 0x37, 0x97,
	0xd8, 0xef, 0xcd, 0xcc, 0xef, 0x8f, 0xdf, 0xb3, 0x1f, 0x94, 0x22, 0x1a, 0x0e, 0x69, 0xe8, 0x24,
	0x7f, 0x41, 0x5b, 0x5d, 0xd4, 0x82, 0x90, 0x0b, 0x8e, 0x9f, 0x9f, 0x87, 0xcd, 0x92, 0xc7, 0xb9,
	0xd7, 0xa7, 0x0e, 0x09, 0x98, 0x43, 0x7c, 0x9f, 0x0b, 0x22, 0x18, 0xf7, 0xa3, 0x64, 0x9d, 0xf9,
	0xba, 0xc7, 0xc4, 0x47, 0x83, 0x76, 0xad, 0xc3, 0x2f, 0x9c, 0x8b, 0x11, 0x13, 0x3d, 0x3e, 0x72,
	0x3c, 0x5e, 0x95, 0xc9, 0xea, 0x90, 0xf4, 0x59, 0x97, 0x08, 0x1e, 0x46, 0xce, 0xe2, 0x32, 0xd9,
	0x67, 0xfd, 0x8a, 0x60, 0xeb, 0x9d, 0x90, 0x8c, 0x4e, 0xcf, 0x1a, 0xf4, 0xe3, 0x01, 0x8d, 0x04,
	0xfe, 0x14, 0x34, 0x16, 0x18, 0xa8, 0x82, 0xec, 0xe2, 0x03, 0x3e, 0x9b, 0x96, 0x7b, 0xc0, 0x3e,
	0xb0, 0xed, 0x66, 0xbd, 0x7a, 0xd2, 0xba, 0x6c, 0xde, 0xab, 0x9e, 0xb4, 0x92, 0xcb, 0x7b, 0xf2,
	0x6f, 0xec, 0x4e, 0x2e, 0xdd, 0x66, 0xbd, 0xfa, 0x9a, 0x8a, 0xba, 0xf7, 0x9b, 0xf5, 0xea, 0xfd,
	0xd6, 0xd1, 0xfb, 0xb5, 0xa3, 0xf1, 0xf1, 0xe4, 0xbf, 0xee, 0x7a, 0xa5, 0xa1, 0xb1, 0x00, 0xdf,
	0x81, 0xdc, 0x05, 0x89, 0x7a, 0x86, 0x56, 0x41, 0x76, 0xfe, 0xc1, 0xee, 0x6c, 0x5a, 0xde, 0xda,
	0xf9, 0x63, 0xfe, 0x43, 0xc6, 0xed, 0x86, 0x4c, 0xe3, 0xdb, 0x90, 0x13, 0xc4, 0x8b, 0x0c, 0xbd,
	0xa2, 0xdb, 0xd7, 0xdc, 0xad, 0xda, 0xdc, 0xa8, 0xda, 0x23, 0xe2, 0x35, 0x64, 0xca, 0xaa, 0xc0,
	0xf6, 0x5c, 0x5b, 0x14, 0x70, 0x3f, 0xa2, 0x78, 0x3b, 0x15, 0x17, 0x63, 0x59, 0xdf, 0x22, 0x38,
	0x78, 0x48, 0xc5, 0x59, 0x48, 0xcf, 0xd9, 0x27, 0xa7, 0x7e, 0xa7, 0x3f, 0xe8, 0x32, 0xdf, 0x7b,
	0x76, 0xcc, 0xb0, 0x0e, 0xa1, 0x94, 0xcd, 0x2f, 0x11, 0x64, 0x7d, 0x83, 0x60, 0xf7, 0xad, 0x8e,
	0x60, 0x43, 0x22, 0xe8, 0x33, 0x44, 0x7b, 0x0f, 0xf0, 0x32, 0x2b, 0x45, 0xf6, 0x7b, 0x04, 0x3b,
	0x0b, 0x35, 0x57, 0xac, 0xde, 0xac, 0x1f, 0x34, 0xd8, 0x5d, 0x22, 0xaf, 0x0a, 0xea, 0x3b, 0x04,
	0x79, 0x16, 0xf8, 0x54, 0x28, 0x05, 0x5f, 0xa1, 0xd9, 0xb4, 0xfc, 0x25, 0x82, 0x2f, 0xd0, 0xd3,
	0x12, 0xe1, 0xd8, 0xcb, 0xab, 0xe4, 0x92, 0x93, 0xd6, 0xe5, 0x71, 0xb3, 0x5e, 0x75, 0x63, 0x85,
	0x09, 0x3b, 0x7c, 0x04, 0x3b, 0x5d, 0x7a, 0x4e, 0x06, 0x7d, 0xf1, 0xa1, 0x47, 0x04, 0x1d, 0x91,
	0xc7, 0x91, 0xa1, 0x55, 0x74, 0xbb, 0xd8, 0x78, 0x51, 0xc5, 0x1f, 0xaa, 0x30, 0x2e, 0x41, 0xb1,
	0x1d, 0x72, 0xd2, 0xed, 0x90, 0x48, 0x18, 0xba, 0x6c, 0x95, 0x34, 0x80, 0x0d, 0x78, 0xce, 0xa7,
	0x42, 0x1a, 0x96, 0x93, 0xb9, 0xf9, 0xed, 0xa2, 0x21, 0xf3, 0x9b, 0x1b, 0xf2, 0x47, 0x0d, 0xae,
	0xbf, 0x1d, 0x52, 0x22, 0xe8, 0x95, 0xac, 0x81, 0x4c, 0x17, 0xf5, 0x7f, 0xe1, 0x62, 0xee, 0x6f,
	0x5c, 0xcc, 0x67, 0xbb, 0x58, 0xd8, 0xec, 0xe2, 0x4d, 0xd8, 0x5b, 0x35, 0x51, 0xb5, 0xd7, 0x14,
	0xc1, 0xae, 0x4a, 0x70, 0xde, 0xbf, 0x6a, 0xde, 0x5a, 0x90, 0x0b, 0x38, 0xef, 0xcb, 0x8a, 0xbb,
	0xe6, 0x6e, 0xa7, 0xc2, 0xa5, 0x18, 0x99, 0x8b, 0x5f, 0x2b, 0xcb, 0x02, 0x95, 0xee, 0x37, 0x41,
	0x7f, 0x44, 0x3c, 0x6c, 0x80, 0xde, 0xa3, 0x8f, 0x95, 0xd2, 0xc2, 0x6c, 0x5a, 0xd6, 0xde, 0x43,
	0x8d, 0x38, 0x84, 0x4b, 0x90, 0x1f, 0x92, 0xfe, 0x80, 0x4a, 0x0a, 0x69, 0x2e, 0x09, 0x5a, 0x3f,
	0x6b, 0x90, 0x8b, 0xcf, 0xc3, 0x9f, 0x23, 0xc8, 0x47, 0x82, 0x84, 0xe2, 0xff, 0x72, 0x2b, 0x41,
	0xc7, 0x9f, 0x21, 0xd0, 0xa9, 0xdf, 0x55, 0x6c, 0x9f, 0x3a, 0x8b, 0x18, 0x1b, 0xdf, 0x84, 0x42,
	0x24, 0x88, 0x18, 0x44, 0xf2, 0x79, 0xe4, 0x1b, 0xea, 0x6e, 0x51, 0x9e, 0xb9, 0x8d, 0xe5, 0xe9,
	0xfe, 0xa4, 0xc3, 0x56, 0x52, 0x99, 0xef, 0xd2, 0x70, 0xc8, 0x3a, 0x14, 0x77, 0xa0, 0x90, 0x7c,
	0x87, 0xf1, 0x7e, 0xba, 0x61, 0x65, 0xea, 0x30, 0x8d, 0xf5, 0x84, 0x7a, 0xba, 0x77, 0x9f, 0xfc,
	0xf2, 0xdb, 0xd7, 0x5a, 0x05, 0x1f, 0xca, 0xc9, 0x27, 0x90, 0x07, 0x3b, 0x63, 0x16, 0x4c, 0x9c,
	0x71, 0x5c, 0x37, 0x13, 0xa7, 0x1b, 0x92, 0x11, 0x0b, 0x30, 0x85, 0xe2, 0xe2, 0xf5, 0x8c, 0xcd,
	0xf4, 0xb8, 0xbf, 0x7e, 0x70, 0xcc, 0x83, 0xcc, 0x9c, 0x42, 0x2b, 0x4b, 0xb4, 0x97, 0xf0, 0xfe,
	0x06, 0x34, 0x3c, 0x80, 0x17, 0x96, 0x9b, 0x0f, 0xdf, 0x4a, 0x4f, 0xcb, 0x78, 0xb3, 0x99, 0x87,
	0x9b, 0xd2, 0xab, 0xea, 0xac, 0x8d, 0xea, 0x3a, 0x72, 0x17, 0x1e, 0x01, 0xa4, 0x95, 0x8f, 0x0f,
	0xd6, 0x4e, 0x4d, 0x1b, 0xde, 0x2c, 0x65, 0x27, 0x15, 0x60, 0x4d, 0x02, 0xda, 0xd6, 0xcb, 0x9b,
	0x00, 0xe3, 0x46, 0x53, 0xa8, 0x6f, 0xa0, 0x57, 0xdd, 0xdf, 0x11, 0x14, 0x4f, 0xcf, 0xe6, 0x4f,
	0xf2, 0x09, 0x82, 0xbd, 0xac, 0x79, 0x04, 0xdf, 0xc9, 0x30, 0x75, 0x7d, 0x9e, 0x32, 0xef, 0xfe,
	0xd3, 0x32, 0xc5, 0xf2, 0x40, 0xb2, 0xbc, 0x81, 0xaf, 0x4b, 0x96, 0x2c, 0x48, 0x18, 0x26, 0x6c,
	0xf1, 0x39, 0x40, 0x3a, 0x5c, 0x2c, 0x7b, 0xb1, 0x36, 0x08, 0x2d, 0x7b, 0x91, 0x31, 0x8f, 0xdc,
	0x92, 0x28, 0xfb, 0xd6, 0x8d, 0x15, 0x14, 0xa2, 0x16, 0xb6, 0x0b, 0x72, 0x44, 0x3e, 0xfe, 0x33,
	0x00, 0x00, 0xff, 0xff, 0x60, 0xdc, 0xce, 0x03, 0xa2, 0x0b, 0x00, 0x00,
}
