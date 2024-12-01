package user_roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the auth service refresh tokens gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.POST:   auth.AddUserRole,
		rest.DELETE: auth.RevokeUserRole,
	},
	ByUserId.String(): {
		rest.GET: auth.GetUserRoles,
	},
}
