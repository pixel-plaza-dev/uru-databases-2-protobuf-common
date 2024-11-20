package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
)

// RESTMap is the map of the Rest API endpoints to the auth service gRPC methods
var RESTMap = map[details.RESTEndpoint]map[details.RESTMethod]details.GRPCMethod{
	LogIn: {
		details.POST: detailsauth.LogIn,
	},
	RefreshToken: {
		details.POST: detailsauth.RefreshToken,
	},
	LogOut: {
		details.POST: detailsauth.LogOut,
	},
	Sessions: {
		details.GET: detailsauth.GetSessions,
	},
	CloseSessionByToken: {
		details.POST: detailsauth.CloseSession,
	},
	CloseSessions: {
		details.POST: detailsauth.CloseSessions,
	},
	Permission: {
		details.POST: detailsauth.AddPermission,
	},
	PermissionById: {
		details.GET:    detailsauth.GetPermission,
		details.DELETE: detailsauth.RevokePermission,
	},
	Permissions: {
		details.GET: detailsauth.GetPermissions,
	},
	RolePermission: {
		details.POST: detailsauth.AddRolePermission,
	},
	RolePermissionById: {
		details.DELETE: detailsauth.RevokeRolePermission,
	},
	RolePermissions: {
		details.GET: detailsauth.GetRolePermissions,
	},
	Role: {
		details.POST: detailsauth.AddRole,
	},
	RoleById: {
		details.DELETE: detailsauth.RevokeRole,
	},
	Roles: {
		details.GET: detailsauth.GetRoles,
	},
	UserRole: {
		details.POST: detailsauth.AddUserRole,
	},
	UserRoleById: {
		details.DELETE: detailsauth.RevokeUserRole,
	},
	UserRoles: {
		details.GET: detailsauth.GetUserRoles,
	},
}
