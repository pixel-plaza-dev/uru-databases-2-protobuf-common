package markets

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service markets endpoints mapping
var (
	AddBusinessMarketCategoryMapper   = typesrest.NewMapper(Relative, grpcshop.AddBusinessMarketCategory)
	GetBusinessMarketCategoriesMapper = typesrest.NewMapper(ByBusinessId, grpcshop.GetBusinessMarketCategories)
)
