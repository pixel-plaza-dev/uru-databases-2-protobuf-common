package business

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[grpc.Method]grpc.Interception{
	AddBusiness:                  grpc.AccessToken,
	GetBusiness:                  grpc.None,
	UpdateBusiness:               grpc.AccessToken,
	SetBusinessProfilePicture:    grpc.AccessToken,
	SuspendBusiness:              grpc.AccessToken,
	ActivateBusiness:             grpc.AccessToken,
	DeleteBusiness:               grpc.AccessToken,
	AddBusinessOwner:             grpc.AccessToken,
	RemoveBusinessOwner:          grpc.AccessToken,
	GetBusinessOwners:            grpc.AccessToken,
	AddBusinessClient:            grpc.AccessToken,
	IsBusinessClient:             grpc.AccessToken,
	AddMarketCategory:            grpc.AccessToken,
	GetMarketCategory:            grpc.None,
	UpdateMarketCategory:         grpc.AccessToken,
	AddBusinessMarketCategory:    grpc.AccessToken,
	GetBusinessMarketCategories:  grpc.None,
	AddStore:                     grpc.AccessToken,
	GetStore:                     grpc.None,
	DeleteStore:                  grpc.AccessToken,
	GetUnoccupiedStores:          grpc.None,
	AddBranch:                    grpc.AccessToken,
	GetBranch:                    grpc.None,
	GetBusinessBranches:          grpc.None,
	UpdateBranch:                 grpc.AccessToken,
	SuspendBranch:                grpc.AccessToken,
	ActivateBranch:               grpc.AccessToken,
	DeleteBranch:                 grpc.AccessToken,
	AddStoreBranchRent:           grpc.AccessToken,
	GetStoreBranchRents:          grpc.AccessToken,
	UpdateStoreBranchRent:        grpc.AccessToken,
	GetUnpaidBranchRents:         grpc.AccessToken,
	GetBusinessUnpaidBranchRents: grpc.AccessToken,
	AddProductCategory:           grpc.AccessToken,
	GetProductCategory:           grpc.None,
	UpdateProductCategory:        grpc.AccessToken,
	AddProduct:                   grpc.AccessToken,
	GetProduct:                   grpc.None,
	UpdateProduct:                grpc.AccessToken,
	SearchProducts:               grpc.None,
	SuspendProduct:               grpc.AccessToken,
	ActivateProduct:              grpc.AccessToken,
	AddBusinessProduct:           grpc.AccessToken,
	GetBusinessProduct:           grpc.None,
	UpdateBusinessProduct:        grpc.AccessToken,
	SearchBusinessProducts:       grpc.None,
	SuspendBusinessProduct:       grpc.AccessToken,
	ActivateBusinessProduct:      grpc.AccessToken,
	AddBranchProduct:             grpc.AccessToken,
	GetBranchProduct:             grpc.None,
	SearchBranchProducts:         grpc.None,
	UpdateBranchProduct:          grpc.AccessToken,
}
