// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	user.proto
	vpn.proto
	network.proto
	auth.proto

It has these top-level messages:
	UserListRequest
	UserCreateRequest
	UserUpdateRequest
	UserDeleteRequest
	UserRenewRequest
	UserGenConfigRequest
	UserResponse
	UserGenConfigResponse
	VPNStatusRequest
	VPNInitRequest
	VPNUpdateRequest
	VPNStatusResponse
	VPNInitResponse
	VPNUpdateResponse
	NetworkCreateRequest
	NetworkListRequest
	NetworkDeleteRequest
	NetworkGetAllTypesRequest
	NetworkAssociateRequest
	NetworkDissociateRequest
	NetworkGetAssociatedUsersRequest
	Network
	NetworkType
	NetworkCreateResponse
	NetworkListResponse
	NetworkDeleteResponse
	NetworkGetAllTypesResponse
	NetworkAssociateResponse
	NetworkDissociateResponse
	NetworkGetAssociatedUsersResponse
	AuthStatusRequest
	AuthAuthenticateRequest
	AuthStatusResponse
	AuthAuthenticateResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type UserUpdateRequest_GWPref int32

const (
	UserUpdateRequest_NOPREF UserUpdateRequest_GWPref = 0
	UserUpdateRequest_NOGW   UserUpdateRequest_GWPref = 1
	UserUpdateRequest_GW     UserUpdateRequest_GWPref = 2
)

var UserUpdateRequest_GWPref_name = map[int32]string{
	0: "NOPREF",
	1: "NOGW",
	2: "GW",
}
var UserUpdateRequest_GWPref_value = map[string]int32{
	"NOPREF": 0,
	"NOGW":   1,
	"GW":     2,
}

func (x UserUpdateRequest_GWPref) String() string {
	return proto.EnumName(UserUpdateRequest_GWPref_name, int32(x))
}
func (UserUpdateRequest_GWPref) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type UserUpdateRequest_StaticPref int32

const (
	UserUpdateRequest_NOPREFSTATIC UserUpdateRequest_StaticPref = 0
	UserUpdateRequest_NOSTATIC     UserUpdateRequest_StaticPref = 1
	UserUpdateRequest_STATIC       UserUpdateRequest_StaticPref = 2
)

var UserUpdateRequest_StaticPref_name = map[int32]string{
	0: "NOPREFSTATIC",
	1: "NOSTATIC",
	2: "STATIC",
}
var UserUpdateRequest_StaticPref_value = map[string]int32{
	"NOPREFSTATIC": 0,
	"NOSTATIC":     1,
	"STATIC":       2,
}

func (x UserUpdateRequest_StaticPref) String() string {
	return proto.EnumName(UserUpdateRequest_StaticPref_name, int32(x))
}
func (UserUpdateRequest_StaticPref) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1}
}

type UserUpdateRequest_AdminPref int32

const (
	UserUpdateRequest_NOPREFADMIN UserUpdateRequest_AdminPref = 0
	UserUpdateRequest_NOADMIN     UserUpdateRequest_AdminPref = 1
	UserUpdateRequest_ADMIN       UserUpdateRequest_AdminPref = 2
)

var UserUpdateRequest_AdminPref_name = map[int32]string{
	0: "NOPREFADMIN",
	1: "NOADMIN",
	2: "ADMIN",
}
var UserUpdateRequest_AdminPref_value = map[string]int32{
	"NOPREFADMIN": 0,
	"NOADMIN":     1,
	"ADMIN":       2,
}

func (x UserUpdateRequest_AdminPref) String() string {
	return proto.EnumName(UserUpdateRequest_AdminPref_name, int32(x))
}
func (UserUpdateRequest_AdminPref) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 2}
}

type UserListRequest struct {
}

func (m *UserListRequest) Reset()                    { *m = UserListRequest{} }
func (m *UserListRequest) String() string            { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()               {}
func (*UserListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UserCreateRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	NoGW     bool   `protobuf:"varint,3,opt,name=no_gW,json=noGW" json:"no_gW,omitempty"`
	HostId   uint32 `protobuf:"varint,4,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	IsAdmin  bool   `protobuf:"varint,5,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
}

func (m *UserCreateRequest) Reset()                    { *m = UserCreateRequest{} }
func (m *UserCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*UserCreateRequest) ProtoMessage()               {}
func (*UserCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserCreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserCreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserCreateRequest) GetNoGW() bool {
	if m != nil {
		return m.NoGW
	}
	return false
}

func (m *UserCreateRequest) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserCreateRequest) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

type UserUpdateRequest struct {
	Username   string                       `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password   string                       `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Gwpref     UserUpdateRequest_GWPref     `protobuf:"varint,3,opt,name=gwpref,enum=pb.UserUpdateRequest_GWPref" json:"gwpref,omitempty"`
	HostId     uint32                       `protobuf:"varint,4,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	StaticPref UserUpdateRequest_StaticPref `protobuf:"varint,5,opt,name=static_pref,json=staticPref,enum=pb.UserUpdateRequest_StaticPref" json:"static_pref,omitempty"`
	AdminPref  UserUpdateRequest_AdminPref  `protobuf:"varint,6,opt,name=admin_pref,json=adminPref,enum=pb.UserUpdateRequest_AdminPref" json:"admin_pref,omitempty"`
}

func (m *UserUpdateRequest) Reset()                    { *m = UserUpdateRequest{} }
func (m *UserUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UserUpdateRequest) ProtoMessage()               {}
func (*UserUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserUpdateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserUpdateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserUpdateRequest) GetGwpref() UserUpdateRequest_GWPref {
	if m != nil {
		return m.Gwpref
	}
	return UserUpdateRequest_NOPREF
}

func (m *UserUpdateRequest) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserUpdateRequest) GetStaticPref() UserUpdateRequest_StaticPref {
	if m != nil {
		return m.StaticPref
	}
	return UserUpdateRequest_NOPREFSTATIC
}

func (m *UserUpdateRequest) GetAdminPref() UserUpdateRequest_AdminPref {
	if m != nil {
		return m.AdminPref
	}
	return UserUpdateRequest_NOPREFADMIN
}

type UserDeleteRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UserDeleteRequest) Reset()                    { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()               {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserDeleteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserRenewRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UserRenewRequest) Reset()                    { *m = UserRenewRequest{} }
func (m *UserRenewRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRenewRequest) ProtoMessage()               {}
func (*UserRenewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserRenewRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserGenConfigRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UserGenConfigRequest) Reset()                    { *m = UserGenConfigRequest{} }
func (m *UserGenConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*UserGenConfigRequest) ProtoMessage()               {}
func (*UserGenConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserGenConfigRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserResponse struct {
	Users []*UserResponse_User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserResponse) GetUsers() []*UserResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserResponse_User struct {
	Username           string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	ServerSerialNumber string `protobuf:"bytes,2,opt,name=server_serial_number,json=serverSerialNumber" json:"server_serial_number,omitempty"`
	Cert               string `protobuf:"bytes,3,opt,name=cert" json:"cert,omitempty"`
	CreatedAt          string `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	IpNet              string `protobuf:"bytes,5,opt,name=ip_net,json=ipNet" json:"ip_net,omitempty"`
	NoGw               bool   `protobuf:"varint,6,opt,name=no_gw,json=noGw" json:"no_gw,omitempty"`
	HostId             uint32 `protobuf:"varint,7,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	IsAdmin            bool   `protobuf:"varint,8,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
}

func (m *UserResponse_User) Reset()                    { *m = UserResponse_User{} }
func (m *UserResponse_User) String() string            { return proto.CompactTextString(m) }
func (*UserResponse_User) ProtoMessage()               {}
func (*UserResponse_User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *UserResponse_User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResponse_User) GetServerSerialNumber() string {
	if m != nil {
		return m.ServerSerialNumber
	}
	return ""
}

func (m *UserResponse_User) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *UserResponse_User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserResponse_User) GetIpNet() string {
	if m != nil {
		return m.IpNet
	}
	return ""
}

func (m *UserResponse_User) GetNoGw() bool {
	if m != nil {
		return m.NoGw
	}
	return false
}

func (m *UserResponse_User) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserResponse_User) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

type UserGenConfigResponse struct {
	ClientConfig string `protobuf:"bytes,1,opt,name=client_config,json=clientConfig" json:"client_config,omitempty"`
}

func (m *UserGenConfigResponse) Reset()                    { *m = UserGenConfigResponse{} }
func (m *UserGenConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*UserGenConfigResponse) ProtoMessage()               {}
func (*UserGenConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserGenConfigResponse) GetClientConfig() string {
	if m != nil {
		return m.ClientConfig
	}
	return ""
}

func init() {
	proto.RegisterType((*UserListRequest)(nil), "pb.UserListRequest")
	proto.RegisterType((*UserCreateRequest)(nil), "pb.UserCreateRequest")
	proto.RegisterType((*UserUpdateRequest)(nil), "pb.UserUpdateRequest")
	proto.RegisterType((*UserDeleteRequest)(nil), "pb.UserDeleteRequest")
	proto.RegisterType((*UserRenewRequest)(nil), "pb.UserRenewRequest")
	proto.RegisterType((*UserGenConfigRequest)(nil), "pb.UserGenConfigRequest")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
	proto.RegisterType((*UserResponse_User)(nil), "pb.UserResponse.User")
	proto.RegisterType((*UserGenConfigResponse)(nil), "pb.UserGenConfigResponse")
	proto.RegisterEnum("pb.UserUpdateRequest_GWPref", UserUpdateRequest_GWPref_name, UserUpdateRequest_GWPref_value)
	proto.RegisterEnum("pb.UserUpdateRequest_StaticPref", UserUpdateRequest_StaticPref_name, UserUpdateRequest_StaticPref_value)
	proto.RegisterEnum("pb.UserUpdateRequest_AdminPref", UserUpdateRequest_AdminPref_name, UserUpdateRequest_AdminPref_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Renew(ctx context.Context, in *UserRenewRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GenConfig(ctx context.Context, in *UserGenConfigRequest, opts ...grpc.CallOption) (*UserGenConfigResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Renew(ctx context.Context, in *UserRenewRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GenConfig(ctx context.Context, in *UserGenConfigRequest, opts ...grpc.CallOption) (*UserGenConfigResponse, error) {
	out := new(UserGenConfigResponse)
	err := grpc.Invoke(ctx, "/pb.UserService/GenConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	List(context.Context, *UserListRequest) (*UserResponse, error)
	Create(context.Context, *UserCreateRequest) (*UserResponse, error)
	Update(context.Context, *UserUpdateRequest) (*UserResponse, error)
	Delete(context.Context, *UserDeleteRequest) (*UserResponse, error)
	Renew(context.Context, *UserRenewRequest) (*UserResponse, error)
	GenConfig(context.Context, *UserGenConfigRequest) (*UserGenConfigResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Renew(ctx, req.(*UserRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GenConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGenConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GenConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GenConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GenConfig(ctx, req.(*UserGenConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _UserService_Renew_Handler,
		},
		{
			MethodName: "GenConfig",
			Handler:    _UserService_GenConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xd1, 0x6e, 0xda, 0x4a,
	0x10, 0x86, 0x63, 0x03, 0x06, 0x0f, 0x24, 0x38, 0x1b, 0x50, 0x1c, 0x94, 0xe8, 0x20, 0x1f, 0xe9,
	0x08, 0xe5, 0x48, 0x70, 0x0e, 0xcd, 0x45, 0x15, 0x55, 0x95, 0x50, 0xd2, 0xa2, 0x54, 0x8d, 0x89,
	0x4c, 0x22, 0x2e, 0x2d, 0x83, 0x37, 0xd4, 0x12, 0x59, 0xbb, 0xbb, 0x26, 0xdc, 0xf7, 0x15, 0xda,
	0x57, 0xe8, 0x23, 0xf4, 0x49, 0xfa, 0x0a, 0x95, 0x7a, 0xd5, 0x77, 0xa8, 0x76, 0xd7, 0x36, 0x21,
	0x25, 0x15, 0x17, 0xbd, 0x9b, 0xd9, 0x99, 0xff, 0x63, 0x3c, 0x3b, 0xb3, 0x00, 0xcc, 0x19, 0xa6,
	0xed, 0x88, 0x86, 0x71, 0x88, 0xd4, 0x68, 0xdc, 0x38, 0x9c, 0x86, 0xe1, 0x74, 0x86, 0x3b, 0x5e,
	0x14, 0x74, 0x3c, 0x42, 0xc2, 0xd8, 0x8b, 0x83, 0x90, 0x30, 0x99, 0x61, 0xed, 0x42, 0xf5, 0x86,
	0x61, 0xfa, 0x36, 0x60, 0xb1, 0x83, 0xdf, 0xcf, 0x31, 0x8b, 0xad, 0x4f, 0x0a, 0xec, 0xf2, 0xb3,
	0x33, 0x8a, 0xbd, 0x18, 0x27, 0xa7, 0xa8, 0x01, 0x25, 0x0e, 0x26, 0xde, 0x1d, 0x36, 0x95, 0xa6,
	0xd2, 0xd2, 0x9d, 0xcc, 0xe7, 0xb1, 0xc8, 0x63, 0x6c, 0x11, 0x52, 0xdf, 0x54, 0x65, 0x2c, 0xf5,
	0xd1, 0x1e, 0x14, 0x48, 0xe8, 0x4e, 0x47, 0x66, 0xae, 0xa9, 0xb4, 0x4a, 0x4e, 0x9e, 0x84, 0xfd,
	0x11, 0xda, 0x87, 0xe2, 0xbb, 0x90, 0xc5, 0x6e, 0xe0, 0x9b, 0xf9, 0xa6, 0xd2, 0xda, 0x76, 0x34,
	0xee, 0x5e, 0xf8, 0xe8, 0x00, 0x4a, 0x01, 0x73, 0x3d, 0xff, 0x2e, 0x20, 0x66, 0x41, 0x08, 0x8a,
	0x01, 0xeb, 0x71, 0xd7, 0xfa, 0x92, 0x93, 0x65, 0xdd, 0x44, 0xfe, 0x1f, 0x28, 0xeb, 0x04, 0xb4,
	0xe9, 0x22, 0xa2, 0xf8, 0x56, 0xd4, 0xb5, 0xd3, 0x3d, 0x6c, 0x47, 0xe3, 0xf6, 0x2f, 0xf8, 0x76,
	0x7f, 0x74, 0x45, 0xf1, 0xad, 0x93, 0xe4, 0x3e, 0x5d, 0x77, 0x0f, 0xca, 0x8c, 0x37, 0x76, 0xe2,
	0x0a, 0x66, 0x41, 0x30, 0x9b, 0xeb, 0x99, 0x43, 0x91, 0x28, 0xb8, 0xc0, 0x32, 0x1b, 0xbd, 0x04,
	0x10, 0xdf, 0x2d, 0x09, 0x9a, 0x20, 0xfc, 0xb5, 0x9e, 0x20, 0x1a, 0x22, 0x00, 0xba, 0x97, 0x9a,
	0xd6, 0x3f, 0xa0, 0xc9, 0x6a, 0x11, 0x80, 0x66, 0x0f, 0xae, 0x9c, 0x57, 0xaf, 0x8d, 0x2d, 0x54,
	0x82, 0xbc, 0x3d, 0xe8, 0x8f, 0x0c, 0x05, 0x69, 0xa0, 0xf6, 0x47, 0x86, 0x6a, 0x3d, 0x07, 0x58,
	0x56, 0x80, 0x0c, 0xa8, 0xc8, 0xdc, 0xe1, 0x75, 0xef, 0xfa, 0xe2, 0xcc, 0xd8, 0x42, 0x15, 0x28,
	0xd9, 0x83, 0xc4, 0x53, 0x38, 0x2b, 0xb1, 0x55, 0xeb, 0x04, 0xf4, 0xec, 0x97, 0x51, 0x15, 0xca,
	0x52, 0xd8, 0x3b, 0xbf, 0xbc, 0xb0, 0x8d, 0x2d, 0x54, 0x86, 0xa2, 0x3d, 0x90, 0x8e, 0x82, 0x74,
	0x28, 0x48, 0x53, 0xb5, 0x3a, 0xf2, 0xda, 0xce, 0xf1, 0x0c, 0x6f, 0x74, 0x6d, 0x56, 0x1b, 0x0c,
	0x2e, 0x70, 0x30, 0xc1, 0x8b, 0x4d, 0xf2, 0xbb, 0x50, 0xe3, 0xf9, 0x7d, 0x4c, 0xce, 0x42, 0x72,
	0x1b, 0x4c, 0x37, 0xd1, 0x7c, 0x56, 0xa1, 0x22, 0x7f, 0x84, 0x45, 0x21, 0x61, 0x18, 0xfd, 0x0b,
	0x05, 0x1e, 0x64, 0xa6, 0xd2, 0xcc, 0xb5, 0xca, 0xdd, 0x7a, 0xda, 0xf8, 0x34, 0x41, 0x3a, 0x32,
	0xa7, 0xf1, 0x5d, 0x81, 0x3c, 0xf7, 0x7f, 0x3b, 0x7d, 0xff, 0x41, 0x8d, 0x61, 0x7a, 0x8f, 0xa9,
	0xcb, 0x30, 0x0d, 0xbc, 0x99, 0x4b, 0xe6, 0x77, 0x63, 0x4c, 0x93, 0x49, 0x44, 0x32, 0x36, 0x14,
	0x21, 0x5b, 0x44, 0x10, 0x82, 0xfc, 0x04, 0xd3, 0x58, 0x4c, 0xa4, 0xee, 0x08, 0x1b, 0x1d, 0x01,
	0x4c, 0xc4, 0x1e, 0xfa, 0xae, 0x17, 0x8b, 0xa1, 0xd3, 0x1d, 0x3d, 0x39, 0xe9, 0xc5, 0xa8, 0x0e,
	0x5a, 0x10, 0xb9, 0x04, 0xc7, 0x62, 0xe4, 0x74, 0xa7, 0x10, 0x44, 0x36, 0x8e, 0xd3, 0xa5, 0x5b,
	0x88, 0x31, 0x92, 0x4b, 0xb7, 0x78, 0x38, 0xbc, 0xc5, 0x27, 0x97, 0xae, 0xb4, 0xba, 0x74, 0x2f,
	0xa0, 0xfe, 0xa8, 0xb7, 0x49, 0xbf, 0xfe, 0x86, 0xed, 0xc9, 0x2c, 0xc0, 0x24, 0x76, 0x27, 0x22,
	0x90, 0x7c, 0x7e, 0x45, 0x1e, 0xca, 0xe4, 0xee, 0x8f, 0x1c, 0x94, 0xb9, 0x7c, 0x88, 0xe9, 0x7d,
	0x30, 0xc1, 0xe8, 0x1c, 0xf2, 0xfc, 0xa1, 0x41, 0x7b, 0x69, 0x77, 0x1f, 0x3c, 0x3b, 0x0d, 0xe3,
	0x71, 0xcb, 0xad, 0xfa, 0x87, 0xaf, 0xdf, 0x3e, 0xaa, 0x55, 0xb4, 0xdd, 0xb9, 0xff, 0xbf, 0xc3,
	0xfb, 0xda, 0x99, 0x71, 0xf5, 0x25, 0x68, 0xf2, 0x69, 0x42, 0xd9, 0x2d, 0xad, 0x3c, 0x55, 0x6b,
	0x48, 0x0d, 0x41, 0xaa, 0x59, 0xd5, 0x8c, 0x24, 0x5b, 0x78, 0xaa, 0x1c, 0x73, 0x9c, 0xdc, 0xae,
	0x25, 0x6e, 0x65, 0xdb, 0x36, 0xc2, 0xcd, 0x85, 0x22, 0xc1, 0xc9, 0x51, 0x5f, 0xe2, 0x56, 0x46,
	0x7f, 0x23, 0x9c, 0x2f, 0x14, 0x1c, 0xf7, 0x06, 0x0a, 0x62, 0x11, 0x50, 0x6d, 0x29, 0x5b, 0xee,
	0xc5, 0x1a, 0xd8, 0x81, 0x80, 0xed, 0x59, 0x3b, 0x19, 0x8c, 0x72, 0x01, 0x67, 0xb9, 0xa0, 0x67,
	0x17, 0x89, 0xcc, 0x54, 0xf9, 0x78, 0x6f, 0x1a, 0x07, 0x6b, 0x22, 0x09, 0xfc, 0x48, 0xc0, 0xf7,
	0x2d, 0x94, 0xc1, 0xa7, 0x98, 0xc8, 0x01, 0x38, 0x55, 0x8e, 0xc7, 0x9a, 0xf8, 0x4f, 0x79, 0xf6,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x46, 0x61, 0x32, 0xd9, 0x83, 0x06, 0x00, 0x00,
}
