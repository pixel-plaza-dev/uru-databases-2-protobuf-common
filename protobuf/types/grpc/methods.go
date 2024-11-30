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
