package user_roles

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the auth service user roles REST endpoints
const BaseURI = "/user-roles"

// Auth service user roles REST endpoints
var (
	Relative = rest2.NewEndpoint("/")
	ByUserId = rest2.NewEndpoint("/", rest2.UserId)
)
