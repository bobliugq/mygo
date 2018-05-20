// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inf.proto

package helloworld

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

// enum type default is int32
type Ua int32

const (
	Ua_Unkown  Ua = 0
	Ua_Web     Ua = 1
	Ua_H5      Ua = 2
	Ua_Ios     Ua = 3
	Ua_Android Ua = 4
	Ua_Pc      Ua = 5
	Ua_Wx      Ua = 6
)

var Ua_name = map[int32]string{
	0: "Unkown",
	1: "Web",
	2: "H5",
	3: "Ios",
	4: "Android",
	5: "Pc",
	6: "Wx",
}
var Ua_value = map[string]int32{
	"Unkown":  0,
	"Web":     1,
	"H5":      2,
	"Ios":     3,
	"Android": 4,
	"Pc":      5,
	"Wx":      6,
}

func (x Ua) String() string {
	return proto.EnumName(Ua_name, int32(x))
}
func (Ua) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_inf_e77f91c0f5c1c9da, []int{0}
}

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ua                   int32    `protobuf:"varint,2,opt,name=ua" json:"ua,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_inf_e77f91c0f5c1c9da, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetUa() int32 {
	if m != nil {
		return m.Ua
	}
	return 0
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_inf_e77f91c0f5c1c9da, []int{1}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (dst *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(dst, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

// error ua
type ErrorUa struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorUa) Reset()         { *m = ErrorUa{} }
func (m *ErrorUa) String() string { return proto.CompactTextString(m) }
func (*ErrorUa) ProtoMessage()    {}
func (*ErrorUa) Descriptor() ([]byte, []int) {
	return fileDescriptor_inf_e77f91c0f5c1c9da, []int{2}
}
func (m *ErrorUa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorUa.Unmarshal(m, b)
}
func (m *ErrorUa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorUa.Marshal(b, m, deterministic)
}
func (dst *ErrorUa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorUa.Merge(dst, src)
}
func (m *ErrorUa) XXX_Size() int {
	return xxx_messageInfo_ErrorUa.Size(m)
}
func (m *ErrorUa) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorUa.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorUa proto.InternalMessageInfo

func (m *ErrorUa) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ErrorUa) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type User struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Job                  string   `protobuf:"bytes,2,opt,name=job" json:"job,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_inf_e77f91c0f5c1c9da, []int{3}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *User) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type GetUserReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Level                int32    `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReq) Reset()         { *m = GetUserReq{} }
func (m *GetUserReq) String() string { return proto.CompactTextString(m) }
func (*GetUserReq) ProtoMessage()    {}
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_inf_e77f91c0f5c1c9da, []int{4}
}
func (m *GetUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReq.Unmarshal(m, b)
}
func (m *GetUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReq.Marshal(b, m, deterministic)
}
func (dst *GetUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReq.Merge(dst, src)
}
func (m *GetUserReq) XXX_Size() int {
	return xxx_messageInfo_GetUserReq.Size(m)
}
func (m *GetUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReq proto.InternalMessageInfo

func (m *GetUserReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GetUserReq) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type GetUserReply struct {
	UserInfo             *User    `protobuf:"bytes,1,opt,name=userInfo" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReply) Reset()         { *m = GetUserReply{} }
func (m *GetUserReply) String() string { return proto.CompactTextString(m) }
func (*GetUserReply) ProtoMessage()    {}
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_inf_e77f91c0f5c1c9da, []int{5}
}
func (m *GetUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReply.Unmarshal(m, b)
}
func (m *GetUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReply.Marshal(b, m, deterministic)
}
func (dst *GetUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReply.Merge(dst, src)
}
func (m *GetUserReply) XXX_Size() int {
	return xxx_messageInfo_GetUserReply.Size(m)
}
func (m *GetUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReply proto.InternalMessageInfo

func (m *GetUserReply) GetUserInfo() *User {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*ErrorUa)(nil), "helloworld.ErrorUa")
	proto.RegisterType((*User)(nil), "helloworld.User")
	proto.RegisterType((*GetUserReq)(nil), "helloworld.GetUserReq")
	proto.RegisterType((*GetUserReply)(nil), "helloworld.GetUserReply")
	proto.RegisterEnum("helloworld.Ua", Ua_name, Ua_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	GetUserInfo(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetUserInfo(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	GetUserInfo(context.Context, *GetUserReq) (*GetUserReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetUserInfo(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Greeter_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inf.proto",
}

func init() { proto.RegisterFile("inf.proto", fileDescriptor_inf_e77f91c0f5c1c9da) }

var fileDescriptor_inf_e77f91c0f5c1c9da = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6b, 0xea, 0x40,
	0x14, 0x35, 0x1f, 0x26, 0x7a, 0x95, 0xc7, 0x70, 0x79, 0x3c, 0x82, 0x2b, 0xc9, 0x4a, 0x1e, 0x0f,
	0x17, 0xbe, 0x96, 0x42, 0x29, 0x05, 0x17, 0x45, 0xdd, 0x49, 0x4a, 0x70, 0x3d, 0x9a, 0x6b, 0x6b,
	0x3b, 0x66, 0x74, 0x92, 0xd4, 0xfa, 0x1f, 0xfa, 0xa3, 0xcb, 0x4c, 0x63, 0x8c, 0xe0, 0x2a, 0xe7,
	0xde, 0x7b, 0xce, 0x99, 0x39, 0x73, 0x03, 0xed, 0x4d, 0xba, 0x1e, 0xee, 0x94, 0xcc, 0x25, 0xc2,
	0x2b, 0x09, 0x21, 0x0f, 0x52, 0x89, 0x24, 0x1c, 0x41, 0x77, 0xaa, 0xab, 0x88, 0xf6, 0x05, 0x65,
	0x39, 0x22, 0xb8, 0x29, 0xdf, 0x52, 0x60, 0xf5, 0xad, 0x41, 0x3b, 0x32, 0x18, 0x7f, 0x81, 0x5d,
	0xf0, 0xc0, 0xee, 0x5b, 0x83, 0x66, 0x64, 0x17, 0x3c, 0xbc, 0x07, 0x28, 0x35, 0x3b, 0x71, 0xc4,
	0x00, 0xfc, 0x2d, 0x65, 0x19, 0x7f, 0x39, 0x89, 0x4e, 0xa5, 0xf6, 0x5a, 0xc9, 0x84, 0x4a, 0xa5,
	0xc1, 0xe1, 0x1d, 0xf8, 0x4f, 0x4a, 0x49, 0x15, 0xf3, 0x6a, 0x6c, 0x9d, 0xc7, 0x75, 0x33, 0xfb,
	0xc2, 0x2c, 0x9c, 0x83, 0x1b, 0x67, 0xa4, 0x90, 0x81, 0x53, 0x6c, 0x12, 0x23, 0x72, 0x22, 0x0d,
	0x75, 0xe7, 0x4d, 0x2e, 0x4b, 0xbe, 0x86, 0x55, 0x08, 0xa7, 0x16, 0x82, 0x81, 0xa3, 0x5d, 0x5d,
	0x73, 0x98, 0x86, 0xe1, 0x0d, 0xc0, 0x84, 0x72, 0x6d, 0x1a, 0xd1, 0xfe, 0x8a, 0xef, 0x6f, 0x68,
	0x0a, 0xfa, 0x20, 0x51, 0xde, 0xff, 0xa7, 0x08, 0x1f, 0xa0, 0x5b, 0xa9, 0x74, 0xfc, 0x7f, 0xd0,
	0x2a, 0x32, 0x52, 0xb3, 0x74, 0x2d, 0x8d, 0xb8, 0x33, 0x62, 0xc3, 0xf3, 0xfb, 0x0e, 0x0d, 0xb1,
	0x62, 0xfc, 0x9d, 0x80, 0x1d, 0x73, 0x04, 0xf0, 0xe2, 0xf4, 0x5d, 0x1e, 0x52, 0xd6, 0x40, 0x1f,
	0x9c, 0x05, 0x2d, 0x99, 0x85, 0x1e, 0xd8, 0xd3, 0x5b, 0x66, 0xeb, 0xc6, 0x4c, 0x66, 0xcc, 0xc1,
	0x0e, 0xf8, 0xe3, 0x34, 0x51, 0x72, 0x93, 0x30, 0x57, 0x4f, 0xe7, 0x2b, 0xd6, 0xd4, 0xdf, 0xc5,
	0x27, 0xf3, 0x46, 0x5f, 0x16, 0xf8, 0x13, 0x45, 0x94, 0x93, 0xc2, 0x47, 0x68, 0x3d, 0xf3, 0xa3,
	0x59, 0x09, 0x06, 0xf5, 0xc3, 0xeb, 0x9b, 0xed, 0xfd, 0xb9, 0x32, 0xd9, 0x89, 0x63, 0xd8, 0xc0,
	0x31, 0x74, 0xca, 0x48, 0xfa, 0x8e, 0x78, 0x41, 0x3c, 0xbf, 0x50, 0x2f, 0xb8, 0xda, 0x37, 0x16,
	0x4b, 0xcf, 0xfc, 0x59, 0xff, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x25, 0xc7, 0x51, 0x9e, 0x66,
	0x02, 0x00, 0x00,
}
