package usernames

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// User service usernames REST endpoints
var (
	ByUserId = typesrest.NewEndpoint(
		"",
		typesrest.UserId,
	)
	ExistsByUsername = typesrest.NewEndpoint(
		"exists",
		typesrest.Username,
	)
)
