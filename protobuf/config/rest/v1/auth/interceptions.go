package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the auth service gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	LogIn.String(): {
		rest.POST: auth.LogIn,
	},
	LogOut.String(): {
		rest.POST: auth.LogOut,
	},
}
