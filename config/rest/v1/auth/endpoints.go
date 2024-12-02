package auth

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service REST endpoints
var (
	LogIn  = typesrest.NewEndpoint("log-in")
	LogOut = typesrest.NewEndpoint("log-out")
)
