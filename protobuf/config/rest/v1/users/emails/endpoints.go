package emails

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// User service emails REST endpoints
var (
	Relative         = rest.NewEndpoint("")
	ByEmail          = rest.NewEndpoint("", rest.Email)
	Primary          = rest.NewEndpoint("primary")
	SendVerification = rest.NewEndpoint("send-verification")
	VerifyByToken    = rest.NewEndpoint(
		"verify", rest.Token,
	)
)
