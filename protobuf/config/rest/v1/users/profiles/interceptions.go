package profiles

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the user service profiles gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.GET: user.GetMyProfile,
	},
	ByUsername.String(): {
		rest.GET: user.GetProfile,
	},
}
