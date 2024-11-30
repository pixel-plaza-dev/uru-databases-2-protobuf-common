package phone_numbers

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the user service phone numbers REST endpoints
const BaseURI = "/phone-numbers"

// User service phone numbers REST endpoints
var (
	Relative         = api.NewRESTEndpoint("/")
	SendVerification = api.NewRESTEndpoint("/send-verification")
	VerifyByToken    = api.NewRESTEndpoint(
		"/verify", api.Token,
	)
)
