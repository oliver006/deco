package deco

import "testing"

func TestAdminREDisconnectCloudMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminREDisconnectCloudNotify",
			path:      "/admin/re_disconnect_cloud",
			form:      "",
			operation: "notify",
			call:      (*Client).AdminREDisconnectCloudNotify,
		},
	})
}
