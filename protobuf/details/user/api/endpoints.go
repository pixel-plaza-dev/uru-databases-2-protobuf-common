package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// User service REST endpoints
var (
	SignUp          = types.NewRESTEndpoint("/sign-up")
	Profile         = types.NewRESTEndpoint("/profile")
	FullProfile     = types.NewRESTEndpoint("/full-profile")
	GetIdByUsername = types.NewRESTEndpoint(
		"/get-id",
		types.Username,
	)
	Password        = types.NewRESTEndpoint("/password")
	Username        = types.NewRESTEndpoint("/username")
	GetUsernameById = types.NewRESTEndpoint(
		"/get-username",
		types.Id,
	)
	UsernameExistsByUsername = types.NewRESTEndpoint(
		"/username-exists",
		types.Username,
	)
	Email                 = types.NewRESTEndpoint("/email")
	EmailByEmail          = types.NewRESTEndpoint("/email", types.Email)
	Emails                = types.NewRESTEndpoint("/emails")
	SendVerificationEmail = types.NewRESTEndpoint("/send-verification-email")
	VerifyEmailByToken    = types.NewRESTEndpoint(
		"/verify-email", types.Token,
	)
	PhoneNumber              = types.NewRESTEndpoint("/phone-number")
	SendVerificationSMS      = types.NewRESTEndpoint("/send-verification-sms")
	VerifyPhoneNumberByToken = types.NewRESTEndpoint(
		"/verify-phone-number", types.Token,
	)
	ForgotPassword       = types.NewRESTEndpoint("/forgot-password")
	ResetPasswordByToken = types.NewRESTEndpoint(
		"/reset-password", types.Token,
	)
	DeleteAccount = types.NewRESTEndpoint("/delete-account")
)
