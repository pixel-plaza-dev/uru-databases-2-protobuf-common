package accounts

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the payment service accounts REST endpoints
var Base = typesrest.NewEndpoint("accounts")

// Payment service accounts REST endpoints
var (
	Relative      = typesrest.NewRelativeEndpoint()
	Active        = typesrest.NewEndpoint("active")
	Activate      = typesrest.NewEndpoint("activate")
	Suspended     = typesrest.NewEndpoint("suspended")
	Suspend       = typesrest.NewEndpoint("suspend")
	VerifyPayment = typesrest.NewEndpoint("verify")
)
