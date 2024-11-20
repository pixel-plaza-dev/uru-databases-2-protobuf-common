package user

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
)

// User service gRPC methods
const (
	SignUp                details.GRPCMethod = "SignUp"
	IsPasswordCorrect     details.GRPCMethod = "IsPasswordCorrect"
	UpdateProfile         details.GRPCMethod = "UpdateProfile"
	GetProfile            details.GRPCMethod = "GetProfile"
	GetFullProfile        details.GRPCMethod = "GetFullProfile"
	ChangePassword        details.GRPCMethod = "ChangePassword"
	ChangeUsername        details.GRPCMethod = "ChangeUsername"
	AddEmail              details.GRPCMethod = "AddEmail"
	DeleteEmail           details.GRPCMethod = "DeleteEmail"
	SendVerificationEmail details.GRPCMethod = "SendVerificationEmail"
	VerifyEmail           details.GRPCMethod = "VerifyEmail"
	GetPrimaryEmail       details.GRPCMethod = "GetPrimaryEmail"
	ChangePrimaryEmail    details.GRPCMethod = "ChangePrimaryEmail"
	GetActiveEmails       details.GRPCMethod = "GetActiveEmails"
	ChangePhoneNumber     details.GRPCMethod = "ChangePhoneNumber"
	GetPhoneNumber        details.GRPCMethod = "GetPhoneNumber"
	SendVerificationSMS   details.GRPCMethod = "SendVerificationSMS"
	VerifyPhoneNumber     details.GRPCMethod = "VerifyPhoneNumber"
	ForgotPassword        details.GRPCMethod = "ForgotPassword"
	ResetPassword         details.GRPCMethod = "ResetPassword"
	DeleteUser            details.GRPCMethod = "DeleteUser"
)
