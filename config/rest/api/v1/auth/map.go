package auth

import (
	accesstokens "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/access-tokens"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/permissions"
	refreshtokens "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/refresh-tokens"
	rolepermissions "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/role-permissions"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/roles"
	userroles "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/user-roles"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// ChildrenMaps is the map of the REST API endpoints of the auth service
var ChildrenMaps = map[string]*rest.Map{
	AccessTokens.String():    accesstokens.Map,
	RefreshTokens.String():   refreshtokens.Map,
	Permissions.String():     permissions.Map,
	RolePermissions.String(): rolepermissions.Map,
	Roles.String():           roles.Map,
	UserRoles.String():       userroles.Map,
}

// Map is the map of the REST API endpoints of the auth service
var Map = rest.NewMap(
	&Interceptions,
	&ChildrenMaps,
)
