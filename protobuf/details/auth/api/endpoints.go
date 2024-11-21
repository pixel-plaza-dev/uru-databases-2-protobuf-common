package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// Auth service REST endpoints
var (
	LogIn              = types.NewRESTEndpoint("/log-in")
	LogOut             = types.NewRESTEndpoint("/log-out")
	AccessTokenByToken = types.NewRESTEndpoint(
		"/access-token",
		types.Token,
	)
	RefreshToken        = types.NewRESTEndpoint("/refresh-token")
	RefreshTokenByToken = types.NewRESTEndpoint(
		"/refresh-token",
		types.Token,
	)
	Sessions           = types.NewRESTEndpoint("/sessions")
	SessionByToken     = types.NewRESTEndpoint("/session", types.Token)
	Permission         = types.NewRESTEndpoint("/permission")
	PermissionById     = types.NewRESTEndpoint("/permission/:" + types.Id.String())
	Permissions        = types.NewRESTEndpoint("/permissions")
	RolePermissionById = types.NewRESTEndpoint(
		"/role-permission", types.Id,
	)
	Role             = types.NewRESTEndpoint("/role")
	RoleById         = types.NewRESTEndpoint("/role/:" + types.Id.String())
	Roles            = types.NewRESTEndpoint("/roles")
	UserRoleByUserId = types.NewRESTEndpoint(
		"/user-role", types.UserId,
	)
	UserRoleById      = types.NewRESTEndpoint("/user-role/:" + types.Id.String())
	UserRolesByUserId = types.NewRESTEndpoint(
		"/user-roles", types.UserId,
	)
)
