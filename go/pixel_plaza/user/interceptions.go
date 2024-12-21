package user

import (
	compiled "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	"github.com/ralvarezdev/go-jwt/token/interception"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[string]interception.Interception{
	compiled.User_SignUp_FullMethodName:                interception.None,
	compiled.User_UsernameExists_FullMethodName:        interception.None,
	compiled.User_GetUserIdByUsername_FullMethodName:   interception.None,
	compiled.User_GetUsernameByUserId_FullMethodName:   interception.None,
	compiled.User_IsPasswordCorrect_FullMethodName:     interception.None,
	compiled.User_GetProfile_FullMethodName:            interception.None,
	compiled.User_UpdateUser_FullMethodName:            interception.AccessToken,
	compiled.User_SetProfilePicture_FullMethodName:     interception.AccessToken,
	compiled.User_GetMyProfile_FullMethodName:          interception.AccessToken,
	compiled.User_ChangePassword_FullMethodName:        interception.AccessToken,
	compiled.User_ChangeUsername_FullMethodName:        interception.AccessToken,
	compiled.User_AddEmail_FullMethodName:              interception.AccessToken,
	compiled.User_DeleteEmail_FullMethodName:           interception.AccessToken,
	compiled.User_SendVerificationEmail_FullMethodName: interception.AccessToken,
	compiled.User_VerifyEmail_FullMethodName:           interception.AccessToken,
	compiled.User_GetPrimaryEmail_FullMethodName:       interception.AccessToken,
	compiled.User_ChangePrimaryEmail_FullMethodName:    interception.AccessToken,
	compiled.User_GetActiveEmails_FullMethodName:       interception.AccessToken,
	compiled.User_ChangePhoneNumber_FullMethodName:     interception.AccessToken,
	compiled.User_GetPhoneNumber_FullMethodName:        interception.AccessToken,
	compiled.User_SendVerificationSMS_FullMethodName:   interception.AccessToken,
	compiled.User_VerifyPhoneNumber_FullMethodName:     interception.AccessToken,
	compiled.User_ForgotPassword_FullMethodName:        interception.None,
	compiled.User_ResetPassword_FullMethodName:         interception.None,
	compiled.User_DeleteUser_FullMethodName:            interception.AccessToken,
}
