package users

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the user service REST endpoints
const BaseURI = "/users"

// Users service REST endpoints
var (
	Relative     = rest2.NewEndpoint("/")
	SignUp       = rest2.NewEndpoint("/sign-up")
	Password     = rest2.NewEndpoint("/password")
	Username     = rest2.NewEndpoint("/username")
	IdByUsername = rest2.NewEndpoint(
		"/id",
		rest2.Username,
	)
	ForgotPassword       = rest2.NewEndpoint("/forgot-password")
	ResetPasswordByToken = rest2.NewEndpoint(
		"/reset-password", rest2.Token,
	)
	DeleteAccount = rest2.NewEndpoint("/delete-account")
)
