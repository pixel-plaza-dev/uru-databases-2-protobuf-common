package access_tokens

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the auth service access tokens REST endpoints
var Base = typesrest.NewBaseEndpoint("access-tokens")

// Auth service access tokens REST endpoints
var (
	ValidByJwtId = typesrest.NewEndpoint(
		"valid",
		typesrest.JwtId,
	)
)
