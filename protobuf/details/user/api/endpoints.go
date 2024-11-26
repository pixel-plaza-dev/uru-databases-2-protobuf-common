package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// User service REST endpoints
var (
	SignUp       = types.NewRESTEndpoint("/sign-up")
	Profile      = types.NewRESTEndpoint("/profile")
	MyProfile    = types.NewRESTEndpoint("/my-profile")
	IdByUsername = types.NewRESTEndpoint(
		"/id-by-username",
		types.Username,
	)
	Password     = types.NewRESTEndpoint("/password")
	Username     = types.NewRESTEndpoint("/username")
	UsernameById = types.NewRESTEndpoint(
		"/username-by-id",
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
