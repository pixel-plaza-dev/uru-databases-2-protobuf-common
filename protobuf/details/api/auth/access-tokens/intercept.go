package access_tokens

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// RESTMap is the map of the REST API endpoints to the auth service access tokens gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Valid.String(): {
		api.GET: detailsauth.IsAccessTokenValid,
	},
}
