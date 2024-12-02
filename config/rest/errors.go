package rest

import (
	"errors"
)

var (
	MissingInterceptions        = errors.New("missing interceptions")
	MissingEndpointInterception = errors.New("missing endpoint interception")
	MissingChildrenMap          = errors.New("missing children map")
	FailedToTraverse            = errors.New("failed to traverse")
)
