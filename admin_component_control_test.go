package deco

import "testing"

func TestAdminComponentControlMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminComponentControlSwitchListRead",
			path:      "/admin/component_control",
			form:      "switch_list",
			operation: "read",
			call:      (*Client).AdminComponentControlSwitchListRead,
		},
	})
}
