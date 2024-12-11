package branch_rents

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the payment service branch rents REST endpoints
var Base = typesrest.NewEndpoint("branch-rents")

// Payment service branch rents REST endpoints
var (
	Relative       = typesrest.NewRelativeEndpoint()
	ByBranchRentId = typesrest.NewRelativeEndpoint(typesrest.BranchRentId)
)
