package deco

import "testing"

func TestAdminCWMPMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminCWMPCWMPInfoRead",
			path:      "/admin/cwmp",
			form:      "cwmp_info",
			operation: "read",
			call:      (*Client).AdminCWMPCWMPInfoRead,
		},
		{
			name:      "AdminCWMPCWMPInfoWrite",
			path:      "/admin/cwmp",
			form:      "cwmp_info",
			operation: "write",
			call:      (*Client).AdminCWMPCWMPInfoWrite,
		},
	})
}
