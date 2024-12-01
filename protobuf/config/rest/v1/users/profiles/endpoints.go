package profiles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// User service profiles REST endpoints
var (
	Relative   = rest.NewEndpoint("")
	ByUsername = rest.NewEndpoint("", rest.Username)
)
