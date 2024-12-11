package main

import (
	"fmt"
	configgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	configgrpcorder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/order"
	configgrpcpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/payment"
	configgrpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	utilsparsertags "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/utils/parser/tags"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/utils/path"
)

func main() {
	// Get the path of the compiled protobuf file in Go
	userPath := path.GetProtoGoPath("pixel_plaza", "user")
	authPath := path.GetProtoGoPath("pixel_plaza", "auth")
	orderPath := path.GetProtoGoPath("pixel_plaza", "order")
	paymentPath := path.GetProtoGoPath("pixel_plaza", "payment")
	businessPath := path.GetProtoGoPath("pixel_plaza", "business")

	// Create the GoFileStructFields map
	goFileStructFields := utilsparsertags.GoFileStructFields{
		userPath: {
			configgrpcuser.DeleteEmail.RequestString(): []string{
				"Email",
			},
			configgrpcuser.VerifyEmail.RequestString(): []string{
				"Token",
			},
			configgrpcuser.VerifyPhoneNumber.RequestString(): []string{
				"Token",
			},
			configgrpcuser.GetProfile.RequestString(): []string{
				"Username",
			},
			configgrpcuser.GetUsernameByUserId.RequestString(): []string{
				"UserId",
			},
			configgrpcuser.GetUserIdByUsername.RequestString(): []string{
				"Username",
			},
			configgrpcuser.ResetPassword.RequestString(): []string{
				"Token",
			},
		},
		authPath: {
			configgrpcauth.IsAccessTokenValid.RequestString(): []string{
				"JwtId",
			},
			configgrpcauth.GetPermission.RequestString(): []string{
				"PermissionId",
			},
			configgrpcauth.RevokePermission.RequestString(): []string{
				"PermissionId",
			},
			configgrpcauth.IsRefreshTokenValid.RequestString(): []string{
				"JwtId",
			},
			configgrpcauth.GetRefreshTokenInformation.RequestString(): []string{
				"JwtId",
			},
			configgrpcauth.RevokeRefreshToken.RequestString(): []string{
				"JwtId",
			},
			configgrpcauth.RevokeRolePermission.RequestString(): []string{
				"RoleId",
			},
			configgrpcauth.AddRolePermission.RequestString(): []string{
				"RoleId",
			},
			configgrpcauth.GetRolePermissions.RequestString(): []string{
				"RoleId",
			},
			configgrpcauth.RevokeRole.RequestString(): []string{
				"RoleId",
			},
			configgrpcauth.GetUserRoles.RequestString(): []string{
				"UserId",
			},
			configgrpcauth.AddUserRole.RequestString(): []string{
				"UserId",
			},
			configgrpcauth.RevokeUserRole.RequestString(): []string{
				"UserId",
			},
		},
		orderPath: {
			configgrpcorder.GetCart.RequestString(): []string{
				"CartId",
			},
			configgrpcorder.GetCartTotal.RequestString(): []string{
				"CartId",
			},
			configgrpcorder.GetOrder.RequestString(): []string{
				"OrderId",
			},
		},
		paymentPath: {
			configgrpcpayment.AddBranchRentPayment.RequestString(): []string{
				"BranchRentId",
			},
			configgrpcpayment.GetBranchRentPayments.RequestString(): []string{
				"BranchRentId",
			},
			configgrpcpayment.PayForBranchRent.RequestString(): []string{
				"BranchRentId",
			},
			configgrpcpayment.AddOrderPayment.RequestString(): []string{
				"OrderId",
			},
			configgrpcpayment.GetOrderPayments.RequestString(): []string{
				"OrderId",
			},
			configgrpcpayment.PayForOrder.RequestString(): []string{
				"OrderId",
			},
		},
	}

	// Hide the JSON tags in the specified files
	err := utilsparsertags.HideFilesJSONTags(&goFileStructFields)
	if err != nil {
		panic(err)
	}
	fmt.Println("Structs JSON tags hidden successfully")
}
