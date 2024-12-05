package profiles

import (
	grpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service profiles endpoints mapping
var (
	GetMyProfileMapper = typesrest.NewMapper(Relative, grpcuser.GetMyProfile)
	GetProfileMapper   = typesrest.NewMapper(ByUsername, grpcuser.GetProfile)
)
