package api

import (
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// RESTMap is the map of the Rest API endpoints to the auth service gRPC methods
var RESTMap = map[*types.RESTEndpoint]map[types.RESTMethod]types.GRPCMethod{
	LogIn: {
		types.POST: detailsauth.LogIn,
	},
	RefreshToken: {
		types.POST: detailsauth.RefreshToken,
	},
	LogOut: {
		types.POST: detailsauth.LogOut,
	},
	Sessions: {
		types.GET: detailsauth.GetSessions,
	},
	CloseSessionByToken: {
		types.POST: detailsauth.CloseSession,
	},
	CloseSessions: {
		types.POST: detailsauth.CloseSessions,
	},
	Permission: {
		types.POST: detailsauth.AddPermission,
	},
	PermissionById: {
		types.GET:    detailsauth.GetPermission,
		types.DELETE: detailsauth.RevokePermission,
	},
	Permissions: {
		types.GET: detailsauth.GetPermissions,
	},
	RolePermission: {
		types.POST: detailsauth.AddRolePermission,
	},
	RolePermissionById: {
		types.DELETE: detailsauth.RevokeRolePermission,
	},
	RolePermissions: {
		types.GET: detailsauth.GetRolePermissions,
	},
	Role: {
		types.POST: detailsauth.AddRole,
	},
	RoleById: {
		types.DELETE: detailsauth.RevokeRole,
	},
	Roles: {
		types.GET: detailsauth.GetRoles,
	},
	UserRole: {
		types.POST: detailsauth.AddUserRole,
	},
	UserRoleById: {
		types.DELETE: detailsauth.RevokeUserRole,
	},
	UserRoles: {
		types.GET: detailsauth.GetUserRoles,
	},
}
