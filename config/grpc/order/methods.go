package order

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Order service gRPC methods
var (
	AddProductToCart      = grpc.NewMethod("AddProductToCart")
	RemoveProductFromCart = grpc.NewMethod("RemoveProductFromCart")
	GetCart               = grpc.NewMethod("GetCart")
	GetCurrentCart        = grpc.NewMethod("GetCurrentCart")
	GetCarts              = grpc.NewMethod("GetCarts")
	GetCartTotal          = grpc.NewMethod("GetCartTotal")
	PlaceOrder            = grpc.NewMethod("PlaceOrder")
	GetOrders             = grpc.NewMethod("GetOrders")
	GetOrder              = grpc.NewMethod("GetOrder")
)
