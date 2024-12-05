package auth

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service endpoints mapping
var (
	LogInMapper  = typesrest.NewMapper(LogIn, grpcauth.LogIn)
	LogOutMapper = typesrest.NewMapper(LogOut, grpcauth.LogOut)
)
