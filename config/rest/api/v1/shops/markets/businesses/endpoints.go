package businesses

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service business markets REST endpoints
var Base = typesrest.NewBaseEndpoint("businesses")

// Shop service business markets REST endpoints
var (
	Relative     = typesrest.NewRelativeEndpoint()
	ByBusinessId = typesrest.NewRelativeEndpoint(typesrest.BusinessId)
)
