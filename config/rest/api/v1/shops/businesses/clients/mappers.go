package clients

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service endpoints clients mapping
var (
	AddBusinessClientMapper = typesrest.NewMapper(ByBusinessId, grpcshop.AddBusinessClient)
	IsBusinessClientMapper  = typesrest.NewMapper(ByBusinessId, grpcshop.IsBusinessClient)
)
