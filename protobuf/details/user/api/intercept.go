package api

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user"
)

// RESTMap is the map of the Rest API endpoints to the user service gRPC methods
var RESTMap = map[string]map[types.RESTMethod]types.GRPCMethod{
	SignUp.String(): {
		types.POST: detailsuser.SignUp,
	},
	Profile.String(): {
		types.GET:   detailsuser.GetProfile,
		types.PATCH: detailsuser.UpdateProfile,
	},
	FullProfile.String(): {
		types.GET: detailsuser.GetFullProfile,
	},
	IdByUsername.String(): {
		types.GET: detailsuser.GetUserIdByUsername,
	},
	Password.String(): {
		types.PUT: detailsuser.ChangePassword,
	},
	Username.String(): {
		types.PUT: detailsuser.ChangeUsername,
	},
	UsernameByUsername.String(): {
		types.GET: detailsuser.UsernameExists,
	},
	UsernameById.String(): {
		types.GET: detailsuser.GetUsernameByUserId,
	},
	Email.String(): {
		types.GET:  detailsuser.GetPrimaryEmail,
		types.PUT:  detailsuser.ChangePrimaryEmail,
		types.POST: detailsuser.AddEmail,
	},
	EmailByEmail.String(): {
		types.DELETE: detailsuser.DeleteEmail,
	},
	Emails.String(): {
		types.GET: detailsuser.GetActiveEmails,
	},
	SendVerificationEmail.String(): {
		types.POST: detailsuser.SendVerificationEmail,
	},
	VerifyEmailByToken.String(): {
		types.POST: detailsuser.VerifyEmail,
	},
	PhoneNumber.String(): {
		types.PUT: detailsuser.ChangePhoneNumber,
		types.GET: detailsuser.GetPhoneNumber,
	},
	SendVerificationSMS.String(): {
		types.POST: detailsuser.SendVerificationSMS,
	},
	VerifyPhoneNumberByToken.String(): {
		types.POST: detailsuser.VerifyPhoneNumber,
	},
	ForgotPassword.String(): {
		types.POST: detailsuser.ForgotPassword,
	},
	ResetPasswordByToken.String(): {
		types.POST: detailsuser.ResetPassword,
	},
	DeleteAccount.String(): {
		types.DELETE: detailsuser.DeleteUser,
	},
}
