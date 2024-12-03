package permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Interceptions is the map of the REST API endpoints to the auth service permissions gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.GET:  auth.GetPermissions,
		rest.POST: auth.AddPermission,
	},
	ByPermissionId.String(): {
		rest.GET:    auth.GetPermission,
		rest.DELETE: auth.RevokePermission,
	},
}