package users

import (
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users/emails"
	phonenumbers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users/phone-numbers"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users/profiles"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users/usernames"
	"github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
)

// ChildrenMaps is the map of the REST API endpoints of the auth service
var ChildrenMaps = map[string]*rest.Map{
	Emails.String():       emails.Map,
	PhoneNumbers.String(): phonenumbers.Map,
	Profiles.String():     profiles.Map,
	Usernames.String():    usernames.Map,
}

// Map is the map of the REST API endpoints of the auth service
var Map = rest.NewMap(
	&Interceptions,
	&ChildrenMaps,
)
