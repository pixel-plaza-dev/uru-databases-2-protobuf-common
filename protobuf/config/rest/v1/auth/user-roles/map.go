package user_roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest"
)

// Map is the map of the REST API endpoints of the auth service user roles
var Map = rest.NewMap(
	&Interceptions,
	nil,
)
