package clients

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service clients REST endpoints
var Base = typesrest.NewBaseEndpoint("clients")

// Shop service clients REST endpoints
var (
	ByBusinessId = typesrest.NewRelativeEndpoint(typesrest.BusinessId)
)
