package roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the auth service roles REST endpoints
const BaseURI = "/roles"

// Auth service roles REST endpoints
var (
	Relative = api.NewRESTEndpoint("/")
	ById     = api.NewRESTEndpoint("/" + api.Id.String())
)
