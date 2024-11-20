package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
)

// User service REST endpoints
const (
	SignUp                = details.RESTEndpoint("/sign-up")
	Profile               = details.RESTEndpoint("/profile")
	FullProfile           = details.RESTEndpoint("/full-profile")
	Password              = details.RESTEndpoint("/password")
	Username              = details.RESTEndpoint("/username")
	Email                 = details.RESTEndpoint("/email")
	EmailByMail           = details.RESTEndpoint("/email/:" + details.Email)
	Emails                = details.RESTEndpoint("/emails")
	SendVerificationEmail = details.RESTEndpoint("/send-verification-email")
	VerifyEmailByToken    = details.RESTEndpoint(
		"/verify-email/:" + details.Token,
	)
	PhoneNumber              = details.RESTEndpoint("/phone-number")
	SendVerificationSMS      = details.RESTEndpoint("/send-verification-sms")
	VerifyPhoneNumberByToken = details.RESTEndpoint(
		"/verify-phone-number/:" + details.Token,
	)
	ForgotPassword       = details.RESTEndpoint("/forgot-password")
	ResetPasswordByToken = details.RESTEndpoint(
		"/reset-password/:" + details.Token,
	)
	DeleteAccount = details.RESTEndpoint("/delete-account")
)
