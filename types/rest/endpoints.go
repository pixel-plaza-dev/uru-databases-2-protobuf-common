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

// NewRelativeEndpoint creates a new relative Endpoint
func NewRelativeEndpoint(params ...Param) *Endpoint {
	return &Endpoint{Name: "", Params: params}
}

// NewBaseEndpoint creates a new base Endpoint
func NewBaseEndpoint(name string) *Endpoint {
	return &Endpoint{Name: name}
}

// String returns the string representation of the Endpoint
func (r Endpoint) String() string {
	formattedParams := ""
	for _, param := range r.Params {
		formattedParams += "/:" + param.String()
	}
	return "/" + r.Name + formattedParams
}

// Path returns the path of the Endpoint
func (r Endpoint) Path() string {
	return r.String()
}
