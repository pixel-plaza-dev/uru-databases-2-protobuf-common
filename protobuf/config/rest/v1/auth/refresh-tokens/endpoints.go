package refresh_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Auth service refresh tokens REST endpoints
var (
	Relative = rest.NewEndpoint("")
	ByJwtId  = rest.NewEndpoint(
		"",
		rest.JwtId,
	)
	Valid = rest.NewEndpoint(
		"valid",
		rest.JwtId,
	)
)
