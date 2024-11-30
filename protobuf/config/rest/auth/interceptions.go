package auth

import (
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Map is the map of the REST API endpoints to the auth service gRPC methods
var Map = map[string]map[rest.Method]grpc.Method{
	LogIn.String(): {
		rest.POST: detailsauth.LogIn,
	},
	LogOut.String(): {
		rest.POST: detailsauth.LogOut,
	},
}
