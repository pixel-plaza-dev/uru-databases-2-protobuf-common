package access_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Auth service access tokens REST endpoints
var (
	Valid = rest.NewEndpoint(
		"valid",
		rest.JwtId,
	)
)
