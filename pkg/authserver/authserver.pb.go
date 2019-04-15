// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/authserver/authserver.proto

package authserver

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

//*
// CreateTokenRequest represents the request of CreateToken.
type CreateTokenRequest struct {
	// user is the username which want to be authenticate.
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// password is the credential of given user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// orig_host is the hostname for which JWT token is valid.
	OrigHost             string   `protobuf:"bytes,3,opt,name=orig_host,json=origHost,proto3" json:"orig_host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTokenRequest) Reset()         { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()    {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b970e5b30bc31cb8, []int{0}
}

func (m *CreateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenRequest.Unmarshal(m, b)
}
func (m *CreateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenRequest.Marshal(b, m, deterministic)
}
func (m *CreateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenRequest.Merge(m, src)
}
func (m *CreateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTokenRequest.Size(m)
}
func (m *CreateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenRequest proto.InternalMessageInfo

func (m *CreateTokenRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *CreateTokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateTokenRequest) GetOrigHost() string {
	if m != nil {
		return m.OrigHost
	}
	return ""
}

//*
// Token represents the response of CreateToken.
type Token struct {
	// token is the JWT token.
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_b970e5b30bc31cb8, []int{1}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTokenRequest)(nil), "authserver.CreateTokenRequest")
	proto.RegisterType((*Token)(nil), "authserver.Token")
}

func init() { proto.RegisterFile("pkg/authserver/authserver.proto", fileDescriptor_b970e5b30bc31cb8) }

var fileDescriptor_b970e5b30bc31cb8 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x4f, 0x0b, 0x82, 0x40,
	0x10, 0xc5, 0xb3, 0x32, 0x72, 0x3a, 0x35, 0x74, 0x10, 0xa3, 0x3f, 0x78, 0xea, 0x64, 0x50, 0x9f,
	0xa0, 0xba, 0x74, 0x0c, 0xe9, 0x1e, 0x1b, 0x0d, 0x1a, 0x42, 0x6b, 0x3b, 0x63, 0x7d, 0xfd, 0x70,
	0x0d, 0x5b, 0xe8, 0xf6, 0xe6, 0xfd, 0xe0, 0xf1, 0x1b, 0x58, 0x94, 0x45, 0xb6, 0x56, 0x95, 0xe4,
	0x4c, 0xe6, 0x45, 0xc6, 0x89, 0x49, 0x69, 0xb4, 0x68, 0x84, 0x5f, 0x13, 0x2b, 0xc0, 0x83, 0x21,
	0x25, 0x74, 0xd6, 0x05, 0x3d, 0x52, 0x7a, 0x56, 0xc4, 0x82, 0x08, 0xfd, 0x8a, 0xc9, 0x84, 0xde,
	0xd2, 0x5b, 0x05, 0xa9, 0xcd, 0x18, 0xc1, 0xb0, 0x54, 0xcc, 0x6f, 0x6d, 0x6e, 0x61, 0xd7, 0xf6,
	0xed, 0x8d, 0x53, 0x08, 0xb4, 0xb9, 0x67, 0x97, 0x5c, 0xb3, 0x84, 0xbd, 0x06, 0xd6, 0xc5, 0x51,
	0xb3, 0xc4, 0x33, 0xf0, 0xed, 0x38, 0x4e, 0xc0, 0x97, 0x3a, 0x7c, 0x67, 0x9b, 0x63, 0x73, 0x02,
	0xd8, 0xb5, 0x3e, 0xb8, 0x87, 0x91, 0xe3, 0x83, 0xf3, 0xc4, 0xb1, 0xff, 0x17, 0x8d, 0xc6, 0x2e,
	0xb7, 0x24, 0xee, 0x5c, 0x07, 0xf6, 0xcd, 0xed, 0x27, 0x00, 0x00, 0xff, 0xff, 0x31, 0xf6, 0x1e,
	0x8f, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthserverClient is the client API for Authserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthserverClient interface {
	// CreateToken creates and returns new JWT token for requested identity.
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*Token, error)
}

type authserverClient struct {
	cc *grpc.ClientConn
}

func NewAuthserverClient(cc *grpc.ClientConn) AuthserverClient {
	return &authserverClient{cc}
}

func (c *authserverClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/authserver.Authserver/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthserverServer is the server API for Authserver service.
type AuthserverServer interface {
	// CreateToken creates and returns new JWT token for requested identity.
	CreateToken(context.Context, *CreateTokenRequest) (*Token, error)
}

func RegisterAuthserverServer(s *grpc.Server, srv AuthserverServer) {
	s.RegisterService(&_Authserver_serviceDesc, srv)
}

func _Authserver_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthserverServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authserver.Authserver/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthserverServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authserver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authserver.Authserver",
	HandlerType: (*AuthserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Authserver_CreateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/authserver/authserver.proto",
}
