package access_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the auth service access tokens gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Valid.String(): {
		rest.GET: auth.IsAccessTokenValid,
	},
}
