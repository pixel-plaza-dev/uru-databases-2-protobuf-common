package role_permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the auth service role permissions REST endpoints
const BaseURI = "/refresh-tokens"

// Auth service role permissions REST endpoints
var (
	ById = rest.NewEndpoint(
		"/", rest.Id,
	)
)
