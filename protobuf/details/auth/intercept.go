package auth

import "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details"

// MethodsToIntercept is the list of methods to intercept
var MethodsToIntercept = map[string]details.Interception{
	"LogIn":                details.None,
	"IsAccessTokenValid":   details.AccessToken,
	"IsRefreshTokenValid":  details.RefreshToken,
	"RefreshToken":         details.RefreshToken,
	"LogOut":               details.AccessToken,
	"GetSessions":          details.AccessToken,
	"CloseSessions":        details.AccessToken,
	"AddPermission":        details.AccessToken,
	"RevokePermission":     details.AccessToken,
	"GetPermission":        details.AccessToken,
	"GetPermissions":       details.AccessToken,
	"AddRolePermission":    details.AccessToken,
	"RevokeRolePermission": details.AccessToken,
	"GetRolePermissions":   details.AccessToken,
	"AddRole":              details.AccessToken,
	"RevokeRole":           details.AccessToken,
	"GetRoles":             details.AccessToken,
	"AddUserRole":          details.AccessToken,
	"RevokeUserRole":       details.AccessToken,
	"GetUserRoles":         details.AccessToken,
}
