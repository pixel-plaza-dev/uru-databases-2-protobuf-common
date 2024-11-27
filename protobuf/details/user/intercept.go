package user

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
)

// GRPCInterceptions is the list of gRPC methods to intercept
var GRPCInterceptions = map[types.GRPCMethod]types.Interception{
	SignUp:                types.None,
	UsernameExists:        types.None,
	GetUserIdByUsername:   types.None,
	GetUsernameByUserId:   types.None,
	IsPasswordCorrect:     types.None,
	UpdateProfile:         types.AccessToken,
	GetProfile:            types.AccessToken,
	GetMyProfile:          types.AccessToken,
	ChangePassword:        types.AccessToken,
	ChangeUsername:        types.AccessToken,
	AddEmail:              types.AccessToken,
	DeleteEmail:           types.AccessToken,
	SendVerificationEmail: types.AccessToken,
	VerifyEmail:           types.AccessToken,
	GetPrimaryEmail:       types.AccessToken,
	ChangePrimaryEmail:    types.AccessToken,
	GetActiveEmails:       types.AccessToken,
	ChangePhoneNumber:     types.AccessToken,
	GetPhoneNumber:        types.AccessToken,
	SendVerificationSMS:   types.AccessToken,
	VerifyPhoneNumber:     types.AccessToken,
	ForgotPassword:        types.None,
	ResetPassword:         types.None,
	DeleteUser:            types.AccessToken,
}
