package carts

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the orders service carts REST endpoints
var Base = typesrest.NewBaseEndpoint("carts")

// Orders service carts REST endpoints
var (
	Relative      = typesrest.NewRelativeEndpoint()
	ByCartId      = typesrest.NewRelativeEndpoint(typesrest.CartId)
	TotalByCartId = typesrest.NewEndpoint("total", typesrest.CartId)
)
