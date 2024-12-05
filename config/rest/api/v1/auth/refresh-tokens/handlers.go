package refresh_tokens

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service refresh tokens endpoints handlers
var (
	RefreshTokenHandler                = typesrest.NewHandler(Relative, grpcauth.RefreshToken)
	GetRefreshTokensInformationHandler = typesrest.NewHandler(Relative, grpcauth.GetRefreshTokensInformation)
	RevokeRefreshTokensHandler         = typesrest.NewHandler(Relative, grpcauth.RevokeRefreshTokens)
	GetRefreshTokenInformationHandler  = typesrest.NewHandler(ByJwtId, grpcauth.GetRefreshTokenInformation)
	RevokeRefreshTokenHandler          = typesrest.NewHandler(ByJwtId, grpcauth.RevokeRefreshToken)
	IsRefreshTokenValidHandler         = typesrest.NewHandler(ValidByJwtId, grpcauth.IsRefreshTokenValid)
)
