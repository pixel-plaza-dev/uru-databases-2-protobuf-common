package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the auth service REST endpoints
const BaseURI = "/auth"

// Auth service REST endpoints
var (
	LogIn  = rest.NewEndpoint("/log-in")
	LogOut = rest.NewEndpoint("/log-out")
)
