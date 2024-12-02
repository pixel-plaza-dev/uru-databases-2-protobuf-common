package users

import (
	configusers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Interceptions is the map of the REST API endpoints to the user service gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.PATCH: configusers.UpdateUser,
	},
	SignUp.String(): {
		rest.POST: configusers.SignUp,
	},
	UserIdByUsername.String(): {
		rest.GET: configusers.GetUserIdByUsername,
	},
	Password.String(): {
		rest.PUT: configusers.ChangePassword,
	},
	Username.String(): {
		rest.PUT: configusers.ChangeUsername,
	},
	ForgotPassword.String(): {
		rest.POST: configusers.ForgotPassword,
	},
	ResetPasswordByToken.String(): {
		rest.POST: configusers.ResetPassword,
	},
	DeleteAccount.String(): {
		rest.DELETE: configusers.DeleteUser,
	},
}
