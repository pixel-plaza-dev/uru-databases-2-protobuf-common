package roles

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the auth service roles REST endpoints
var Base = typesrest.NewEndpoint("roles")

// Auth service roles REST endpoints
var (
	Relative = typesrest.NewEndpoint("")
	ByRoleId = typesrest.NewEndpoint("", typesrest.RoleId)
)
