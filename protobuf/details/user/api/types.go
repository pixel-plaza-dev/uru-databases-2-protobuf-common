package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
)

// User service REST endpoints
var (
	SignUp                = details.NewRESTEndpoint("/sign-up")
	Profile               = details.NewRESTEndpoint("/profile")
	FullProfile           = details.NewRESTEndpoint("/full-profile")
	Password              = details.NewRESTEndpoint("/password")
	Username              = details.NewRESTEndpoint("/username")
	Email                 = details.NewRESTEndpoint("/email")
	EmailByMail           = details.NewRESTEndpoint("/email", details.Email)
	Emails                = details.NewRESTEndpoint("/emails")
	SendVerificationEmail = details.NewRESTEndpoint("/send-verification-email")
	VerifyEmailByToken    = details.NewRESTEndpoint(
		"/verify-email", details.Token,
	)
	PhoneNumber              = details.NewRESTEndpoint("/phone-number")
	SendVerificationSMS      = details.NewRESTEndpoint("/send-verification-sms")
	VerifyPhoneNumberByToken = details.NewRESTEndpoint(
		"/verify-phone-number", details.Token,
	)
	ForgotPassword       = details.NewRESTEndpoint("/forgot-password")
	ResetPasswordByToken = details.NewRESTEndpoint(
		"/reset-password", details.Token,
	)
	DeleteAccount = details.NewRESTEndpoint("/delete-account")
)
