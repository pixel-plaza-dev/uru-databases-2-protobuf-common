package permissions

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the auth service permissions REST endpoints
const BaseURI = "/permissions"

// Auth service permissions REST endpoints
var (
	Relative = rest2.NewEndpoint("/")
	ById     = rest2.NewEndpoint("/" + rest2.Id.String())
)
