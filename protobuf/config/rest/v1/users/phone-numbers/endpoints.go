package phone_numbers

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// User service phone numbers REST endpoints
var (
	Relative         = rest.NewEndpoint("")
	SendVerification = rest.NewEndpoint("send-verification")
	VerifyByToken    = rest.NewEndpoint(
		"verify", rest.Token,
	)
)
