package deco

import "testing"

func TestAdminSystemMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminSystemEnvarRead",
			path:      "/admin/system",
			form:      "envar",
			operation: "read",
			call:      (*Client).AdminSystemEnvarRead,
		},
		{
			name:      "AdminSystemEnvarWrite",
			path:      "/admin/system",
			form:      "envar",
			operation: "write",
			call:      (*Client).AdminSystemEnvarWrite,
		},
		{
			name:      "AdminSystemLogoutLogout",
			path:      "/admin/system",
			form:      "logout",
			operation: "logout",
			call:      (*Client).AdminSystemLogoutLogout,
		},
	})
}
