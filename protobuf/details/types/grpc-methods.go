package types

type (
	// GRPCMethod is the gRPC method Name
	GRPCMethod struct {
		Name string
	}
)

// NewGRPCMethod creates a new GRPCMethod
func NewGRPCMethod(name string) GRPCMethod {
	return GRPCMethod{name}
}

// String returns the string representation of the GRPCMethod
func (g GRPCMethod) String() string {
	return g.Name
}
