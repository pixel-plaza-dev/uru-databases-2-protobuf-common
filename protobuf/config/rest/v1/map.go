package v1

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/users"
)

// ChildrenMaps is the map of the REST API endpoints of the API Gateway router version 1
var ChildrenMaps = map[string]*rest.Map{
	Auth.String():  auth.Map,
	Users.String(): users.Map,
}

// Map is the map of the REST API endpoints of the API Gateway router version 1
var Map = rest.NewMap(
	nil,
	&ChildrenMaps,
)
