package rest

type (
	// Endpoint is the REST endpoint
	Endpoint struct {
		Name   string
		Params []Param
	}
)

// NewEndpoint creates a new Endpoint
func NewEndpoint(name string, params ...Param) *Endpoint {
	return &Endpoint{Name: name, Params: params}
}

// String returns the string representation of the Endpoint
func (r Endpoint) String() string {
	formattedParams := ""
	for _, param := range r.Params {
		formattedParams += "/:" + param.String()
	}
	return r.Name + formattedParams
}
