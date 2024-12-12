package businesses

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service endpoints business revisions mapping
var (
	OpenAdminRevisionToBusinessMapper        = typesrest.NewMapper(ByBusinessId, grpcshop.OpenAdminRevisionToBusiness)
	OpenAdminRevisionToBusinessProductMapper = typesrest.NewMapper(
		ProductsByProductId,
		grpcshop.OpenAdminRevisionToBusinessProduct,
	)
)
