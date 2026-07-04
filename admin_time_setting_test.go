package deco

import "testing"

func TestAdminTimeSettingMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminTimeSettingDstUpdateNotify",
			path:      "/admin/time_setting",
			form:      "dst_update",
			operation: "notify",
			call:      (*Client).AdminTimeSettingDstUpdateNotify,
		},
		{
			name:      "AdminTimeSettingDstUpdateRequest",
			path:      "/admin/time_setting",
			form:      "dst_update",
			operation: "request",
			call:      (*Client).AdminTimeSettingDstUpdateRequest,
		},
	})
}
