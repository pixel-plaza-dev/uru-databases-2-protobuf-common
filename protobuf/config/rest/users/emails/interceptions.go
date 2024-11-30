package emails

import (
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Map is the map of the REST API endpoints to the user service emails gRPC methods
var Map = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.GET:  detailsuser.GetActiveEmails,
		rest.POST: detailsuser.AddEmail,
	},
	Primary.String(): {
		rest.GET: detailsuser.GetPrimaryEmail,
		rest.PUT: detailsuser.ChangePrimaryEmail,
	},
	ByEmail.String(): {
		rest.DELETE: detailsuser.DeleteEmail,
	},
	SendVerification.String(): {
		rest.POST: detailsuser.SendVerificationEmail,
	},
	VerifyByToken.String(): {
		rest.POST: detailsuser.VerifyEmail,
	},
}
