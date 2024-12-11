package products

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service products REST endpoints
var Base = typesrest.NewBaseEndpoint("products")

// Shop service products REST endpoints
var (
	Relative            = typesrest.NewRelativeEndpoint()
	ByProductId         = typesrest.NewRelativeEndpoint(typesrest.ProductId)
	Search              = typesrest.NewEndpoint("search")
	SuspendByProductId  = typesrest.NewEndpoint("suspend", typesrest.ProductId)
	ActivateByProductId = typesrest.NewEndpoint("activate", typesrest.ProductId)
)
