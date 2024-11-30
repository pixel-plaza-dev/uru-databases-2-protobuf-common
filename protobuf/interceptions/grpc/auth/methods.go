package auth

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc"
)

// Auth service gRPC methods
var (
	LogIn                       = grpc.NewMethod("LogIn")
	IsAccessTokenValid          = grpc.NewMethod("IsAccessTokenValid")
	IsRefreshTokenValid         = grpc.NewMethod("IsRefreshTokenValid")
	RefreshToken                = grpc.NewMethod("RefreshToken")
	LogOut                      = grpc.NewMethod("LogOut")
	GetRefreshTokenInformation  = grpc.NewMethod("GetRefreshTokenInformation")
	GetRefreshTokensInformation = grpc.NewMethod("GetRefreshTokensInformation")
	RevokeRefreshToken          = grpc.NewMethod("RevokeRefreshToken")
	RevokeRefreshTokens         = grpc.NewMethod("RevokeRefreshTokens")
	AddPermission               = grpc.NewMethod("AddPermission")
	RevokePermission            = grpc.NewMethod("RevokePermission")
	GetPermission               = grpc.NewMethod("GetPermission")
	GetPermissions              = grpc.NewMethod("GetPermissions")
	AddRolePermission           = grpc.NewMethod("AddRolePermission")
	RevokeRolePermission        = grpc.NewMethod("RevokeRolePermission")
	GetRolePermissions          = grpc.NewMethod("GetRolePermissions")
	AddRole                     = grpc.NewMethod("AddRole")
	RevokeRole                  = grpc.NewMethod("RevokeRole")
	GetRoles                    = grpc.NewMethod("GetRoles")
	AddUserRole                 = grpc.NewMethod("AddUserRole")
	RevokeUserRole              = grpc.NewMethod("RevokeUserRole")
	GetUserRoles                = grpc.NewMethod("GetUserRoles")
)
