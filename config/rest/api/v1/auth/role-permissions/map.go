package role_permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Map is the map of the REST API endpoints of the auth service role permissions
var Map = rest.NewMap(
	&Interceptions,
	nil,
)
