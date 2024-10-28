// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: protobuf/user.proto

package user

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
	UsersService_SignUp_FullMethodName                = "/users_service.UsersService/SignUp"
	UsersService_UpdateProfile_FullMethodName         = "/users_service.UsersService/UpdateProfile"
	UsersService_ChangeUsername_FullMethodName        = "/users_service.UsersService/ChangeUsername"
	UsersService_ChangePassword_FullMethodName        = "/users_service.UsersService/ChangePassword"
	UsersService_ChangeEmail_FullMethodName           = "/users_service.UsersService/ChangeEmail"
	UsersService_VerifyEmail_FullMethodName           = "/users_service.UsersService/VerifyEmail"
	UsersService_GetActiveEmails_FullMethodName       = "/users_service.UsersService/GetActiveEmails"
	UsersService_ChangePhoneNumber_FullMethodName     = "/users_service.UsersService/ChangePhoneNumber"
	UsersService_VerifyPhoneNumber_FullMethodName     = "/users_service.UsersService/VerifyPhoneNumber"
	UsersService_GetActivePhoneNumbers_FullMethodName = "/users_service.UsersService/GetActivePhoneNumbers"
	UsersService_ForgotPassword_FullMethodName        = "/users_service.UsersService/ForgotPassword"
	UsersService_ResetPassword_FullMethodName         = "/users_service.UsersService/ResetPassword"
	UsersService_DeleteUser_FullMethodName            = "/users_service.UsersService/DeleteUser"
)

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
	ChangeUsername(ctx context.Context, in *ChangeUsernameRequest, opts ...grpc.CallOption) (*ChangeUsernameResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...grpc.CallOption) (*ChangeEmailResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	GetActiveEmails(ctx context.Context, in *GetActiveEmailsRequest, opts ...grpc.CallOption) (*GetActiveEmailsResponse, error)
	ChangePhoneNumber(ctx context.Context, in *ChangePhoneNumberRequest, opts ...grpc.CallOption) (*ChangePhoneNumberResponse, error)
	VerifyPhoneNumber(ctx context.Context, in *VerifyPhoneNumberRequest, opts ...grpc.CallOption) (*VerifyPhoneNumberResponse, error)
	GetActivePhoneNumbers(ctx context.Context, in *GetActivePhoneNumbersRequest, opts ...grpc.CallOption) (*GetActivePhoneNumbersResponse, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, UsersService_SignUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProfileResponse)
	err := c.cc.Invoke(ctx, UsersService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ChangeUsername(ctx context.Context, in *ChangeUsernameRequest, opts ...grpc.CallOption) (*ChangeUsernameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeUsernameResponse)
	err := c.cc.Invoke(ctx, UsersService_ChangeUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, UsersService_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...grpc.CallOption) (*ChangeEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeEmailResponse)
	err := c.cc.Invoke(ctx, UsersService_ChangeEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, UsersService_VerifyEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetActiveEmails(ctx context.Context, in *GetActiveEmailsRequest, opts ...grpc.CallOption) (*GetActiveEmailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActiveEmailsResponse)
	err := c.cc.Invoke(ctx, UsersService_GetActiveEmails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ChangePhoneNumber(ctx context.Context, in *ChangePhoneNumberRequest, opts ...grpc.CallOption) (*ChangePhoneNumberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePhoneNumberResponse)
	err := c.cc.Invoke(ctx, UsersService_ChangePhoneNumber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) VerifyPhoneNumber(ctx context.Context, in *VerifyPhoneNumberRequest, opts ...grpc.CallOption) (*VerifyPhoneNumberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPhoneNumberResponse)
	err := c.cc.Invoke(ctx, UsersService_VerifyPhoneNumber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetActivePhoneNumbers(ctx context.Context, in *GetActivePhoneNumbersRequest, opts ...grpc.CallOption) (*GetActivePhoneNumbersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActivePhoneNumbersResponse)
	err := c.cc.Invoke(ctx, UsersService_GetActivePhoneNumbers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, UsersService_ForgotPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, UsersService_ResetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UsersService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility.
type UsersServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	ChangeUsername(context.Context, *ChangeUsernameRequest) (*ChangeUsernameResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	ChangeEmail(context.Context, *ChangeEmailRequest) (*ChangeEmailResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	GetActiveEmails(context.Context, *GetActiveEmailsRequest) (*GetActiveEmailsResponse, error)
	ChangePhoneNumber(context.Context, *ChangePhoneNumberRequest) (*ChangePhoneNumberResponse, error)
	VerifyPhoneNumber(context.Context, *VerifyPhoneNumberRequest) (*VerifyPhoneNumberResponse, error)
	GetActivePhoneNumbers(context.Context, *GetActivePhoneNumbersRequest) (*GetActivePhoneNumbersResponse, error)
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServiceServer struct{}

func (UnimplementedUsersServiceServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedUsersServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUsersServiceServer) ChangeUsername(context.Context, *ChangeUsernameRequest) (*ChangeUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUsername not implemented")
}
func (UnimplementedUsersServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUsersServiceServer) ChangeEmail(context.Context, *ChangeEmailRequest) (*ChangeEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeEmail not implemented")
}
func (UnimplementedUsersServiceServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedUsersServiceServer) GetActiveEmails(context.Context, *GetActiveEmailsRequest) (*GetActiveEmailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveEmails not implemented")
}
func (UnimplementedUsersServiceServer) ChangePhoneNumber(context.Context, *ChangePhoneNumberRequest) (*ChangePhoneNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePhoneNumber not implemented")
}
func (UnimplementedUsersServiceServer) VerifyPhoneNumber(context.Context, *VerifyPhoneNumberRequest) (*VerifyPhoneNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPhoneNumber not implemented")
}
func (UnimplementedUsersServiceServer) GetActivePhoneNumbers(context.Context, *GetActivePhoneNumbersRequest) (*GetActivePhoneNumbersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivePhoneNumbers not implemented")
}
func (UnimplementedUsersServiceServer) ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedUsersServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUsersServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}
func (UnimplementedUsersServiceServer) testEmbeddedByValue()                      {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	// If the following call pancis, it indicates UnimplementedUsersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ChangeUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ChangeUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_ChangeUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ChangeUsername(ctx, req.(*ChangeUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ChangeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ChangeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_ChangeEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ChangeEmail(ctx, req.(*ChangeEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetActiveEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveEmailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetActiveEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_GetActiveEmails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetActiveEmails(ctx, req.(*GetActiveEmailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ChangePhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ChangePhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_ChangePhoneNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ChangePhoneNumber(ctx, req.(*ChangePhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_VerifyPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).VerifyPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_VerifyPhoneNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).VerifyPhoneNumber(ctx, req.(*VerifyPhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetActivePhoneNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivePhoneNumbersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetActivePhoneNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_GetActivePhoneNumbers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetActivePhoneNumbers(ctx, req.(*GetActivePhoneNumbersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_ForgotPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users_service.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _UsersService_SignUp_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UsersService_UpdateProfile_Handler,
		},
		{
			MethodName: "ChangeUsername",
			Handler:    _UsersService_ChangeUsername_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UsersService_ChangePassword_Handler,
		},
		{
			MethodName: "ChangeEmail",
			Handler:    _UsersService_ChangeEmail_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _UsersService_VerifyEmail_Handler,
		},
		{
			MethodName: "GetActiveEmails",
			Handler:    _UsersService_GetActiveEmails_Handler,
		},
		{
			MethodName: "ChangePhoneNumber",
			Handler:    _UsersService_ChangePhoneNumber_Handler,
		},
		{
			MethodName: "VerifyPhoneNumber",
			Handler:    _UsersService_VerifyPhoneNumber_Handler,
		},
		{
			MethodName: "GetActivePhoneNumbers",
			Handler:    _UsersService_GetActivePhoneNumbers_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _UsersService_ForgotPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UsersService_ResetPassword_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UsersService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/user.proto",
}