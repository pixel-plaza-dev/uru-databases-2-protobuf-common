package user

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[grpc.Method]grpc.Interception{
	SignUp:                grpc.None,
	UsernameExists:        grpc.None,
	GetUserIdByUsername:   grpc.None,
	GetUsernameByUserId:   grpc.None,
	IsPasswordCorrect:     grpc.None,
	GetProfile:            grpc.None,
	UpdateUser:            grpc.AccessToken,
	GetMyProfile:          grpc.AccessToken,
	ChangePassword:        grpc.AccessToken,
	ChangeUsername:        grpc.AccessToken,
	AddEmail:              grpc.AccessToken,
	DeleteEmail:           grpc.AccessToken,
	SendVerificationEmail: grpc.AccessToken,
	VerifyEmail:           grpc.AccessToken,
	GetPrimaryEmail:       grpc.AccessToken,
	ChangePrimaryEmail:    grpc.AccessToken,
	GetActiveEmails:       grpc.AccessToken,
	ChangePhoneNumber:     grpc.AccessToken,
	GetPhoneNumber:        grpc.AccessToken,
	SendVerificationSMS:   grpc.AccessToken,
	VerifyPhoneNumber:     grpc.AccessToken,
	ForgotPassword:        grpc.None,
	ResetPassword:         grpc.None,
	DeleteUser:            grpc.AccessToken,
}
