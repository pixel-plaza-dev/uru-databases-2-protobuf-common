package businesses

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service business products REST endpoints
var Base = typesrest.NewBaseEndpoint("business")

// Shop service business products REST endpoints
var (
	Relative            = typesrest.NewRelativeEndpoint()
	ByProductId         = typesrest.NewRelativeEndpoint(typesrest.ProductId)
	Search              = typesrest.NewEndpoint("search")
	ActivateByProductId = typesrest.NewEndpoint("activate", typesrest.ProductId)
	SuspendByProductId  = typesrest.NewEndpoint("suspend", typesrest.ProductId)
)
