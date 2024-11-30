package phone_numbers

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
)

// Map is the map of the REST API endpoints to the user service phone numbers gRPC methods
var Map = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.PUT: detailsuser.ChangePhoneNumber,
		rest.GET: detailsuser.GetPhoneNumber,
	},
	SendVerification.String(): {
		rest.POST: detailsuser.SendVerificationSMS,
	},
	VerifyByToken.String(): {
		rest.POST: detailsuser.VerifyPhoneNumber,
	},
}
