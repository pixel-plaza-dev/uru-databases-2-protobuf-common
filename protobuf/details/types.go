package details

type (
	// Interception is the type of interception
	Interception int

	// GRPCMethod is the gRPC method Name
	GRPCMethod struct {
		Name string
	}

	// RESTEndpoint is the REST endpoint
	RESTEndpoint struct {
		Name   string
		Params []RESTParam
	}

	// RESTMethod is the type of REST request method
	RESTMethod struct {
		Name string
	}

	// RESTParam is the type of REST request parameter
	RESTParam struct {
		Name string
	}
)

// NewRESTEndpoint creates a new RESTEndpoint
func NewRESTEndpoint(name string, params ...RESTParam) RESTEndpoint {
	return RESTEndpoint{Name: name, Params: params}
}

// String returns the string representation of the RESTEndpoint
func (r RESTEndpoint) String() string {
	formattedParams := ""
	for _, param := range r.Params {
		formattedParams += "/:" + param.String()
	}
	return r.Name + formattedParams
}

// NewGRPCMethod creates a new GRPCMethod
func NewGRPCMethod(name string) GRPCMethod {
	return GRPCMethod{name}
}

// String returns the string representation of the GRPCMethod
func (g GRPCMethod) String() string {
	return g.Name
}

// NewRESTMethod creates a new RESTMethod
func NewRESTMethod(name string) RESTMethod {
	return RESTMethod{name}
}

// String returns the string representation of the RESTMethod
func (r RESTMethod) String() string {
	return r.Name
}

// NewRESTParam creates a new RESTParam
func NewRESTParam(name string) RESTParam {
	return RESTParam{name}
}

// String returns the string representation of the RESTParam
func (r RESTParam) String() string {
	return r.Name
}

// Interception values
const (
	RefreshToken Interception = iota
	AccessToken
	None
)

// RESTMethod values
var (
	GET    = NewRESTMethod("GET")
	POST   = NewRESTMethod("POST")
	PUT    = NewRESTMethod("PUT")
	DELETE = NewRESTMethod("DELETE")
	PATCH  = NewRESTMethod("PATCH")
)

// RESTParam values
var (
	Email  = NewRESTParam("email")
	Token  = NewRESTParam("token")
	Id     = NewRESTParam("id")
	UserId = NewRESTParam("user_id")
)
