package payment

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Payment service gRPC methods
var (
	AddOrderPayment             = grpc.NewMethod("AddOrderPayment")
	GetOrderPayments            = grpc.NewMethod("GetOrderPayments")
	AddBranchRentPayment        = grpc.NewMethod("AddBranchRentPayment")
	GetBranchRentPayments       = grpc.NewMethod("GetBranchRentPayments")
	GetBranchRentsPayments      = grpc.NewMethod("GetBranchRentsPayments")
	VerifyPayment               = grpc.NewMethod("VerifyPayment")
	AddPaymentAccount           = grpc.NewMethod("AddPaymentAccount")
	GetPaymentAccounts          = grpc.NewMethod("GetPaymentAccounts")
	GetActivePaymentAccounts    = grpc.NewMethod("GetActivePaymentAccounts")
	GetSuspendedPaymentAccounts = grpc.NewMethod("GetSuspendedPaymentAccounts")
	SuspendPaymentAccount       = grpc.NewMethod("SuspendPaymentAccount")
	ActivatePaymentAccount      = grpc.NewMethod("ActivatePaymentAccount")
	PayForOrder                 = grpc.NewMethod("PayForOrder")
	PayForBranchRent            = grpc.NewMethod("PayForBranchRent")
)
