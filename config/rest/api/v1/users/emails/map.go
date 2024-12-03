package emails

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Map is the map of the REST API endpoints of the user service emails
var Map = rest.NewMap(
	&Interceptions,
	nil,
)
