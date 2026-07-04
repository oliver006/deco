package deco

// AdminTimeSettingDstUpdateNotify calls /admin/time_setting?form=dst_update with operation notify.
func (c *Client) AdminTimeSettingDstUpdateNotify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/time_setting", "dst_update", "notify", params)
}

// AdminTimeSettingDstUpdateRequest calls /admin/time_setting?form=dst_update with operation request.
func (c *Client) AdminTimeSettingDstUpdateRequest(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/time_setting", "dst_update", "request", params)
}
