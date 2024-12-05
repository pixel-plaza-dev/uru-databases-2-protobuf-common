package phone_numbers

import (
	grpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Users service phone numbers endpoints mapping
var (
	ChangePhoneNumberMapper   = typesrest.NewMapper(Relative, grpcuser.ChangePhoneNumber)
	GetPhoneNumberMapper      = typesrest.NewMapper(Relative, grpcuser.GetPhoneNumber)
	SendVerificationSMSMapper = typesrest.NewMapper(SendVerification, grpcuser.SendVerificationSMS)
	VerifyPhoneNumberMapper   = typesrest.NewMapper(VerifyByToken, grpcuser.VerifyPhoneNumber)
)
