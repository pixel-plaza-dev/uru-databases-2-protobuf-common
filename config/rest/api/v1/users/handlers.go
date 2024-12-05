package users

import (
	configusers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service endpoints handlers
var (
	UpdateUserHandler          = rest.NewHandler(Relative, configusers.UpdateUser)
	SignUpHandler              = rest.NewHandler(SignUp, configusers.SignUp)
	GetUserIdByUsernameHandler = rest.NewHandler(UserIdByUsername, configusers.GetUserIdByUsername)
	ChangePasswordHandler      = rest.NewHandler(Password, configusers.ChangePassword)
	ChangeUsernameHandler      = rest.NewHandler(Username, configusers.ChangeUsername)
	ForgotPasswordHandler      = rest.NewHandler(ForgotPassword, configusers.ForgotPassword)
	PostResetPasswordHandler   = rest.NewHandler(ResetPasswordByToken, configusers.ResetPassword)
	DeleteAccountHandler       = rest.NewHandler(DeleteAccount, configusers.DeleteUser)
)
