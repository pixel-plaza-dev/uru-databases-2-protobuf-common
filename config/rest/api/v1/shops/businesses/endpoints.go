package businesses

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the shop service REST endpoints
var Base = typesrest.NewBaseEndpoint("shops")

// Shop service REST endpoints
var (
	Relative                   = typesrest.NewRelativeEndpoint()
	ByBusinessId               = typesrest.NewRelativeEndpoint(typesrest.BusinessId)
	ProfilePictureByBusinessId = typesrest.NewEndpoint("profile-picture", typesrest.BusinessId)
)
