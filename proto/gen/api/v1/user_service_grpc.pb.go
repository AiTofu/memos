// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/v1/user_service.proto

package apiv1

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_ListUsers_FullMethodName             = "/memos.api.v1.UserService/ListUsers"
	UserService_GetUser_FullMethodName               = "/memos.api.v1.UserService/GetUser"
	UserService_GetUserByUsername_FullMethodName     = "/memos.api.v1.UserService/GetUserByUsername"
	UserService_GetUserAvatarBinary_FullMethodName   = "/memos.api.v1.UserService/GetUserAvatarBinary"
	UserService_CreateUser_FullMethodName            = "/memos.api.v1.UserService/CreateUser"
	UserService_UpdateUser_FullMethodName            = "/memos.api.v1.UserService/UpdateUser"
	UserService_DeleteUser_FullMethodName            = "/memos.api.v1.UserService/DeleteUser"
	UserService_ListAllUserStats_FullMethodName      = "/memos.api.v1.UserService/ListAllUserStats"
	UserService_GetUserStats_FullMethodName          = "/memos.api.v1.UserService/GetUserStats"
	UserService_GetUserSetting_FullMethodName        = "/memos.api.v1.UserService/GetUserSetting"
	UserService_UpdateUserSetting_FullMethodName     = "/memos.api.v1.UserService/UpdateUserSetting"
	UserService_ListUserAccessTokens_FullMethodName  = "/memos.api.v1.UserService/ListUserAccessTokens"
	UserService_CreateUserAccessToken_FullMethodName = "/memos.api.v1.UserService/CreateUserAccessToken"
	UserService_DeleteUserAccessToken_FullMethodName = "/memos.api.v1.UserService/DeleteUserAccessToken"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// ListUsers returns a list of users.
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// GetUser gets a user by name.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	// GetUserByUsername gets a user by username.
	GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*User, error)
	// GetUserAvatarBinary gets the avatar of a user.
	GetUserAvatarBinary(ctx context.Context, in *GetUserAvatarBinaryRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// UpdateUser updates a user.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// DeleteUser deletes a user.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListAllUserStats returns all user stats.
	ListAllUserStats(ctx context.Context, in *ListAllUserStatsRequest, opts ...grpc.CallOption) (*ListAllUserStatsResponse, error)
	// GetUserStats returns the stats of a user.
	GetUserStats(ctx context.Context, in *GetUserStatsRequest, opts ...grpc.CallOption) (*UserStats, error)
	// GetUserSetting gets the setting of a user.
	GetUserSetting(ctx context.Context, in *GetUserSettingRequest, opts ...grpc.CallOption) (*UserSetting, error)
	// UpdateUserSetting updates the setting of a user.
	UpdateUserSetting(ctx context.Context, in *UpdateUserSettingRequest, opts ...grpc.CallOption) (*UserSetting, error)
	// ListUserAccessTokens returns a list of access tokens for a user.
	ListUserAccessTokens(ctx context.Context, in *ListUserAccessTokensRequest, opts ...grpc.CallOption) (*ListUserAccessTokensResponse, error)
	// CreateUserAccessToken creates a new access token for a user.
	CreateUserAccessToken(ctx context.Context, in *CreateUserAccessTokenRequest, opts ...grpc.CallOption) (*UserAccessToken, error)
	// DeleteUserAccessToken deletes an access token for a user.
	DeleteUserAccessToken(ctx context.Context, in *DeleteUserAccessTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, UserService_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_GetUserByUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserAvatarBinary(ctx context.Context, in *GetUserAvatarBinaryRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, UserService_GetUserAvatarBinary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListAllUserStats(ctx context.Context, in *ListAllUserStatsRequest, opts ...grpc.CallOption) (*ListAllUserStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAllUserStatsResponse)
	err := c.cc.Invoke(ctx, UserService_ListAllUserStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserStats(ctx context.Context, in *GetUserStatsRequest, opts ...grpc.CallOption) (*UserStats, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserStats)
	err := c.cc.Invoke(ctx, UserService_GetUserStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserSetting(ctx context.Context, in *GetUserSettingRequest, opts ...grpc.CallOption) (*UserSetting, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserSetting)
	err := c.cc.Invoke(ctx, UserService_GetUserSetting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserSetting(ctx context.Context, in *UpdateUserSettingRequest, opts ...grpc.CallOption) (*UserSetting, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserSetting)
	err := c.cc.Invoke(ctx, UserService_UpdateUserSetting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUserAccessTokens(ctx context.Context, in *ListUserAccessTokensRequest, opts ...grpc.CallOption) (*ListUserAccessTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserAccessTokensResponse)
	err := c.cc.Invoke(ctx, UserService_ListUserAccessTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUserAccessToken(ctx context.Context, in *CreateUserAccessTokenRequest, opts ...grpc.CallOption) (*UserAccessToken, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserAccessToken)
	err := c.cc.Invoke(ctx, UserService_CreateUserAccessToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserAccessToken(ctx context.Context, in *DeleteUserAccessTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteUserAccessToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	// ListUsers returns a list of users.
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// GetUser gets a user by name.
	GetUser(context.Context, *GetUserRequest) (*User, error)
	// GetUserByUsername gets a user by username.
	GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*User, error)
	// GetUserAvatarBinary gets the avatar of a user.
	GetUserAvatarBinary(context.Context, *GetUserAvatarBinaryRequest) (*httpbody.HttpBody, error)
	// CreateUser creates a new user.
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// UpdateUser updates a user.
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// DeleteUser deletes a user.
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	// ListAllUserStats returns all user stats.
	ListAllUserStats(context.Context, *ListAllUserStatsRequest) (*ListAllUserStatsResponse, error)
	// GetUserStats returns the stats of a user.
	GetUserStats(context.Context, *GetUserStatsRequest) (*UserStats, error)
	// GetUserSetting gets the setting of a user.
	GetUserSetting(context.Context, *GetUserSettingRequest) (*UserSetting, error)
	// UpdateUserSetting updates the setting of a user.
	UpdateUserSetting(context.Context, *UpdateUserSettingRequest) (*UserSetting, error)
	// ListUserAccessTokens returns a list of access tokens for a user.
	ListUserAccessTokens(context.Context, *ListUserAccessTokensRequest) (*ListUserAccessTokensResponse, error)
	// CreateUserAccessToken creates a new access token for a user.
	CreateUserAccessToken(context.Context, *CreateUserAccessTokenRequest) (*UserAccessToken, error)
	// DeleteUserAccessToken deletes an access token for a user.
	DeleteUserAccessToken(context.Context, *DeleteUserAccessTokenRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUserServiceServer) GetUserAvatarBinary(context.Context, *GetUserAvatarBinaryRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAvatarBinary not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) ListAllUserStats(context.Context, *ListAllUserStatsRequest) (*ListAllUserStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllUserStats not implemented")
}
func (UnimplementedUserServiceServer) GetUserStats(context.Context, *GetUserStatsRequest) (*UserStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStats not implemented")
}
func (UnimplementedUserServiceServer) GetUserSetting(context.Context, *GetUserSettingRequest) (*UserSetting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSetting not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserSetting(context.Context, *UpdateUserSettingRequest) (*UserSetting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSetting not implemented")
}
func (UnimplementedUserServiceServer) ListUserAccessTokens(context.Context, *ListUserAccessTokensRequest) (*ListUserAccessTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserAccessTokens not implemented")
}
func (UnimplementedUserServiceServer) CreateUserAccessToken(context.Context, *CreateUserAccessTokenRequest) (*UserAccessToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserAccessToken not implemented")
}
func (UnimplementedUserServiceServer) DeleteUserAccessToken(context.Context, *DeleteUserAccessTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAccessToken not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
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
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByUsername(ctx, req.(*GetUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserAvatarBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAvatarBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserAvatarBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserAvatarBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserAvatarBinary(ctx, req.(*GetUserAvatarBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
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
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
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
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
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
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListAllUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllUserStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListAllUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListAllUserStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListAllUserStats(ctx, req.(*ListAllUserStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserStats(ctx, req.(*GetUserStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserSetting(ctx, req.(*GetUserSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserSetting(ctx, req.(*UpdateUserSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUserAccessTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserAccessTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUserAccessTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListUserAccessTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUserAccessTokens(ctx, req.(*ListUserAccessTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUserAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUserAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUserAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUserAccessToken(ctx, req.(*CreateUserAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUserAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserAccessToken(ctx, req.(*DeleteUserAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memos.api.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _UserService_GetUserByUsername_Handler,
		},
		{
			MethodName: "GetUserAvatarBinary",
			Handler:    _UserService_GetUserAvatarBinary_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "ListAllUserStats",
			Handler:    _UserService_ListAllUserStats_Handler,
		},
		{
			MethodName: "GetUserStats",
			Handler:    _UserService_GetUserStats_Handler,
		},
		{
			MethodName: "GetUserSetting",
			Handler:    _UserService_GetUserSetting_Handler,
		},
		{
			MethodName: "UpdateUserSetting",
			Handler:    _UserService_UpdateUserSetting_Handler,
		},
		{
			MethodName: "ListUserAccessTokens",
			Handler:    _UserService_ListUserAccessTokens_Handler,
		},
		{
			MethodName: "CreateUserAccessToken",
			Handler:    _UserService_CreateUserAccessToken_Handler,
		},
		{
			MethodName: "DeleteUserAccessToken",
			Handler:    _UserService_DeleteUserAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/user_service.proto",
}
