package api

type (
	// RESTParam is the type of REST request parameter
	RESTParam struct {
		Name string
	}
)

// NewRESTParam creates a new RESTParam
func NewRESTParam(name string) RESTParam {
	return RESTParam{name}
}

// String returns the string representation of the RESTParam
func (r RESTParam) String() string {
	return r.Name
}

// RESTParam values
var (
	Email    = NewRESTParam("email")
	Username = NewRESTParam("username")
	Token    = NewRESTParam("token")
	Id       = NewRESTParam("id")
	UserId   = NewRESTParam("user_id")
	JwtId    = NewRESTParam("jwt_id")
)
