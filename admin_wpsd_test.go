package deco

import "testing"

func TestAdminWPSDMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminWPSDMain",
			path:      "/admin/wpsd",
			form:      "",
			operation: "main",
			call:      (*Client).AdminWPSDMain,
		},
	})
}
