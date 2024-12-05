package refresh_tokens

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service refresh tokens endpoints mapping
var (
	RefreshTokenMapper                = typesrest.NewMapper(Relative, grpcauth.RefreshToken)
	GetRefreshTokensInformationMapper = typesrest.NewMapper(Relative, grpcauth.GetRefreshTokensInformation)
	RevokeRefreshTokensMapper         = typesrest.NewMapper(Relative, grpcauth.RevokeRefreshTokens)
	GetRefreshTokenInformationMapper  = typesrest.NewMapper(ByJwtId, grpcauth.GetRefreshTokenInformation)
	RevokeRefreshTokenMapper          = typesrest.NewMapper(ByJwtId, grpcauth.RevokeRefreshToken)
	IsRefreshTokenValidMapper         = typesrest.NewMapper(ValidByJwtId, grpcauth.IsRefreshTokenValid)
)
