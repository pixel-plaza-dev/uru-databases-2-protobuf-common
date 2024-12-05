package role_permissions

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service role permissions endpoints handlers
var (
	RevokeRolePermissionHandler = typesrest.NewHandler(ByRoleId, grpcauth.RevokeRolePermission)
)
