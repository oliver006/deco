package deco

import "testing"

func TestAdminLogExportMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminLogExportFeedbackLogBuild",
			path:      "/admin/log_export",
			form:      "feedback_log",
			operation: "build",
			call:      (*Client).AdminLogExportFeedbackLogBuild,
		},
		{
			name:      "AdminLogExportFeedbackLogRead",
			path:      "/admin/log_export",
			form:      "feedback_log",
			operation: "read",
			call:      (*Client).AdminLogExportFeedbackLogRead,
		},
		{
			name:      "AdminLogExportFeedbackLogRemove",
			path:      "/admin/log_export",
			form:      "feedback_log",
			operation: "remove",
			call:      (*Client).AdminLogExportFeedbackLogRemove,
		},
		{
			name:      "AdminLogExportSaveLogSave",
			path:      "/admin/log_export",
			form:      "save_log",
			operation: "save",
			call:      (*Client).AdminLogExportSaveLogSave,
		},
		{
			name:      "AdminLogExportTypesRead",
			path:      "/admin/log_export",
			form:      "types",
			operation: "read",
			call:      (*Client).AdminLogExportTypesRead,
		},
	})
}
