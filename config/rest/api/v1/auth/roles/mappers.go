package roles

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service roles endpoints mapping
var (
	AddRoleMapper            = typesrest.NewMapper(Relative, grpcauth.AddRole)
	GetRolesMapper           = typesrest.NewMapper(Relative, grpcauth.GetRoles)
	AddRolePermissionMapper  = typesrest.NewMapper(ByRoleId, grpcauth.AddRolePermission)
	GetRolePermissionsMapper = typesrest.NewMapper(ByRoleId, grpcauth.GetRolePermissions)
	RevokeRoleMapper         = typesrest.NewMapper(ByRoleId, grpcauth.RevokeRole)
)
