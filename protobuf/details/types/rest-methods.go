package types

type (
	// RESTMethod is the type of REST request method
	RESTMethod int
)

// GetRESTMethod returns the RESTMethod from the string
func GetRESTMethod(method string) RESTMethod {
	switch method {
	case "":
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "PATCH":
		return PATCH
	case "DELETE":
		return DELETE
	default:
		return INVALID
	}
	return INVALID
}

// RESTMethod values
const (
	GET RESTMethod = iota
	POST
	PUT
	PATCH
	DELETE
	INVALID
)
