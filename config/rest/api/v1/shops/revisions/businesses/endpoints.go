package businesses

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service businesses revisions REST endpoints
var Base = typesrest.NewBaseEndpoint("businesses")

// Shop service REST businesses revisions endpoints
var (
	ByBusinessId        = typesrest.NewRelativeEndpoint(typesrest.BusinessId)
	ProductsByProductId = typesrest.NewEndpoint("products", typesrest.ProductId)
)
