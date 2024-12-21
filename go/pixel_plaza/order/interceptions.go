package order

import (
	compiled "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/order"
	"github.com/ralvarezdev/go-jwt/token/interception"
)

// Interceptions is the list of gRPC methods to intercept
var Interceptions = map[string]interception.Interception{
	compiled.Order_AddProductToCart_FullMethodName:      interception.AccessToken,
	compiled.Order_RemoveProductFromCart_FullMethodName: interception.AccessToken,
	compiled.Order_GetCart_FullMethodName:               interception.AccessToken,
	compiled.Order_GetCurrentCart_FullMethodName:        interception.AccessToken,
	compiled.Order_GetCarts_FullMethodName:              interception.AccessToken,
	compiled.Order_GetCartTotal_FullMethodName:          interception.AccessToken,
	compiled.Order_PlaceOrder_FullMethodName:            interception.AccessToken,
	compiled.Order_GetOrders_FullMethodName:             interception.AccessToken,
	compiled.Order_GetOrder_FullMethodName:              interception.AccessToken,
}
