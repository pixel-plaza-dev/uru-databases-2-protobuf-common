package stores

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service stores endpoints mapping
var (
	AddStoreMapper            = typesrest.NewMapper(Relative, grpcshop.AddStore)
	GetStoreMapper            = typesrest.NewMapper(ByStoreId, grpcshop.GetStore)
	DeleteStoreMapper         = typesrest.NewMapper(ByStoreId, grpcshop.DeleteStore)
	GetUnoccupiedStoresMapper = typesrest.NewMapper(Unoccupied, grpcshop.GetUnoccupiedStores)
)
