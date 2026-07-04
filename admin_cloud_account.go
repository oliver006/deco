package deco

// AdminCloudAccountCheckCloudConnectionRead calls /admin/cloud_account?form=check_cloud_connection with operation read.
func (c *Client) AdminCloudAccountCheckCloudConnectionRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "check_cloud_connection", "read", params)
}

// AdminCloudAccountCheckCloudVersionRead calls /admin/cloud_account?form=check_cloud_version with operation read.
func (c *Client) AdminCloudAccountCheckCloudVersionRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "check_cloud_version", "read", params)
}

// AdminCloudAccountCheckConnectionRead calls /admin/cloud_account?form=check_connection with operation read.
func (c *Client) AdminCloudAccountCheckConnectionRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "check_connection", "read", params)
}

// AdminCloudAccountCheckDeviceRead calls /admin/cloud_account?form=check_device with operation read.
func (c *Client) AdminCloudAccountCheckDeviceRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "check_device", "read", params)
}

// AdminCloudAccountCheckInternetRead calls /admin/cloud_account?form=check_internet with operation read.
func (c *Client) AdminCloudAccountCheckInternetRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "check_internet", "read", params)
}

// AdminCloudAccountCheckLoginRead calls /admin/cloud_account?form=check_login with operation read.
func (c *Client) AdminCloudAccountCheckLoginRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "check_login", "read", params)
}

// AdminCloudAccountCheckSupportRead calls /admin/cloud_account?form=check_support with operation read.
func (c *Client) AdminCloudAccountCheckSupportRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "check_support", "read", params)
}

// AdminCloudAccountCheckUpgradeRead calls /admin/cloud_account?form=check_upgrade with operation read.
func (c *Client) AdminCloudAccountCheckUpgradeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "check_upgrade", "read", params)
}

// AdminCloudAccountCloudUnbindWrite calls /admin/cloud_account?form=cloud_unbind with operation write.
func (c *Client) AdminCloudAccountCloudUnbindWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "cloud_unbind", "write", params)
}

// AdminCloudAccountCloudUpgradeLoad calls /admin/cloud_account?form=cloud_upgrade with operation load.
func (c *Client) AdminCloudAccountCloudUpgradeLoad(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "cloud_upgrade", "load", params)
}

// AdminCloudAccountCloudUpgradeRead calls /admin/cloud_account?form=cloud_upgrade with operation read.
func (c *Client) AdminCloudAccountCloudUpgradeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "cloud_upgrade", "read", params)
}

// AdminCloudAccountCloudUpgradeUpgrade calls /admin/cloud_account?form=cloud_upgrade with operation upgrade.
func (c *Client) AdminCloudAccountCloudUpgradeUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "cloud_upgrade", "upgrade", params)
}

// AdminCloudAccountDetectUpgradeStatusRead calls /admin/cloud_account?form=detect_upgrade_status with operation read.
func (c *Client) AdminCloudAccountDetectUpgradeStatusRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "detect_upgrade_status", "read", params)
}

// AdminCloudAccountGetDeviceInfoRead calls /admin/cloud_account?form=get_deviceInfo with operation read.
func (c *Client) AdminCloudAccountGetDeviceInfoRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "get_deviceInfo", "read", params)
}

// AdminCloudAccountGetTokenRead calls /admin/cloud_account?form=get_token with operation read.
func (c *Client) AdminCloudAccountGetTokenRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "get_token", "read", params)
}

// AdminCloudAccountModifyCloudPwdWrite calls /admin/cloud_account?form=modify_cloud_pwd with operation write.
func (c *Client) AdminCloudAccountModifyCloudPwdWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "modify_cloud_pwd", "write", params)
}

// AdminCloudAccountSetShowFlagWrite calls /admin/cloud_account?form=set_show_flag with operation write.
func (c *Client) AdminCloudAccountSetShowFlagWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "set_show_flag", "write", params)
}

// AdminCloudAccountTmpCmdBindOwner calls /admin/cloud_account?form=tmp_cmd with operation bind_owner.
func (c *Client) AdminCloudAccountTmpCmdBindOwner(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "tmp_cmd", "bind_owner", params)
}

// AdminCloudAccountTmpCmdCloudPassThrough calls /admin/cloud_account?form=tmp_cmd with operation cloud_pass_through.
func (c *Client) AdminCloudAccountTmpCmdCloudPassThrough(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "tmp_cmd", "cloud_pass_through", params)
}

// AdminCloudAccountTmpCmdGetDevInfo calls /admin/cloud_account?form=tmp_cmd with operation get_dev_info.
func (c *Client) AdminCloudAccountTmpCmdGetDevInfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "tmp_cmd", "get_dev_info", params)
}

// AdminCloudAccountTmpCmdSetDevInfo calls /admin/cloud_account?form=tmp_cmd with operation set_dev_info.
func (c *Client) AdminCloudAccountTmpCmdSetDevInfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "tmp_cmd", "set_dev_info", params)
}

// AdminCloudAccountTmpCmdUnbindOwner calls /admin/cloud_account?form=tmp_cmd with operation unbind_owner.
func (c *Client) AdminCloudAccountTmpCmdUnbindOwner(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "tmp_cmd", "unbind_owner", params)
}

// AdminCloudAccountUserLoginRead calls /admin/cloud_account?form=user_login with operation read.
func (c *Client) AdminCloudAccountUserLoginRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "user_login", "read", params)
}

// AdminCloudAccountUserLoginWrite calls /admin/cloud_account?form=user_login with operation write.
func (c *Client) AdminCloudAccountUserLoginWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud_account", "user_login", "write", params)
}
