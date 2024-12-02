package rest

type (
	// Method is the type of REST request method
	Method int
)

// GetMethod returns the Method from the string
func GetMethod(method string) Method {
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

// Method values
const (
	GET Method = iota
	POST
	PUT
	PATCH
	DELETE
	INVALID
)
