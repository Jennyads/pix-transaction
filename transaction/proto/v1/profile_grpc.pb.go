// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: profile.proto

package profile

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_CreateUser_FullMethodName = "/profile.proto.v2.UserService/CreateUser"
	UserService_FindUser_FullMethodName   = "/profile.proto.v2.UserService/FindUser"
	UserService_UpdateUser_FullMethodName = "/profile.proto.v2.UserService/UpdateUser"
	UserService_ListUsers_FullMethodName  = "/profile.proto.v2.UserService/ListUsers"
	UserService_DeleteUser_FullMethodName = "/profile.proto.v2.UserService/DeleteUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	FindUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	ListUsers(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUser, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_FindUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUser, error) {
	out := new(ListUser)
	err := c.cc.Invoke(ctx, UserService_ListUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *User) (*empty.Empty, error)
	FindUser(context.Context, *UserRequest) (*User, error)
	UpdateUser(context.Context, *User) (*empty.Empty, error)
	ListUsers(context.Context, *ListUserRequest) (*ListUser, error)
	DeleteUser(context.Context, *UserRequest) (*empty.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) FindUser(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) ListUsers(context.Context, *ListUserRequest) (*ListUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *UserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.proto.v2.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "FindUser",
			Handler:    _UserService_FindUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}

const (
	AccountService_CreateAccount_FullMethodName   = "/profile.proto.v2.AccountService/CreateAccount"
	AccountService_FindAccount_FullMethodName     = "/profile.proto.v2.AccountService/FindAccount"
	AccountService_UpdateAccount_FullMethodName   = "/profile.proto.v2.AccountService/UpdateAccount"
	AccountService_ListAccounts_FullMethodName    = "/profile.proto.v2.AccountService/ListAccounts"
	AccountService_DeleteAccount_FullMethodName   = "/profile.proto.v2.AccountService/DeleteAccount"
	AccountService_IsAccountActive_FullMethodName = "/profile.proto.v2.AccountService/IsAccountActive"
	AccountService_FindByKey_FullMethodName       = "/profile.proto.v2.AccountService/FindByKey"
)

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountResponse, error)
	FindAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	UpdateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*empty.Empty, error)
	ListAccounts(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccount, error)
	DeleteAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	IsAccountActive(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	FindByKey(ctx context.Context, in *FindByKeyRequest, opts ...grpc.CallOption) (*AccountResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) FindAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_FindAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountService_UpdateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ListAccounts(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccount, error) {
	out := new(ListAccount)
	err := c.cc.Invoke(ctx, AccountService_ListAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountService_DeleteAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) IsAccountActive(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, AccountService_IsAccountActive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) FindByKey(ctx context.Context, in *FindByKeyRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, AccountService_FindByKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	CreateAccount(context.Context, *Account) (*AccountResponse, error)
	FindAccount(context.Context, *AccountRequest) (*AccountResponse, error)
	UpdateAccount(context.Context, *Account) (*empty.Empty, error)
	ListAccounts(context.Context, *ListAccountRequest) (*ListAccount, error)
	DeleteAccount(context.Context, *AccountRequest) (*empty.Empty, error)
	IsAccountActive(context.Context, *AccountRequest) (*wrappers.BoolValue, error)
	FindByKey(context.Context, *FindByKeyRequest) (*AccountResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *Account) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) FindAccount(context.Context, *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAccount not implemented")
}
func (UnimplementedAccountServiceServer) UpdateAccount(context.Context, *Account) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountServiceServer) ListAccounts(context.Context, *ListAccountRequest) (*ListAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccount(context.Context, *AccountRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountServiceServer) IsAccountActive(context.Context, *AccountRequest) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAccountActive not implemented")
}
func (UnimplementedAccountServiceServer) FindByKey(context.Context, *FindByKeyRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByKey not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_FindAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).FindAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_FindAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).FindAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_ListAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ListAccounts(ctx, req.(*ListAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_IsAccountActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).IsAccountActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_IsAccountActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).IsAccountActive(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_FindByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).FindByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_FindByKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).FindByKey(ctx, req.(*FindByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.proto.v2.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "FindAccount",
			Handler:    _AccountService_FindAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _AccountService_ListAccounts_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountService_DeleteAccount_Handler,
		},
		{
			MethodName: "IsAccountActive",
			Handler:    _AccountService_IsAccountActive_Handler,
		},
		{
			MethodName: "FindByKey",
			Handler:    _AccountService_FindByKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}

const (
	KeysService_CreateKey_FullMethodName = "/profile.proto.v2.KeysService/CreateKey"
	KeysService_UpdateKey_FullMethodName = "/profile.proto.v2.KeysService/UpdateKey"
	KeysService_ListKey_FullMethodName   = "/profile.proto.v2.KeysService/ListKey"
	KeysService_DeleteKey_FullMethodName = "/profile.proto.v2.KeysService/DeleteKey"
)

// KeysServiceClient is the client API for KeysService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeysServiceClient interface {
	CreateKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*empty.Empty, error)
	ListKey(ctx context.Context, in *ListKeyRequest, opts ...grpc.CallOption) (*ListKeys, error)
	DeleteKey(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type keysServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeysServiceClient(cc grpc.ClientConnInterface) KeysServiceClient {
	return &keysServiceClient{cc}
}

func (c *keysServiceClient) CreateKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, KeysService_CreateKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysServiceClient) UpdateKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, KeysService_UpdateKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysServiceClient) ListKey(ctx context.Context, in *ListKeyRequest, opts ...grpc.CallOption) (*ListKeys, error) {
	out := new(ListKeys)
	err := c.cc.Invoke(ctx, KeysService_ListKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysServiceClient) DeleteKey(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, KeysService_DeleteKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeysServiceServer is the server API for KeysService service.
// All implementations must embed UnimplementedKeysServiceServer
// for forward compatibility
type KeysServiceServer interface {
	CreateKey(context.Context, *Key) (*empty.Empty, error)
	UpdateKey(context.Context, *Key) (*empty.Empty, error)
	ListKey(context.Context, *ListKeyRequest) (*ListKeys, error)
	DeleteKey(context.Context, *KeyRequest) (*empty.Empty, error)
	mustEmbedUnimplementedKeysServiceServer()
}

// UnimplementedKeysServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeysServiceServer struct {
}

func (UnimplementedKeysServiceServer) CreateKey(context.Context, *Key) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKey not implemented")
}
func (UnimplementedKeysServiceServer) UpdateKey(context.Context, *Key) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKey not implemented")
}
func (UnimplementedKeysServiceServer) ListKey(context.Context, *ListKeyRequest) (*ListKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKey not implemented")
}
func (UnimplementedKeysServiceServer) DeleteKey(context.Context, *KeyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (UnimplementedKeysServiceServer) mustEmbedUnimplementedKeysServiceServer() {}

// UnsafeKeysServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeysServiceServer will
// result in compilation errors.
type UnsafeKeysServiceServer interface {
	mustEmbedUnimplementedKeysServiceServer()
}

func RegisterKeysServiceServer(s grpc.ServiceRegistrar, srv KeysServiceServer) {
	s.RegisterService(&KeysService_ServiceDesc, srv)
}

func _KeysService_CreateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).CreateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeysService_CreateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).CreateKey(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeysService_UpdateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).UpdateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeysService_UpdateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).UpdateKey(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeysService_ListKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).ListKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeysService_ListKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).ListKey(ctx, req.(*ListKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeysService_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeysService_DeleteKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).DeleteKey(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeysService_ServiceDesc is the grpc.ServiceDesc for KeysService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeysService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.proto.v2.KeysService",
	HandlerType: (*KeysServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKey",
			Handler:    _KeysService_CreateKey_Handler,
		},
		{
			MethodName: "UpdateKey",
			Handler:    _KeysService_UpdateKey_Handler,
		},
		{
			MethodName: "ListKey",
			Handler:    _KeysService_ListKey_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _KeysService_DeleteKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}

const (
	PixTransactionService_SendPix_FullMethodName    = "/profile.proto.v2.PixTransactionService/SendPix"
	PixTransactionService_PixWebhook_FullMethodName = "/profile.proto.v2.PixTransactionService/PixWebhook"
)

// PixTransactionServiceClient is the client API for PixTransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PixTransactionServiceClient interface {
	SendPix(ctx context.Context, in *PixTransaction, opts ...grpc.CallOption) (*empty.Empty, error)
	PixWebhook(ctx context.Context, in *Webhook, opts ...grpc.CallOption) (*empty.Empty, error)
}

type pixTransactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPixTransactionServiceClient(cc grpc.ClientConnInterface) PixTransactionServiceClient {
	return &pixTransactionServiceClient{cc}
}

func (c *pixTransactionServiceClient) SendPix(ctx context.Context, in *PixTransaction, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, PixTransactionService_SendPix_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pixTransactionServiceClient) PixWebhook(ctx context.Context, in *Webhook, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, PixTransactionService_PixWebhook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PixTransactionServiceServer is the server API for PixTransactionService service.
// All implementations must embed UnimplementedPixTransactionServiceServer
// for forward compatibility
type PixTransactionServiceServer interface {
	SendPix(context.Context, *PixTransaction) (*empty.Empty, error)
	PixWebhook(context.Context, *Webhook) (*empty.Empty, error)
	mustEmbedUnimplementedPixTransactionServiceServer()
}

// UnimplementedPixTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPixTransactionServiceServer struct {
}

func (UnimplementedPixTransactionServiceServer) SendPix(context.Context, *PixTransaction) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPix not implemented")
}
func (UnimplementedPixTransactionServiceServer) PixWebhook(context.Context, *Webhook) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PixWebhook not implemented")
}
func (UnimplementedPixTransactionServiceServer) mustEmbedUnimplementedPixTransactionServiceServer() {}

// UnsafePixTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PixTransactionServiceServer will
// result in compilation errors.
type UnsafePixTransactionServiceServer interface {
	mustEmbedUnimplementedPixTransactionServiceServer()
}

func RegisterPixTransactionServiceServer(s grpc.ServiceRegistrar, srv PixTransactionServiceServer) {
	s.RegisterService(&PixTransactionService_ServiceDesc, srv)
}

func _PixTransactionService_SendPix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PixTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PixTransactionServiceServer).SendPix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PixTransactionService_SendPix_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PixTransactionServiceServer).SendPix(ctx, req.(*PixTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _PixTransactionService_PixWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Webhook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PixTransactionServiceServer).PixWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PixTransactionService_PixWebhook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PixTransactionServiceServer).PixWebhook(ctx, req.(*Webhook))
	}
	return interceptor(ctx, in, info, handler)
}

// PixTransactionService_ServiceDesc is the grpc.ServiceDesc for PixTransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PixTransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.proto.v2.PixTransactionService",
	HandlerType: (*PixTransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPix",
			Handler:    _PixTransactionService_SendPix_Handler,
		},
		{
			MethodName: "PixWebhook",
			Handler:    _PixTransactionService_PixWebhook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
