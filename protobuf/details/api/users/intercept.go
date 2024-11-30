package users

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user"
)

// RESTMap is the map of the Rest API endpoints to the user service gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Relative.String(): {
		api.PATCH: detailsuser.UpdateUser,
	},
	SignUp.String(): {
		api.POST: detailsuser.SignUp,
	},
	IdByUsername.String(): {
		api.GET: detailsuser.GetUserIdByUsername,
	},
	Password.String(): {
		api.PUT: detailsuser.ChangePassword,
	},
	Username.String(): {
		api.PUT: detailsuser.ChangeUsername,
	},
	ForgotPassword.String(): {
		api.POST: detailsuser.ForgotPassword,
	},
	ResetPasswordByToken.String(): {
		api.POST: detailsuser.ResetPassword,
	},
	DeleteAccount.String(): {
		api.DELETE: detailsuser.DeleteUser,
	},
}
