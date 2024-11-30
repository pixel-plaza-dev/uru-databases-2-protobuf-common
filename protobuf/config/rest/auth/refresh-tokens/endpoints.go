package refresh_tokens

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the auth service refresh tokens REST endpoints
const BaseURI = "/refresh-tokens"

// Auth service refresh tokens REST endpoints
var (
	Relative = rest2.NewEndpoint("/")
	ByJwtId  = rest2.NewEndpoint(
		"/",
		rest2.JwtId,
	)
	Valid = rest2.NewEndpoint(
		"/valid",
		rest2.JwtId,
	)
)
