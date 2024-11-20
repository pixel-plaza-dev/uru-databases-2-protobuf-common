package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user"
)

// RESTMap is the map of the Rest API endpoints to the user service gRPC methods
var RESTMap = map[*types.RESTEndpoint]map[types.RESTMethod]types.GRPCMethod{
	SignUp: {
		types.POST: detailsuser.SignUp,
	},
	Profile: {
		types.GET:   detailsuser.GetProfile,
		types.PATCH: detailsuser.UpdateProfile,
	},
	FullProfile: {
		types.GET: detailsuser.GetFullProfile,
	},
	Password: {
		types.PUT: detailsuser.ChangePassword,
	},
	Username: {
		types.PUT: detailsuser.ChangeUsername,
	},
	Email: {
		types.GET:  detailsuser.GetPrimaryEmail,
		types.PUT:  detailsuser.ChangePrimaryEmail,
		types.POST: detailsuser.AddEmail,
	},
	EmailByMail: {
		types.DELETE: detailsuser.DeleteEmail,
	},
	Emails: {
		types.GET: detailsuser.GetActiveEmails,
	},
	SendVerificationEmail: {
		types.POST: detailsuser.SendVerificationEmail,
	},
	VerifyEmailByToken: {
		types.POST: detailsuser.VerifyEmail,
	},
	PhoneNumber: {
		types.PUT: detailsuser.ChangePhoneNumber,
		types.GET: detailsuser.GetPhoneNumber,
	},
	SendVerificationSMS: {
		types.POST: detailsuser.SendVerificationSMS,
	},
	VerifyPhoneNumberByToken: {
		types.POST: detailsuser.VerifyPhoneNumber,
	},
	ForgotPassword: {
		types.POST: detailsuser.ForgotPassword,
	},
	ResetPasswordByToken: {
		types.POST: detailsuser.ResetPassword,
	},
	DeleteAccount: {
		types.DELETE: detailsuser.DeleteUser,
	},
}
