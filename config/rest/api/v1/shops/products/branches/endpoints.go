package branches

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service branch products REST endpoints
var Base = typesrest.NewBaseEndpoint("branches")

// Shop service branch products REST endpoints
var (
	ByBranchId = typesrest.NewRelativeEndpoint(typesrest.BranchId)
	Search     = typesrest.NewEndpoint("search")
)
