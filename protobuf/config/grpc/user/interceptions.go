package user

import (
	typesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[typesgrpc.Method]typesgrpc.Interception{
	SignUp:                  typesgrpc.None,
	UsernameExists:          typesgrpc.None,
	GetUserIdByUsername:     typesgrpc.None,
	GetUsernameByUserId:     typesgrpc.None,
	GetUserSharedIdByUserId: typesgrpc.None,
	GetUserIdByUserSharedId: typesgrpc.None,
	IsPasswordCorrect:       typesgrpc.None,
	GetProfile:              typesgrpc.None,
	UpdateUser:              typesgrpc.AccessToken,
	GetMyProfile:            typesgrpc.AccessToken,
	ChangePassword:          typesgrpc.AccessToken,
	ChangeUsername:          typesgrpc.AccessToken,
	AddEmail:                typesgrpc.AccessToken,
	DeleteEmail:             typesgrpc.AccessToken,
	SendVerificationEmail:   typesgrpc.AccessToken,
	VerifyEmail:             typesgrpc.AccessToken,
	GetPrimaryEmail:         typesgrpc.AccessToken,
	ChangePrimaryEmail:      typesgrpc.AccessToken,
	GetActiveEmails:         typesgrpc.AccessToken,
	ChangePhoneNumber:       typesgrpc.AccessToken,
	GetPhoneNumber:          typesgrpc.AccessToken,
	SendVerificationSMS:     typesgrpc.AccessToken,
	VerifyPhoneNumber:       typesgrpc.AccessToken,
	ForgotPassword:          typesgrpc.None,
	ResetPassword:           typesgrpc.None,
	DeleteUser:              typesgrpc.AccessToken,
}
