package path

import (
	"strings"
)

// GetProtoGoPath returns path of the compiled protobuf file in Go
func GetProtoGoPath(protoPackage, protoService string) string {
	return strings.Join([]string{"protobuf/compiled", protoPackage, protoService, protoService + ".pb.go"}, "/")
}
