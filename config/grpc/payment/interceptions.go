package payment

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[grpc.Method]grpc.Interception{
	AddOrderPayment:             grpc.AccessToken,
	GetOrderPayments:            grpc.AccessToken,
	AddBranchRentPayment:        grpc.AccessToken,
	GetBranchRentPayments:       grpc.AccessToken,
	GetBranchRentsPayments:      grpc.AccessToken,
	VerifyPayment:               grpc.AccessToken,
	AddPaymentAccount:           grpc.AccessToken,
	GetPaymentAccounts:          grpc.AccessToken,
	GetActivePaymentAccounts:    grpc.AccessToken,
	GetSuspendedPaymentAccounts: grpc.AccessToken,
	SuspendPaymentAccount:       grpc.AccessToken,
	ActivatePaymentAccount:      grpc.AccessToken,
	PayForOrder:                 grpc.AccessToken,
	PayForBranchRent:            grpc.AccessToken,
}
