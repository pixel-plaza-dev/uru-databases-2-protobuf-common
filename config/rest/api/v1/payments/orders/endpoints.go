package orders

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the payment service orders REST endpoints
var Base = typesrest.NewBaseEndpoint("orders")

// Payment service orders REST endpoints
var (
	ByOrderId    = typesrest.NewRelativeEndpoint(typesrest.OrderId)
	PayByOrderId = typesrest.NewEndpoint("pay", typesrest.OrderId)
)
