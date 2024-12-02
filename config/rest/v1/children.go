package v1

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// API Gateway router version 1 children REST endpoints
var (
	Auth  = rest.NewEndpoint("auth")
	Users = rest.NewEndpoint("users")
)
