package deco

// AdminNetworkDHCPDialRead calls /admin/network?form=dhcp_dial with operation read.
func (c *Client) AdminNetworkDHCPDialRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "dhcp_dial", "read", params)
}

// AdminNetworkDHCPDialWrite calls /admin/network?form=dhcp_dial with operation write.
func (c *Client) AdminNetworkDHCPDialWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "dhcp_dial", "write", params)
}

// AdminNetworkIGMPSettingRead calls /admin/network?form=igmp_setting with operation read.
func (c *Client) AdminNetworkIGMPSettingRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "igmp_setting", "read", params)
}

// AdminNetworkIGMPSettingWrite calls /admin/network?form=igmp_setting with operation write.
func (c *Client) AdminNetworkIGMPSettingWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "igmp_setting", "write", params)
}

// AdminNetworkInternetRead calls /admin/network?form=internet with operation read.
func (c *Client) AdminNetworkInternetRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "internet", "read", params)
}

// AdminNetworkInternetWrite calls /admin/network?form=internet with operation write.
func (c *Client) AdminNetworkInternetWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "internet", "write", params)
}

// AdminNetworkIPv6Read calls /admin/network?form=ipv6 with operation read.
func (c *Client) AdminNetworkIPv6Read(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "ipv6", "read", params)
}

// AdminNetworkIPv6Write calls /admin/network?form=ipv6 with operation write.
func (c *Client) AdminNetworkIPv6Write(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "ipv6", "write", params)
}

// AdminNetworkLANIPRead calls /admin/network?form=lan_ip with operation read.
func (c *Client) AdminNetworkLANIPRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "lan_ip", "read", params)
}

// AdminNetworkLANIPWrite calls /admin/network?form=lan_ip with operation write.
func (c *Client) AdminNetworkLANIPWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "lan_ip", "write", params)
}

// AdminNetworkLANIPv4Read calls /admin/network?form=lan_ipv4 with operation read.
func (c *Client) AdminNetworkLANIPv4Read(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "lan_ipv4", "read", params)
}

// AdminNetworkLANIPv4Write calls /admin/network?form=lan_ipv4 with operation write.
func (c *Client) AdminNetworkLANIPv4Write(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "lan_ipv4", "write", params)
}

// AdminNetworkMACCloneRead calls /admin/network?form=mac_clone with operation read.
func (c *Client) AdminNetworkMACCloneRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "mac_clone", "read", params)
}

// AdminNetworkMACCloneWrite calls /admin/network?form=mac_clone with operation write.
func (c *Client) AdminNetworkMACCloneWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "mac_clone", "write", params)
}

// AdminNetworkPerformanceRead calls /admin/network?form=performance with operation read.
func (c *Client) AdminNetworkPerformanceRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "performance", "read", params)
}

// AdminNetworkVLANRead calls /admin/network?form=vlan with operation read.
func (c *Client) AdminNetworkVLANRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "vlan", "read", params)
}

// AdminNetworkVLANWrite calls /admin/network?form=vlan with operation write.
func (c *Client) AdminNetworkVLANWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "vlan", "write", params)
}

// AdminNetworkWANIPv4Connect calls /admin/network?form=wan_ipv4 with operation connect.
func (c *Client) AdminNetworkWANIPv4Connect(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "wan_ipv4", "connect", params)
}

// AdminNetworkWANIPv4Disconnect calls /admin/network?form=wan_ipv4 with operation disconnect.
func (c *Client) AdminNetworkWANIPv4Disconnect(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "wan_ipv4", "disconnect", params)
}

// AdminNetworkWANIPv4Read calls /admin/network?form=wan_ipv4 with operation read.
func (c *Client) AdminNetworkWANIPv4Read(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "wan_ipv4", "read", params)
}

// AdminNetworkWANIPv4Write calls /admin/network?form=wan_ipv4 with operation write.
func (c *Client) AdminNetworkWANIPv4Write(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "wan_ipv4", "write", params)
}

// AdminNetworkWANModeRead calls /admin/network?form=wan_mode with operation read.
func (c *Client) AdminNetworkWANModeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "wan_mode", "read", params)
}

// AdminNetworkWANModeWrite calls /admin/network?form=wan_mode with operation write.
func (c *Client) AdminNetworkWANModeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/network", "wan_mode", "write", params)
}
