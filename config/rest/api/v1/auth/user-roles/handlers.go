package user_roles

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service user roles endpoints handlers
var (
	AddUserRoleHandler    = typesrest.NewHandler(Relative, grpcauth.AddUserRole)
	RevokeUserRoleHandler = typesrest.NewHandler(Relative, grpcauth.RevokeUserRole)
	GetUserRolesHandler   = typesrest.NewHandler(ByUserId, grpcauth.GetUserRoles)
)
