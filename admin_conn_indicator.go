package deco

// AdminConnIndicatorInternetDown calls /admin/conn-indicator?form=internet with operation down.
func (c *Client) AdminConnIndicatorInternetDown(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/conn-indicator", "internet", "down", params)
}

// AdminConnIndicatorInternetSync calls /admin/conn-indicator?form=internet with operation sync.
func (c *Client) AdminConnIndicatorInternetSync(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/conn-indicator", "internet", "sync", params)
}

// AdminConnIndicatorInternetUp calls /admin/conn-indicator?form=internet with operation up.
func (c *Client) AdminConnIndicatorInternetUp(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/conn-indicator", "internet", "up", params)
}
