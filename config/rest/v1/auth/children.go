package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service children REST endpoints
var (
	AccessTokens    = rest.NewEndpoint("access-tokens")
	RefreshTokens   = rest.NewEndpoint("refresh-tokens")
	Permissions     = rest.NewEndpoint("permissions")
	RolePermissions = rest.NewEndpoint("role-permissions")
	Roles           = rest.NewEndpoint("roles")
	UserRoles       = rest.NewEndpoint("user-roles")
)
