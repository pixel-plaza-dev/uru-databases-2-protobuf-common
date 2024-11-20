package api

// User REST endpoints
const (
	SignUp                   = "/sign-up"
	Profile                  = "/profile"
	FullProfile              = "/full-profile"
	Password                 = "/password"
	Username                 = "/username"
	Email                    = "/email"
	EmailByMail              = "/email/:email"
	Emails                   = "/emails"
	SendVerificationEmail    = "/send-verification-email"
	VerifyEmailByToken       = "/verify-email/:token"
	PhoneNumber              = "/phone-number"
	SendVerificationSMS      = "/send-verification-sms"
	VerifyPhoneNumberByToken = "/verify-phone-number/:token"
	ForgotPassword           = "/forgot-password"
	ResetPasswordByToken     = "/reset-password/:token"
	DeleteAccount            = "/delete-account"
)
