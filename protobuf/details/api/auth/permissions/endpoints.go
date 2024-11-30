package permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the auth service permissions REST endpoints
const BaseURI = "/permissions"

// Auth service permissions REST endpoints
var (
	Relative = api.NewRESTEndpoint("/")
	ById     = api.NewRESTEndpoint("/" + api.Id.String())
)
