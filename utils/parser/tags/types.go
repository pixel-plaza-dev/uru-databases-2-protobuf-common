package tags

type (
	// StructFields map the struct name to the fields name
	StructFields map[string][]string

	// GoFileStructFields map the Go file path to the struct fields
	GoFileStructFields map[string]*StructFields
)
