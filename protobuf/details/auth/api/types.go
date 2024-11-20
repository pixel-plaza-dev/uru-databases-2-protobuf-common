package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
)

// Auth service REST endpoints
const (
	LogIn               = details.RESTEndpoint("/log-in")
	LogOut              = details.RESTEndpoint("/log-out")
	RefreshToken        = details.RESTEndpoint("/refresh-token")
	Sessions            = details.RESTEndpoint("/sessions")
	CloseSessionByToken = details.RESTEndpoint(
		"/close-session/:" + details.
			Token,
	)
	CloseSessions      = details.RESTEndpoint("/close-sessions")
	Permission         = details.RESTEndpoint("/permission")
	PermissionById     = details.RESTEndpoint("/permission/:" + details.Id)
	Permissions        = details.RESTEndpoint("/permissions")
	RolePermission     = details.RESTEndpoint("/role-permission")
	RolePermissionById = details.RESTEndpoint("/role-permission/:" + details.Id)
	RolePermissions    = details.RESTEndpoint("/role-permissions")
	Role               = details.RESTEndpoint("/role")
	RoleById           = details.RESTEndpoint("/role/:" + details.Id)
	Roles              = details.RESTEndpoint("/roles")
	UserRole           = details.RESTEndpoint("/user-role/:" + details.UserId)
	UserRoleById       = details.RESTEndpoint("/user-role/:" + details.Id)
	UserRoles          = details.RESTEndpoint("/user-roles/:" + details.UserId)
)
