package orders

import (
	grpcorder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/order"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Order service endpoints mapping
var (
	GetOrdersMapper = typesrest.NewMapper(Relative, grpcorder.GetOrders)
	GetOrderMapper  = typesrest.NewMapper(ByOrderId, grpcorder.GetOrder)
)
