package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// Auth service gRPC methods
var (
	LogIn                = types.NewGRPCMethod("LogIn")
	IsAccessTokenValid   = types.NewGRPCMethod("IsAccessTokenValid")
	IsRefreshTokenValid  = types.NewGRPCMethod("IsRefreshTokenValid")
	RefreshToken         = types.NewGRPCMethod("RefreshToken")
	LogOut               = types.NewGRPCMethod("LogOut")
	GetSessions          = types.NewGRPCMethod("GetSessions")
	CloseSession         = types.NewGRPCMethod("CloseSession")
	CloseSessions        = types.NewGRPCMethod("CloseSessions")
	AddPermission        = types.NewGRPCMethod("AddPermission")
	RevokePermission     = types.NewGRPCMethod("RevokePermission")
	GetPermission        = types.NewGRPCMethod("GetPermission")
	GetPermissions       = types.NewGRPCMethod("GetPermissions")
	AddRolePermission    = types.NewGRPCMethod("AddRolePermission")
	RevokeRolePermission = types.NewGRPCMethod("RevokeRolePermission")
	GetRolePermissions   = types.NewGRPCMethod("GetRolePermissions")
	AddRole              = types.NewGRPCMethod("AddRole")
	RevokeRole           = types.NewGRPCMethod("RevokeRole")
	GetRoles             = types.NewGRPCMethod("GetRoles")
	AddUserRole          = types.NewGRPCMethod("AddUserRole")
	RevokeUserRole       = types.NewGRPCMethod("RevokeUserRole")
	GetUserRoles         = types.NewGRPCMethod("GetUserRoles")
)
