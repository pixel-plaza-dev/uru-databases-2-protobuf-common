package revisions

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service endpoints revisions mapping
var (
	UpdateAdminRevisionMapper        = typesrest.NewMapper(Relative, grpcshop.UpdateAdminRevision)
	CloseAdminRevisionMapper         = typesrest.NewMapper(Relative, grpcshop.CloseAdminRevision)
	OpenAdminRevisionToBranchMapper  = typesrest.NewMapper(BranchesByBranchId, grpcshop.OpenAdminRevisionToBranch)
	OpenAdminRevisionToProductMapper = typesrest.NewMapper(ProductsByProductId, grpcshop.OpenAdminRevisionToProduct)
)
