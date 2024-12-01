package permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Auth service permissions REST endpoints
var (
	Relative = rest.NewEndpoint("")
	ById     = rest.NewEndpoint("", rest.Id)
)
