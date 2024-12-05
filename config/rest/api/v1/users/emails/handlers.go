package emails

import (
	grpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service emails endpoints handlers
var (
	GetActiveEmailsHandler       = typesrest.NewHandler(Relative, grpcuser.GetActiveEmails)
	AddEmailHandler              = typesrest.NewHandler(Relative, grpcuser.AddEmail)
	GetPrimaryEmailHandler       = typesrest.NewHandler(Primary, grpcuser.GetPrimaryEmail)
	ChangePrimaryEmailHandler    = typesrest.NewHandler(Primary, grpcuser.ChangePrimaryEmail)
	DeleteEmailHandler           = typesrest.NewHandler(ByEmail, grpcuser.DeleteEmail)
	SendVerificationEmailHandler = typesrest.NewHandler(SendVerification, grpcuser.SendVerificationEmail)
	VerifyEmailHandler           = typesrest.NewHandler(VerifyByToken, grpcuser.VerifyEmail)
)
