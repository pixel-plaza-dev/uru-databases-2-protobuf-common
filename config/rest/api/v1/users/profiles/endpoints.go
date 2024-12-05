package profiles

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the user service profiles REST endpoints
var Base = typesrest.NewEndpoint("profiles")

// User service profiles REST endpoints
var (
	Relative   = typesrest.NewEndpoint("")
	ByUsername = typesrest.NewEndpoint("", typesrest.Username)
)
