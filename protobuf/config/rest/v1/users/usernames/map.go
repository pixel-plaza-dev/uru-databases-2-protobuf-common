package usernames

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest"
)

// Map is the map of the REST API endpoints of the user service usernames
var Map = rest.NewMap(
	&Interceptions,
	nil,
)
