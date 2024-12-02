package refresh_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
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
	ValidByJwtId.String(): {
		rest.GET: auth.IsRefreshTokenValid,
	},
}
