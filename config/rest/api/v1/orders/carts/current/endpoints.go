package carts

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the orders service current cart REST endpoints
var Base = typesrest.NewBaseEndpoint("current")

// Orders service current cart REST endpoints
var (
	Relative = typesrest.NewRelativeEndpoint()
	Checkout = typesrest.NewEndpoint("checkout")
)
