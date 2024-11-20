package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// User service REST endpoints
var (
	SignUp                = types.NewRESTEndpoint("/sign-up")
	Profile               = types.NewRESTEndpoint("/profile")
	FullProfile           = types.NewRESTEndpoint("/full-profile")
	Password              = types.NewRESTEndpoint("/password")
	Username              = types.NewRESTEndpoint("/username")
	Email                 = types.NewRESTEndpoint("/email")
	EmailByMail           = types.NewRESTEndpoint("/email", types.Email)
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