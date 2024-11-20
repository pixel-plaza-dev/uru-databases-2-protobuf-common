package details

type (
	// Interception is the type of interception
	Interception int

	// GRPCMethod is the gRPC method name
	GRPCMethod string

	// RESTEndpoint is the REST endpoint
	RESTEndpoint string

	// RESTMethod is the type of REST request method
	RESTMethod string
)

// Interception values
const (
	RefreshToken Interception = iota
	AccessToken
	None
)

// RESTMethod values
const (
	GET    RESTMethod = "GET"
	POST   RESTMethod = "POST"
	PUT    RESTMethod = "PUT"
	DELETE RESTMethod = "DELETE"
	PATCH  RESTMethod = "PATCH"
)
