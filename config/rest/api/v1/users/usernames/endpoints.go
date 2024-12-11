package usernames

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the user service usernames REST endpoints
var Base = typesrest.NewBaseEndpoint("usernames")

// User service usernames REST endpoints
var (
	ByUserId = typesrest.NewRelativeEndpoint(
		typesrest.UserId,
	)
	ExistsByUsername = typesrest.NewEndpoint(
		"exists",
		typesrest.Username,
	)
)
