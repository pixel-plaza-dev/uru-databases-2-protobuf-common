package orders

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the order service REST endpoints
var Base = typesrest.NewBaseEndpoint("orders")

// Order service REST endpoints
var (
	Relative  = typesrest.NewRelativeEndpoint()
	ByOrderId = typesrest.NewRelativeEndpoint(typesrest.OrderId)
)
