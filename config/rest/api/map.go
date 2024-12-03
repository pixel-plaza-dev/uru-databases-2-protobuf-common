package v1

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// ChildrenMaps is the map of the REST API endpoints of the API Gateway router versions
var ChildrenMaps = map[string]*rest.Map{
	V1.String(): v1.Map,
}

// Map is the map of the REST API endpoints of the API Gateway router versions
var Map = rest.NewMap(
	nil,
	&ChildrenMaps,
)
