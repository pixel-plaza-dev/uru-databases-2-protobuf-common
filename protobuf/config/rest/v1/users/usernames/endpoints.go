package usernames

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// User service usernames REST endpoints
var (
	ById = rest.NewEndpoint(
		"",
		rest.Id,
	)
	ExistsByUsername = rest.NewEndpoint(
		"exists",
		rest.Username,
	)
)
