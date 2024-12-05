package rest

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
)

type (
	// Handler is a handler for a REST endpoint
	Handler struct {
		Endpoint   *Endpoint
		GRPCMethod grpc.Method
	}
)

// NewHandler creates a new handler
func NewHandler(
	endpoint *Endpoint,
	grpcMethod grpc.Method,
) *Handler {
	return &Handler{
		Endpoint:   endpoint,
		GRPCMethod: grpcMethod,
	}
}
