package payment

import (
	compiled "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/payment"
	"github.com/ralvarezdev/go-jwt/token/interception"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[string]interception.Interception{
	compiled.Payment_AddOrderPayment_FullMethodName:             interception.AccessToken,
	compiled.Payment_GetOrderPayments_FullMethodName:            interception.AccessToken,
	compiled.Payment_AddBranchRentPayment_FullMethodName:        interception.AccessToken,
	compiled.Payment_GetBranchRentPayments_FullMethodName:       interception.AccessToken,
	compiled.Payment_GetBranchRentsPayments_FullMethodName:      interception.AccessToken,
	compiled.Payment_VerifyPayment_FullMethodName:               interception.AccessToken,
	compiled.Payment_AddPaymentAccount_FullMethodName:           interception.AccessToken,
	compiled.Payment_GetPaymentAccounts_FullMethodName:          interception.AccessToken,
	compiled.Payment_GetActivePaymentAccounts_FullMethodName:    interception.AccessToken,
	compiled.Payment_GetSuspendedPaymentAccounts_FullMethodName: interception.AccessToken,
	compiled.Payment_SuspendPaymentAccount_FullMethodName:       interception.AccessToken,
	compiled.Payment_ActivatePaymentAccount_FullMethodName:      interception.AccessToken,
	compiled.Payment_PayForOrder_FullMethodName:                 interception.AccessToken,
	compiled.Payment_PayForBranchRent_FullMethodName:            interception.AccessToken,
}
