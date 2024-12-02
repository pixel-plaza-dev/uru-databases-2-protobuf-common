package emails

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// User service emails REST endpoints
var (
	Relative         = typesrest.NewEndpoint("")
	ByEmail          = typesrest.NewEndpoint("", typesrest.Email)
	Primary          = typesrest.NewEndpoint("primary")
	SendVerification = typesrest.NewEndpoint("send-verification")
	VerifyByToken    = typesrest.NewEndpoint(
		"verify", typesrest.Token,
	)
)
