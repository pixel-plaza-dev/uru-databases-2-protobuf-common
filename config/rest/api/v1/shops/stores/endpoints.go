package stores

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service stores REST endpoints
var Base = typesrest.NewBaseEndpoint("stores")

// Shop service stores REST endpoints
var (
	Relative   = typesrest.NewRelativeEndpoint()
	ByStoreId  = typesrest.NewRelativeEndpoint(typesrest.StoreId)
	Unoccupied = typesrest.NewEndpoint("unoccupied")
)
