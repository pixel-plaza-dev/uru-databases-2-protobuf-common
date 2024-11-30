package usernames

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// Map is the map of the REST API endpoints to the user service usernames gRPC methods
var Map = map[string]map[rest.Method]grpc.Method{
	ById.String(): {
		rest.GET: detailsuser.GetUsernameByUserId,
	},
	ExistsByUsername.String(): {
		rest.GET: detailsuser.UsernameExists,
	},
}
