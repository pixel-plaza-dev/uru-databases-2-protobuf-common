package refresh_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the auth service refresh tokens REST endpoints
const BaseURI = "/refresh-tokens"

// Auth service refresh tokens REST endpoints
var (
	Relative = rest.NewEndpoint("/")
	ByJwtId  = rest.NewEndpoint(
		"/",
		rest.JwtId,
	)
	Valid = rest.NewEndpoint(
		"/valid",
		rest.JwtId,
	)
)
