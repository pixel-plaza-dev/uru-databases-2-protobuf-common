package profiles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the user service profiles REST endpoints
const BaseURI = "/profiles"

// User service profiles REST endpoints
var (
	Relative   = rest.NewEndpoint("/")
	ByUsername = rest.NewEndpoint("/", rest.Username)
)
