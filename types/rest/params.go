package rest

type (
	// Param is the type of REST request parameter
	Param struct {
		Name string
	}
)

// NewParam creates a new Param
func NewParam(name string) Param {
	return Param{name}
}

// String returns the string representation of the Param
func (r Param) String() string {
	return r.Name
}

// Param values
var (
	Email           = NewParam("email")
	Username        = NewParam("username")
	Token           = NewParam("token")
	Id              = NewParam("id")
	PermissionId    = NewParam("permission-id")
	RoleId          = NewParam("role-id")
	UserId          = NewParam("user-id")
	JwtId           = NewParam("jwt-id")
	AccountId       = NewParam("account-id")
	OrderId         = NewParam("order-id")
	CartId          = NewParam("cart-id")
	BranchRentId    = NewParam("branch-rent-id")
	BranchId        = NewParam("branch-id")
	StoreId         = NewParam("store-id")
	BusinessId      = NewParam("business-id")
	CategoryId      = NewParam("category-id")
	ProductId       = NewParam("product-id")
	RevisionId      = NewParam("revision-id")
	BranchProductId = NewParam("branch-product-id")
)
