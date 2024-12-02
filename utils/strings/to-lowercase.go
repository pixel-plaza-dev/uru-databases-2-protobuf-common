package strings

import (
	"strings"
	"unicode"
)

// ToSnakeCase convert string to lower case
func ToSnakeCase(s string) string {
	var result string

	for _, char := range s {
		if unicode.IsUpper(char) {
			if len(result) > 0 {
				result += "_"
			}
			result += strings.ToLower(string(char))
		} else {
			result += string(char)
		}
	}

	return result
}
