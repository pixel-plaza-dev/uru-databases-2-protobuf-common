package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// Auth service REST endpoints
var (
	LogIn              = types.NewRESTEndpoint("/log-in")
	LogOut             = types.NewRESTEndpoint("/log-out")
	RefreshToken       = types.NewRESTEndpoint("/refresh-token")
	Sessions           = types.NewRESTEndpoint("/sessions")
	SessionByToken     = types.NewRESTEndpoint("/session", types.Token)
	Permission         = types.NewRESTEndpoint("/permission")
	PermissionById     = types.NewRESTEndpoint("/permission/:" + types.Id.String())
	Permissions        = types.NewRESTEndpoint("/permissions")
	RolePermission     = types.NewRESTEndpoint("/role-permission")
	RolePermissionById = types.NewRESTEndpoint(
		"/role-permission", types.Id,
	)
	RolePermissions = types.NewRESTEndpoint("/role-permissions")
	Role            = types.NewRESTEndpoint("/role")
	RoleById        = types.NewRESTEndpoint("/role/:" + types.Id.String())
	Roles           = types.NewRESTEndpoint("/roles")
	UserRole        = types.NewRESTEndpoint(
		"/user-role", types.UserId,
	)
	UserRoleById = types.NewRESTEndpoint("/user-role/:" + types.Id.String())
	UserRoles    = types.NewRESTEndpoint(
		"/user-roles", types.UserId,
	)
)
