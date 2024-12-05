package auth

import (
	grpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Auth service endpoints handlers
var (
	LogInHandler  = typesrest.NewHandler(LogIn, grpcauth.LogIn)
	LogOutHandler = typesrest.NewHandler(LogOut, grpcauth.LogOut)
)
