package role_permissions

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Auth service role permissions REST endpoints
var (
	ById = rest.NewEndpoint(
		"", rest.Id,
	)
)
