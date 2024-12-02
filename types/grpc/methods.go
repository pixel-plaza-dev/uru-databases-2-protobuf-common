package grpc

type (
	// Method is the gRPC method Name
	Method struct {
		Name string
	}
)

// NewMethod creates a new Method
func NewMethod(name string) Method {
	return Method{name}
}

// String returns the string representation of the Method
func (g Method) String() string {
	return g.Name
}

// RequestString returns the string representation of the Method for the request
func (g Method) RequestString() string {
	return g.Name + "Request"
}

// ResponseString returns the string representation of the Method for the response
func (g Method) ResponseString() string {
	return g.Name + "Response"
}
