package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[grpc.Method]grpc.Interception{
	LogIn:                       grpc.None,
	IsAccessTokenValid:          grpc.None,
	IsRefreshTokenValid:         grpc.None,
	RefreshToken:                grpc.RefreshToken,
	LogOut:                      grpc.AccessToken,
	GetRefreshTokenInformation:  grpc.AccessToken,
	GetRefreshTokensInformation: grpc.AccessToken,
	RevokeRefreshToken:          grpc.AccessToken,
	RevokeRefreshTokens:         grpc.AccessToken,
	AddPermission:               grpc.AccessToken,
	RevokePermission:            grpc.AccessToken,
	GetPermission:               grpc.AccessToken,
	GetPermissions:              grpc.AccessToken,
	AddRolePermission:           grpc.AccessToken,
	RevokeRolePermission:        grpc.AccessToken,
	AddRole:                     grpc.AccessToken,
	RevokeRole:                  grpc.AccessToken,
	GetRoles:                    grpc.AccessToken,
	AddUserRole:                 grpc.AccessToken,
	RevokeUserRole:              grpc.AccessToken,
	GetUserRoles:                grpc.AccessToken,
}
