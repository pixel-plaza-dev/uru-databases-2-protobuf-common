package usernames

import (
	grpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service usernames endpoints handlers
var (
	GetUsernameByUserIdMapper = typesrest.NewMapper(ByUserId, grpcuser.GetUsernameByUserId)
	UsernameExistsMapper      = typesrest.NewMapper(ExistsByUsername, grpcuser.UsernameExists)
)
