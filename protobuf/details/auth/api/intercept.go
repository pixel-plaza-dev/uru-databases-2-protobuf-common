package api

import (
	detailsauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// RESTMap is the map of the Rest API endpoints to the auth service gRPC methods
var RESTMap = map[string]map[types.RESTMethod]types.GRPCMethod{
	LogIn.String(): {
		types.POST: detailsauth.LogIn,
	},
	AccessTokenByToken.String(): {
		types.GET: detailsauth.IsAccessTokenValid,
	},
	RefreshTokenByToken.String(): {
		types.GET: detailsauth.IsRefreshTokenValid,
	},
	RefreshToken.String(): {
		types.POST: detailsauth.RefreshToken,
	},
	LogOut.String(): {
		types.POST: detailsauth.LogOut,
	},
	Sessions.String(): {
		types.GET:    detailsauth.GetSessions,
		types.DELETE: detailsauth.CloseSessions,
	},
	SessionByToken.String(): {
		types.POST: detailsauth.CloseSession,
	},
	Permission.String(): {
		types.POST: detailsauth.AddPermission,
	},
	PermissionById.String(): {
		types.GET:    detailsauth.GetPermission,
		types.DELETE: detailsauth.RevokePermission,
	},
	Permissions.String(): {
		types.GET: detailsauth.GetPermissions,
	},
	RolePermissionById.String(): {
		types.DELETE: detailsauth.RevokeRolePermission,
	},
	Role.String(): {
		types.POST: detailsauth.AddRole,
	},
	RoleById.String(): {
		types.POST:   detailsauth.AddRolePermission,
		types.GET:    detailsauth.GetRolePermissions,
		types.DELETE: detailsauth.RevokeRole,
	},
	Roles.String(): {
		types.GET: detailsauth.GetRoles,
	},
	UserRoleByUserId.String(): {
		types.POST:   detailsauth.AddUserRole,
		types.DELETE: detailsauth.RevokeUserRole,
	},
	UserRolesByUserId.String(): {
		types.GET: detailsauth.GetUserRoles,
	},
}
