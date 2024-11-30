package profiles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user"
)

// RESTMap is the map of the REST API endpoints to the user service profiles gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Relative.String(): {
		api.GET: detailsuser.GetMyProfile,
	},
	ByUsername.String(): {
		api.GET: detailsuser.GetProfile,
	},
}
