package rest

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

type (
	// Mapper is the mapper for a REST endpoint and a gRPC method
	Mapper struct {
		Endpoint   *Endpoint
		GRPCMethod grpc.Method
	}
)

// NewMapper creates a new mapper
func NewMapper(
	endpoint *Endpoint,
	grpcMethod grpc.Method,
) *Mapper {
	return &Mapper{
		Endpoint:   endpoint,
		GRPCMethod: grpcMethod,
	}
}

// Path returns the path of the REST endpoint
func (m Mapper) Path() string {
	return m.Endpoint.Path()
}
