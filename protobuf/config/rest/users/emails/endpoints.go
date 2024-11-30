package emails

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the user service emails REST endpoints
const BaseURI = "/emails"

// User service emails REST endpoints
var (
	Relative         = rest2.NewEndpoint("/")
	ByEmail          = rest2.NewEndpoint("/", rest2.Email)
	Primary          = rest2.NewEndpoint("/primary")
	SendVerification = rest2.NewEndpoint("/send-verification")
	VerifyByToken    = rest2.NewEndpoint(
		"/verify", rest2.Token,
	)
)
