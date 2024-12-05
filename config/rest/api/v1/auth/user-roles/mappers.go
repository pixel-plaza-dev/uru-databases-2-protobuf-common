package user_roles

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service user roles endpoints mapping
var (
	AddUserRoleMapper    = typesrest.NewMapper(Relative, grpcauth.AddUserRole)
	RevokeUserRoleMapper = typesrest.NewMapper(Relative, grpcauth.RevokeUserRole)
	GetUserRolesMapper   = typesrest.NewMapper(ByUserId, grpcauth.GetUserRoles)
)
