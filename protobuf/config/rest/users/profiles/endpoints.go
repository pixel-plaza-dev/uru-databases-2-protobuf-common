package profiles

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the user service profiles REST endpoints
const BaseURI = "/profiles"

// User service profiles REST endpoints
var (
	Relative   = rest2.NewEndpoint("/")
	ByUsername = rest2.NewEndpoint("/", rest2.Username)
)
