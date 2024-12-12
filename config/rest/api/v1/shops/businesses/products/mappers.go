package products

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service businesses products endpoints mapping
var (
	AddBusinessProductMapper     = typesrest.NewMapper(Relative, grpcshop.AddBusinessProduct)
	GetBusinessProductMapper     = typesrest.NewMapper(ByProductId, grpcshop.GetBusinessProduct)
	UpdateBusinessProductMapper  = typesrest.NewMapper(ByProductId, grpcshop.UpdateBusinessProduct)
	SearchBusinessProductsMapper = typesrest.NewMapper(Search, grpcshop.SearchBusinessProducts)
)
