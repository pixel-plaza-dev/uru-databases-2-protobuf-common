package types

type (
	// RESTEndpoint is the REST endpoint
	RESTEndpoint struct {
		Name   string
		Params []RESTParam
	}
)

// NewRESTEndpoint creates a new RESTEndpoint
func NewRESTEndpoint(name string, params ...RESTParam) *RESTEndpoint {
	return &RESTEndpoint{Name: name, Params: params}
}

// String returns the string representation of the RESTEndpoint
func (r RESTEndpoint) String() string {
	formattedParams := ""
	for _, param := range r.Params {
		formattedParams += "/:" + param.String()
	}
	return r.Name + formattedParams
}
