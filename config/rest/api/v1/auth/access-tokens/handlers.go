package access_tokens

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service access tokens endpoints handlers
var (
	IsAccessTokenValidHandler = typesrest.NewHandler(ValidByJwtId, grpcauth.IsAccessTokenValid)
)
