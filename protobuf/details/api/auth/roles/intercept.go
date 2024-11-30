package roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// RESTMap is the map of the REST API endpoints to the auth service roles gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Relative.String(): {
		api.POST: detailsauth.AddRole,
		api.GET:  detailsauth.GetRoles,
	},
	ById.String(): {
		api.POST:   detailsauth.AddRolePermission,
		api.GET:    detailsauth.GetRolePermissions,
		api.DELETE: detailsauth.RevokeRole,
	},
}
