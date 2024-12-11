package categories

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service market categories endpoints mapping
var (
	AddMarketCategoryMapper    = typesrest.NewMapper(Relative, grpcshop.AddMarketCategory)
	GetMarketCategoryMapper    = typesrest.NewMapper(ByCategoryId, grpcshop.GetMarketCategory)
	UpdateMarketCategoryMapper = typesrest.NewMapper(ByCategoryId, grpcshop.UpdateMarketCategory)
)
