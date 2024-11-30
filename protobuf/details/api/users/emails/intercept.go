package emails

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/api"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	detailsuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user"
)

// RESTMap is the map of the REST API endpoints to the user service emails gRPC methods
var RESTMap = map[string]map[api.RESTMethod]types.GRPCMethod{
	Relative.String(): {
		api.GET:  detailsuser.GetActiveEmails,
		api.POST: detailsuser.AddEmail,
	},
	Primary.String(): {
		api.GET: detailsuser.GetPrimaryEmail,
		api.PUT: detailsuser.ChangePrimaryEmail,
	},
	ByEmail.String(): {
		api.DELETE: detailsuser.DeleteEmail,
	},
	SendVerification.String(): {
		api.POST: detailsuser.SendVerificationEmail,
	},
	VerifyByToken.String(): {
		api.POST: detailsuser.VerifyEmail,
	},
}
