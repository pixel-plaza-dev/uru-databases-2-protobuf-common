package phone_numbers

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the user service phone numbers REST endpoints
var Base = typesrest.NewBaseEndpoint("phone-numbers")

// User service phone numbers REST endpoints
var (
	Relative         = typesrest.NewRelativeEndpoint()
	SendVerification = typesrest.NewEndpoint("send-verification")
	VerifyByToken    = typesrest.NewEndpoint(
		"verify", typesrest.Token,
	)
)
