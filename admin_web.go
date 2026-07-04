package deco

// AdminWebExtraComponentInfoGet calls /admin/web?form=extra_component_info with operation get.
func (c *Client) AdminWebExtraComponentInfoGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/web", "extra_component_info", "get", params)
}
