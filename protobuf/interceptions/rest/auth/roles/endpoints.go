package roles

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the auth service roles REST endpoints
const BaseURI = "/roles"

// Auth service roles REST endpoints
var (
	Relative = rest2.NewEndpoint("/")
	ById     = rest2.NewEndpoint("/" + rest2.Id.String())
)
