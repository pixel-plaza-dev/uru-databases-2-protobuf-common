package rents

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service stores rents endpoints mapping
var (
	AddBranchRentMapper                = typesrest.NewMapper(Relative, grpcshop.AddBranchRent)
	GetBranchRentsMapper               = typesrest.NewMapper(ByBranchId, grpcshop.GetBranchRents)
	UpdateBranchRentMapper             = typesrest.NewMapper(ByBranchRentId, grpcshop.UpdateBranchRent)
	GetUnpaidBranchRentsMapper         = typesrest.NewMapper(BranchUnpaidByBranchId, grpcshop.GetUnpaidBranchRents)
	GetBusinessUnpaidBranchRentsMapper = typesrest.NewMapper(
		BusinessUnpaidByBusinessId,
		grpcshop.GetBusinessUnpaidBranchRents,
	)
)
