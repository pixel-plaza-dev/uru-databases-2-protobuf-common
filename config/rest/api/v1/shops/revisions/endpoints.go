package revisions

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service revisions REST endpoints
var Base = typesrest.NewBaseEndpoint("revisions")

// Shop service REST revisions endpoints
var (
	Relative            = typesrest.NewRelativeEndpoint(typesrest.RevisionId)
	BranchesByBranchId  = typesrest.NewEndpoint("branches", typesrest.BranchId)
	ProductsByProductId = typesrest.NewEndpoint("products", typesrest.ProductId)
)
