package permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest"
)

// Map is the map of the REST API endpoints of the auth service permissions
var Map = rest.NewMap(
	&Interceptions,
	nil,
)
