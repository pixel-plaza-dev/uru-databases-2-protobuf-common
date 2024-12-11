package order

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[grpc.Method]grpc.Interception{
	AddProductToCart:      grpc.AccessToken,
	RemoveProductFromCart: grpc.AccessToken,
	GetCart:               grpc.AccessToken,
	GetCurrentCart:        grpc.AccessToken,
	GetCarts:              grpc.AccessToken,
	GetCartTotal:          grpc.AccessToken,
	PlaceOrder:            grpc.AccessToken,
	GetOrders:             grpc.AccessToken,
	GetOrder:              grpc.AccessToken,
}
