package rest

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"strings"
)

type (
	// Mapper is the interface for the REST map
	Mapper interface {
		Traverse(relativeURI string, restMethod rest.Method) (*grpc.Method, error)
	}

	// Map struct for the REST endpoints
	Map struct {
		Interceptions *map[string]map[rest.Method]grpc.Method
		ChildrenMaps  *map[string]*Map
	}
)

// NewMap creates a new REST map
func NewMap(interceptions *map[string]map[rest.Method]grpc.Method, childrenMaps *map[string]*Map) *Map {
	return &Map{
		Interceptions: interceptions,
		ChildrenMaps:  childrenMaps,
	}
}

// Traverse traverses the REST map
func (m *Map) Traverse(relativeURI string, restMethod rest.Method) (*grpc.Method, error) {
	var firstHalfUri, secondHalfUri string

	// Get the first half of the relative URI
	firstHalfUriIndex := -1
	if len(relativeURI) > 1 {
		firstHalfUriIndex = strings.Index(relativeURI[1:], "/") + 1
	}

	// Check if the URI is the endpoint or a URI with a parameter
	isEndpoint := false
	if firstHalfUriIndex == -1 || relativeURI[firstHalfUriIndex+1] == ':' {
		firstHalfUri = relativeURI
		isEndpoint = true
	} else {
		firstHalfUri = relativeURI[:firstHalfUriIndex]
		secondHalfUri = relativeURI[firstHalfUriIndex:]
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
			return nil, MissingEndpointInterception
		}
	} else if isEndpoint {
		return nil, MissingInterceptions
	}

	// Check if the first half URI is in the children maps
	if m.ChildrenMaps != nil {
		// Check if the first half URI is in the children maps. If it is, traverse the child map
		if childMap, ok := (*m.ChildrenMaps)[firstHalfUri]; ok {
			return childMap.Traverse(secondHalfUri, restMethod)
		}
	} else {
		return nil, MissingChildrenMap
	}
	return nil, FailedToTraverse
}
