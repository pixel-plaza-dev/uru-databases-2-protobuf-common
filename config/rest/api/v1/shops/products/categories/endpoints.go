package categories

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service product categories REST endpoints
var Base = typesrest.NewBaseEndpoint("categories")

// Shop service product categories REST endpoints
var (
	Relative     = typesrest.NewRelativeEndpoint()
	ByCategoryId = typesrest.NewRelativeEndpoint(typesrest.CategoryId)
)
