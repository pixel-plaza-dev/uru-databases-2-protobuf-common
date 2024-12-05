package usernames

import (
	grpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service usernames endpoints handlers
var (
	GetUsernameByUserIdHandler = typesrest.NewHandler(ByUserId, grpcuser.GetUsernameByUserId)
	UsernameExistsHandler      = typesrest.NewHandler(ExistsByUsername, grpcuser.UsernameExists)
)
