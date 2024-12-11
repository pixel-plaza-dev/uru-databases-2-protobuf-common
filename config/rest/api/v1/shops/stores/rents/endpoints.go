package rents

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service stores rents REST endpoints
var Base = typesrest.NewBaseEndpoint("rents")

// Shop service stores rents REST endpoints
var (
	Relative       = typesrest.NewRelativeEndpoint()
	ByBranchId     = typesrest.NewRelativeEndpoint(typesrest.BranchId)
	ByBranchRentId = typesrest.NewRelativeEndpoint(typesrest.BranchRentId)
	ByBusinessId   = typesrest.NewRelativeEndpoint(typesrest.BusinessId)
)
