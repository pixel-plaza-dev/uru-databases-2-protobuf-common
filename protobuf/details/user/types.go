package user

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
)

// User service gRPC methods
var (
	SignUp                = details.NewGRPCMethod("SignUp")
	IsPasswordCorrect     = details.NewGRPCMethod("IsPasswordCorrect")
	UpdateProfile         = details.NewGRPCMethod("UpdateProfile")
	GetProfile            = details.NewGRPCMethod("GetProfile")
	GetFullProfile        = details.NewGRPCMethod("GetFullProfile")
	ChangePassword        = details.NewGRPCMethod("ChangePassword")
	ChangeUsername        = details.NewGRPCMethod("ChangeUsername")
	AddEmail              = details.NewGRPCMethod("AddEmail")
	DeleteEmail           = details.NewGRPCMethod("DeleteEmail")
	SendVerificationEmail = details.NewGRPCMethod("SendVerificationEmail")
	VerifyEmail           = details.NewGRPCMethod("VerifyEmail")
	GetPrimaryEmail       = details.NewGRPCMethod("GetPrimaryEmail")
	ChangePrimaryEmail    = details.NewGRPCMethod("ChangePrimaryEmail")
	GetActiveEmails       = details.NewGRPCMethod("GetActiveEmails")
	ChangePhoneNumber     = details.NewGRPCMethod("ChangePhoneNumber")
	GetPhoneNumber        = details.NewGRPCMethod("GetPhoneNumber")
	SendVerificationSMS   = details.NewGRPCMethod("SendVerificationSMS")
	VerifyPhoneNumber     = details.NewGRPCMethod("VerifyPhoneNumber")
	ForgotPassword        = details.NewGRPCMethod("ForgotPassword")
	ResetPassword         = details.NewGRPCMethod("ResetPassword")
	DeleteUser            = details.NewGRPCMethod("DeleteUser")
)
