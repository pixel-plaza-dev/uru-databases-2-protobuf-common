package emails

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/interceptions/rest"
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
