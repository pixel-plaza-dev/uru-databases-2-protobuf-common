package shop

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Business service gRPC methods
var (
	AddBusiness                  = grpc.NewMethod("AddBusiness")
	GetBusiness                  = grpc.NewMethod("GetBusiness")
	UpdateBusiness               = grpc.NewMethod("UpdateBusiness")
	SetBusinessProfilePicture    = grpc.NewMethod("SetBusinessProfilePicture")
	SuspendBusiness              = grpc.NewMethod("SuspendBusiness")
	ActivateBusiness             = grpc.NewMethod("ActivateBusiness")
	DeleteBusiness               = grpc.NewMethod("DeleteBusiness")
	AddBusinessOwner             = grpc.NewMethod("AddBusinessOwner")
	RemoveBusinessOwner          = grpc.NewMethod("RemoveBusinessOwner")
	GetBusinessOwners            = grpc.NewMethod("GetBusinessOwners")
	AddBusinessClient            = grpc.NewMethod("AddBusinessClient")
	IsBusinessClient             = grpc.NewMethod("IsBusinessClient")
	AddMarketCategory            = grpc.NewMethod("AddMarketCategory")
	GetMarketCategory            = grpc.NewMethod("GetMarketCategory")
	UpdateMarketCategory         = grpc.NewMethod("UpdateMarketCategory")
	AddBusinessMarketCategory    = grpc.NewMethod("AddBusinessMarketCategory")
	GetBusinessMarketCategories  = grpc.NewMethod("GetBusinessMarketCategories")
	AddStore                     = grpc.NewMethod("AddStore")
	GetStore                     = grpc.NewMethod("GetStore")
	DeleteStore                  = grpc.NewMethod("DeleteStore")
	GetUnoccupiedStores          = grpc.NewMethod("GetUnoccupiedStores")
	AddBranch                    = grpc.NewMethod("AddBranch")
	GetBranch                    = grpc.NewMethod("GetBranch")
	GetBusinessBranches          = grpc.NewMethod("GetBusinessBranches")
	UpdateBranch                 = grpc.NewMethod("UpdateBranch")
	SuspendBranch                = grpc.NewMethod("SuspendBranch")
	ActivateBranch               = grpc.NewMethod("ActivateBranch")
	DeleteBranch                 = grpc.NewMethod("DeleteBranch")
	AddBranchRent                = grpc.NewMethod("AddBranchRent")
	GetBranchRents               = grpc.NewMethod("GetBranchRents")
	UpdateBranchRent             = grpc.NewMethod("UpdateBranchRent")
	GetUnpaidBranchRents         = grpc.NewMethod("GetUnpaidBranchRents")
	GetBusinessUnpaidBranchRents = grpc.NewMethod("GetBusinessUnpaidBranchRents")
	AddProductCategory           = grpc.NewMethod("AddProductCategory")
	GetProductCategory           = grpc.NewMethod("GetProductCategory")
	UpdateProductCategory        = grpc.NewMethod("UpdateProductCategory")
	AddProduct                   = grpc.NewMethod("AddProduct")
	GetProduct                   = grpc.NewMethod("GetProduct")
	UpdateProduct                = grpc.NewMethod("UpdateProduct")
	SearchProducts               = grpc.NewMethod("SearchProducts")
	SuspendProduct               = grpc.NewMethod("SuspendProduct")
	ActivateProduct              = grpc.NewMethod("ActivateProduct")
	AddBusinessProduct           = grpc.NewMethod("AddBusinessProduct")
	GetBusinessProduct           = grpc.NewMethod("GetBusinessProduct")
	UpdateBusinessProduct        = grpc.NewMethod("UpdateBusinessProduct")
	SearchBusinessProducts       = grpc.NewMethod("SearchBusinessProducts")
	SuspendBusinessProduct       = grpc.NewMethod("SuspendBusinessProduct")
	ActivateBusinessProduct      = grpc.NewMethod("ActivateBusinessProduct")
	AddBranchProduct             = grpc.NewMethod("AddBranchProduct")
	GetBranchProduct             = grpc.NewMethod("GetBranchProduct")
	SearchBranchProducts         = grpc.NewMethod("SearchBranchProducts")
	UpdateBranchProduct          = grpc.NewMethod("UpdateBranchProduct")
)