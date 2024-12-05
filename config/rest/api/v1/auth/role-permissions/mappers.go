package role_permissions

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service role permissions endpoints mapping
var (
	RevokeRolePermissionMapper = typesrest.NewMapper(ByRoleId, grpcauth.RevokeRolePermission)
)
