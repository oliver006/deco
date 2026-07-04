package deco

import "testing"

func TestAdminDebugMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminDebugAnonymous",
			path:      "/admin/debug",
			form:      "",
			operation: ".anonymous",
			call:      (*Client).AdminDebugAnonymous,
		},
	})
}
