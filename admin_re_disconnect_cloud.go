package deco

// AdminREDisconnectCloudNotify calls /admin/re_disconnect_cloud (no form) with operation notify.
func (c *Client) AdminREDisconnectCloudNotify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/re_disconnect_cloud", "", "notify", params)
}
