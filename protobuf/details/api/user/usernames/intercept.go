package usernames

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user"
)

// RESTMap is the map of the REST API endpoints to the user service usernames gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	ById.String(): {
		api.GET: detailsuser.GetUsernameByUserId,
	},
	ExistsByUsername.String(): {
		api.GET: detailsuser.UsernameExists,
	},
}
