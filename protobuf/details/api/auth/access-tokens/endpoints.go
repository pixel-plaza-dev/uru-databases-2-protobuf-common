package access_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the auth service access tokens REST endpoints
const BaseURI = "/access-tokens"

// Auth service access tokens REST endpoints
var (
	Valid = api.NewRESTEndpoint(
		"/valid",
		api.JwtId,
	)
)
