package user_roles

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the auth service user roles REST endpoints
var Base = typesrest.NewBaseEndpoint("user-roles")

// Auth service user roles REST endpoints
var (
	Relative = typesrest.NewRelativeEndpoint()
	ByUserId = typesrest.NewRelativeEndpoint(typesrest.UserId)
)
