package deco

// AdminCloudAccountNotify calls /admin/cloud?form=account with operation notify.
func (c *Client) AdminCloudAccountNotify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "account", "notify", params)
}

// AdminCloudAccountWrite calls /admin/cloud?form=account with operation write.
func (c *Client) AdminCloudAccountWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "account", "write", params)
}

// AdminCloudDDNSGet calls /admin/cloud?form=ddns with operation get.
func (c *Client) AdminCloudDDNSGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "ddns", "get", params)
}

// AdminCloudDDNSSet calls /admin/cloud?form=ddns with operation set.
func (c *Client) AdminCloudDDNSSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "ddns", "set", params)
}

// AdminCloudFirmwareDownload calls /admin/cloud?form=firmware with operation download.
func (c *Client) AdminCloudFirmwareDownload(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "firmware", "download", params)
}

// AdminCloudFirmwareSyncCheckFirmware calls /admin/cloud?form=firmware with operation sync_check_firmware.
func (c *Client) AdminCloudFirmwareSyncCheckFirmware(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "firmware", "sync_check_firmware", params)
}

// AdminCloudFirmwareUpload calls /admin/cloud?form=firmware with operation upload.
func (c *Client) AdminCloudFirmwareUpload(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "firmware", "upload", params)
}

// AdminCloudFirmwareStatusCheck calls /admin/cloud?form=firmware_status with operation check.
func (c *Client) AdminCloudFirmwareStatusCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "firmware_status", "check", params)
}

// AdminCloudFirmwareStatusCheckUpgrade calls /admin/cloud?form=firmware_status with operation check_upgrade.
func (c *Client) AdminCloudFirmwareStatusCheckUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "firmware_status", "check_upgrade", params)
}

// AdminCloudFirmwareStatusLocalUpgrade calls /admin/cloud?form=firmware_status with operation local_upgrade.
func (c *Client) AdminCloudFirmwareStatusLocalUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "firmware_status", "local_upgrade", params)
}

// AdminCloudFirmwareStatusRead calls /admin/cloud?form=firmware_status with operation read.
func (c *Client) AdminCloudFirmwareStatusRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "firmware_status", "read", params)
}

// AdminCloudFirmwareStatusUpgrade calls /admin/cloud?form=firmware_status with operation upgrade.
func (c *Client) AdminCloudFirmwareStatusUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "firmware_status", "upgrade", params)
}

// AdminCloudGroupAdd calls /admin/cloud?form=group with operation add.
func (c *Client) AdminCloudGroupAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "group", "add", params)
}

// AdminCloudGroupCreate calls /admin/cloud?form=group with operation create.
func (c *Client) AdminCloudGroupCreate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "group", "create", params)
}

// AdminCloudManagerGet calls /admin/cloud?form=manager with operation get.
func (c *Client) AdminCloudManagerGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "manager", "get", params)
}

// AdminCloudManagerSet calls /admin/cloud?form=manager with operation set.
func (c *Client) AdminCloudManagerSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "manager", "set", params)
}

// AdminCloudMessageGet calls /admin/cloud?form=message with operation get.
func (c *Client) AdminCloudMessageGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "message", "get", params)
}

// AdminCloudMessageIoTRead calls /admin/cloud?form=message with operation iot_read.
func (c *Client) AdminCloudMessageIoTRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "message", "iot_read", params)
}

// AdminCloudMessagePush calls /admin/cloud?form=message with operation push.
func (c *Client) AdminCloudMessagePush(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "message", "push", params)
}

// AdminCloudMessageRead calls /admin/cloud?form=message with operation read.
func (c *Client) AdminCloudMessageRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "message", "read", params)
}

// AdminCloudMessageRemove calls /admin/cloud?form=message with operation remove.
func (c *Client) AdminCloudMessageRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "message", "remove", params)
}

// AdminCloudMessageReport calls /admin/cloud?form=message with operation report.
func (c *Client) AdminCloudMessageReport(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "message", "report", params)
}

// AdminCloudMessageSet calls /admin/cloud?form=message with operation set.
func (c *Client) AdminCloudMessageSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "message", "set", params)
}

// AdminCloudMessageUpdate calls /admin/cloud?form=message with operation update.
func (c *Client) AdminCloudMessageUpdate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "message", "update", params)
}

// AdminCloudMiniFirmwareCheck calls /admin/cloud?form=mini_firmware with operation check.
func (c *Client) AdminCloudMiniFirmwareCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "mini_firmware", "check", params)
}

// AdminCloudMiniFirmwareCheckUpgrade calls /admin/cloud?form=mini_firmware with operation check_upgrade.
func (c *Client) AdminCloudMiniFirmwareCheckUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "mini_firmware", "check_upgrade", params)
}

// AdminCloudMiniFirmwareDownload calls /admin/cloud?form=mini_firmware with operation download.
func (c *Client) AdminCloudMiniFirmwareDownload(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "mini_firmware", "download", params)
}

// AdminCloudMiniFirmwareLocalUpgrade calls /admin/cloud?form=mini_firmware with operation local_upgrade.
func (c *Client) AdminCloudMiniFirmwareLocalUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "mini_firmware", "local_upgrade", params)
}

// AdminCloudMiniFirmwareRead calls /admin/cloud?form=mini_firmware with operation read.
func (c *Client) AdminCloudMiniFirmwareRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "mini_firmware", "read", params)
}

// AdminCloudMiniFirmwareUpgrade calls /admin/cloud?form=mini_firmware with operation upgrade.
func (c *Client) AdminCloudMiniFirmwareUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "mini_firmware", "upgrade", params)
}

// AdminCloudMiniFirmwareUpload calls /admin/cloud?form=mini_firmware with operation upload.
func (c *Client) AdminCloudMiniFirmwareUpload(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "mini_firmware", "upload", params)
}

// AdminCloudNicknameRead calls /admin/cloud?form=nickname with operation read.
func (c *Client) AdminCloudNicknameRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "nickname", "read", params)
}

// AdminCloudNicknameWrite calls /admin/cloud?form=nickname with operation write.
func (c *Client) AdminCloudNicknameWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "nickname", "write", params)
}

// AdminCloudProxyWrite calls /admin/cloud?form=proxy with operation write.
func (c *Client) AdminCloudProxyWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "proxy", "write", params)
}

// AdminCloudSystemBind calls /admin/cloud?form=system with operation bind.
func (c *Client) AdminCloudSystemBind(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "system", "bind", params)
}

// AdminCloudSystemRemove calls /admin/cloud?form=system with operation remove.
func (c *Client) AdminCloudSystemRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "system", "remove", params)
}

// AdminCloudSystemRemoveAll calls /admin/cloud?form=system with operation remove_all.
func (c *Client) AdminCloudSystemRemoveAll(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "system", "remove_all", params)
}

// AdminCloudSystemUnbind calls /admin/cloud?form=system with operation unbind.
func (c *Client) AdminCloudSystemUnbind(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/cloud", "system", "unbind", params)
}
