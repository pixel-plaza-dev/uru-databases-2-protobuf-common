package users

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the user service REST endpoints
const BaseURI = "/users"

// Users service REST endpoints
var (
	Relative     = api.NewRESTEndpoint("/")
	SignUp       = api.NewRESTEndpoint("/sign-up")
	Password     = api.NewRESTEndpoint("/password")
	Username     = api.NewRESTEndpoint("/username")
	IdByUsername = api.NewRESTEndpoint(
		"/id",
		api.Username,
	)
	ForgotPassword       = api.NewRESTEndpoint("/forgot-password")
	ResetPasswordByToken = api.NewRESTEndpoint(
		"/reset-password", api.Token,
	)
	DeleteAccount = api.NewRESTEndpoint("/delete-account")
)
