package branches

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service branch products endpoints mapping
var (
	AddBranchProductMapper     = typesrest.NewMapper(ByBranchId, grpcshop.AddBranchProduct)
	GetBranchProductMapper     = typesrest.NewMapper(ByBranchId, grpcshop.GetBranchProduct)
	UpdateBranchProductMapper  = typesrest.NewMapper(ByBranchId, grpcshop.UpdateBranchProduct)
	SearchBranchProductsMapper = typesrest.NewMapper(Search, grpcshop.SearchBranchProducts)
)
