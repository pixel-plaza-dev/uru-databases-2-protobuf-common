package role_permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Interceptions is the map of the REST API endpoints to the auth service refresh tokens gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	ByRoleId.String(): {
		rest.DELETE: auth.RevokeRolePermission,
	},
}
