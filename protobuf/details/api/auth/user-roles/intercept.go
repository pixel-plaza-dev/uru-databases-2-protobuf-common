package user_roles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// RESTMap is the map of the REST API endpoints to the auth service refresh tokens gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Relative.String(): {
		api.POST:   detailsauth.AddUserRole,
		api.DELETE: detailsauth.RevokeUserRole,
	},
	ByUserId.String(): {
		api.GET: detailsauth.GetUserRoles,
	},
}
