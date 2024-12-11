package owners

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service owners REST endpoints
var Base = typesrest.NewBaseEndpoint("owners")

// Shop service owners REST endpoints
var (
	ByBusinessId = typesrest.NewRelativeEndpoint(typesrest.BusinessId)
)
