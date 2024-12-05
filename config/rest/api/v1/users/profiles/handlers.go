package profiles

import (
	grpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service profiles endpoints handlers
var (
	GetMyProfileHandler = typesrest.NewHandler(Relative, grpcuser.GetMyProfile)
	GetProfileHandler   = typesrest.NewHandler(ByUsername, grpcuser.GetProfile)
)
