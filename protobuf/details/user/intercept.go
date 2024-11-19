package user

// MethodsToIntercept is a map of methods to not intercept
var MethodsToIntercept = map[string]bool{
	"SignUp":                false,
	"IsPasswordCorrect":     false,
	"UpdateProfile":         true,
	"GetProfile":            true,
	"GetFullProfile":        true,
	"ChangePassword":        true,
	"ChangeUsername":        true,
	"AddEmail":              true,
	"DeleteEmail":           true,
	"SendVerificationEmail": true,
	"VerifyEmail":           true,
	"GetPrimaryEmail":       true,
	"ChangePrimaryEmail":    true,
	"GetActiveEmails":       true,
	"ChangePhoneNumber":     true,
	"GetPhoneNumber":        true,
	"VerifyPhoneNumber":     true,
	"ForgotPassword":        false,
	"ResetPassword":         false,
	"DeleteUser":            true,
}
