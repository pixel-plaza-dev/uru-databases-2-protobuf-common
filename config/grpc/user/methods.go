package user

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// User service gRPC methods
var (
	SignUp                  = grpc.NewMethod("SignUp")
	UsernameExists          = grpc.NewMethod("UsernameExists")
	GetUserIdByUsername     = grpc.NewMethod("GetUserIdByUsername")
	GetUsernameByUserId     = grpc.NewMethod("GetUsernameByUserId")
	GetUserSharedIdByUserId = grpc.NewMethod("GetUserSharedIdByUserId")
	GetUserIdByUserSharedId = grpc.NewMethod("GetUserIdByUserSharedId")
	IsPasswordCorrect       = grpc.NewMethod("IsPasswordCorrect")
	UpdateUser              = grpc.NewMethod("UpdateUser")
	GetProfile              = grpc.NewMethod("GetProfile")
	GetMyProfile            = grpc.NewMethod("GetMyProfile")
	ChangePassword          = grpc.NewMethod("ChangePassword")
	ChangeUsername          = grpc.NewMethod("ChangeUsername")
	AddEmail                = grpc.NewMethod("AddEmail")
	DeleteEmail             = grpc.NewMethod("DeleteEmail")
	SendVerificationEmail   = grpc.NewMethod("SendVerificationEmail")
	VerifyEmail             = grpc.NewMethod("VerifyEmail")
	GetPrimaryEmail         = grpc.NewMethod("GetPrimaryEmail")
	ChangePrimaryEmail      = grpc.NewMethod("ChangePrimaryEmail")
	GetActiveEmails         = grpc.NewMethod("GetActiveEmails")
	ChangePhoneNumber       = grpc.NewMethod("ChangePhoneNumber")
	GetPhoneNumber          = grpc.NewMethod("GetPhoneNumber")
	SendVerificationSMS     = grpc.NewMethod("SendVerificationSMS")
	VerifyPhoneNumber       = grpc.NewMethod("VerifyPhoneNumber")
	ForgotPassword          = grpc.NewMethod("ForgotPassword")
	ResetPassword           = grpc.NewMethod("ResetPassword")
	DeleteUser              = grpc.NewMethod("DeleteUser")
)
