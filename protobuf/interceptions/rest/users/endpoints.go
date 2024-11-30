package users

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the user service REST endpoints
const BaseURI = "/users"

// Users service REST endpoints
var (
	Relative     = rest.NewEndpoint("/")
	SignUp       = rest.NewEndpoint("/sign-up")
	Password     = rest.NewEndpoint("/password")
	Username     = rest.NewEndpoint("/username")
	IdByUsername = rest.NewEndpoint(
		"/id",
		rest.Username,
	)
	ForgotPassword       = rest.NewEndpoint("/forgot-password")
	ResetPasswordByToken = rest.NewEndpoint(
		"/reset-password", rest.Token,
	)
	DeleteAccount = rest.NewEndpoint("/delete-account")
)
