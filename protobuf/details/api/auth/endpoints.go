package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the auth service REST endpoints
const BaseURI = "/auth"

// Auth service REST endpoints
var (
	LogIn  = api.NewRESTEndpoint("/log-in")
	LogOut = api.NewRESTEndpoint("/log-out")
)
