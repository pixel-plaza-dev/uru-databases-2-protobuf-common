package user_roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the auth service user roles REST endpoints
const BaseURI = "/user-roles"

// Auth service user roles REST endpoints
var (
	Relative = api.NewRESTEndpoint("/")
	ByUserId = api.NewRESTEndpoint("/", api.UserId)
)
