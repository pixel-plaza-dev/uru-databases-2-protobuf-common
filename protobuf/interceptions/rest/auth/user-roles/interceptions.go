package user_roles

import (
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Map is the map of the REST API endpoints to the auth service refresh tokens gRPC methods
var Map = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.POST:   detailsauth.AddUserRole,
		rest.DELETE: detailsauth.RevokeUserRole,
	},
	ByUserId.String(): {
		rest.GET: detailsauth.GetUserRoles,
	},
}
