package deco

// AdminSyncCheckFirmware calls /admin/sync (no form) with operation check_firmware.
func (c *Client) AdminSyncCheckFirmware(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "check_firmware", params)
}

// AdminSyncForceUpgrade calls /admin/sync (no form) with operation force_upgrade.
func (c *Client) AdminSyncForceUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "force_upgrade", params)
}

// AdminSyncForceUpgradeLTE calls /admin/sync (no form) with operation force_upgrade_lte.
func (c *Client) AdminSyncForceUpgradeLTE(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "force_upgrade_lte", params)
}

// AdminSyncSyncCheck calls /admin/sync (no form) with operation sync_check.
func (c *Client) AdminSyncSyncCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_check", params)
}

// AdminSyncSyncConfig calls /admin/sync (no form) with operation sync_config.
func (c *Client) AdminSyncSyncConfig(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_config", params)
}

// AdminSyncSyncDetectSlave calls /admin/sync (no form) with operation sync_detect_slave.
func (c *Client) AdminSyncSyncDetectSlave(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_detect_slave", params)
}

// AdminSyncSyncDownloadBigfirm calls /admin/sync (no form) with operation sync_download_bigfirm.
func (c *Client) AdminSyncSyncDownloadBigfirm(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_download_bigfirm", params)
}

// AdminSyncSyncDownloadStatusBigfirm calls /admin/sync (no form) with operation sync_download_status_bigfirm.
func (c *Client) AdminSyncSyncDownloadStatusBigfirm(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_download_status_bigfirm", params)
}

// AdminSyncSyncEMMCCheck calls /admin/sync (no form) with operation sync_emmc_check.
func (c *Client) AdminSyncSyncEMMCCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_emmc_check", params)
}

// AdminSyncSyncEMMCConfig calls /admin/sync (no form) with operation sync_emmc_config.
func (c *Client) AdminSyncSyncEMMCConfig(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_emmc_config", params)
}

// AdminSyncSyncFirmware calls /admin/sync (no form) with operation sync_firmware.
func (c *Client) AdminSyncSyncFirmware(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_firmware", params)
}

// AdminSyncSyncGetCfg calls /admin/sync (no form) with operation sync_get_cfg.
func (c *Client) AdminSyncSyncGetCfg(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_get_cfg", params)
}

// AdminSyncSyncGetInfo calls /admin/sync (no form) with operation sync_get_info.
func (c *Client) AdminSyncSyncGetInfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_get_info", params)
}

// AdminSyncSyncISPProfile calls /admin/sync (no form) with operation sync_isp_profile.
func (c *Client) AdminSyncSyncISPProfile(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_isp_profile", params)
}

// AdminSyncSyncSubconfig calls /admin/sync (no form) with operation sync_subconfig.
func (c *Client) AdminSyncSyncSubconfig(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_subconfig", params)
}

// AdminSyncSyncUpdateDevList calls /admin/sync (no form) with operation sync_update_dev_list.
func (c *Client) AdminSyncSyncUpdateDevList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_update_dev_list", params)
}

// AdminSyncSyncUpgrade calls /admin/sync (no form) with operation sync_upgrade.
func (c *Client) AdminSyncSyncUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/sync", "", "sync_upgrade", params)
}
