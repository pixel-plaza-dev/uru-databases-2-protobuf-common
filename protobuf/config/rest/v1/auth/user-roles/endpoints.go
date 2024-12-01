package user_roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Auth service user roles REST endpoints
var (
	Relative = rest.NewEndpoint("")
	ByUserId = rest.NewEndpoint("", rest.UserId)
)
