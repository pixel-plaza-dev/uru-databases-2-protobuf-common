// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/pixel_plaza/payment.proto

package payment

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
	Payment_AddOrderPayment_FullMethodName             = "/pixel_plaza.Payment/AddOrderPayment"
	Payment_GetOrderPayments_FullMethodName            = "/pixel_plaza.Payment/GetOrderPayments"
	Payment_AddBranchRentPayment_FullMethodName        = "/pixel_plaza.Payment/AddBranchRentPayment"
	Payment_GetBranchRentPayments_FullMethodName       = "/pixel_plaza.Payment/GetBranchRentPayments"
	Payment_GetBranchRentsPayments_FullMethodName      = "/pixel_plaza.Payment/GetBranchRentsPayments"
	Payment_VerifyPayment_FullMethodName               = "/pixel_plaza.Payment/VerifyPayment"
	Payment_AddPaymentAccount_FullMethodName           = "/pixel_plaza.Payment/AddPaymentAccount"
	Payment_GetPaymentAccounts_FullMethodName          = "/pixel_plaza.Payment/GetPaymentAccounts"
	Payment_GetActivePaymentAccounts_FullMethodName    = "/pixel_plaza.Payment/GetActivePaymentAccounts"
	Payment_GetSuspendedPaymentAccounts_FullMethodName = "/pixel_plaza.Payment/GetSuspendedPaymentAccounts"
	Payment_SuspendPaymentAccount_FullMethodName       = "/pixel_plaza.Payment/SuspendPaymentAccount"
	Payment_ActivatePaymentAccount_FullMethodName      = "/pixel_plaza.Payment/ActivatePaymentAccount"
	Payment_PayForOrder_FullMethodName                 = "/pixel_plaza.Payment/PayForOrder"
	Payment_PayForBranchRent_FullMethodName            = "/pixel_plaza.Payment/PayForBranchRent"
)

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	AddOrderPayment(ctx context.Context, in *AddOrderPaymentRequest, opts ...grpc.CallOption) (*AddOrderPaymentResponse, error)
	GetOrderPayments(ctx context.Context, in *GetOrderPaymentsRequest, opts ...grpc.CallOption) (*GetOrderPaymentsResponse, error)
	AddBranchRentPayment(ctx context.Context, in *AddBranchRentPaymentRequest, opts ...grpc.CallOption) (*AddBranchRentPaymentResponse, error)
	GetBranchRentPayments(ctx context.Context, in *GetBranchRentPaymentsRequest, opts ...grpc.CallOption) (*GetBranchRentPaymentsResponse, error)
	GetBranchRentsPayments(ctx context.Context, in *GetBranchRentPaymentsRequest, opts ...grpc.CallOption) (*GetBranchRentPaymentsResponse, error)
	VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error)
	AddPaymentAccount(ctx context.Context, in *AddPaymentAccountRequest, opts ...grpc.CallOption) (*AddPaymentAccountResponse, error)
	GetPaymentAccounts(ctx context.Context, in *GetPaymentAccountsRequest, opts ...grpc.CallOption) (*GetPaymentAccountsResponse, error)
	GetActivePaymentAccounts(ctx context.Context, in *GetActivePaymentAccountsRequest, opts ...grpc.CallOption) (*GetActivePaymentAccountsResponse, error)
	GetSuspendedPaymentAccounts(ctx context.Context, in *GetSuspendedPaymentAccountsRequest, opts ...grpc.CallOption) (*GetSuspendedPaymentAccountsResponse, error)
	SuspendPaymentAccount(ctx context.Context, in *SuspendPaymentAccountRequest, opts ...grpc.CallOption) (*SuspendPaymentAccountResponse, error)
	ActivatePaymentAccount(ctx context.Context, in *ActivatePaymentAccountRequest, opts ...grpc.CallOption) (*ActivatePaymentAccountResponse, error)
	PayForOrder(ctx context.Context, in *PayForOrderRequest, opts ...grpc.CallOption) (*PayForOrderResponse, error)
	PayForBranchRent(ctx context.Context, in *PayForBranchRentRequest, opts ...grpc.CallOption) (*PayForBranchRentResponse, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) AddOrderPayment(ctx context.Context, in *AddOrderPaymentRequest, opts ...grpc.CallOption) (*AddOrderPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddOrderPaymentResponse)
	err := c.cc.Invoke(ctx, Payment_AddOrderPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetOrderPayments(ctx context.Context, in *GetOrderPaymentsRequest, opts ...grpc.CallOption) (*GetOrderPaymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderPaymentsResponse)
	err := c.cc.Invoke(ctx, Payment_GetOrderPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) AddBranchRentPayment(ctx context.Context, in *AddBranchRentPaymentRequest, opts ...grpc.CallOption) (*AddBranchRentPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBranchRentPaymentResponse)
	err := c.cc.Invoke(ctx, Payment_AddBranchRentPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetBranchRentPayments(ctx context.Context, in *GetBranchRentPaymentsRequest, opts ...grpc.CallOption) (*GetBranchRentPaymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBranchRentPaymentsResponse)
	err := c.cc.Invoke(ctx, Payment_GetBranchRentPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetBranchRentsPayments(ctx context.Context, in *GetBranchRentPaymentsRequest, opts ...grpc.CallOption) (*GetBranchRentPaymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBranchRentPaymentsResponse)
	err := c.cc.Invoke(ctx, Payment_GetBranchRentsPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) VerifyPayment(ctx context.Context, in *VerifyPaymentRequest, opts ...grpc.CallOption) (*VerifyPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPaymentResponse)
	err := c.cc.Invoke(ctx, Payment_VerifyPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) AddPaymentAccount(ctx context.Context, in *AddPaymentAccountRequest, opts ...grpc.CallOption) (*AddPaymentAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPaymentAccountResponse)
	err := c.cc.Invoke(ctx, Payment_AddPaymentAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetPaymentAccounts(ctx context.Context, in *GetPaymentAccountsRequest, opts ...grpc.CallOption) (*GetPaymentAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentAccountsResponse)
	err := c.cc.Invoke(ctx, Payment_GetPaymentAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetActivePaymentAccounts(ctx context.Context, in *GetActivePaymentAccountsRequest, opts ...grpc.CallOption) (*GetActivePaymentAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActivePaymentAccountsResponse)
	err := c.cc.Invoke(ctx, Payment_GetActivePaymentAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetSuspendedPaymentAccounts(ctx context.Context, in *GetSuspendedPaymentAccountsRequest, opts ...grpc.CallOption) (*GetSuspendedPaymentAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSuspendedPaymentAccountsResponse)
	err := c.cc.Invoke(ctx, Payment_GetSuspendedPaymentAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) SuspendPaymentAccount(ctx context.Context, in *SuspendPaymentAccountRequest, opts ...grpc.CallOption) (*SuspendPaymentAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuspendPaymentAccountResponse)
	err := c.cc.Invoke(ctx, Payment_SuspendPaymentAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) ActivatePaymentAccount(ctx context.Context, in *ActivatePaymentAccountRequest, opts ...grpc.CallOption) (*ActivatePaymentAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivatePaymentAccountResponse)
	err := c.cc.Invoke(ctx, Payment_ActivatePaymentAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) PayForOrder(ctx context.Context, in *PayForOrderRequest, opts ...grpc.CallOption) (*PayForOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayForOrderResponse)
	err := c.cc.Invoke(ctx, Payment_PayForOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) PayForBranchRent(ctx context.Context, in *PayForBranchRentRequest, opts ...grpc.CallOption) (*PayForBranchRentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayForBranchRentResponse)
	err := c.cc.Invoke(ctx, Payment_PayForBranchRent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility.
type PaymentServer interface {
	AddOrderPayment(context.Context, *AddOrderPaymentRequest) (*AddOrderPaymentResponse, error)
	GetOrderPayments(context.Context, *GetOrderPaymentsRequest) (*GetOrderPaymentsResponse, error)
	AddBranchRentPayment(context.Context, *AddBranchRentPaymentRequest) (*AddBranchRentPaymentResponse, error)
	GetBranchRentPayments(context.Context, *GetBranchRentPaymentsRequest) (*GetBranchRentPaymentsResponse, error)
	GetBranchRentsPayments(context.Context, *GetBranchRentPaymentsRequest) (*GetBranchRentPaymentsResponse, error)
	VerifyPayment(context.Context, *VerifyPaymentRequest) (*VerifyPaymentResponse, error)
	AddPaymentAccount(context.Context, *AddPaymentAccountRequest) (*AddPaymentAccountResponse, error)
	GetPaymentAccounts(context.Context, *GetPaymentAccountsRequest) (*GetPaymentAccountsResponse, error)
	GetActivePaymentAccounts(context.Context, *GetActivePaymentAccountsRequest) (*GetActivePaymentAccountsResponse, error)
	GetSuspendedPaymentAccounts(context.Context, *GetSuspendedPaymentAccountsRequest) (*GetSuspendedPaymentAccountsResponse, error)
	SuspendPaymentAccount(context.Context, *SuspendPaymentAccountRequest) (*SuspendPaymentAccountResponse, error)
	ActivatePaymentAccount(context.Context, *ActivatePaymentAccountRequest) (*ActivatePaymentAccountResponse, error)
	PayForOrder(context.Context, *PayForOrderRequest) (*PayForOrderResponse, error)
	PayForBranchRent(context.Context, *PayForBranchRentRequest) (*PayForBranchRentResponse, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServer struct{}

func (UnimplementedPaymentServer) AddOrderPayment(context.Context, *AddOrderPaymentRequest) (*AddOrderPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrderPayment not implemented")
}
func (UnimplementedPaymentServer) GetOrderPayments(context.Context, *GetOrderPaymentsRequest) (*GetOrderPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderPayments not implemented")
}
func (UnimplementedPaymentServer) AddBranchRentPayment(context.Context, *AddBranchRentPaymentRequest) (*AddBranchRentPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBranchRentPayment not implemented")
}
func (UnimplementedPaymentServer) GetBranchRentPayments(context.Context, *GetBranchRentPaymentsRequest) (*GetBranchRentPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranchRentPayments not implemented")
}
func (UnimplementedPaymentServer) GetBranchRentsPayments(context.Context, *GetBranchRentPaymentsRequest) (*GetBranchRentPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranchRentsPayments not implemented")
}
func (UnimplementedPaymentServer) VerifyPayment(context.Context, *VerifyPaymentRequest) (*VerifyPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPayment not implemented")
}
func (UnimplementedPaymentServer) AddPaymentAccount(context.Context, *AddPaymentAccountRequest) (*AddPaymentAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPaymentAccount not implemented")
}
func (UnimplementedPaymentServer) GetPaymentAccounts(context.Context, *GetPaymentAccountsRequest) (*GetPaymentAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentAccounts not implemented")
}
func (UnimplementedPaymentServer) GetActivePaymentAccounts(context.Context, *GetActivePaymentAccountsRequest) (*GetActivePaymentAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivePaymentAccounts not implemented")
}
func (UnimplementedPaymentServer) GetSuspendedPaymentAccounts(context.Context, *GetSuspendedPaymentAccountsRequest) (*GetSuspendedPaymentAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuspendedPaymentAccounts not implemented")
}
func (UnimplementedPaymentServer) SuspendPaymentAccount(context.Context, *SuspendPaymentAccountRequest) (*SuspendPaymentAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuspendPaymentAccount not implemented")
}
func (UnimplementedPaymentServer) ActivatePaymentAccount(context.Context, *ActivatePaymentAccountRequest) (*ActivatePaymentAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivatePaymentAccount not implemented")
}
func (UnimplementedPaymentServer) PayForOrder(context.Context, *PayForOrderRequest) (*PayForOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayForOrder not implemented")
}
func (UnimplementedPaymentServer) PayForBranchRent(context.Context, *PayForBranchRentRequest) (*PayForBranchRentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayForBranchRent not implemented")
}
func (UnimplementedPaymentServer) mustEmbedUnimplementedPaymentServer() {}
func (UnimplementedPaymentServer) testEmbeddedByValue()                 {}

// UnsafePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServer will
// result in compilation errors.
type UnsafePaymentServer interface {
	mustEmbedUnimplementedPaymentServer()
}

func RegisterPaymentServer(s grpc.ServiceRegistrar, srv PaymentServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Payment_ServiceDesc, srv)
}

func _Payment_AddOrderPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).AddOrderPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_AddOrderPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).AddOrderPayment(ctx, req.(*AddOrderPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetOrderPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetOrderPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetOrderPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetOrderPayments(ctx, req.(*GetOrderPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_AddBranchRentPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBranchRentPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).AddBranchRentPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_AddBranchRentPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).AddBranchRentPayment(ctx, req.(*AddBranchRentPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetBranchRentPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchRentPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetBranchRentPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetBranchRentPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetBranchRentPayments(ctx, req.(*GetBranchRentPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetBranchRentsPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchRentPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetBranchRentsPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetBranchRentsPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetBranchRentsPayments(ctx, req.(*GetBranchRentPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_VerifyPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).VerifyPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_VerifyPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).VerifyPayment(ctx, req.(*VerifyPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_AddPaymentAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPaymentAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).AddPaymentAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_AddPaymentAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).AddPaymentAccount(ctx, req.(*AddPaymentAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetPaymentAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetPaymentAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetPaymentAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetPaymentAccounts(ctx, req.(*GetPaymentAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetActivePaymentAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivePaymentAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetActivePaymentAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetActivePaymentAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetActivePaymentAccounts(ctx, req.(*GetActivePaymentAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetSuspendedPaymentAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSuspendedPaymentAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetSuspendedPaymentAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetSuspendedPaymentAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetSuspendedPaymentAccounts(ctx, req.(*GetSuspendedPaymentAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_SuspendPaymentAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendPaymentAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).SuspendPaymentAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_SuspendPaymentAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).SuspendPaymentAccount(ctx, req.(*SuspendPaymentAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_ActivatePaymentAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivatePaymentAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).ActivatePaymentAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_ActivatePaymentAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).ActivatePaymentAccount(ctx, req.(*ActivatePaymentAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_PayForOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayForOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).PayForOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_PayForOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).PayForOrder(ctx, req.(*PayForOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_PayForBranchRent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayForBranchRentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).PayForBranchRent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_PayForBranchRent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).PayForBranchRent(ctx, req.(*PayForBranchRentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_ServiceDesc is the grpc.ServiceDesc for Payment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pixel_plaza.Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrderPayment",
			Handler:    _Payment_AddOrderPayment_Handler,
		},
		{
			MethodName: "GetOrderPayments",
			Handler:    _Payment_GetOrderPayments_Handler,
		},
		{
			MethodName: "AddBranchRentPayment",
			Handler:    _Payment_AddBranchRentPayment_Handler,
		},
		{
			MethodName: "GetBranchRentPayments",
			Handler:    _Payment_GetBranchRentPayments_Handler,
		},
		{
			MethodName: "GetBranchRentsPayments",
			Handler:    _Payment_GetBranchRentsPayments_Handler,
		},
		{
			MethodName: "VerifyPayment",
			Handler:    _Payment_VerifyPayment_Handler,
		},
		{
			MethodName: "AddPaymentAccount",
			Handler:    _Payment_AddPaymentAccount_Handler,
		},
		{
			MethodName: "GetPaymentAccounts",
			Handler:    _Payment_GetPaymentAccounts_Handler,
		},
		{
			MethodName: "GetActivePaymentAccounts",
			Handler:    _Payment_GetActivePaymentAccounts_Handler,
		},
		{
			MethodName: "GetSuspendedPaymentAccounts",
			Handler:    _Payment_GetSuspendedPaymentAccounts_Handler,
		},
		{
			MethodName: "SuspendPaymentAccount",
			Handler:    _Payment_SuspendPaymentAccount_Handler,
		},
		{
			MethodName: "ActivatePaymentAccount",
			Handler:    _Payment_ActivatePaymentAccount_Handler,
		},
		{
			MethodName: "PayForOrder",
			Handler:    _Payment_PayForOrder_Handler,
		},
		{
			MethodName: "PayForBranchRent",
			Handler:    _Payment_PayForBranchRent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pixel_plaza/payment.proto",
}