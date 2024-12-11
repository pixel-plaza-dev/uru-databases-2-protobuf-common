package accounts

import (
	grpcpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/payment"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Payment service accounts endpoints mapping
var (
	AddPaymentAccountMapper           = typesrest.NewMapper(Relative, grpcpayment.AddPaymentAccount)
	GetPaymentAccountsMapper          = typesrest.NewMapper(Relative, grpcpayment.GetPaymentAccounts)
	GetActivePaymentAccountsMapper    = typesrest.NewMapper(Active, grpcpayment.GetActivePaymentAccounts)
	ActivatePaymentAccountMapper      = typesrest.NewMapper(Activate, grpcpayment.ActivatePaymentAccount)
	GetSuspendedPaymentAccountsMapper = typesrest.NewMapper(Suspended, grpcpayment.GetSuspendedPaymentAccounts)
	SuspendPaymentAccountMapper       = typesrest.NewMapper(Suspend, grpcpayment.SuspendPaymentAccount)
	VerifyPaymentMapper               = typesrest.NewMapper(VerifyPayment, grpcpayment.VerifyPayment)
)
