package users

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// User service children REST endpoints
var (
	Emails       = rest.NewEndpoint("emails")
	PhoneNumbers = rest.NewEndpoint("phone-numbers")
	Profiles     = rest.NewEndpoint("profiles")
	Usernames    = rest.NewEndpoint("usernames")
)
