package refresh_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// RESTMap is the map of the REST API endpoints to the auth service refresh tokens gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Relative.String(): {
		api.POST:   detailsauth.RefreshToken,
		api.GET:    detailsauth.GetRefreshTokensInformation,
		api.DELETE: detailsauth.RevokeRefreshTokens,
	},
	ByJwtId.String(): {
		api.GET:    detailsauth.GetRefreshTokenInformation,
		api.DELETE: detailsauth.RevokeRefreshToken,
	},
	Valid.String(): {
		api.GET: detailsauth.IsRefreshTokenValid,
	},
}
