package user_roles

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service user roles REST endpoints
var (
	Relative = typesrest.NewEndpoint("")
	ByUserId = typesrest.NewEndpoint("", typesrest.UserId)
)
