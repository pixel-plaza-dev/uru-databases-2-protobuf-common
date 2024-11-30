package permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc"
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// Map is the map of the REST API endpoints to the auth service permissions gRPC methods
var Map = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.GET:  detailsauth.GetPermissions,
		rest.POST: detailsauth.AddPermission,
	},
	ById.String(): {
		rest.GET:    detailsauth.GetPermission,
		rest.DELETE: detailsauth.RevokePermission,
	},
}
