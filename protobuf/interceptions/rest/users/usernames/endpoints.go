package usernames

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// BaseURI is the base URI for the user service usernames REST endpoints
const BaseURI = "/usernames"

// User service usernames REST endpoints
var (
	ById = rest.NewEndpoint(
		"/",
		rest.Id,
	)
	ExistsByUsername = rest.NewEndpoint(
		"/exists",
		rest.Username,
	)
)
