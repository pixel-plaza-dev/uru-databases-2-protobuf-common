package usernames

import (
	rest2 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// BaseURI is the base URI for the user service usernames REST endpoints
const BaseURI = "/usernames"

// User service usernames REST endpoints
var (
	ById = rest2.NewEndpoint(
		"/",
		rest2.Id,
	)
	ExistsByUsername = rest2.NewEndpoint(
		"/exists",
		rest2.Username,
	)
)
