package emails

import (
	grpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service emails endpoints mapping
var (
	GetActiveEmailsMapper       = typesrest.NewMapper(Relative, grpcuser.GetActiveEmails)
	AddEmailMapper              = typesrest.NewMapper(Relative, grpcuser.AddEmail)
	GetPrimaryEmailMapper       = typesrest.NewMapper(Primary, grpcuser.GetPrimaryEmail)
	ChangePrimaryEmailMapper    = typesrest.NewMapper(Primary, grpcuser.ChangePrimaryEmail)
	DeleteEmailMapper           = typesrest.NewMapper(ByEmail, grpcuser.DeleteEmail)
	SendVerificationEmailMapper = typesrest.NewMapper(SendVerification, grpcuser.SendVerificationEmail)
	VerifyEmailMapper           = typesrest.NewMapper(VerifyByToken, grpcuser.VerifyEmail)
)
