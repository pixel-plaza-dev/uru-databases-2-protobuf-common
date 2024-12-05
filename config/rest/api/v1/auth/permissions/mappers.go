package permissions

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service permissions endpoints mapping
var (
	GetPermissionsMapper   = typesrest.NewMapper(Relative, grpcauth.GetPermissions)
	AddPermissionMapper    = typesrest.NewMapper(Relative, grpcauth.AddPermission)
	GetPermissionMapper    = typesrest.NewMapper(ByPermissionId, grpcauth.GetPermission)
	RevokePermissionMapper = typesrest.NewMapper(ByPermissionId, grpcauth.RevokePermission)
)
