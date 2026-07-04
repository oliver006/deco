package deco

// AdminWPSDMain calls /admin/wpsd (no form) with operation main.
func (c *Client) AdminWPSDMain(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wpsd", "", "main", params)
}
