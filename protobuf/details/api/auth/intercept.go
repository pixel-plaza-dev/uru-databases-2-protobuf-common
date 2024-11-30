package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// RESTMap is the map of the Rest API endpoints to the auth service gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	LogIn.String(): {
		api.POST: detailsauth.LogIn,
	},
	LogOut.String(): {
		api.POST: detailsauth.LogOut,
	},
}
