package phone_numbers

import (
	grpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service phone numbers endpoints handlers
var (
	ChangePhoneNumberHandler   = typesrest.NewHandler(Relative, grpcuser.ChangePhoneNumber)
	GetPhoneNumberHandler      = typesrest.NewHandler(Relative, grpcuser.GetPhoneNumber)
	SendVerificationSMSHandler = typesrest.NewHandler(SendVerification, grpcuser.SendVerificationSMS)
	VerifyPhoneNumberHandler   = typesrest.NewHandler(VerifyByToken, grpcuser.VerifyPhoneNumber)
)
