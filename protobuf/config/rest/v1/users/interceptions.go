package users

import (
	configusers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the user service gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.PATCH: configusers.UpdateUser,
	},
	SignUp.String(): {
		rest.POST: configusers.SignUp,
	},
	IdByUsername.String(): {
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
