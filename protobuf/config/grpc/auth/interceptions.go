package auth

import (
	typesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[typesgrpc.Method]typesgrpc.Interception{
	LogIn:                       typesgrpc.None,
	IsAccessTokenValid:          typesgrpc.None,
	IsRefreshTokenValid:         typesgrpc.None,
	RefreshToken:                typesgrpc.RefreshToken,
	LogOut:                      typesgrpc.AccessToken,
	GetRefreshTokenInformation:  typesgrpc.AccessToken,
	GetRefreshTokensInformation: typesgrpc.AccessToken,
	RevokeRefreshToken:          typesgrpc.AccessToken,
	RevokeRefreshTokens:         typesgrpc.AccessToken,
	AddPermission:               typesgrpc.AccessToken,
	RevokePermission:            typesgrpc.AccessToken,
	GetPermission:               typesgrpc.AccessToken,
	GetPermissions:              typesgrpc.AccessToken,
	AddRolePermission:           typesgrpc.AccessToken,
	RevokeRolePermission:        typesgrpc.AccessToken,
	AddRole:                     typesgrpc.AccessToken,
	RevokeRole:                  typesgrpc.AccessToken,
	GetRoles:                    typesgrpc.AccessToken,
	AddUserRole:                 typesgrpc.AccessToken,
	RevokeUserRole:              typesgrpc.AccessToken,
	GetUserRoles:                typesgrpc.AccessToken,
}
