package roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the auth service roles REST endpoints
const BaseURI = "/roles"

// Auth service roles REST endpoints
var (
	Relative = rest.NewEndpoint("/")
	ById     = rest.NewEndpoint("/" + rest.Id.String())
)
