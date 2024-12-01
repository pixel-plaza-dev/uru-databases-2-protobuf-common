package profiles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest"
)

// Map is the map of the REST API endpoints of the user service profiles
var Map = rest.NewMap(
	&Interceptions,
	nil,
)
