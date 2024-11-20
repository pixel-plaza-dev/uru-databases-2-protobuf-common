package user

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// User service gRPC methods
var (
	SignUp                = types.NewGRPCMethod("SignUp")
	UsernameExists        = types.NewGRPCMethod("UsernameExists")
	GetUserIdByUsername   = types.NewGRPCMethod("GetUserIdByUsername")
	GetUsernameByUserId   = types.NewGRPCMethod("GetUsernameByUserId")
	IsPasswordCorrect     = types.NewGRPCMethod("IsPasswordCorrect")
	UpdateProfile         = types.NewGRPCMethod("UpdateProfile")
	GetProfile            = types.NewGRPCMethod("GetProfile")
	GetFullProfile        = types.NewGRPCMethod("GetFullProfile")
	ChangePassword        = types.NewGRPCMethod("ChangePassword")
	ChangeUsername        = types.NewGRPCMethod("ChangeUsername")
	AddEmail              = types.NewGRPCMethod("AddEmail")
	DeleteEmail           = types.NewGRPCMethod("DeleteEmail")
	SendVerificationEmail = types.NewGRPCMethod("SendVerificationEmail")
	VerifyEmail           = types.NewGRPCMethod("VerifyEmail")
	GetPrimaryEmail       = types.NewGRPCMethod("GetPrimaryEmail")
	ChangePrimaryEmail    = types.NewGRPCMethod("ChangePrimaryEmail")
	GetActiveEmails       = types.NewGRPCMethod("GetActiveEmails")
	ChangePhoneNumber     = types.NewGRPCMethod("ChangePhoneNumber")
	GetPhoneNumber        = types.NewGRPCMethod("GetPhoneNumber")
	SendVerificationSMS   = types.NewGRPCMethod("SendVerificationSMS")
	VerifyPhoneNumber     = types.NewGRPCMethod("VerifyPhoneNumber")
	ForgotPassword        = types.NewGRPCMethod("ForgotPassword")
	ResetPassword         = types.NewGRPCMethod("ResetPassword")
	DeleteUser            = types.NewGRPCMethod("DeleteUser")
)
