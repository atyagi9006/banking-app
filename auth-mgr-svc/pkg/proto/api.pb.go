// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenRequest) Reset()         { *m = TokenRequest{} }
func (m *TokenRequest) String() string { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()    {}
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *TokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenRequest.Unmarshal(m, b)
}
func (m *TokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenRequest.Marshal(b, m, deterministic)
}
func (m *TokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenRequest.Merge(m, src)
}
func (m *TokenRequest) XXX_Size() int {
	return xxx_messageInfo_TokenRequest.Size(m)
}
func (m *TokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenRequest proto.InternalMessageInfo

func (m *TokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyTokenResponse struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTokenResponse) Reset()         { *m = VerifyTokenResponse{} }
func (m *VerifyTokenResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenResponse) ProtoMessage()    {}
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *VerifyTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenResponse.Unmarshal(m, b)
}
func (m *VerifyTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenResponse.Marshal(b, m, deterministic)
}
func (m *VerifyTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenResponse.Merge(m, src)
}
func (m *VerifyTokenResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenResponse.Size(m)
}
func (m *VerifyTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenResponse proto.InternalMessageInfo

func (m *VerifyTokenResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *VerifyTokenResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type GenerateTokenRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateTokenRequest) Reset()         { *m = GenerateTokenRequest{} }
func (m *GenerateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateTokenRequest) ProtoMessage()    {}
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *GenerateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateTokenRequest.Unmarshal(m, b)
}
func (m *GenerateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateTokenRequest.Marshal(b, m, deterministic)
}
func (m *GenerateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateTokenRequest.Merge(m, src)
}
func (m *GenerateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateTokenRequest.Size(m)
}
func (m *GenerateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateTokenRequest proto.InternalMessageInfo

func (m *GenerateTokenRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GenerateTokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GenerateTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateTokenResponse) Reset()         { *m = GenerateTokenResponse{} }
func (m *GenerateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateTokenResponse) ProtoMessage()    {}
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *GenerateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateTokenResponse.Unmarshal(m, b)
}
func (m *GenerateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateTokenResponse.Marshal(b, m, deterministic)
}
func (m *GenerateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateTokenResponse.Merge(m, src)
}
func (m *GenerateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateTokenResponse.Size(m)
}
func (m *GenerateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateTokenResponse proto.InternalMessageInfo

func (m *GenerateTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RegisterUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role                 string   `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserRequest) Reset()         { *m = RegisterUserRequest{} }
func (m *RegisterUserRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterUserRequest) ProtoMessage()    {}
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *RegisterUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserRequest.Unmarshal(m, b)
}
func (m *RegisterUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserRequest.Marshal(b, m, deterministic)
}
func (m *RegisterUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserRequest.Merge(m, src)
}
func (m *RegisterUserRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterUserRequest.Size(m)
}
func (m *RegisterUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserRequest proto.InternalMessageInfo

func (m *RegisterUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterUserRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Role                 string   `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type GetUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// empty
type EmptyMessageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyMessageResponse) Reset()         { *m = EmptyMessageResponse{} }
func (m *EmptyMessageResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyMessageResponse) ProtoMessage()    {}
func (*EmptyMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *EmptyMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyMessageResponse.Unmarshal(m, b)
}
func (m *EmptyMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyMessageResponse.Marshal(b, m, deterministic)
}
func (m *EmptyMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyMessageResponse.Merge(m, src)
}
func (m *EmptyMessageResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyMessageResponse.Size(m)
}
func (m *EmptyMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyMessageResponse proto.InternalMessageInfo

type UserIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdRequest) Reset()         { *m = UserIdRequest{} }
func (m *UserIdRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdRequest) ProtoMessage()    {}
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *UserIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdRequest.Unmarshal(m, b)
}
func (m *UserIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdRequest.Marshal(b, m, deterministic)
}
func (m *UserIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdRequest.Merge(m, src)
}
func (m *UserIdRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdRequest.Size(m)
}
func (m *UserIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdRequest proto.InternalMessageInfo

func (m *UserIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenRequest)(nil), "proto.TokenRequest")
	proto.RegisterType((*VerifyTokenResponse)(nil), "proto.VerifyTokenResponse")
	proto.RegisterType((*GenerateTokenRequest)(nil), "proto.GenerateTokenRequest")
	proto.RegisterType((*GenerateTokenResponse)(nil), "proto.GenerateTokenResponse")
	proto.RegisterType((*RegisterUserRequest)(nil), "proto.RegisterUserRequest")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*GetUserRequest)(nil), "proto.GetUserRequest")
	proto.RegisterType((*EmptyMessageResponse)(nil), "proto.EmptyMessageResponse")
	proto.RegisterType((*UserIdRequest)(nil), "proto.UserIdRequest")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4b, 0x4b, 0xeb, 0x40,
	0x14, 0xc7, 0xd3, 0xd0, 0xde, 0x7b, 0x7b, 0xfa, 0x58, 0x4c, 0xd3, 0x4b, 0xc9, 0x2d, 0x5c, 0x19,
	0x5c, 0xb8, 0xb1, 0x05, 0x05, 0xc1, 0x95, 0x2d, 0x28, 0x55, 0xb0, 0x9b, 0xfa, 0xd8, 0xb8, 0x1a,
	0xcd, 0x31, 0x0e, 0xa6, 0x99, 0x38, 0x33, 0x55, 0xfa, 0x29, 0xfd, 0x4a, 0x92, 0x74, 0x92, 0xa6,
	0x65, 0x0a, 0xe2, 0x2a, 0x39, 0xe7, 0xfc, 0xe7, 0x77, 0x9e, 0x50, 0x67, 0x09, 0x1f, 0x24, 0x52,
	0x68, 0x41, 0x6a, 0xd9, 0xc7, 0xef, 0x87, 0x42, 0x84, 0x11, 0x0e, 0x59, 0xc2, 0x87, 0x2c, 0x8e,
	0x85, 0x66, 0x9a, 0x8b, 0x58, 0xad, 0x44, 0x74, 0x1f, 0x9a, 0xb7, 0xe2, 0x15, 0xe3, 0x19, 0xbe,
	0x2d, 0x50, 0x69, 0xe2, 0x41, 0x4d, 0xa7, 0x76, 0xaf, 0xb2, 0x57, 0x39, 0xa8, 0xcf, 0x56, 0x06,
	0x3d, 0x83, 0xce, 0x3d, 0x4a, 0xfe, 0xbc, 0x34, 0x5a, 0x95, 0x88, 0x58, 0x61, 0x2a, 0xc6, 0x39,
	0xe3, 0x51, 0x2e, 0xce, 0x0c, 0x42, 0xa0, 0x2a, 0x45, 0x84, 0x3d, 0x37, 0x73, 0x66, 0xff, 0xf4,
	0x12, 0xbc, 0x09, 0xc6, 0x28, 0x99, 0xc6, 0xed, 0x74, 0x16, 0x82, 0x0f, 0x7f, 0x12, 0xa6, 0xd4,
	0x87, 0x90, 0x81, 0xa1, 0x14, 0x36, 0x3d, 0x84, 0xee, 0x16, 0x69, 0x5d, 0x8c, 0xa5, 0xf2, 0x07,
	0xe8, 0xcc, 0x30, 0xe4, 0x4a, 0xa3, 0xbc, 0x53, 0x28, 0x7f, 0x9c, 0xb7, 0xe8, 0xaa, 0x5a, 0xea,
	0x6a, 0x04, 0xd5, 0x14, 0x4a, 0xda, 0xe0, 0xf2, 0xc0, 0xa0, 0x5c, 0x1e, 0xac, 0xe9, 0xae, 0x6d,
	0x2e, 0xb5, 0x12, 0xe1, 0x04, 0xda, 0x13, 0xd4, 0xe5, 0xca, 0xbe, 0xc5, 0xa2, 0x7f, 0xc1, 0xbb,
	0x98, 0x27, 0x7a, 0x39, 0x45, 0xa5, 0x58, 0x88, 0xf9, 0x10, 0xe8, 0x7f, 0x68, 0xa5, 0xb0, 0xab,
	0x60, 0x07, 0xee, 0xe8, 0xd3, 0x85, 0xf6, 0x78, 0xa1, 0x5f, 0xa6, 0xa1, 0xbc, 0x41, 0xf9, 0xce,
	0x9f, 0x90, 0x5c, 0x43, 0x6b, 0x63, 0xa2, 0xe4, 0xdf, 0xea, 0x36, 0x06, 0xb6, 0x8d, 0xf9, 0x7d,
	0x7b, 0xd0, 0xe4, 0x77, 0xc8, 0x08, 0x1a, 0xa5, 0x53, 0x21, 0x1d, 0x23, 0xdf, 0x60, 0xf8, 0xc6,
	0x69, 0xb9, 0x29, 0xea, 0x90, 0x53, 0x68, 0x96, 0x57, 0x46, 0x72, 0xb5, 0x65, 0x8f, 0x7e, 0xc3,
	0xc4, 0x52, 0x1f, 0x75, 0xc8, 0x18, 0xe0, 0x1c, 0x23, 0xd4, 0x98, 0x3d, 0xf4, 0x4a, 0xc1, 0x62,
	0x22, 0x7e, 0xde, 0x9d, 0x75, 0x7e, 0x0e, 0x19, 0xc2, 0x6f, 0xb3, 0x11, 0xd2, 0x2d, 0x5a, 0xd5,
	0xbb, 0x73, 0x3e, 0xfe, 0xca, 0xac, 0xe3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xf0, 0x6f,
	0x08, 0x7a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthMgrServiceClient is the client API for AuthMgrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthMgrServiceClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
	VerifyToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*EmptyMessageResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
}

type authMgrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthMgrServiceClient(cc grpc.ClientConnInterface) AuthMgrServiceClient {
	return &authMgrServiceClient{cc}
}

func (c *authMgrServiceClient) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	out := new(GenerateTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthMgrService/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMgrServiceClient) VerifyToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthMgrService/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMgrServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.AuthMgrService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMgrServiceClient) DeleteUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*EmptyMessageResponse, error) {
	out := new(EmptyMessageResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthMgrService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authMgrServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.AuthMgrService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthMgrServiceServer is the server API for AuthMgrService service.
type AuthMgrServiceServer interface {
	GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error)
	VerifyToken(context.Context, *TokenRequest) (*VerifyTokenResponse, error)
	RegisterUser(context.Context, *RegisterUserRequest) (*User, error)
	DeleteUser(context.Context, *UserIdRequest) (*EmptyMessageResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
}

// UnimplementedAuthMgrServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthMgrServiceServer struct {
}

func (*UnimplementedAuthMgrServiceServer) GenerateToken(ctx context.Context, req *GenerateTokenRequest) (*GenerateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (*UnimplementedAuthMgrServiceServer) VerifyToken(ctx context.Context, req *TokenRequest) (*VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (*UnimplementedAuthMgrServiceServer) RegisterUser(ctx context.Context, req *RegisterUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (*UnimplementedAuthMgrServiceServer) DeleteUser(ctx context.Context, req *UserIdRequest) (*EmptyMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedAuthMgrServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterAuthMgrServiceServer(s *grpc.Server, srv AuthMgrServiceServer) {
	s.RegisterService(&_AuthMgrService_serviceDesc, srv)
}

func _AuthMgrService_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMgrServiceServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthMgrService/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMgrServiceServer).GenerateToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMgrService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMgrServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthMgrService/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMgrServiceServer).VerifyToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMgrService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMgrServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthMgrService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMgrServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMgrService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMgrServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthMgrService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMgrServiceServer).DeleteUser(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthMgrService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMgrServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthMgrService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMgrServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthMgrService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthMgrService",
	HandlerType: (*AuthMgrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _AuthMgrService_GenerateToken_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _AuthMgrService_VerifyToken_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _AuthMgrService_RegisterUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthMgrService_DeleteUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthMgrService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
