package permissions

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the auth service permissions REST endpoints
var Base = typesrest.NewBaseEndpoint("permissions")

// Auth service permissions REST endpoints
var (
	Relative       = typesrest.NewRelativeEndpoint()
	ByPermissionId = typesrest.NewRelativeEndpoint(typesrest.PermissionId)
)
