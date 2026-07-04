package deco

import "testing"

func TestAdminWebMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminWebExtraComponentInfoGet",
			path:      "/admin/web",
			form:      "extra_component_info",
			operation: "get",
			call:      (*Client).AdminWebExtraComponentInfoGet,
		},
	})
}
