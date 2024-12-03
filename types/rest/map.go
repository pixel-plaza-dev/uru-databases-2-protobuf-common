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
	var firstPartUri, secondPartUri string

	// Get the second part index of the relative URI
	secondPartUriIndex := strings.Index(relativeURI[1:], "/")
	if secondPartUriIndex != -1 {
		secondPartUriIndex++
	}

	// Check if the URI is an endpoint or a URI with a parameter
	isEndpoint := false
	if secondPartUriIndex == -1 {
		firstPartUri = "/"
		isEndpoint = true
	} else if relativeURI[secondPartUriIndex+1] == ':' {
		firstPartUri = relativeURI[secondPartUriIndex:]
		isEndpoint = true
	} else {
		firstPartUri = relativeURI[:secondPartUriIndex]
		secondPartUri = relativeURI[secondPartUriIndex:]
	}

	// Check if debug is enabled
	if debug {
		// Print if the URI is an endpoint, first half URI and second half URI
		fmt.Printf("Is endpoint: %t\n", isEndpoint)
		fmt.Printf("First half URI: %s\n", firstPartUri)
		fmt.Printf("Second half URI: %s\n", secondPartUri)
	}

	// Check if the first part of the URI is in the interceptions map
	if m.Interceptions != nil && isEndpoint {
		if methodMap, ok := (*m.Interceptions)[firstPartUri]; ok {
			if grpcMethod, ok := methodMap[restMethod]; ok {
				return &grpcMethod, nil
			}
		}
	}

	// Check if the URI is an endpoint
	if isEndpoint {
		return nil, fmt.Errorf(MissingEndpointInterception, relativeURI, restMethod)
	}

	if m.ChildrenMaps != nil {
		// Check if the first part of the URI is in the children maps
		// If it is, traverse the child map
		if childMap, ok := (*m.ChildrenMaps)[firstPartUri]; ok {
			return childMap.Traverse(debug, secondPartUri, restMethod)
		}
		return nil, fmt.Errorf(FailedToTraverse, relativeURI, restMethod)
	}

	return nil, fmt.Errorf(MissingChildrenMap, relativeURI, restMethod)
}
