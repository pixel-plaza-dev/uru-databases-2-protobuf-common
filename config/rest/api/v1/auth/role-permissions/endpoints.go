package role_permissions

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the auth service role permissions REST endpoints
var Base = typesrest.NewBaseEndpoint("role-permissions")

// Auth service role permissions REST endpoints
var (
	ByRoleId = typesrest.NewRelativeEndpoint(
		typesrest.RoleId,
	)
)
