// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: protobuf/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Auth_LogIn_FullMethodName                = "/auth.Auth/LogIn"
	Auth_RefreshToken_FullMethodName         = "/auth.Auth/RefreshToken"
	Auth_LogOut_FullMethodName               = "/auth.Auth/LogOut"
	Auth_CloseSessions_FullMethodName        = "/auth.Auth/CloseSessions"
	Auth_AddPermission_FullMethodName        = "/auth.Auth/AddPermission"
	Auth_RevokePermission_FullMethodName     = "/auth.Auth/RevokePermission"
	Auth_GetPermissions_FullMethodName       = "/auth.Auth/GetPermissions"
	Auth_AddRolePermission_FullMethodName    = "/auth.Auth/AddRolePermission"
	Auth_RevokeRolePermission_FullMethodName = "/auth.Auth/RevokeRolePermission"
	Auth_GetRolePermissions_FullMethodName   = "/auth.Auth/GetRolePermissions"
	Auth_AddRole_FullMethodName              = "/auth.Auth/AddRole"
	Auth_RevokeRole_FullMethodName           = "/auth.Auth/RevokeRole"
	Auth_GetRoles_FullMethodName             = "/auth.Auth/GetRoles"
	Auth_AddUserRole_FullMethodName          = "/auth.Auth/AddUserRole"
	Auth_RevokeUserRole_FullMethodName       = "/auth.Auth/RevokeUserRole"
	Auth_GetUserRoles_FullMethodName         = "/auth.Auth/GetUserRoles"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error)
	CloseSessions(ctx context.Context, in *CloseSessionsRequest, opts ...grpc.CallOption) (*CloseSessionsResponse, error)
	AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error)
	RevokePermission(ctx context.Context, in *RevokePermissionRequest, opts ...grpc.CallOption) (*RevokePermissionResponse, error)
	GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error)
	AddRolePermission(ctx context.Context, in *AddRolePermissionRequest, opts ...grpc.CallOption) (*AddRolePermissionResponse, error)
	RevokeRolePermission(ctx context.Context, in *RevokeRolePermissionRequest, opts ...grpc.CallOption) (*RevokeRolePermissionResponse, error)
	GetRolePermissions(ctx context.Context, in *GetRolePermissionsRequest, opts ...grpc.CallOption) (*GetRolePermissionsResponse, error)
	AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error)
	RevokeRole(ctx context.Context, in *RevokeRoleRequest, opts ...grpc.CallOption) (*RevokeRoleResponse, error)
	GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error)
	AddUserRole(ctx context.Context, in *AddUserRoleRequest, opts ...grpc.CallOption) (*AddUserRoleResponse, error)
	RevokeUserRole(ctx context.Context, in *RevokeUserRoleRequest, opts ...grpc.CallOption) (*RevokeUserRoleResponse, error)
	GetUserRoles(ctx context.Context, in *GetUserRolesRequest, opts ...grpc.CallOption) (*GetUserRolesResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, Auth_LogIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, Auth_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogOutResponse)
	err := c.cc.Invoke(ctx, Auth_LogOut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CloseSessions(ctx context.Context, in *CloseSessionsRequest, opts ...grpc.CallOption) (*CloseSessionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseSessionsResponse)
	err := c.cc.Invoke(ctx, Auth_CloseSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPermissionResponse)
	err := c.cc.Invoke(ctx, Auth_AddPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RevokePermission(ctx context.Context, in *RevokePermissionRequest, opts ...grpc.CallOption) (*RevokePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokePermissionResponse)
	err := c.cc.Invoke(ctx, Auth_RevokePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPermissionsResponse)
	err := c.cc.Invoke(ctx, Auth_GetPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddRolePermission(ctx context.Context, in *AddRolePermissionRequest, opts ...grpc.CallOption) (*AddRolePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddRolePermissionResponse)
	err := c.cc.Invoke(ctx, Auth_AddRolePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RevokeRolePermission(ctx context.Context, in *RevokeRolePermissionRequest, opts ...grpc.CallOption) (*RevokeRolePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeRolePermissionResponse)
	err := c.cc.Invoke(ctx, Auth_RevokeRolePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetRolePermissions(ctx context.Context, in *GetRolePermissionsRequest, opts ...grpc.CallOption) (*GetRolePermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRolePermissionsResponse)
	err := c.cc.Invoke(ctx, Auth_GetRolePermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddRole(ctx context.Context, in *AddRoleRequest, opts ...grpc.CallOption) (*AddRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddRoleResponse)
	err := c.cc.Invoke(ctx, Auth_AddRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RevokeRole(ctx context.Context, in *RevokeRoleRequest, opts ...grpc.CallOption) (*RevokeRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeRoleResponse)
	err := c.cc.Invoke(ctx, Auth_RevokeRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRolesResponse)
	err := c.cc.Invoke(ctx, Auth_GetRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddUserRole(ctx context.Context, in *AddUserRoleRequest, opts ...grpc.CallOption) (*AddUserRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserRoleResponse)
	err := c.cc.Invoke(ctx, Auth_AddUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RevokeUserRole(ctx context.Context, in *RevokeUserRoleRequest, opts ...grpc.CallOption) (*RevokeUserRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeUserRoleResponse)
	err := c.cc.Invoke(ctx, Auth_RevokeUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserRoles(ctx context.Context, in *GetUserRolesRequest, opts ...grpc.CallOption) (*GetUserRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserRolesResponse)
	err := c.cc.Invoke(ctx, Auth_GetUserRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility.
type AuthServer interface {
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error)
	CloseSessions(context.Context, *CloseSessionsRequest) (*CloseSessionsResponse, error)
	AddPermission(context.Context, *AddPermissionRequest) (*AddPermissionResponse, error)
	RevokePermission(context.Context, *RevokePermissionRequest) (*RevokePermissionResponse, error)
	GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error)
	AddRolePermission(context.Context, *AddRolePermissionRequest) (*AddRolePermissionResponse, error)
	RevokeRolePermission(context.Context, *RevokeRolePermissionRequest) (*RevokeRolePermissionResponse, error)
	GetRolePermissions(context.Context, *GetRolePermissionsRequest) (*GetRolePermissionsResponse, error)
	AddRole(context.Context, *AddRoleRequest) (*AddRoleResponse, error)
	RevokeRole(context.Context, *RevokeRoleRequest) (*RevokeRoleResponse, error)
	GetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error)
	AddUserRole(context.Context, *AddUserRoleRequest) (*AddUserRoleResponse, error)
	RevokeUserRole(context.Context, *RevokeUserRoleRequest) (*RevokeUserRoleResponse, error)
	GetUserRoles(context.Context, *GetUserRolesRequest) (*GetUserRolesResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) LogIn(context.Context, *LogInRequest) (*LogInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedAuthServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServer) LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedAuthServer) CloseSessions(context.Context, *CloseSessionsRequest) (*CloseSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSessions not implemented")
}
func (UnimplementedAuthServer) AddPermission(context.Context, *AddPermissionRequest) (*AddPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermission not implemented")
}
func (UnimplementedAuthServer) RevokePermission(context.Context, *RevokePermissionRequest) (*RevokePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokePermission not implemented")
}
func (UnimplementedAuthServer) GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedAuthServer) AddRolePermission(context.Context, *AddRolePermissionRequest) (*AddRolePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRolePermission not implemented")
}
func (UnimplementedAuthServer) RevokeRolePermission(context.Context, *RevokeRolePermissionRequest) (*RevokeRolePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRolePermission not implemented")
}
func (UnimplementedAuthServer) GetRolePermissions(context.Context, *GetRolePermissionsRequest) (*GetRolePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolePermissions not implemented")
}
func (UnimplementedAuthServer) AddRole(context.Context, *AddRoleRequest) (*AddRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedAuthServer) RevokeRole(context.Context, *RevokeRoleRequest) (*RevokeRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRole not implemented")
}
func (UnimplementedAuthServer) GetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedAuthServer) AddUserRole(context.Context, *AddUserRoleRequest) (*AddUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserRole not implemented")
}
func (UnimplementedAuthServer) RevokeUserRole(context.Context, *RevokeUserRoleRequest) (*RevokeUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeUserRole not implemented")
}
func (UnimplementedAuthServer) GetUserRoles(context.Context, *GetUserRolesRequest) (*GetUserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRoles not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}
func (UnimplementedAuthServer) testEmbeddedByValue()              {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LogIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LogOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LogOut(ctx, req.(*LogOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CloseSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CloseSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_CloseSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CloseSessions(ctx, req.(*CloseSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddPermission(ctx, req.(*AddPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RevokePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RevokePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RevokePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RevokePermission(ctx, req.(*RevokePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetPermissions(ctx, req.(*GetPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddRolePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddRolePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddRolePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddRolePermission(ctx, req.(*AddRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RevokeRolePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RevokeRolePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RevokeRolePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RevokeRolePermission(ctx, req.(*RevokeRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetRolePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetRolePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetRolePermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetRolePermissions(ctx, req.(*GetRolePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddRole(ctx, req.(*AddRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RevokeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RevokeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RevokeRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RevokeRole(ctx, req.(*RevokeRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetRoles(ctx, req.(*GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddUserRole(ctx, req.(*AddUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RevokeUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RevokeUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RevokeUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RevokeUserRole(ctx, req.(*RevokeUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetUserRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserRoles(ctx, req.(*GetUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogIn",
			Handler:    _Auth_LogIn_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Auth_RefreshToken_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _Auth_LogOut_Handler,
		},
		{
			MethodName: "CloseSessions",
			Handler:    _Auth_CloseSessions_Handler,
		},
		{
			MethodName: "AddPermission",
			Handler:    _Auth_AddPermission_Handler,
		},
		{
			MethodName: "RevokePermission",
			Handler:    _Auth_RevokePermission_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _Auth_GetPermissions_Handler,
		},
		{
			MethodName: "AddRolePermission",
			Handler:    _Auth_AddRolePermission_Handler,
		},
		{
			MethodName: "RevokeRolePermission",
			Handler:    _Auth_RevokeRolePermission_Handler,
		},
		{
			MethodName: "GetRolePermissions",
			Handler:    _Auth_GetRolePermissions_Handler,
		},
		{
			MethodName: "AddRole",
			Handler:    _Auth_AddRole_Handler,
		},
		{
			MethodName: "RevokeRole",
			Handler:    _Auth_RevokeRole_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _Auth_GetRoles_Handler,
		},
		{
			MethodName: "AddUserRole",
			Handler:    _Auth_AddUserRole_Handler,
		},
		{
			MethodName: "RevokeUserRole",
			Handler:    _Auth_RevokeUserRole_Handler,
		},
		{
			MethodName: "GetUserRoles",
			Handler:    _Auth_GetUserRoles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/auth.proto",
}
