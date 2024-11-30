package permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the auth service permissions REST endpoints
const BaseURI = "/permissions"

// Auth service permissions REST endpoints
var (
	Relative = rest.NewEndpoint("/")
	ById     = rest.NewEndpoint("/" + rest.Id.String())
)
