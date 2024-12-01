package refresh_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the auth service refresh tokens gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.POST:   auth.RefreshToken,
		rest.GET:    auth.GetRefreshTokensInformation,
		rest.DELETE: auth.RevokeRefreshTokens,
	},
	ByJwtId.String(): {
		rest.GET:    auth.GetRefreshTokenInformation,
		rest.DELETE: auth.RevokeRefreshToken,
	},
	Valid.String(): {
		rest.GET: auth.IsRefreshTokenValid,
	},
}
