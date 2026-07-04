package deco

// AdminWirelessBandwidthEnhanceRead calls /admin/wireless?form=bandwidth_enhance with operation read.
func (c *Client) AdminWirelessBandwidthEnhanceRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "bandwidth_enhance", "read", params)
}

// AdminWirelessBandwidthEnhanceWrite calls /admin/wireless?form=bandwidth_enhance with operation write.
func (c *Client) AdminWirelessBandwidthEnhanceWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "bandwidth_enhance", "write", params)
}

// AdminWirelessBeamformingRead calls /admin/wireless?form=beamforming with operation read.
func (c *Client) AdminWirelessBeamformingRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "beamforming", "read", params)
}

// AdminWirelessBeamformingWrite calls /admin/wireless?form=beamforming with operation write.
func (c *Client) AdminWirelessBeamformingWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "beamforming", "write", params)
}

// AdminWirelessBridgeCheck calls /admin/wireless?form=bridge with operation check.
func (c *Client) AdminWirelessBridgeCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "bridge", "check", params)
}

// AdminWirelessBridgeRead calls /admin/wireless?form=bridge with operation read.
func (c *Client) AdminWirelessBridgeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "bridge", "read", params)
}

// AdminWirelessIeee80211rRead calls /admin/wireless?form=ieee80211r with operation read.
func (c *Client) AdminWirelessIeee80211rRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "ieee80211r", "read", params)
}

// AdminWirelessIeee80211rWrite calls /admin/wireless?form=ieee80211r with operation write.
func (c *Client) AdminWirelessIeee80211rWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "ieee80211r", "write", params)
}

// AdminWirelessOperationModeRead calls /admin/wireless?form=operation_mode with operation read.
func (c *Client) AdminWirelessOperationModeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "operation_mode", "read", params)
}

// AdminWirelessOperationModeWrite calls /admin/wireless?form=operation_mode with operation write.
func (c *Client) AdminWirelessOperationModeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "operation_mode", "write", params)
}

// AdminWirelessPowerRead calls /admin/wireless?form=power with operation read.
func (c *Client) AdminWirelessPowerRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "power", "read", params)
}

// AdminWirelessPowerWrite calls /admin/wireless?form=power with operation write.
func (c *Client) AdminWirelessPowerWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "power", "write", params)
}

// AdminWirelessWLANRead calls /admin/wireless?form=wlan with operation read.
func (c *Client) AdminWirelessWLANRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "wlan", "read", params)
}

// AdminWirelessWLANWrite calls /admin/wireless?form=wlan with operation write.
func (c *Client) AdminWirelessWLANWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/wireless", "wlan", "write", params)
}
