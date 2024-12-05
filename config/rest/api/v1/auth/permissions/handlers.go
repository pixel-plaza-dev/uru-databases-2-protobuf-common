package permissions

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service permissions endpoints handlers
var (
	GetPermissionsHandler   = typesrest.NewHandler(Relative, grpcauth.GetPermissions)
	AddPermissionHandler    = typesrest.NewHandler(Relative, grpcauth.AddPermission)
	GetPermissionHandler    = typesrest.NewHandler(ByPermissionId, grpcauth.GetPermission)
	RevokePermissionHandler = typesrest.NewHandler(ByPermissionId, grpcauth.RevokePermission)
)
