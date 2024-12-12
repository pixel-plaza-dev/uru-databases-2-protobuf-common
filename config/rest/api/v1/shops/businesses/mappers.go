package businesses

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service endpoints mapping
var (
	AddBusinessMapper               = typesrest.NewMapper(Relative, grpcshop.AddBusiness)
	GetBusinessMapper               = typesrest.NewMapper(ByBusinessId, grpcshop.GetBusiness)
	UpdateBusinessMapper            = typesrest.NewMapper(ByBusinessId, grpcshop.UpdateBusiness)
	SetBusinessProfilePictureMapper = typesrest.NewMapper(
		ProfilePictureByBusinessId,
		grpcshop.SetBusinessProfilePicture,
	)
	SuspendBusinessMapper  = typesrest.NewMapper(SuspendByBusinessId, grpcshop.SuspendBusiness)
	ActivateBusinessMapper = typesrest.NewMapper(ActivateByBusinessId, grpcshop.ActivateBusiness)
	DeleteBusinessMapper   = typesrest.NewMapper(ByBusinessId, grpcshop.DeleteBusiness)
)
