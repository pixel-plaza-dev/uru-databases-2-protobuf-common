package branches

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service branches REST endpoints
var Base = typesrest.NewBaseEndpoint("branches")

// Shop service branches REST endpoints
var (
	Relative                   = typesrest.NewRelativeEndpoint()
	ByBranchId                 = typesrest.NewRelativeEndpoint(typesrest.BranchId)
	ByBusinessId               = typesrest.NewEndpoint("business-id", typesrest.BusinessId)
	CloseTemporarilyByBranchId = typesrest.NewEndpoint("close-temporarily", typesrest.BranchId)
	OpenByBranchId             = typesrest.NewEndpoint("open", typesrest.BranchId)
)
