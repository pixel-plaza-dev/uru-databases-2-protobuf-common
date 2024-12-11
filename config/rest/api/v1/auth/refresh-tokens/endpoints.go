package refresh_tokens

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the auth service refresh tokens REST endpoints
var Base = typesrest.NewBaseEndpoint("refresh-tokens")

// Auth service refresh tokens REST endpoints
var (
	Relative = typesrest.NewRelativeEndpoint()
	ByJwtId  = typesrest.NewRelativeEndpoint(

		typesrest.JwtId,
	)
	ValidByJwtId = typesrest.NewEndpoint(
		"valid",
		typesrest.JwtId,
	)
)
