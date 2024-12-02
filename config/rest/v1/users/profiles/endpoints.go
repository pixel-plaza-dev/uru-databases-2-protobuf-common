package profiles

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// User service profiles REST endpoints
var (
	Relative   = typesrest.NewEndpoint("")
	ByUsername = typesrest.NewEndpoint("", typesrest.Username)
)
