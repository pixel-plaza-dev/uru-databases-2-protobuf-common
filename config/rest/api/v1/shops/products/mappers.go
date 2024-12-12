package products

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service products endpoints mapping
var (
	AddProductMapper     = typesrest.NewMapper(Relative, grpcshop.AddProduct)
	GetProductMapper     = typesrest.NewMapper(ByProductId, grpcshop.GetProduct)
	UpdateProductMapper  = typesrest.NewMapper(ByProductId, grpcshop.UpdateProduct)
	SearchProductsMapper = typesrest.NewMapper(Search, grpcshop.SearchProducts)
)
