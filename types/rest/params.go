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
	Email        = NewParam("email")
	Username     = NewParam("username")
	Token        = NewParam("token")
	Id           = NewParam("id")
	PermissionId = NewParam("permission-id")
	RoleId       = NewParam("role-id")
	UserId       = NewParam("user-id")
	JwtId        = NewParam("jwt-id")
)
