package auth

// MethodsToIntercept is the list of methods to intercept
var MethodsToIntercept = map[string]bool{
	"LogIn":                false,
	"IsAccessTokenValid":   true,
	"RefreshToken":         true,
	"LogOut":               true,
	"GetSessions":          true,
	"CloseSessions":        true,
	"AddPermission":        true,
	"RevokePermission":     true,
	"GetPermission":        true,
	"GetPermissions":       true,
	"AddRolePermission":    true,
	"RevokeRolePermission": true,
	"GetRolePermissions":   true,
	"AddRole":              true,
	"RevokeRole":           true,
	"GetRoles":             true,
	"AddUserRole":          true,
	"RevokeUserRole":       true,
	"GetUserRoles":         true,
}
