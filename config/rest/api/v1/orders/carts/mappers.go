package carts

import (
	grpcorder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/order"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Order service carts endpoints mapping
var (
	GetCartsMapper     = typesrest.NewMapper(Relative, grpcorder.GetCarts)
	GetCartMapper      = typesrest.NewMapper(ByCartId, grpcorder.GetCart)
	GetCartTotalMapper = typesrest.NewMapper(TotalByCartId, grpcorder.GetCartTotal)
)
