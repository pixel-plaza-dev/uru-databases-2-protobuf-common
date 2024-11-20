package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
)

// Auth service REST endpoints
const (
	LogIn               details.RESTEndpoint = "/log-in"
	LogOut              details.RESTEndpoint = "/log-out"
	RefreshToken        details.RESTEndpoint = "/refresh-token"
	Sessions            details.RESTEndpoint = "/sessions"
	CloseSessionByToken details.RESTEndpoint = "/close-session/:token"
	CloseSessions       details.RESTEndpoint = "/close-sessions"
	Permission          details.RESTEndpoint = "/permission"
	PermissionById      details.RESTEndpoint = "/permission/:id"
	Permissions         details.RESTEndpoint = "/permissions"
	RolePermission      details.RESTEndpoint = "/role-permission"
	RolePermissionById  details.RESTEndpoint = "/role-permission/:id"
	RolePermissions     details.RESTEndpoint = "/role-permissions"
	Role                details.RESTEndpoint = "/role"
	RoleById            details.RESTEndpoint = "/role/:id"
	Roles               details.RESTEndpoint = "/roles"
	UserRole            details.RESTEndpoint = "/user-role/:user_id"
	UserRoleById        details.RESTEndpoint = "/user-role/:id"
	UserRoles           details.RESTEndpoint = "/user-roles/:user_id"
)
