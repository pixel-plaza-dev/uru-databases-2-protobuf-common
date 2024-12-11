package owners

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service endpoints owners mapping
var (
	AddBusinessOwnerMapper    = typesrest.NewMapper(ByBusinessId, grpcshop.AddBusinessOwner)
	RemoveBusinessOwnerMapper = typesrest.NewMapper(ByBusinessId, grpcshop.RemoveBusinessOwner)
	GetBusinessOwnersMapper   = typesrest.NewMapper(ByBusinessId, grpcshop.GetBusinessOwners)
)
