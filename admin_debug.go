package deco

// AdminDebugAnonymous calls /admin/debug (no form) with operation .anonymous.
func (c *Client) AdminDebugAnonymous(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/debug", "", ".anonymous", params)
}
