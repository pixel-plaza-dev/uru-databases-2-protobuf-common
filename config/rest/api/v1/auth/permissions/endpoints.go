package permissions

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service permissions REST endpoints
var (
	Relative       = typesrest.NewEndpoint("")
	ByPermissionId = typesrest.NewEndpoint("", typesrest.PermissionId)
)
