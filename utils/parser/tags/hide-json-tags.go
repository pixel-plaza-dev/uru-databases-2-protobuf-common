package tags

// HideFileJSONTags hides the JSON tags in the specified file
func HideFileJSONTags(filePath string, structFields *StructFields) error {
	// Check if the StructFields is nil
	if structFields == nil {
		return NilStructFieldsError
	}

	// Generate the StructJSONTag
	structJSONTag := StructJSONTag{}

	// Loop through the struct fields
	for structName, fields := range *structFields {
		fieldJSONTag := FieldJSONTag{}
		for _, field := range fields {
			fieldJSONTag[field] = "-"
		}
		structJSONTag[structName] = fieldJSONTag
	}
	// Overwrite the JSON tags
	return OverwriteJSONTags(filePath, &structJSONTag)
}

// HideFilesJSONTags hides the JSON tags in the specified files
func HideFilesJSONTags(goFileStructFields *GoFileStructFields) error {
	// Check if the GoFileStructFields is nil
	if goFileStructFields == nil {
		return NilGoStructFieldsError
	}

	// Loop through the file paths
	for filePath, structFields := range *goFileStructFields {
		// Hide the JSON tags in the file
		if err := HideFileJSONTags(filePath, structFields); err != nil {
			return err
		}
	}
	return nil
}
