package phone_numbers

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user"
)

// RESTMap is the map of the REST API endpoints to the user service phone numbers gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Relative.String(): {
		api.PUT: detailsuser.ChangePhoneNumber,
		api.GET: detailsuser.GetPhoneNumber,
	},
	SendVerification.String(): {
		api.POST: detailsuser.SendVerificationSMS,
	},
	VerifyByToken.String(): {
		api.POST: detailsuser.VerifyPhoneNumber,
	},
}
