package usernames

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
)

// BaseURI is the base URI for the user service usernames REST endpoints
const BaseURI = "/usernames"

// User service usernames REST endpoints
var (
	ById = api.NewRESTEndpoint(
		"/",
		api.Id,
	)
	ExistsByUsername = api.NewRESTEndpoint(
		"/exists",
		api.Username,
	)
)
