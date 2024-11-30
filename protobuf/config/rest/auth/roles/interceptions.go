package roles

import (
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Map is the map of the REST API endpoints to the auth service roles gRPC methods
var Map = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.POST: detailsauth.AddRole,
		rest.GET:  detailsauth.GetRoles,
	},
	ById.String(): {
		rest.POST:   detailsauth.AddRolePermission,
		rest.GET:    detailsauth.GetRolePermissions,
		rest.DELETE: detailsauth.RevokeRole,
	},
}
