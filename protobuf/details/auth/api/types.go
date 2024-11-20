package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
)

// Auth service REST endpoints
var (
	LogIn               = details.NewRESTEndpoint("/log-in")
	LogOut              = details.NewRESTEndpoint("/log-out")
	RefreshToken        = details.NewRESTEndpoint("/refresh-token")
	Sessions            = details.NewRESTEndpoint("/sessions")
	CloseSessionByToken = details.NewRESTEndpoint(
		"/close-session", details.Token,
	)
	CloseSessions      = details.NewRESTEndpoint("/close-sessions")
	Permission         = details.NewRESTEndpoint("/permission")
	PermissionById     = details.NewRESTEndpoint("/permission/:" + details.Id.String())
	Permissions        = details.NewRESTEndpoint("/permissions")
	RolePermission     = details.NewRESTEndpoint("/role-permission")
	RolePermissionById = details.NewRESTEndpoint(
		"/role-permission", details.Id,
	)
	RolePermissions = details.NewRESTEndpoint("/role-permissions")
	Role            = details.NewRESTEndpoint("/role")
	RoleById        = details.NewRESTEndpoint("/role/:" + details.Id.String())
	Roles           = details.NewRESTEndpoint("/roles")
	UserRole        = details.NewRESTEndpoint(
		"/user-role", details.UserId,
	)
	UserRoleById = details.NewRESTEndpoint("/user-role/:" + details.Id.String())
	UserRoles    = details.NewRESTEndpoint(
		"/user-roles", details.UserId,
	)
)
