package access_tokens

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service access tokens endpoints mapping
var (
	IsAccessTokenValidMapper = typesrest.NewMapper(ValidByJwtId, grpcauth.IsAccessTokenValid)
)
