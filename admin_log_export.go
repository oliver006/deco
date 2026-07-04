package deco

// AdminLogExportFeedbackLogBuild calls /admin/log_export?form=feedback_log with operation build.
func (c *Client) AdminLogExportFeedbackLogBuild(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/log_export", "feedback_log", "build", params)
}

// AdminLogExportFeedbackLogRead calls /admin/log_export?form=feedback_log with operation read.
func (c *Client) AdminLogExportFeedbackLogRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/log_export", "feedback_log", "read", params)
}

// AdminLogExportFeedbackLogRemove calls /admin/log_export?form=feedback_log with operation remove.
func (c *Client) AdminLogExportFeedbackLogRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/log_export", "feedback_log", "remove", params)
}

// AdminLogExportSaveLogSave calls /admin/log_export?form=save_log with operation save.
func (c *Client) AdminLogExportSaveLogSave(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/log_export", "save_log", "save", params)
}

// AdminLogExportTypesRead calls /admin/log_export?form=types with operation read.
func (c *Client) AdminLogExportTypesRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/log_export", "types", "read", params)
}
