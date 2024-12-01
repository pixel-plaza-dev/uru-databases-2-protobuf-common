package roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the auth service roles gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.POST: auth.AddRole,
		rest.GET:  auth.GetRoles,
	},
	ById.String(): {
		rest.POST:   auth.AddRolePermission,
		rest.GET:    auth.GetRolePermissions,
		rest.DELETE: auth.RevokeRole,
	},
}
