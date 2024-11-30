package access_tokens

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the auth service access tokens REST endpoints
const BaseURI = "/access-tokens"

// Auth service access tokens REST endpoints
var (
	Valid = rest2.NewEndpoint(
		"/valid",
		rest2.JwtId,
	)
)
