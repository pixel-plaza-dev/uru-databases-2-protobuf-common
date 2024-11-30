package emails

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the user service emails REST endpoints
const BaseURI = "/emails"

// User service emails REST endpoints
var (
	Relative         = rest.NewEndpoint("/")
	ByEmail          = rest.NewEndpoint("/", rest.Email)
	Primary          = rest.NewEndpoint("/primary")
	SendVerification = rest.NewEndpoint("/send-verification")
	VerifyByToken    = rest.NewEndpoint(
		"/verify", rest.Token,
	)
)
