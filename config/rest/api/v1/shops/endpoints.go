package shops

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
	SuspendByBusinessId        = typesrest.NewEndpoint("suspend", typesrest.BusinessId)
	ActivateByBusinessId       = typesrest.NewEndpoint("activate", typesrest.BusinessId)
)
