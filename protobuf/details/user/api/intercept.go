package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user"
)

// RESTMap is the map of the Rest API endpoints to the user service gRPC methods
var RESTMap = map[details.RESTEndpoint]map[details.RESTMethod]details.GRPCMethod{
	SignUp: {
		details.POST: detailsuser.SignUp,
	},
	Profile: {
		details.GET:   detailsuser.GetProfile,
		details.PATCH: detailsuser.UpdateProfile,
	},
	FullProfile: {
		details.GET: detailsuser.GetFullProfile,
	},
	Password: {
		details.PUT: detailsuser.ChangePassword,
	},
	Username: {
		details.PUT: detailsuser.ChangeUsername,
	},
	Email: {
		details.GET:  detailsuser.GetPrimaryEmail,
		details.PUT:  detailsuser.ChangePrimaryEmail,
		details.POST: detailsuser.AddEmail,
	},
	EmailByMail: {
		details.DELETE: detailsuser.DeleteEmail,
	},
	Emails: {
		details.GET: detailsuser.GetActiveEmails,
	},
	SendVerificationEmail: {
		details.POST: detailsuser.SendVerificationEmail,
	},
	VerifyEmailByToken: {
		details.POST: detailsuser.VerifyEmail,
	},
	PhoneNumber: {
		details.PUT: detailsuser.ChangePhoneNumber,
		details.GET: detailsuser.GetPhoneNumber,
	},
	SendVerificationSMS: {
		details.POST: detailsuser.SendVerificationSMS,
	},
	VerifyPhoneNumberByToken: {
		details.POST: detailsuser.VerifyPhoneNumber,
	},
	ForgotPassword: {
		details.POST: detailsuser.ForgotPassword,
	},
	ResetPasswordByToken: {
		details.POST: detailsuser.ResetPassword,
	},
	DeleteAccount: {
		details.DELETE: detailsuser.DeleteUser,
	},
}
