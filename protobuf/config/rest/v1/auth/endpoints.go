package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Auth service REST endpoints
var (
	LogIn  = rest.NewEndpoint("log-in")
	LogOut = rest.NewEndpoint("log-out")
)
