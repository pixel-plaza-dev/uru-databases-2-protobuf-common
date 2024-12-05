package roles

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service roles endpoints handlers
var (
	AddRoleHandler            = typesrest.NewHandler(Relative, grpcauth.AddRole)
	GetRolesHandler           = typesrest.NewHandler(Relative, grpcauth.GetRoles)
	AddRolePermissionHandler  = typesrest.NewHandler(ByRoleId, grpcauth.AddRolePermission)
	GetRolePermissionsHandler = typesrest.NewHandler(ByRoleId, grpcauth.GetRolePermissions)
	RevokeRoleHandler         = typesrest.NewHandler(ByRoleId, grpcauth.RevokeRole)
)
