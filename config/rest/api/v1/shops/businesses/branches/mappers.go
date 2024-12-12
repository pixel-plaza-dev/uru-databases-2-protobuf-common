package branches

import (
	grpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Shop service branches endpoints mapping
var (
	AddBranchMapper              = typesrest.NewMapper(Relative, grpcshop.AddBranch)
	GetBranchMapper              = typesrest.NewMapper(ByBranchId, grpcshop.GetBranch)
	GetBusinessBranchesMapper    = typesrest.NewMapper(ByBusinessId, grpcshop.GetBusinessBranches)
	UpdateBranchMapper           = typesrest.NewMapper(ByBranchId, grpcshop.UpdateBranch)
	CloseTemporarilyBranchMapper = typesrest.NewMapper(CloseTemporarilyByBranchId, grpcshop.CloseTemporarilyBranch)
	OpenBranchMapper             = typesrest.NewMapper(OpenByBranchId, grpcshop.OpenBranch)
	DeleteBranchMapper           = typesrest.NewMapper(ByBranchId, grpcshop.DeleteBranch)
)
