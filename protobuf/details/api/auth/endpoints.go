package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// Auth service REST endpoints
var (
	LogIn  = api.NewRESTEndpoint("/log-in")
	LogOut = api.NewRESTEndpoint("/log-out")
)
