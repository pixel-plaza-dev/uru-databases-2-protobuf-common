package rest

import (
	"fmt"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"strings"
)

type (
	// Mapper is the interface for the REST map
	Mapper interface {
		Traverse(debug bool, relativeURI string, restMethod Method) (*grpc.Method, error)
	}

	// Map struct for the REST endpoints
	Map struct {
		Interceptions *map[string]map[Method]grpc.Method
		ChildrenMaps  *map[string]*Map
	}
)

// NewMap creates a new REST map
func NewMap(interceptions *map[string]map[Method]grpc.Method, childrenMaps *map[string]*Map) *Map {
	return &Map{
		Interceptions: interceptions,
		ChildrenMaps:  childrenMaps,
	}
}

// Traverse traverses the REST map
func (m *Map) Traverse(debug bool, relativeURI string, restMethod Method) (*grpc.Method, error) {
	// Check if the whole relative URI is in the interceptions map
	if m.Interceptions != nil {
		if methodMap, ok := (*m.Interceptions)[relativeURI]; ok {
			if grpcMethod, ok := methodMap[restMethod]; ok {
				return &grpcMethod, nil
			}
		}
	}

	// Get the index of the possible URI
	var possibleBaseUri, possibleUri string
	possibleUriIndex := strings.Index(relativeURI[1:], "/")
	if possibleUriIndex != -1 {
		possibleUriIndex++
	}

	// Check if the URI is an endpoint or a URI with a parameter
	if possibleUriIndex == -1 {
		possibleUri = "/"
	} else {
		possibleBaseUri = relativeURI[:possibleUriIndex]
		possibleUri = relativeURI[possibleUriIndex:]
	}

	// Check if debug is enabled
	if debug {
		// Print the possible base URI and URI
		fmt.Printf("Possible base URI: %s\n", possibleBaseUri)
		fmt.Printf("Possible URI: %s\n", possibleUri)
	}

	if m.Interceptions != nil {
		// Check if the possible URI is in the interceptions map
		if methodMap, ok := (*m.Interceptions)[possibleUri]; ok {
			if grpcMethod, ok := methodMap[restMethod]; ok {
				return &grpcMethod, nil
			}
		}
	}

	if m.ChildrenMaps != nil {
		// Check if the possible base URI is in the children maps
		// If it is, traverse the child map
		if childMap, ok := (*m.ChildrenMaps)[possibleBaseUri]; ok {
			return childMap.Traverse(debug, possibleUri, restMethod)
		}
	}

	return nil, fmt.Errorf(MissingEndpointInterception, relativeURI, restMethod)
}
