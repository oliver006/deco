package deco

// AdminTipcControllerNewdeviceSync calls /admin/tipc-controller?form=newdevice with operation sync.
func (c *Client) AdminTipcControllerNewdeviceSync(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/tipc-controller", "newdevice", "sync", params)
}

// AdminTipcControllerNewdeviceWrite calls /admin/tipc-controller?form=newdevice with operation write.
func (c *Client) AdminTipcControllerNewdeviceWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/tipc-controller", "newdevice", "write", params)
}
