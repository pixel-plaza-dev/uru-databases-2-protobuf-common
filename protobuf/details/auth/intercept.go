package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// GRPCInterceptions is the list of gRPC methods to intercept
var GRPCInterceptions = map[types.GRPCMethod]types.Interception{
	LogIn:                types.None,
	IsAccessTokenValid:   types.None,
	IsRefreshTokenValid:  types.None,
	RefreshToken:         types.RefreshToken,
	LogOut:               types.AccessToken,
	GetSessions:          types.AccessToken,
	CloseSession:         types.AccessToken,
	CloseSessions:        types.AccessToken,
	AddPermission:        types.AccessToken,
	RevokePermission:     types.AccessToken,
	GetPermission:        types.AccessToken,
	GetPermissions:       types.AccessToken,
	AddRolePermission:    types.AccessToken,
	RevokeRolePermission: types.AccessToken,
	AddRole:              types.AccessToken,
	RevokeRole:           types.AccessToken,
	GetRoles:             types.AccessToken,
	AddUserRole:          types.AccessToken,
	RevokeUserRole:       types.AccessToken,
	GetUserRoles:         types.AccessToken,
}
