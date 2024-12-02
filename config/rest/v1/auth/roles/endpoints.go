package roles

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service roles REST endpoints
var (
	Relative = typesrest.NewEndpoint("")
	ByRoleId = typesrest.NewEndpoint("", typesrest.RoleId)
)
