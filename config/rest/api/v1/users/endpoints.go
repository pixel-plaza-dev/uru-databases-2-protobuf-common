package users

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the user service REST endpoints
var Base = typesrest.NewEndpoint("users")

// Users service REST endpoints
var (
	Relative         = typesrest.NewEndpoint("")
	SignUp           = typesrest.NewEndpoint("sign-up")
	Password         = typesrest.NewEndpoint("password")
	Username         = typesrest.NewEndpoint("username")
	UserIdByUsername = typesrest.NewEndpoint(
		"user-id",
		typesrest.Username,
	)
	ForgotPassword       = typesrest.NewEndpoint("forgot-password")
	ResetPasswordByToken = typesrest.NewEndpoint(
		"reset-password", typesrest.Token,
	)
	DeleteAccount = typesrest.NewEndpoint("delete-account")
)
