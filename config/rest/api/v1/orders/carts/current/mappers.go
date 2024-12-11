package carts

import (
	grpcorder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/order"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Order service current cart endpoints mapping
var (
	GetCurrentCart              = typesrest.NewMapper(Relative, grpcorder.GetCurrentCart)
	AddProductToCartMapper      = typesrest.NewMapper(Relative, grpcorder.AddProductToCart)
	RemoveProductFromCartMapper = typesrest.NewMapper(Relative, grpcorder.RemoveProductFromCart)
	PlaceOrderMapper            = typesrest.NewMapper(Checkout, grpcorder.PlaceOrder)
)
