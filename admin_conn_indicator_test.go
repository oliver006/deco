package deco

import "testing"

func TestAdminConnIndicatorMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminConnIndicatorInternetDown",
			path:      "/admin/conn-indicator",
			form:      "internet",
			operation: "down",
			call:      (*Client).AdminConnIndicatorInternetDown,
		},
		{
			name:      "AdminConnIndicatorInternetSync",
			path:      "/admin/conn-indicator",
			form:      "internet",
			operation: "sync",
			call:      (*Client).AdminConnIndicatorInternetSync,
		},
		{
			name:      "AdminConnIndicatorInternetUp",
			path:      "/admin/conn-indicator",
			form:      "internet",
			operation: "up",
			call:      (*Client).AdminConnIndicatorInternetUp,
		},
	})
}
