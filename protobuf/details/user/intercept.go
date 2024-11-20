package user

import "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"

// GRPCInterceptions is the list of gRPC methods to intercept
var GRPCInterceptions = map[details.GRPCMethod]details.Interception{
	SignUp:                details.None,
	IsPasswordCorrect:     details.None,
	UpdateProfile:         details.AccessToken,
	GetProfile:            details.AccessToken,
	GetFullProfile:        details.AccessToken,
	ChangePassword:        details.AccessToken,
	ChangeUsername:        details.AccessToken,
	AddEmail:              details.AccessToken,
	DeleteEmail:           details.AccessToken,
	SendVerificationEmail: details.AccessToken,
	VerifyEmail:           details.AccessToken,
	GetPrimaryEmail:       details.AccessToken,
	ChangePrimaryEmail:    details.AccessToken,
	GetActiveEmails:       details.AccessToken,
	ChangePhoneNumber:     details.AccessToken,
	GetPhoneNumber:        details.AccessToken,
	SendVerificationSMS:   details.AccessToken,
	VerifyPhoneNumber:     details.AccessToken,
	ForgotPassword:        details.None,
	ResetPassword:         details.None,
	DeleteUser:            details.AccessToken,
}
