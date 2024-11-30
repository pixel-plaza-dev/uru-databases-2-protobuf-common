package emails

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the user service emails REST endpoints
const BaseURI = "/emails"

// User service emails REST endpoints
var (
	Relative         = api.NewRESTEndpoint("/")
	ByEmail          = api.NewRESTEndpoint("/", api.Email)
	Primary          = api.NewRESTEndpoint("/primary")
	SendVerification = api.NewRESTEndpoint("/send-verification")
	VerifyByToken    = api.NewRESTEndpoint(
		"/verify", api.Token,
	)
)
