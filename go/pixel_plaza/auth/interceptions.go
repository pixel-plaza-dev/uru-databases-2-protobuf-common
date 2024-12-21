package auth

import (
	compiled "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	"github.com/ralvarezdev/go-jwt/token/interception"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[string]interception.Interception{
	compiled.Auth_LogIn_FullMethodName:                       interception.None,
	compiled.Auth_IsAccessTokenValid_FullMethodName:          interception.None,
	compiled.Auth_IsRefreshTokenValid_FullMethodName:         interception.None,
	compiled.Auth_RefreshToken_FullMethodName:                interception.RefreshToken,
	compiled.Auth_LogOut_FullMethodName:                      interception.AccessToken,
	compiled.Auth_GetRefreshTokenInformation_FullMethodName:  interception.AccessToken,
	compiled.Auth_GetRefreshTokensInformation_FullMethodName: interception.AccessToken,
	compiled.Auth_RevokeRefreshToken_FullMethodName:          interception.AccessToken,
	compiled.Auth_RevokeRefreshTokens_FullMethodName:         interception.AccessToken,
	compiled.Auth_AddPermission_FullMethodName:               interception.AccessToken,
	compiled.Auth_RevokePermission_FullMethodName:            interception.AccessToken,
	compiled.Auth_GetPermission_FullMethodName:               interception.AccessToken,
	compiled.Auth_GetPermissions_FullMethodName:              interception.AccessToken,
	compiled.Auth_AddRolePermission_FullMethodName:           interception.AccessToken,
	compiled.Auth_RevokeRolePermission_FullMethodName:        interception.AccessToken,
	compiled.Auth_AddRole_FullMethodName:                     interception.AccessToken,
	compiled.Auth_RevokeRole_FullMethodName:                  interception.AccessToken,
	compiled.Auth_GetRoles_FullMethodName:                    interception.AccessToken,
	compiled.Auth_AddUserRole_FullMethodName:                 interception.AccessToken,
	compiled.Auth_RevokeUserRole_FullMethodName:              interception.AccessToken,
	compiled.Auth_GetUserRoles_FullMethodName:                interception.AccessToken,
}
