package access_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the auth service access tokens REST endpoints
const BaseURI = "/access-tokens"

// Auth service access tokens REST endpoints
var (
	Valid = rest.NewEndpoint(
		"/valid",
		rest.JwtId,
	)
)
