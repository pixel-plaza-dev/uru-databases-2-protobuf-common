package types

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
	Email  = NewRESTParam("email")
	Token  = NewRESTParam("token")
	Id     = NewRESTParam("id")
	UserId = NewRESTParam("user_id")
)
