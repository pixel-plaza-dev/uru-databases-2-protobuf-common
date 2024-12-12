package orders

import (
	grpcpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/payment"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Payment service orders endpoints mapping
var (
	AddOrderPaymentMapper  = typesrest.NewMapper(ByOrderId, grpcpayment.AddOrderPayment)
	GetOrderPaymentsMapper = typesrest.NewMapper(ByOrderId, grpcpayment.GetOrderPayments)
	PayForOrderMapper      = typesrest.NewMapper(PayByOrderId, grpcpayment.PayForOrder)
)
