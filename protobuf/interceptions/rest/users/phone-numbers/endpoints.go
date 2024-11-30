package phone_numbers

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the user service phone numbers REST endpoints
const BaseURI = "/phone-numbers"

// User service phone numbers REST endpoints
var (
	Relative         = rest.NewEndpoint("/")
	SendVerification = rest.NewEndpoint("/send-verification")
	VerifyByToken    = rest.NewEndpoint(
		"/verify", rest.Token,
	)
)
