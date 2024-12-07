package tags

import (
	"fmt"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/utils/parser"
	"go/ast"
	"go/token"
	"regexp"
)

type (
	// FieldJSONTag map the field name to the new JSON tag
	FieldJSONTag map[string]string

	// StructJSONTag map the struct name to the field JSON tag
	StructJSONTag map[string]FieldJSONTag
)

// OverwriteJSONTags overwrite the given structs fields JSON tags from the given Go file
func OverwriteJSONTags(filePath string, structJSONTag *StructJSONTag) error {
	// Check if the struct JSON tag is nil
	if structJSONTag == nil {
		return NilStructJSONTagError
	}

	// Compile the regex pattern to match any content inside json tag
	regex, err := regexp.Compile(`json:"[^"]*"`)
	if err != nil {
		return fmt.Errorf("failed to compile regex: %w", err)
	}

	// Parse the Go file
	fileSet, node, err := parser.ParseGoFile(filePath)
	if err != nil {
		return err
	}

	// Traverse the AST to find the struct and field
	err = parser.TraverseAST(
		node, func(n ast.Node) bool {
			// Check if the node is a type spec
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}

			// Check if the node is a struct
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				return true
			}

			// Check if the struct name is in the map
			fieldJSONTag, ok := (*structJSONTag)[ts.Name.Name]
			if !ok {
				return true
			}

			// Iterate the struct fields
			for _, field := range st.Fields.List {
				for _, name := range field.Names {
					// Check if the field name is in the map and get the new JSON tag
					newJSONTag, ok := fieldJSONTag[name.Name]
					if !ok {
						continue
					}

					// Print the struct and field name
					// fmt.Printf("Struct: %s, Field: %s\n", ts.Name.Name, name.Name)

					// Modify the JSON tag
					if field.Tag != nil {
						tag := field.Tag.Value

						// Replace the matched content
						newTag := regex.ReplaceAllString(tag, fmt.Sprintf(`json:"%s"`, newJSONTag))

						field.Tag.Value = newTag
					} else {
						field.Tag = &ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf("`json:\"%s\"`", newJSONTag),
						}
					}

					// Remove the field from the map
					delete(fieldJSONTag, name.Name)
				}
			}

			// Check if the struct has fields to update
			numFields := len(fieldJSONTag)
			if numFields == 0 {
				delete(*structJSONTag, ts.Name.Name)
				return false
			}
			return true
		},
	)
	if err != nil {
		return err
	}

	// Check if all structs have been updated
	if len(*structJSONTag) > 0 {
		// Print the structs fields that haven't been updated
		fmt.Println("The following structs fields haven't been updated:")
		for structName := range *structJSONTag {
			for fieldName := range (*structJSONTag)[structName] {
				fmt.Printf("Struct: %s, Field: %s\n", structName, fieldName)
			}
		}

		return fmt.Errorf("failed to update all structs")
	}

	// Write the modified AST back to the file
	if err := parser.WriteGoFile(filePath, fileSet, node); err != nil {
		return err
	}

	return nil
}
