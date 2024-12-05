package users

import (
	configusers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service endpoints mapping
var (
	UpdateUserMapper          = rest.NewMapper(Relative, configusers.UpdateUser)
	SignUpMapper              = rest.NewMapper(SignUp, configusers.SignUp)
	GetUserIdByUsernameMapper = rest.NewMapper(UserIdByUsername, configusers.GetUserIdByUsername)
	ChangePasswordMapper      = rest.NewMapper(Password, configusers.ChangePassword)
	ChangeUsernameMapper      = rest.NewMapper(Username, configusers.ChangeUsername)
	ForgotPasswordMapper      = rest.NewMapper(ForgotPassword, configusers.ForgotPassword)
	PostResetPasswordMapper   = rest.NewMapper(ResetPasswordByToken, configusers.ResetPassword)
	DeleteAccountMapper       = rest.NewMapper(DeleteAccount, configusers.DeleteUser)
)
