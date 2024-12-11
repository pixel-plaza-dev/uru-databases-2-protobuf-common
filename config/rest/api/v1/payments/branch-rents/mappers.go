package branch_rents

import (
	grpcpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/payment"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Payment service branch rents endpoints mapping
var (
	GetBranchRentsPaymentsMapper = typesrest.NewMapper(Relative, grpcpayment.GetBranchRentsPayments)
	AddBranchRentPaymentMapper   = typesrest.NewMapper(ByBranchRentId, grpcpayment.AddBranchRentPayment)
	GetBranchRentPaymentsMapper  = typesrest.NewMapper(ByBranchRentId, grpcpayment.GetBranchRentPayments)
	PayForBranchRentMapper       = typesrest.NewMapper(ByBranchRentId, grpcpayment.PayForBranchRent)
)
