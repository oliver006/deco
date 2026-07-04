package deco

// AdminCWMPCWMPInfoRead calls /admin/cwmp?form=cwmp_info with operation read.
func (c *Client) AdminCWMPCWMPInfoRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cwmp", "cwmp_info", "read", params)
}

// AdminCWMPCWMPInfoWrite calls /admin/cwmp?form=cwmp_info with operation write.
func (c *Client) AdminCWMPCWMPInfoWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cwmp", "cwmp_info", "write", params)
}
