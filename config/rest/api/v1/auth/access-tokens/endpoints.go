package access_tokens

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service access tokens REST endpoints
var (
	ValidByJwtId = typesrest.NewEndpoint(
		"valid",
		typesrest.JwtId,
	)
)
