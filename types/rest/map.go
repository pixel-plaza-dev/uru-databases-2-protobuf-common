package rest

import (
	"fmt"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"strings"
)

type (
	// Mapper is the interface for the REST map
	Mapper interface {
		Traverse(relativeURI string, restMethod Method) (*grpc.Method, error)
	}

	// Map struct for the REST endpoints
	Map struct {
		Debug         bool
		Interceptions *map[string]map[Method]grpc.Method
		ChildrenMaps  *map[string]*Map
	}
)

// NewMap creates a new REST map
func NewMap(debug bool, interceptions *map[string]map[Method]grpc.Method, childrenMaps *map[string]*Map) *Map {
	return &Map{
		Debug:         debug,
		Interceptions: interceptions,
		ChildrenMaps:  childrenMaps,
	}
}

// Traverse traverses the REST map
func (m *Map) Traverse(relativeURI string, restMethod Method) (*grpc.Method, error) {
	var firstHalfUri, secondHalfUri string

	// Get the first half of the relative URI
	firstHalfUriIndex := -1
	if len(relativeURI) > 1 {
		firstHalfUriIndex = strings.Index(relativeURI[1:], "/") + 1
	}

	// Check if the URI is the endpoint or a URI with a parameter
	isEndpoint := false
	if firstHalfUriIndex == -1 || len(relativeURI) == firstHalfUriIndex+1 || relativeURI[firstHalfUriIndex+1] == ':' {
		firstHalfUri = relativeURI
		isEndpoint = true
	} else {
		firstHalfUri = relativeURI[:firstHalfUriIndex]
		secondHalfUri = relativeURI[firstHalfUriIndex:]
	}

	if m.Debug {
		// Print if the URI is an endpoint. first half URI and second half URI if debug is enabled
		fmt.Printf("Is endpoint: %t\n", isEndpoint)
		fmt.Printf("First half URI: %s\n", firstHalfUri)
		fmt.Printf("Second half URI: %s\n", secondHalfUri)
	}

	// Check if the first half URI is in the interceptions map
	if m.Interceptions != nil {
		if methodMap, ok := (*m.Interceptions)[firstHalfUri]; ok {
			if grpcMethod, ok := methodMap[restMethod]; ok {
				return &grpcMethod, nil
			}
		}

		// Check if the URI is an endpoint
		if isEndpoint {
			return nil, fmt.Errorf(MissingEndpointInterception, relativeURI)
		}
	} else if isEndpoint {
		return nil, fmt.Errorf(MissingInterceptions, relativeURI)
	}

	// Check if the first half URI is in the children maps
	if m.ChildrenMaps != nil {
		// Check if the first half URI is in the children maps. If it is, traverse the child map
		if childMap, ok := (*m.ChildrenMaps)[firstHalfUri]; ok {
			return childMap.Traverse(secondHalfUri, restMethod)
		}
	} else {
		return nil, fmt.Errorf(MissingChildrenMap, relativeURI)
	}
	return nil, fmt.Errorf(FailedToTraverse, relativeURI)
}
