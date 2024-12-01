package emails

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/grpc/user"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
)

// Interceptions is the map of the REST API endpoints to the user service emails gRPC methods
var Interceptions = map[string]map[rest.Method]grpc.Method{
	Relative.String(): {
		rest.GET:  user.GetActiveEmails,
		rest.POST: user.AddEmail,
	},
	Primary.String(): {
		rest.GET: user.GetPrimaryEmail,
		rest.PUT: user.ChangePrimaryEmail,
	},
	ByEmail.String(): {
		rest.DELETE: user.DeleteEmail,
	},
	SendVerification.String(): {
		rest.POST: user.SendVerificationEmail,
	},
	VerifyByToken.String(): {
		rest.POST: user.VerifyEmail,
	},
}
