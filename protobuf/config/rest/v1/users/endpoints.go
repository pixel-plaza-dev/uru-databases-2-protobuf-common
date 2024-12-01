package users

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Users service REST endpoints
var (
	Relative     = typesrest.NewEndpoint("")
	SignUp       = typesrest.NewEndpoint("sign-up")
	Password     = typesrest.NewEndpoint("password")
	Username     = typesrest.NewEndpoint("username")
	IdByUsername = typesrest.NewEndpoint(
		"id",
		typesrest.Username,
	)
	ForgotPassword       = typesrest.NewEndpoint("forgot-password")
	ResetPasswordByToken = typesrest.NewEndpoint(
		"reset-password", typesrest.Token,
	)
	DeleteAccount = typesrest.NewEndpoint("delete-account")
)
