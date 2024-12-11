package categories

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service product categories endpoints mapping
var (
	AddProductCategoryMapper    = typesrest.NewMapper(Relative, grpcshop.AddProductCategory)
	GetProductCategoryMapper    = typesrest.NewMapper(ByCategoryId, grpcshop.GetProductCategory)
	UpdateProductCategoryMapper = typesrest.NewMapper(ByCategoryId, grpcshop.UpdateProductCategory)
)
