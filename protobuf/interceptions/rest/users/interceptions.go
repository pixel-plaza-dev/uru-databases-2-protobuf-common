package users

import (
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Map is the map of the Rest API endpoints to the user service gRPC methods
var Map = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.PATCH: detailsuser.UpdateUser,
	},
	SignUp.String(): {
		rest.POST: detailsuser.SignUp,
	},
	IdByUsername.String(): {
		rest.GET: detailsuser.GetUserIdByUsername,
	},
	Password.String(): {
		rest.PUT: detailsuser.ChangePassword,
	},
	Username.String(): {
		rest.PUT: detailsuser.ChangeUsername,
	},
	ForgotPassword.String(): {
		rest.POST: detailsuser.ForgotPassword,
	},
	ResetPasswordByToken.String(): {
		rest.POST: detailsuser.ResetPassword,
	},
	DeleteAccount.String(): {
		rest.DELETE: detailsuser.DeleteUser,
	},
}
