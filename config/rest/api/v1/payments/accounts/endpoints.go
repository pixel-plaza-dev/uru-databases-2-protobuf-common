package accounts

import (
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Base is the base endpoint for the payment service accounts REST endpoints
var Base = typesrest.NewEndpoint("accounts")

// Payment service accounts REST endpoints
var (
	Relative            = typesrest.NewRelativeEndpoint()
	Active              = typesrest.NewEndpoint("active")
	ActivateByAccountId = typesrest.NewEndpoint("activate", typesrest.AccountId)
	Suspended           = typesrest.NewEndpoint("suspended")
	SuspendByAccountId  = typesrest.NewEndpoint("suspend", typesrest.AccountId)
	VerifyPayment       = typesrest.NewEndpoint("verify")
)
