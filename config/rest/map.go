package rest

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// ChildrenMaps is the map of the REST API endpoints of the API Gateway router
var ChildrenMaps = map[string]*rest.Map{
	Api.String(): api.Map,
}

// Map is the map of the REST API endpoints of the API Gateway router
var Map = rest.NewMap(
	nil,
	&ChildrenMaps,
)
