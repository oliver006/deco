package deco

// AdminDeviceDeviceListRead calls /admin/device?form=device_list with operation read.
func (c *Client) AdminDeviceDeviceListRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "device_list", "read", params)
}

// AdminDeviceDeviceListRemove calls /admin/device?form=device_list with operation remove.
func (c *Client) AdminDeviceDeviceListRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "device_list", "remove", params)
}

// AdminDeviceMiniDeviceListRead calls /admin/device?form=mini_device_list with operation read.
func (c *Client) AdminDeviceMiniDeviceListRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "mini_device_list", "read", params)
}

// AdminDeviceModeRead calls /admin/device?form=mode with operation read.
func (c *Client) AdminDeviceModeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "mode", "read", params)
}

// AdminDeviceSpeedtestClear calls /admin/device?form=speedtest with operation clear.
func (c *Client) AdminDeviceSpeedtestClear(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "speedtest", "clear", params)
}

// AdminDeviceSpeedtestGet calls /admin/device?form=speedtest with operation get.
func (c *Client) AdminDeviceSpeedtestGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "speedtest", "get", params)
}

// AdminDeviceSpeedtestGetServer calls /admin/device?form=speedtest with operation get_server.
func (c *Client) AdminDeviceSpeedtestGetServer(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "speedtest", "get_server", params)
}

// AdminDeviceSpeedtestRead calls /admin/device?form=speedtest with operation read.
func (c *Client) AdminDeviceSpeedtestRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "speedtest", "read", params)
}

// AdminDeviceSpeedtestStop calls /admin/device?form=speedtest with operation stop.
func (c *Client) AdminDeviceSpeedtestStop(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "speedtest", "stop", params)
}

// AdminDeviceSpeedtestWrite calls /admin/device?form=speedtest with operation write.
func (c *Client) AdminDeviceSpeedtestWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "speedtest", "write", params)
}

// AdminDeviceSyncRead calls /admin/device?form=sync with operation read.
func (c *Client) AdminDeviceSyncRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "sync", "read", params)
}

// AdminDeviceSyncWrite calls /admin/device?form=sync with operation write.
func (c *Client) AdminDeviceSyncWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "sync", "write", params)
}

// AdminDeviceSystemFactory calls /admin/device?form=system with operation factory.
func (c *Client) AdminDeviceSystemFactory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "system", "factory", params)
}

// AdminDeviceSystemGateway calls /admin/device?form=system with operation gateway.
func (c *Client) AdminDeviceSystemGateway(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "system", "gateway", params)
}

// AdminDeviceSystemReboot calls /admin/device?form=system with operation reboot.
func (c *Client) AdminDeviceSystemReboot(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "system", "reboot", params)
}

// AdminDeviceTimesettingRead calls /admin/device?form=timesetting with operation read.
func (c *Client) AdminDeviceTimesettingRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "timesetting", "read", params)
}

// AdminDeviceTimesettingWrite calls /admin/device?form=timesetting with operation write.
func (c *Client) AdminDeviceTimesettingWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/device", "timesetting", "write", params)
}
