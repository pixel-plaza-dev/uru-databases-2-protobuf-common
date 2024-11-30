package permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// RESTMap is the map of the REST API endpoints to the auth service permissions gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Relative.String(): {
		api.GET:  detailsauth.GetPermissions,
		api.POST: detailsauth.AddPermission,
	},
	ById.String(): {
		api.GET:    detailsauth.GetPermission,
		api.DELETE: detailsauth.RevokePermission,
	},
}
