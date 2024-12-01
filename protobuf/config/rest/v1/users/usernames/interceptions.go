package usernames

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the user service usernames gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	ById.String(): {
		rest.GET: user.GetUsernameByUserId,
	},
	ExistsByUsername.String(): {
		rest.GET: user.UsernameExists,
	},
}
