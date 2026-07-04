package deco

// AdminComponentControlSwitchListRead calls /admin/component_control?form=switch_list with operation read.
func (c *Client) AdminComponentControlSwitchListRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/component_control", "switch_list", "read", params)
}
