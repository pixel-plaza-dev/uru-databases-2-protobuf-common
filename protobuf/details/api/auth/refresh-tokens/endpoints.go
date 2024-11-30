package refresh_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the auth service refresh tokens REST endpoints
const BaseURI = "/refresh-tokens"

// Auth service refresh tokens REST endpoints
var (
	Relative = api.NewRESTEndpoint("/")
	ByJwtId  = api.NewRESTEndpoint(
		"/",
		api.JwtId,
	)
	Valid = api.NewRESTEndpoint(
		"/valid",
		api.JwtId,
	)
)
