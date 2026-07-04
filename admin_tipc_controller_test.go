package deco

import "testing"

func TestAdminTipcControllerMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminTipcControllerNewdeviceSync",
			path:      "/admin/tipc-controller",
			form:      "newdevice",
			operation: "sync",
			call:      (*Client).AdminTipcControllerNewdeviceSync,
		},
		{
			name:      "AdminTipcControllerNewdeviceWrite",
			path:      "/admin/tipc-controller",
			form:      "newdevice",
			operation: "write",
			call:      (*Client).AdminTipcControllerNewdeviceWrite,
		},
	})
}
