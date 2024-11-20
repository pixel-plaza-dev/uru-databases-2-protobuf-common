package auth

import "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"

// Auth service gRPC methods
var (
	LogIn                = details.NewGRPCMethod("LogIn")
	IsAccessTokenValid   = details.NewGRPCMethod("IsAccessTokenValid")
	IsRefreshTokenValid  = details.NewGRPCMethod("IsRefreshTokenValid")
	RefreshToken         = details.NewGRPCMethod("RefreshToken")
	LogOut               = details.NewGRPCMethod("LogOut")
	GetSessions          = details.NewGRPCMethod("GetSessions")
	CloseSession         = details.NewGRPCMethod("CloseSession")
	CloseSessions        = details.NewGRPCMethod("CloseSessions")
	AddPermission        = details.NewGRPCMethod("AddPermission")
	RevokePermission     = details.NewGRPCMethod("RevokePermission")
	GetPermission        = details.NewGRPCMethod("GetPermission")
	GetPermissions       = details.NewGRPCMethod("GetPermissions")
	AddRolePermission    = details.NewGRPCMethod("AddRolePermission")
	RevokeRolePermission = details.NewGRPCMethod("RevokeRolePermission")
	GetRolePermissions   = details.NewGRPCMethod("GetRolePermissions")
	AddRole              = details.NewGRPCMethod("AddRole")
	RevokeRole           = details.NewGRPCMethod("RevokeRole")
	GetRoles             = details.NewGRPCMethod("GetRoles")
	AddUserRole          = details.NewGRPCMethod("AddUserRole")
	RevokeUserRole       = details.NewGRPCMethod("RevokeUserRole")
	GetUserRoles         = details.NewGRPCMethod("GetUserRoles")
)
