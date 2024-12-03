package phone_numbers

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// Interceptions is the map of the REST API endpoints to the user service phone numbers gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.PUT: user.ChangePhoneNumber,
		rest.GET: user.GetPhoneNumber,
	},
	SendVerification.String(): {
		rest.POST: user.SendVerificationSMS,
	},
	VerifyByToken.String(): {
		rest.POST: user.VerifyPhoneNumber,
	},
}
