package main

import (
	"fmt"
	configgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	configgrpcorder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/order"
	configgrpcpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/payment"
	configgrpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
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
	shopPath := path.GetProtoGoPath("pixel_plaza", "shop")

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
			configgrpcpayment.ActivatePaymentAccount.RequestString(): []string{
				"PaymentAccountId",
			},
			configgrpcpayment.SuspendPaymentAccount.RequestString(): []string{
				"PaymentAccountId",
			},
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
		shopPath: {
			configgrpcshop.UpdateAdminRevision.RequestString(): []string{
				"AdminRevisionId",
			},
			configgrpcshop.CloseAdminRevision.RequestString(): []string{
				"AdminRevisionId",
			},
			configgrpcshop.GetStore.RequestString(): []string{
				"StoreId",
			},
			configgrpcshop.DeleteStore.RequestString(): []string{
				"StoreId",
			},
			configgrpcshop.AddBranchProduct.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.GetBranchProduct.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.UpdateBranchProduct.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.GetBusinessMarketCategories.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.GetProductCategory.RequestString(): []string{
				"ProductCategoryId",
			},
			configgrpcshop.UpdateProductCategory.RequestString(): []string{
				"ProductCategoryId",
			},
			configgrpcshop.GetMarketCategory.RequestString(): []string{
				"MarketCategoryId",
			},
			configgrpcshop.UpdateMarketCategory.RequestString(): []string{
				"MarketCategoryId",
			},
			configgrpcshop.GetProduct.RequestString(): []string{
				"ProductId",
			},
			configgrpcshop.UpdateProduct.RequestString(): []string{
				"ProductId",
			},
			configgrpcshop.OpenAdminRevisionToProduct.RequestString(): []string{
				"ProductId",
			},
			configgrpcshop.GetBusinessProduct.RequestString(): []string{
				"ProductId",
			},
			configgrpcshop.UpdateBusinessProduct.RequestString(): []string{
				"ProductId",
			},
			configgrpcshop.OpenAdminRevisionToBusinessProduct.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.GetBranchRents.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.UpdateBranchRent.RequestString(): []string{
				"BranchRentId",
			},
			configgrpcshop.GetUnpaidBranchRents.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.GetBusinessUnpaidBranchRents.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.GetBranch.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.GetBusinessBranches.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.UpdateBranch.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.OpenBranch.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.CloseTemporarilyBranch.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.OpenAdminRevisionToBranch.RequestString(): []string{
				"BranchId",
			},
			configgrpcshop.AddBusinessOwner.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.RemoveBusinessOwner.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.GetBusinessOwners.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.GetBusiness.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.UpdateBusiness.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.SetBusinessProfilePicture.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.OpenAdminRevisionToBusiness.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.DeleteBusiness.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.AddBusinessClient.RequestString(): []string{
				"BusinessId",
			},
			configgrpcshop.IsBusinessClient.RequestString(): []string{
				"BusinessId",
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
