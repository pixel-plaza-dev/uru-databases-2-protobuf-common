package phone_numbers

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the user service phone numbers REST endpoints
const BaseURI = "/phone-numbers"

// User service phone numbers REST endpoints
var (
	Relative         = rest2.NewEndpoint("/")
	SendVerification = rest2.NewEndpoint("/send-verification")
	VerifyByToken    = rest2.NewEndpoint(
		"/verify", rest2.Token,
	)
)
