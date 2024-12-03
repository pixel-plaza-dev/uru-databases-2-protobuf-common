package role_permissions

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service role permissions REST endpoints
var (
	ByRoleId = typesrest.NewEndpoint(
		"", typesrest.RoleId,
	)
)
