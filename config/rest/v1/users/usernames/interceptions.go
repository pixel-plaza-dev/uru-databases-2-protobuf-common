package usernames

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Interceptions is the map of the REST API endpoints to the user service usernames gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	ByUserId.String(): {
		rest.GET: user.GetUsernameByUserId,
	},
	ExistsByUsername.String(): {
		rest.GET: user.UsernameExists,
	},
}
