package refresh_tokens

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service refresh tokens REST endpoints
var (
	Relative = typesrest.NewEndpoint("")
	ByJwtId  = typesrest.NewEndpoint(
		"",
		typesrest.JwtId,
	)
	ValidByJwtId = typesrest.NewEndpoint(
		"valid",
		typesrest.JwtId,
	)
)
