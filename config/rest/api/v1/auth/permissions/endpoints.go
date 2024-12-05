package permissions

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the auth service permissions REST endpoints
var Base = typesrest.NewEndpoint("permissions")

// Auth service permissions REST endpoints
var (
	Relative       = typesrest.NewEndpoint("")
	ByPermissionId = typesrest.NewEndpoint("", typesrest.PermissionId)
)
