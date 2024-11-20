package auth

import "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"

const (
	LogIn                details.GRPCMethod = "LogIn"
	IsAccessTokenValid   details.GRPCMethod = "IsAccessTokenValid"
	IsRefreshTokenValid  details.GRPCMethod = "IsRefreshTokenValid"
	RefreshToken         details.GRPCMethod = "RefreshToken"
	LogOut               details.GRPCMethod = "LogOut"
	GetSessions          details.GRPCMethod = "GetSessions"
	CloseSession         details.GRPCMethod = "CloseSession"
	CloseSessions        details.GRPCMethod = "CloseSessions"
	AddPermission        details.GRPCMethod = "AddPermission"
	RevokePermission     details.GRPCMethod = "RevokePermission"
	GetPermission        details.GRPCMethod = "GetPermission"
	GetPermissions       details.GRPCMethod = "GetPermissions"
	AddRolePermission    details.GRPCMethod = "AddRolePermission"
	RevokeRolePermission details.GRPCMethod = "RevokeRolePermission"
	GetRolePermissions   details.GRPCMethod = "GetRolePermissions"
	AddRole              details.GRPCMethod = "AddRole"
	RevokeRole           details.GRPCMethod = "RevokeRole"
	GetRoles             details.GRPCMethod = "GetRoles"
	AddUserRole          details.GRPCMethod = "AddUserRole"
	RevokeUserRole       details.GRPCMethod = "RevokeUserRole"
	GetUserRoles         details.GRPCMethod = "GetUserRoles"
)
