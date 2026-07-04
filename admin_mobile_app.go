package deco

// AdminMobileAppAutoTestAllInfoRead calls /admin/mobile_app/auto_test?form=all_info with operation read.
func (c *Client) AdminMobileAppAutoTestAllInfoRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/auto_test", "all_info", "read", params)
}

// AdminMobileAppAutoTestDHCPRead calls /admin/mobile_app/auto_test?form=dhcp with operation read.
func (c *Client) AdminMobileAppAutoTestDHCPRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/auto_test", "dhcp", "read", params)
}

// AdminMobileAppAutoTestNATRead calls /admin/mobile_app/auto_test?form=nat with operation read.
func (c *Client) AdminMobileAppAutoTestNATRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/auto_test", "nat", "read", params)
}

// AdminMobileAppAutoTestTestRead calls /admin/mobile_app/auto_test?form=test with operation read.
func (c *Client) AdminMobileAppAutoTestTestRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/auto_test", "test", "read", params)
}

// AdminMobileAppAutoTestUPnPList calls /admin/mobile_app/auto_test?form=upnp with operation list.
func (c *Client) AdminMobileAppAutoTestUPnPList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/auto_test", "upnp", "list", params)
}

// AdminMobileAppAutoTestUPnPRead calls /admin/mobile_app/auto_test?form=upnp with operation read.
func (c *Client) AdminMobileAppAutoTestUPnPRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/auto_test", "upnp", "read", params)
}

// AdminMobileAppAutoTestWiFiRead calls /admin/mobile_app/auto_test?form=wifi with operation read.
func (c *Client) AdminMobileAppAutoTestWiFiRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/auto_test", "wifi", "read", params)
}

// AdminMobileAppClientAddrReservationAdd calls /admin/mobile_app/client?form=addr_reservation with operation add.
func (c *Client) AdminMobileAppClientAddrReservationAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "addr_reservation", "add", params)
}

// AdminMobileAppClientAddrReservationGetlist calls /admin/mobile_app/client?form=addr_reservation with operation getlist.
func (c *Client) AdminMobileAppClientAddrReservationGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "addr_reservation", "getlist", params)
}

// AdminMobileAppClientAddrReservationModify calls /admin/mobile_app/client?form=addr_reservation with operation modify.
func (c *Client) AdminMobileAppClientAddrReservationModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "addr_reservation", "modify", params)
}

// AdminMobileAppClientAddrReservationRemove calls /admin/mobile_app/client?form=addr_reservation with operation remove.
func (c *Client) AdminMobileAppClientAddrReservationRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "addr_reservation", "remove", params)
}

// AdminMobileAppClientBlackListBlock calls /admin/mobile_app/client?form=black_list with operation block.
func (c *Client) AdminMobileAppClientBlackListBlock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "black_list", "block", params)
}

// AdminMobileAppClientBlackListList calls /admin/mobile_app/client?form=black_list with operation list.
func (c *Client) AdminMobileAppClientBlackListList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "black_list", "list", params)
}

// AdminMobileAppClientBlackListModify calls /admin/mobile_app/client?form=black_list with operation modify.
func (c *Client) AdminMobileAppClientBlackListModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "black_list", "modify", params)
}

// AdminMobileAppClientBlackListUnblock calls /admin/mobile_app/client?form=black_list with operation unblock.
func (c *Client) AdminMobileAppClientBlackListUnblock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "black_list", "unblock", params)
}

// AdminMobileAppClientClientAccessRead calls /admin/mobile_app/client?form=client_access with operation read.
func (c *Client) AdminMobileAppClientClientAccessRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "client_access", "read", params)
}

// AdminMobileAppClientClientAccessWrite calls /admin/mobile_app/client?form=client_access with operation write.
func (c *Client) AdminMobileAppClientClientAccessWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "client_access", "write", params)
}

// AdminMobileAppClientClientIsolationRead calls /admin/mobile_app/client?form=client_isolation with operation read.
func (c *Client) AdminMobileAppClientClientIsolationRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "client_isolation", "read", params)
}

// AdminMobileAppClientClientIsolationWrite calls /admin/mobile_app/client?form=client_isolation with operation write.
func (c *Client) AdminMobileAppClientClientIsolationWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "client_isolation", "write", params)
}

// AdminMobileAppClientClientListRead calls /admin/mobile_app/client?form=client_list with operation read.
func (c *Client) AdminMobileAppClientClientListRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "client_list", "read", params)
}

// AdminMobileAppClientClientListRemove calls /admin/mobile_app/client?form=client_list with operation remove.
func (c *Client) AdminMobileAppClientClientListRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "client_list", "remove", params)
}

// AdminMobileAppClientClientListWrite calls /admin/mobile_app/client?form=client_list with operation write.
func (c *Client) AdminMobileAppClientClientListWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "client_list", "write", params)
}

// AdminMobileAppClientLeaseGet calls /admin/mobile_app/client?form=lease with operation get.
func (c *Client) AdminMobileAppClientLeaseGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "lease", "get", params)
}

// AdminMobileAppClientTrafficStatClient calls /admin/mobile_app/client?form=traffic_stat with operation client.
func (c *Client) AdminMobileAppClientTrafficStatClient(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "traffic_stat", "client", params)
}

// AdminMobileAppClientTrafficStatList calls /admin/mobile_app/client?form=traffic_stat with operation list.
func (c *Client) AdminMobileAppClientTrafficStatList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/client", "traffic_stat", "list", params)
}

// AdminMobileAppCloudAccountNotify calls /admin/mobile_app/cloud?form=account with operation notify.
func (c *Client) AdminMobileAppCloudAccountNotify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "account", "notify", params)
}

// AdminMobileAppCloudAccountSync calls /admin/mobile_app/cloud?form=account with operation sync.
func (c *Client) AdminMobileAppCloudAccountSync(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "account", "sync", params)
}

// AdminMobileAppCloudAccountTransfer calls /admin/mobile_app/cloud?form=account with operation transfer.
func (c *Client) AdminMobileAppCloudAccountTransfer(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "account", "transfer", params)
}

// AdminMobileAppCloudAccountWrite calls /admin/mobile_app/cloud?form=account with operation write.
func (c *Client) AdminMobileAppCloudAccountWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "account", "write", params)
}

// AdminMobileAppCloudAutoUpgradeRead calls /admin/mobile_app/cloud?form=auto_upgrade with operation read.
func (c *Client) AdminMobileAppCloudAutoUpgradeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "auto_upgrade", "read", params)
}

// AdminMobileAppCloudAutoUpgradeWrite calls /admin/mobile_app/cloud?form=auto_upgrade with operation write.
func (c *Client) AdminMobileAppCloudAutoUpgradeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "auto_upgrade", "write", params)
}

// AdminMobileAppCloudDDNSGet calls /admin/mobile_app/cloud?form=ddns with operation get.
func (c *Client) AdminMobileAppCloudDDNSGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "ddns", "get", params)
}

// AdminMobileAppCloudDDNSSet calls /admin/mobile_app/cloud?form=ddns with operation set.
func (c *Client) AdminMobileAppCloudDDNSSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "ddns", "set", params)
}

// AdminMobileAppCloudFirmwareCheck calls /admin/mobile_app/cloud?form=firmware with operation check.
func (c *Client) AdminMobileAppCloudFirmwareCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "firmware", "check", params)
}

// AdminMobileAppCloudFirmwareDownload calls /admin/mobile_app/cloud?form=firmware with operation download.
func (c *Client) AdminMobileAppCloudFirmwareDownload(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "firmware", "download", params)
}

// AdminMobileAppCloudFirmwareGetSync calls /admin/mobile_app/cloud?form=firmware with operation get_sync.
func (c *Client) AdminMobileAppCloudFirmwareGetSync(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "firmware", "get_sync", params)
}

// AdminMobileAppCloudFirmwareRead calls /admin/mobile_app/cloud?form=firmware with operation read.
func (c *Client) AdminMobileAppCloudFirmwareRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "firmware", "read", params)
}

// AdminMobileAppCloudFirmwareSyncCheckFirmware calls /admin/mobile_app/cloud?form=firmware with operation sync_check_firmware.
func (c *Client) AdminMobileAppCloudFirmwareSyncCheckFirmware(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "firmware", "sync_check_firmware", params)
}

// AdminMobileAppCloudFirmwareUpgrade calls /admin/mobile_app/cloud?form=firmware with operation upgrade.
func (c *Client) AdminMobileAppCloudFirmwareUpgrade(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "firmware", "upgrade", params)
}

// AdminMobileAppCloudGroupAdd calls /admin/mobile_app/cloud?form=group with operation add.
func (c *Client) AdminMobileAppCloudGroupAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "group", "add", params)
}

// AdminMobileAppCloudGroupCreate calls /admin/mobile_app/cloud?form=group with operation create.
func (c *Client) AdminMobileAppCloudGroupCreate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "group", "create", params)
}

// AdminMobileAppCloudHomecareServiceGet calls /admin/mobile_app/cloud?form=homecare_service with operation get.
func (c *Client) AdminMobileAppCloudHomecareServiceGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "homecare_service", "get", params)
}

// AdminMobileAppCloudManagerGet calls /admin/mobile_app/cloud?form=manager with operation get.
func (c *Client) AdminMobileAppCloudManagerGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "manager", "get", params)
}

// AdminMobileAppCloudManagerSet calls /admin/mobile_app/cloud?form=manager with operation set.
func (c *Client) AdminMobileAppCloudManagerSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "manager", "set", params)
}

// AdminMobileAppCloudMessageGet calls /admin/mobile_app/cloud?form=message with operation get.
func (c *Client) AdminMobileAppCloudMessageGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "get", params)
}

// AdminMobileAppCloudMessageIoTRead calls /admin/mobile_app/cloud?form=message with operation iot_read.
func (c *Client) AdminMobileAppCloudMessageIoTRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "iot_read", params)
}

// AdminMobileAppCloudMessagePush calls /admin/mobile_app/cloud?form=message with operation push.
func (c *Client) AdminMobileAppCloudMessagePush(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "push", params)
}

// AdminMobileAppCloudMessagePushWeekly calls /admin/mobile_app/cloud?form=message with operation push_weekly.
func (c *Client) AdminMobileAppCloudMessagePushWeekly(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "push_weekly", params)
}

// AdminMobileAppCloudMessageRead calls /admin/mobile_app/cloud?form=message with operation read.
func (c *Client) AdminMobileAppCloudMessageRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "read", params)
}

// AdminMobileAppCloudMessageRemove calls /admin/mobile_app/cloud?form=message with operation remove.
func (c *Client) AdminMobileAppCloudMessageRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "remove", params)
}

// AdminMobileAppCloudMessageReport calls /admin/mobile_app/cloud?form=message with operation report.
func (c *Client) AdminMobileAppCloudMessageReport(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "report", params)
}

// AdminMobileAppCloudMessageSet calls /admin/mobile_app/cloud?form=message with operation set.
func (c *Client) AdminMobileAppCloudMessageSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "set", params)
}

// AdminMobileAppCloudMessageUpdate calls /admin/mobile_app/cloud?form=message with operation update.
func (c *Client) AdminMobileAppCloudMessageUpdate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "message", "update", params)
}

// AdminMobileAppCloudNicknameRead calls /admin/mobile_app/cloud?form=nickname with operation read.
func (c *Client) AdminMobileAppCloudNicknameRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "nickname", "read", params)
}

// AdminMobileAppCloudNicknameWrite calls /admin/mobile_app/cloud?form=nickname with operation write.
func (c *Client) AdminMobileAppCloudNicknameWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "nickname", "write", params)
}

// AdminMobileAppCloudProxyWrite calls /admin/mobile_app/cloud?form=proxy with operation write.
func (c *Client) AdminMobileAppCloudProxyWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "proxy", "write", params)
}

// AdminMobileAppCloudSystemBind calls /admin/mobile_app/cloud?form=system with operation bind.
func (c *Client) AdminMobileAppCloudSystemBind(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "system", "bind", params)
}

// AdminMobileAppCloudSystemRemove calls /admin/mobile_app/cloud?form=system with operation remove.
func (c *Client) AdminMobileAppCloudSystemRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "system", "remove", params)
}

// AdminMobileAppCloudSystemRemoveAll calls /admin/mobile_app/cloud?form=system with operation remove_all.
func (c *Client) AdminMobileAppCloudSystemRemoveAll(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "system", "remove_all", params)
}

// AdminMobileAppCloudSystemUnbind calls /admin/mobile_app/cloud?form=system with operation unbind.
func (c *Client) AdminMobileAppCloudSystemUnbind(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cloud", "system", "unbind", params)
}

// AdminMobileAppComponentListBluetoothRead calls /admin/mobile_app/component_list?form=bluetooth with operation read.
func (c *Client) AdminMobileAppComponentListBluetoothRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/component_list", "bluetooth", "read", params)
}

// AdminMobileAppComponentListMobileRead calls /admin/mobile_app/component_list?form=mobile with operation read.
func (c *Client) AdminMobileAppComponentListMobileRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/component_list", "mobile", "read", params)
}

// AdminMobileAppComponentListProfileRead calls /admin/mobile_app/component_list?form=profile with operation read.
func (c *Client) AdminMobileAppComponentListProfileRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/component_list", "profile", "read", params)
}

// AdminMobileAppCWMPCWMPGet calls /admin/mobile_app/cwmp?form=cwmp with operation get.
func (c *Client) AdminMobileAppCWMPCWMPGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cwmp", "cwmp", "get", params)
}

// AdminMobileAppCWMPCWMPSet calls /admin/mobile_app/cwmp?form=cwmp with operation set.
func (c *Client) AdminMobileAppCWMPCWMPSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/cwmp", "cwmp", "set", params)
}

// AdminMobileAppDDNSDDNSGet calls /admin/mobile_app/ddns?form=ddns with operation get.
func (c *Client) AdminMobileAppDDNSDDNSGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ddns", "ddns", "get", params)
}

// AdminMobileAppDDNSDDNSSet calls /admin/mobile_app/ddns?form=ddns with operation set.
func (c *Client) AdminMobileAppDDNSDDNSSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ddns", "ddns", "set", params)
}

// AdminMobileAppDebugQLogStart calls /admin/mobile_app/debug?form=qlog with operation start.
func (c *Client) AdminMobileAppDebugQLogStart(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/debug", "qlog", "start", params)
}

// AdminMobileAppDebugQLogStop calls /admin/mobile_app/debug?form=qlog with operation stop.
func (c *Client) AdminMobileAppDebugQLogStop(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/debug", "qlog", "stop", params)
}

// AdminMobileAppDebugSimplecomStart calls /admin/mobile_app/debug?form=simplecom with operation start.
func (c *Client) AdminMobileAppDebugSimplecomStart(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/debug", "simplecom", "start", params)
}

// AdminMobileAppDebugSimplecomStop calls /admin/mobile_app/debug?form=simplecom with operation stop.
func (c *Client) AdminMobileAppDebugSimplecomStop(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/debug", "simplecom", "stop", params)
}

// AdminMobileAppDebugTMStart calls /admin/mobile_app/debug?form=tm with operation start.
func (c *Client) AdminMobileAppDebugTMStart(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/debug", "tm", "start", params)
}

// AdminMobileAppDebugTMStop calls /admin/mobile_app/debug?form=tm with operation stop.
func (c *Client) AdminMobileAppDebugTMStop(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/debug", "tm", "stop", params)
}

// AdminMobileAppDebugTty2tcpStart calls /admin/mobile_app/debug?form=tty2tcp with operation start.
func (c *Client) AdminMobileAppDebugTty2tcpStart(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/debug", "tty2tcp", "start", params)
}

// AdminMobileAppDebugTty2tcpStop calls /admin/mobile_app/debug?form=tty2tcp with operation stop.
func (c *Client) AdminMobileAppDebugTty2tcpStop(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/debug", "tty2tcp", "stop", params)
}

// AdminMobileAppDeviceAlertOperationModeSwitchSet calls /admin/mobile_app/device?form=alert_operation_mode_switch with operation set.
func (c *Client) AdminMobileAppDeviceAlertOperationModeSwitchSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "alert_operation_mode_switch", "set", params)
}

// AdminMobileAppDeviceDetectModeRead calls /admin/mobile_app/device?form=detect_mode with operation read.
func (c *Client) AdminMobileAppDeviceDetectModeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "detect_mode", "read", params)
}

// AdminMobileAppDeviceDetectModeWrite calls /admin/mobile_app/device?form=detect_mode with operation write.
func (c *Client) AdminMobileAppDeviceDetectModeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "detect_mode", "write", params)
}

// AdminMobileAppDeviceDeviceListRead calls /admin/mobile_app/device?form=device_list with operation read.
func (c *Client) AdminMobileAppDeviceDeviceListRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "device_list", "read", params)
}

// AdminMobileAppDeviceDeviceListRemove calls /admin/mobile_app/device?form=device_list with operation remove.
func (c *Client) AdminMobileAppDeviceDeviceListRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "device_list", "remove", params)
}

// AdminMobileAppDeviceDevicePreferSetSet calls /admin/mobile_app/device?form=device_prefer_set with operation set.
func (c *Client) AdminMobileAppDeviceDevicePreferSetSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "device_prefer_set", "set", params)
}

// AdminMobileAppDeviceEcoModeRead calls /admin/mobile_app/device?form=eco_mode with operation read.
func (c *Client) AdminMobileAppDeviceEcoModeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "eco_mode", "read", params)
}

// AdminMobileAppDeviceEcoModeWrite calls /admin/mobile_app/device?form=eco_mode with operation write.
func (c *Client) AdminMobileAppDeviceEcoModeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "eco_mode", "write", params)
}

// AdminMobileAppDeviceEnvarWrite calls /admin/mobile_app/device?form=envar with operation write.
func (c *Client) AdminMobileAppDeviceEnvarWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "envar", "write", params)
}

// AdminMobileAppDeviceFixedWANPortGet calls /admin/mobile_app/device?form=fixed_wan_port with operation get.
func (c *Client) AdminMobileAppDeviceFixedWANPortGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "fixed_wan_port", "get", params)
}

// AdminMobileAppDeviceFixedWANPortSet calls /admin/mobile_app/device?form=fixed_wan_port with operation set.
func (c *Client) AdminMobileAppDeviceFixedWANPortSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "fixed_wan_port", "set", params)
}

// AdminMobileAppDeviceLEDRead calls /admin/mobile_app/device?form=led with operation read.
func (c *Client) AdminMobileAppDeviceLEDRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "led", "read", params)
}

// AdminMobileAppDeviceLEDWrite calls /admin/mobile_app/device?form=led with operation write.
func (c *Client) AdminMobileAppDeviceLEDWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "led", "write", params)
}

// AdminMobileAppDeviceRebootRead calls /admin/mobile_app/device?form=reboot with operation read.
func (c *Client) AdminMobileAppDeviceRebootRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "reboot", "read", params)
}

// AdminMobileAppDeviceRebootWrite calls /admin/mobile_app/device?form=reboot with operation write.
func (c *Client) AdminMobileAppDeviceRebootWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "reboot", "write", params)
}

// AdminMobileAppDeviceSignalLevelListGet calls /admin/mobile_app/device?form=signal_level_list with operation get.
func (c *Client) AdminMobileAppDeviceSignalLevelListGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "signal_level_list", "get", params)
}

// AdminMobileAppDeviceSpeedinfoRead calls /admin/mobile_app/device?form=speedinfo with operation read.
func (c *Client) AdminMobileAppDeviceSpeedinfoRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "speedinfo", "read", params)
}

// AdminMobileAppDeviceSpeedtestClear calls /admin/mobile_app/device?form=speedtest with operation clear.
func (c *Client) AdminMobileAppDeviceSpeedtestClear(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "speedtest", "clear", params)
}

// AdminMobileAppDeviceSpeedtestGet calls /admin/mobile_app/device?form=speedtest with operation get.
func (c *Client) AdminMobileAppDeviceSpeedtestGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "speedtest", "get", params)
}

// AdminMobileAppDeviceSpeedtestGetServer calls /admin/mobile_app/device?form=speedtest with operation get_server.
func (c *Client) AdminMobileAppDeviceSpeedtestGetServer(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "speedtest", "get_server", params)
}

// AdminMobileAppDeviceSpeedtestRead calls /admin/mobile_app/device?form=speedtest with operation read.
func (c *Client) AdminMobileAppDeviceSpeedtestRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "speedtest", "read", params)
}

// AdminMobileAppDeviceSpeedtestStop calls /admin/mobile_app/device?form=speedtest with operation stop.
func (c *Client) AdminMobileAppDeviceSpeedtestStop(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "speedtest", "stop", params)
}

// AdminMobileAppDeviceSpeedtestWrite calls /admin/mobile_app/device?form=speedtest with operation write.
func (c *Client) AdminMobileAppDeviceSpeedtestWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "speedtest", "write", params)
}

// AdminMobileAppDeviceSyncRead calls /admin/mobile_app/device?form=sync with operation read.
func (c *Client) AdminMobileAppDeviceSyncRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "sync", "read", params)
}

// AdminMobileAppDeviceSyncWrite calls /admin/mobile_app/device?form=sync with operation write.
func (c *Client) AdminMobileAppDeviceSyncWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "sync", "write", params)
}

// AdminMobileAppDeviceSysmodeRead calls /admin/mobile_app/device?form=sysmode with operation read.
func (c *Client) AdminMobileAppDeviceSysmodeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "sysmode", "read", params)
}

// AdminMobileAppDeviceSysmodeSetBackup calls /admin/mobile_app/device?form=sysmode with operation set_backup.
func (c *Client) AdminMobileAppDeviceSysmodeSetBackup(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "sysmode", "set_backup", params)
}

// AdminMobileAppDeviceSysmodeWrite calls /admin/mobile_app/device?form=sysmode with operation write.
func (c *Client) AdminMobileAppDeviceSysmodeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "sysmode", "write", params)
}

// AdminMobileAppDeviceSystemFactory calls /admin/mobile_app/device?form=system with operation factory.
func (c *Client) AdminMobileAppDeviceSystemFactory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "system", "factory", params)
}

// AdminMobileAppDeviceSystemGateway calls /admin/mobile_app/device?form=system with operation gateway.
func (c *Client) AdminMobileAppDeviceSystemGateway(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "system", "gateway", params)
}

// AdminMobileAppDeviceSystemReboot calls /admin/mobile_app/device?form=system with operation reboot.
func (c *Client) AdminMobileAppDeviceSystemReboot(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "system", "reboot", params)
}

// AdminMobileAppDeviceSystimeRead calls /admin/mobile_app/device?form=systime with operation read.
func (c *Client) AdminMobileAppDeviceSystimeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "systime", "read", params)
}

// AdminMobileAppDeviceSystimeWrite calls /admin/mobile_app/device?form=systime with operation write.
func (c *Client) AdminMobileAppDeviceSystimeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "systime", "write", params)
}

// AdminMobileAppDeviceTimesettingRead calls /admin/mobile_app/device?form=timesetting with operation read.
func (c *Client) AdminMobileAppDeviceTimesettingRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "timesetting", "read", params)
}

// AdminMobileAppDeviceTimesettingWrite calls /admin/mobile_app/device?form=timesetting with operation write.
func (c *Client) AdminMobileAppDeviceTimesettingWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/device", "timesetting", "write", params)
}

// AdminMobileAppDHCPDHCPAPRead calls /admin/mobile_app/dhcp?form=dhcp_ap with operation read.
func (c *Client) AdminMobileAppDHCPDHCPAPRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/dhcp", "dhcp_ap", "read", params)
}

// AdminMobileAppDHCPDHCPAPWrite calls /admin/mobile_app/dhcp?form=dhcp_ap with operation write.
func (c *Client) AdminMobileAppDHCPDHCPAPWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/dhcp", "dhcp_ap", "write", params)
}

// AdminMobileAppDHCPDHCPDialWrite calls /admin/mobile_app/dhcp?form=dhcp_dial with operation write.
func (c *Client) AdminMobileAppDHCPDHCPDialWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/dhcp", "dhcp_dial", "write", params)
}

// AdminMobileAppDHCPDHCPInfoRead calls /admin/mobile_app/dhcp?form=dhcp_info with operation read.
func (c *Client) AdminMobileAppDHCPDHCPInfoRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/dhcp", "dhcp_info", "read", params)
}

// AdminMobileAppDHCPDHCPInfoWrite calls /admin/mobile_app/dhcp?form=dhcp_info with operation write.
func (c *Client) AdminMobileAppDHCPDHCPInfoWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/dhcp", "dhcp_info", "write", params)
}

// AdminMobileAppGAInfoHWVerRead calls /admin/mobile_app/ga_info?form=hw_ver with operation read.
func (c *Client) AdminMobileAppGAInfoHWVerRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ga_info", "hw_ver", "read", params)
}

// AdminMobileAppGAInfoHWVerWrite calls /admin/mobile_app/ga_info?form=hw_ver with operation write.
func (c *Client) AdminMobileAppGAInfoHWVerWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ga_info", "hw_ver", "write", params)
}

// AdminMobileAppGAInfoStatusGet calls /admin/mobile_app/ga_info?form=status with operation get.
func (c *Client) AdminMobileAppGAInfoStatusGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ga_info", "status", "get", params)
}

// AdminMobileAppGAInfoStatusGetInternal calls /admin/mobile_app/ga_info?form=status with operation get_internal.
func (c *Client) AdminMobileAppGAInfoStatusGetInternal(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ga_info", "status", "get_internal", params)
}

// AdminMobileAppIoTAutomationIotautomationAddAction calls /admin/mobile_app/iot_automation?form=iotautomation with operation add_action.
func (c *Client) AdminMobileAppIoTAutomationIotautomationAddAction(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "add_action", params)
}

// AdminMobileAppIoTAutomationIotautomationAddTask calls /admin/mobile_app/iot_automation?form=iotautomation with operation add_task.
func (c *Client) AdminMobileAppIoTAutomationIotautomationAddTask(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "add_task", params)
}

// AdminMobileAppIoTAutomationIotautomationAddTrigger calls /admin/mobile_app/iot_automation?form=iotautomation with operation add_trigger.
func (c *Client) AdminMobileAppIoTAutomationIotautomationAddTrigger(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "add_trigger", params)
}

// AdminMobileAppIoTAutomationIotautomationGetHistory calls /admin/mobile_app/iot_automation?form=iotautomation with operation get_history.
func (c *Client) AdminMobileAppIoTAutomationIotautomationGetHistory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "get_history", params)
}

// AdminMobileAppIoTAutomationIotautomationGetTasklist calls /admin/mobile_app/iot_automation?form=iotautomation with operation get_tasklist.
func (c *Client) AdminMobileAppIoTAutomationIotautomationGetTasklist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "get_tasklist", params)
}

// AdminMobileAppIoTAutomationIotautomationModifyAction calls /admin/mobile_app/iot_automation?form=iotautomation with operation modify_action.
func (c *Client) AdminMobileAppIoTAutomationIotautomationModifyAction(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "modify_action", params)
}

// AdminMobileAppIoTAutomationIotautomationModifyTask calls /admin/mobile_app/iot_automation?form=iotautomation with operation modify_task.
func (c *Client) AdminMobileAppIoTAutomationIotautomationModifyTask(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "modify_task", params)
}

// AdminMobileAppIoTAutomationIotautomationModifyTaskList calls /admin/mobile_app/iot_automation?form=iotautomation with operation modify_task_list.
func (c *Client) AdminMobileAppIoTAutomationIotautomationModifyTaskList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "modify_task_list", params)
}

// AdminMobileAppIoTAutomationIotautomationModifyTrigger calls /admin/mobile_app/iot_automation?form=iotautomation with operation modify_trigger.
func (c *Client) AdminMobileAppIoTAutomationIotautomationModifyTrigger(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "modify_trigger", params)
}

// AdminMobileAppIoTAutomationIotautomationRemoveActionlist calls /admin/mobile_app/iot_automation?form=iotautomation with operation remove_actionlist.
func (c *Client) AdminMobileAppIoTAutomationIotautomationRemoveActionlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "remove_actionlist", params)
}

// AdminMobileAppIoTAutomationIotautomationRemoveHistory calls /admin/mobile_app/iot_automation?form=iotautomation with operation remove_history.
func (c *Client) AdminMobileAppIoTAutomationIotautomationRemoveHistory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "remove_history", params)
}

// AdminMobileAppIoTAutomationIotautomationRemoveTasklist calls /admin/mobile_app/iot_automation?form=iotautomation with operation remove_tasklist.
func (c *Client) AdminMobileAppIoTAutomationIotautomationRemoveTasklist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "remove_tasklist", params)
}

// AdminMobileAppIoTAutomationIotautomationRemoveTriggerlist calls /admin/mobile_app/iot_automation?form=iotautomation with operation remove_triggerlist.
func (c *Client) AdminMobileAppIoTAutomationIotautomationRemoveTriggerlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotautomation", "remove_triggerlist", params)
}

// AdminMobileAppIoTAutomationIotoneclickAddAction calls /admin/mobile_app/iot_automation?form=iotoneclick with operation add_action.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickAddAction(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "add_action", params)
}

// AdminMobileAppIoTAutomationIotoneclickAddScene calls /admin/mobile_app/iot_automation?form=iotoneclick with operation add_scene.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickAddScene(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "add_scene", params)
}

// AdminMobileAppIoTAutomationIotoneclickGetHistory calls /admin/mobile_app/iot_automation?form=iotoneclick with operation get_history.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickGetHistory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "get_history", params)
}

// AdminMobileAppIoTAutomationIotoneclickGetlist calls /admin/mobile_app/iot_automation?form=iotoneclick with operation getlist.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "getlist", params)
}

// AdminMobileAppIoTAutomationIotoneclickModifyAction calls /admin/mobile_app/iot_automation?form=iotoneclick with operation modify_action.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickModifyAction(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "modify_action", params)
}

// AdminMobileAppIoTAutomationIotoneclickModifyScene calls /admin/mobile_app/iot_automation?form=iotoneclick with operation modify_scene.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickModifyScene(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "modify_scene", params)
}

// AdminMobileAppIoTAutomationIotoneclickRemoveActionlist calls /admin/mobile_app/iot_automation?form=iotoneclick with operation remove_actionlist.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickRemoveActionlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "remove_actionlist", params)
}

// AdminMobileAppIoTAutomationIotoneclickRemoveHistory calls /admin/mobile_app/iot_automation?form=iotoneclick with operation remove_history.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickRemoveHistory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "remove_history", params)
}

// AdminMobileAppIoTAutomationIotoneclickRemoveScene calls /admin/mobile_app/iot_automation?form=iotoneclick with operation remove_scene.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickRemoveScene(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "remove_scene", params)
}

// AdminMobileAppIoTAutomationIotoneclickSet calls /admin/mobile_app/iot_automation?form=iotoneclick with operation set.
func (c *Client) AdminMobileAppIoTAutomationIotoneclickSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_automation", "iotoneclick", "set", params)
}

// AdminMobileAppIoTClientMeshClientMeshSet calls /admin/mobile_app/iot_client_mesh?form=client_mesh with operation set.
func (c *Client) AdminMobileAppIoTClientMeshClientMeshSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_client_mesh", "client_mesh", "set", params)
}

// AdminMobileAppIoTCloudIoTCloudReq calls /admin/mobile_app/iot_cloud (no form) with operation iot_cloud_req.
func (c *Client) AdminMobileAppIoTCloudIoTCloudReq(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_cloud", "", "iot_cloud_req", params)
}

// AdminMobileAppIoTDeviceIotdeviceAdd calls /admin/mobile_app/iot_device?form=iotdevice with operation add.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "add", params)
}

// AdminMobileAppIoTDeviceIotdeviceBeginScan calls /admin/mobile_app/iot_device?form=iotdevice with operation begin_scan.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceBeginScan(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "begin_scan", params)
}

// AdminMobileAppIoTDeviceIotdeviceEndScan calls /admin/mobile_app/iot_device?form=iotdevice with operation end_scan.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceEndScan(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "end_scan", params)
}

// AdminMobileAppIoTDeviceIotdeviceGet calls /admin/mobile_app/iot_device?form=iotdevice with operation get.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "get", params)
}

// AdminMobileAppIoTDeviceIotdeviceGetlist calls /admin/mobile_app/iot_device?form=iotdevice with operation getlist.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "getlist", params)
}

// AdminMobileAppIoTDeviceIotdeviceGetlistByMod calls /admin/mobile_app/iot_device?form=iotdevice with operation getlist_by_mod.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceGetlistByMod(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "getlist_by_mod", params)
}

// AdminMobileAppIoTDeviceIotdeviceIdentify calls /admin/mobile_app/iot_device?form=iotdevice with operation identify.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceIdentify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "identify", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerAct calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_act.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerAct(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_act", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerActQr calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_act_qr.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerActQr(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_act_qr", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerBeginAssoc calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_begin_assoc.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerBeginAssoc(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_begin_assoc", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerBeginScan calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_begin_scan.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerBeginScan(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_begin_scan", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerClientMgmt calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_client_mgmt.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerClientMgmt(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_client_mgmt", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerClientNetdevReq calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_client_netdev_req.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerClientNetdevReq(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_client_netdev_req", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerDHCPDGet calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_dhcpd_get.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerDHCPDGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_dhcpd_get", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerFilePull calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_file_pull.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerFilePull(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_file_pull", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerFilePush calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_file_push.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerFilePush(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_file_push", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerGet calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_get.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_get", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerGetlist calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_getlist.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_getlist", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerInquiry calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_inquiry.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerInquiry(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_inquiry", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerModify calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_modify.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_modify", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerRemove calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_remove.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_remove", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerScan calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_scan.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerScan(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_scan", params)
}

// AdminMobileAppIoTDeviceIotdeviceInnerUPnPGet calls /admin/mobile_app/iot_device?form=iotdevice with operation inner_upnp_get.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceInnerUPnPGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "inner_upnp_get", params)
}

// AdminMobileAppIoTDeviceIotdeviceModify calls /admin/mobile_app/iot_device?form=iotdevice with operation modify.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "modify", params)
}

// AdminMobileAppIoTDeviceIotdeviceRemove calls /admin/mobile_app/iot_device?form=iotdevice with operation remove.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "remove", params)
}

// AdminMobileAppIoTDeviceIotdeviceScan calls /admin/mobile_app/iot_device?form=iotdevice with operation scan.
func (c *Client) AdminMobileAppIoTDeviceIotdeviceScan(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotdevice", "scan", params)
}

// AdminMobileAppIoTDeviceIotownerGetlist calls /admin/mobile_app/iot_device?form=iotowner with operation getlist.
func (c *Client) AdminMobileAppIoTDeviceIotownerGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotowner", "getlist", params)
}

// AdminMobileAppIoTDeviceIotprofileGetAndUpdatePwd calls /admin/mobile_app/iot_device?form=iotprofile with operation get_and_update_pwd.
func (c *Client) AdminMobileAppIoTDeviceIotprofileGetAndUpdatePwd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotprofile", "get_and_update_pwd", params)
}

// AdminMobileAppIoTDeviceIotroleInnerProgramMgmt calls /admin/mobile_app/iot_device?form=iotrole with operation inner_program_mgmt.
func (c *Client) AdminMobileAppIoTDeviceIotroleInnerProgramMgmt(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotrole", "inner_program_mgmt", params)
}

// AdminMobileAppIoTDeviceIotspaceAdd calls /admin/mobile_app/iot_device?form=iotspace with operation add.
func (c *Client) AdminMobileAppIoTDeviceIotspaceAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotspace", "add", params)
}

// AdminMobileAppIoTDeviceIotspaceGetlist calls /admin/mobile_app/iot_device?form=iotspace with operation getlist.
func (c *Client) AdminMobileAppIoTDeviceIotspaceGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotspace", "getlist", params)
}

// AdminMobileAppIoTDeviceIotspaceModify calls /admin/mobile_app/iot_device?form=iotspace with operation modify.
func (c *Client) AdminMobileAppIoTDeviceIotspaceModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotspace", "modify", params)
}

// AdminMobileAppIoTDeviceIotspaceRemove calls /admin/mobile_app/iot_device?form=iotspace with operation remove.
func (c *Client) AdminMobileAppIoTDeviceIotspaceRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotspace", "remove", params)
}

// AdminMobileAppIoTDeviceIotspaceRemoveNetworkDevice calls /admin/mobile_app/iot_device?form=iotspace with operation remove_network_device.
func (c *Client) AdminMobileAppIoTDeviceIotspaceRemoveNetworkDevice(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotspace", "remove_network_device", params)
}

// AdminMobileAppIoTDeviceIotspaceSet calls /admin/mobile_app/iot_device?form=iotspace with operation set.
func (c *Client) AdminMobileAppIoTDeviceIotspaceSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotspace", "set", params)
}

// AdminMobileAppIoTDeviceIotspaceSetNetworkDevice calls /admin/mobile_app/iot_device?form=iotspace with operation set_network_device.
func (c *Client) AdminMobileAppIoTDeviceIotspaceSetNetworkDevice(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "iotspace", "set_network_device", params)
}

// AdminMobileAppIoTDeviceZigbeeFormNetwork calls /admin/mobile_app/iot_device?form=zigbee with operation form_network.
func (c *Client) AdminMobileAppIoTDeviceZigbeeFormNetwork(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "zigbee", "form_network", params)
}

// AdminMobileAppIoTDeviceZigbeeInnerNetwork calls /admin/mobile_app/iot_device?form=zigbee with operation inner_network.
func (c *Client) AdminMobileAppIoTDeviceZigbeeInnerNetwork(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "zigbee", "inner_network", params)
}

// AdminMobileAppIoTDeviceZigbeeInnerReportSensor calls /admin/mobile_app/iot_device?form=zigbee with operation inner_report_sensor.
func (c *Client) AdminMobileAppIoTDeviceZigbeeInnerReportSensor(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "zigbee", "inner_report_sensor", params)
}

// AdminMobileAppIoTDeviceZigbeeRemoveAll calls /admin/mobile_app/iot_device?form=zigbee with operation remove_all.
func (c *Client) AdminMobileAppIoTDeviceZigbeeRemoveAll(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "zigbee", "remove_all", params)
}

// AdminMobileAppIoTDeviceZigbeeTestcmd calls /admin/mobile_app/iot_device?form=zigbee with operation testcmd.
func (c *Client) AdminMobileAppIoTDeviceZigbeeTestcmd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iot_device", "zigbee", "testcmd", params)
}

// AdminMobileAppIPTVIPTVGet calls /admin/mobile_app/iptv?form=iptv with operation get.
func (c *Client) AdminMobileAppIPTVIPTVGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iptv", "iptv", "get", params)
}

// AdminMobileAppIPTVIPTVSet calls /admin/mobile_app/iptv?form=iptv with operation set.
func (c *Client) AdminMobileAppIPTVIPTVSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/iptv", "iptv", "set", params)
}

// AdminMobileAppIPv6FirewallClientRead calls /admin/mobile_app/ipv6_firewall?form=client with operation read.
func (c *Client) AdminMobileAppIPv6FirewallClientRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ipv6_firewall", "client", "read", params)
}

// AdminMobileAppIPv6FirewallFirewallModify calls /admin/mobile_app/ipv6_firewall?form=firewall with operation modify.
func (c *Client) AdminMobileAppIPv6FirewallFirewallModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ipv6_firewall", "firewall", "modify", params)
}

// AdminMobileAppIPv6FirewallFirewallRead calls /admin/mobile_app/ipv6_firewall?form=firewall with operation read.
func (c *Client) AdminMobileAppIPv6FirewallFirewallRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ipv6_firewall", "firewall", "read", params)
}

// AdminMobileAppIPv6FirewallFirewallRemove calls /admin/mobile_app/ipv6_firewall?form=firewall with operation remove.
func (c *Client) AdminMobileAppIPv6FirewallFirewallRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ipv6_firewall", "firewall", "remove", params)
}

// AdminMobileAppIPv6FirewallFirewallWrite calls /admin/mobile_app/ipv6_firewall?form=firewall with operation write.
func (c *Client) AdminMobileAppIPv6FirewallFirewallWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/ipv6_firewall", "firewall", "write", params)
}

// AdminMobileAppLogLoad calls /admin/mobile_app/log (no form) with operation load.
func (c *Client) AdminMobileAppLogLoad(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/log", "", "load", params)
}

// AdminMobileAppLogRead calls /admin/mobile_app/log (no form) with operation read.
func (c *Client) AdminMobileAppLogRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/log", "", "read", params)
}

// AdminMobileAppLogWrite calls /admin/mobile_app/log (no form) with operation write.
func (c *Client) AdminMobileAppLogWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/log", "", "write", params)
}

// AdminMobileAppLogExportFeedbackLogBuild calls /admin/mobile_app/log_export?form=feedback_log with operation build.
func (c *Client) AdminMobileAppLogExportFeedbackLogBuild(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/log_export", "feedback_log", "build", params)
}

// AdminMobileAppMsgServerCoordinatorNotify calls /admin/mobile_app/msg_server?form=coordinator with operation notify.
func (c *Client) AdminMobileAppMsgServerCoordinatorNotify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/msg_server", "coordinator", "notify", params)
}

// AdminMobileAppNATAlgRead calls /admin/mobile_app/nat?form=alg with operation read.
func (c *Client) AdminMobileAppNATAlgRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "alg", "read", params)
}

// AdminMobileAppNATAlgWrite calls /admin/mobile_app/nat?form=alg with operation write.
func (c *Client) AdminMobileAppNATAlgWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "alg", "write", params)
}

// AdminMobileAppNATDmzRead calls /admin/mobile_app/nat?form=dmz with operation read.
func (c *Client) AdminMobileAppNATDmzRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "dmz", "read", params)
}

// AdminMobileAppNATDmzWrite calls /admin/mobile_app/nat?form=dmz with operation write.
func (c *Client) AdminMobileAppNATDmzWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "dmz", "write", params)
}

// AdminMobileAppNATPtInsert calls /admin/mobile_app/nat?form=pt with operation insert.
func (c *Client) AdminMobileAppNATPtInsert(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "pt", "insert", params)
}

// AdminMobileAppNATPtLoad calls /admin/mobile_app/nat?form=pt with operation load.
func (c *Client) AdminMobileAppNATPtLoad(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "pt", "load", params)
}

// AdminMobileAppNATPtRemove calls /admin/mobile_app/nat?form=pt with operation remove.
func (c *Client) AdminMobileAppNATPtRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "pt", "remove", params)
}

// AdminMobileAppNATPtUpdate calls /admin/mobile_app/nat?form=pt with operation update.
func (c *Client) AdminMobileAppNATPtUpdate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "pt", "update", params)
}

// AdminMobileAppNATSettingRead calls /admin/mobile_app/nat?form=setting with operation read.
func (c *Client) AdminMobileAppNATSettingRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "setting", "read", params)
}

// AdminMobileAppNATSettingWrite calls /admin/mobile_app/nat?form=setting with operation write.
func (c *Client) AdminMobileAppNATSettingWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "setting", "write", params)
}

// AdminMobileAppNATSipAlgRead calls /admin/mobile_app/nat?form=sip_alg with operation read.
func (c *Client) AdminMobileAppNATSipAlgRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "sip_alg", "read", params)
}

// AdminMobileAppNATSipAlgWrite calls /admin/mobile_app/nat?form=sip_alg with operation write.
func (c *Client) AdminMobileAppNATSipAlgWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "sip_alg", "write", params)
}

// AdminMobileAppNATVsAdd calls /admin/mobile_app/nat?form=vs with operation add.
func (c *Client) AdminMobileAppNATVsAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "vs", "add", params)
}

// AdminMobileAppNATVsBatchRemove calls /admin/mobile_app/nat?form=vs with operation batch_remove.
func (c *Client) AdminMobileAppNATVsBatchRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "vs", "batch_remove", params)
}

// AdminMobileAppNATVsGetlist calls /admin/mobile_app/nat?form=vs with operation getlist.
func (c *Client) AdminMobileAppNATVsGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "vs", "getlist", params)
}

// AdminMobileAppNATVsModify calls /admin/mobile_app/nat?form=vs with operation modify.
func (c *Client) AdminMobileAppNATVsModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "vs", "modify", params)
}

// AdminMobileAppNATVsRead calls /admin/mobile_app/nat?form=vs with operation read.
func (c *Client) AdminMobileAppNATVsRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "vs", "read", params)
}

// AdminMobileAppNATVsRemove calls /admin/mobile_app/nat?form=vs with operation remove.
func (c *Client) AdminMobileAppNATVsRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nat", "vs", "remove", params)
}

// AdminMobileAppNetworkInternetRead calls /admin/mobile_app/network?form=internet with operation read.
func (c *Client) AdminMobileAppNetworkInternetRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "internet", "read", params)
}

// AdminMobileAppNetworkInternetWrite calls /admin/mobile_app/network?form=internet with operation write.
func (c *Client) AdminMobileAppNetworkInternetWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "internet", "write", params)
}

// AdminMobileAppNetworkIPv6Read calls /admin/mobile_app/network?form=ipv6 with operation read.
func (c *Client) AdminMobileAppNetworkIPv6Read(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "ipv6", "read", params)
}

// AdminMobileAppNetworkIPv6Write calls /admin/mobile_app/network?form=ipv6 with operation write.
func (c *Client) AdminMobileAppNetworkIPv6Write(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "ipv6", "write", params)
}

// AdminMobileAppNetworkLANBlockRead calls /admin/mobile_app/network?form=lan_block with operation read.
func (c *Client) AdminMobileAppNetworkLANBlockRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "lan_block", "read", params)
}

// AdminMobileAppNetworkLANBlockWrite calls /admin/mobile_app/network?form=lan_block with operation write.
func (c *Client) AdminMobileAppNetworkLANBlockWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "lan_block", "write", params)
}

// AdminMobileAppNetworkLANIPRead calls /admin/mobile_app/network?form=lan_ip with operation read.
func (c *Client) AdminMobileAppNetworkLANIPRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "lan_ip", "read", params)
}

// AdminMobileAppNetworkLANIPWrite calls /admin/mobile_app/network?form=lan_ip with operation write.
func (c *Client) AdminMobileAppNetworkLANIPWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "lan_ip", "write", params)
}

// AdminMobileAppNetworkLANIPv4Read calls /admin/mobile_app/network?form=lan_ipv4 with operation read.
func (c *Client) AdminMobileAppNetworkLANIPv4Read(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "lan_ipv4", "read", params)
}

// AdminMobileAppNetworkLANIPv4Write calls /admin/mobile_app/network?form=lan_ipv4 with operation write.
func (c *Client) AdminMobileAppNetworkLANIPv4Write(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "lan_ipv4", "write", params)
}

// AdminMobileAppNetworkMACCloneRead calls /admin/mobile_app/network?form=mac_clone with operation read.
func (c *Client) AdminMobileAppNetworkMACCloneRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "mac_clone", "read", params)
}

// AdminMobileAppNetworkMACCloneWrite calls /admin/mobile_app/network?form=mac_clone with operation write.
func (c *Client) AdminMobileAppNetworkMACCloneWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "mac_clone", "write", params)
}

// AdminMobileAppNetworkMACCloneListRead calls /admin/mobile_app/network?form=mac_clone_list with operation read.
func (c *Client) AdminMobileAppNetworkMACCloneListRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "mac_clone_list", "read", params)
}

// AdminMobileAppNetworkRoutesStaticAdd calls /admin/mobile_app/network?form=routes_static with operation add.
func (c *Client) AdminMobileAppNetworkRoutesStaticAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "routes_static", "add", params)
}

// AdminMobileAppNetworkRoutesStaticGetlist calls /admin/mobile_app/network?form=routes_static with operation getlist.
func (c *Client) AdminMobileAppNetworkRoutesStaticGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "routes_static", "getlist", params)
}

// AdminMobileAppNetworkRoutesStaticModify calls /admin/mobile_app/network?form=routes_static with operation modify.
func (c *Client) AdminMobileAppNetworkRoutesStaticModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "routes_static", "modify", params)
}

// AdminMobileAppNetworkRoutesStaticRemove calls /admin/mobile_app/network?form=routes_static with operation remove.
func (c *Client) AdminMobileAppNetworkRoutesStaticRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "routes_static", "remove", params)
}

// AdminMobileAppNetworkRoutesSystemGetlist calls /admin/mobile_app/network?form=routes_system with operation getlist.
func (c *Client) AdminMobileAppNetworkRoutesSystemGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "routes_system", "getlist", params)
}

// AdminMobileAppNetworkUPnPRead calls /admin/mobile_app/network?form=upnp with operation read.
func (c *Client) AdminMobileAppNetworkUPnPRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "upnp", "read", params)
}

// AdminMobileAppNetworkUPnPWrite calls /admin/mobile_app/network?form=upnp with operation write.
func (c *Client) AdminMobileAppNetworkUPnPWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "upnp", "write", params)
}

// AdminMobileAppNetworkVLANRead calls /admin/mobile_app/network?form=vlan with operation read.
func (c *Client) AdminMobileAppNetworkVLANRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "vlan", "read", params)
}

// AdminMobileAppNetworkVLANWrite calls /admin/mobile_app/network?form=vlan with operation write.
func (c *Client) AdminMobileAppNetworkVLANWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "vlan", "write", params)
}

// AdminMobileAppNetworkWANIPv4Connect calls /admin/mobile_app/network?form=wan_ipv4 with operation connect.
func (c *Client) AdminMobileAppNetworkWANIPv4Connect(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "wan_ipv4", "connect", params)
}

// AdminMobileAppNetworkWANIPv4Disconnect calls /admin/mobile_app/network?form=wan_ipv4 with operation disconnect.
func (c *Client) AdminMobileAppNetworkWANIPv4Disconnect(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "wan_ipv4", "disconnect", params)
}

// AdminMobileAppNetworkWANIPv4Read calls /admin/mobile_app/network?form=wan_ipv4 with operation read.
func (c *Client) AdminMobileAppNetworkWANIPv4Read(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "wan_ipv4", "read", params)
}

// AdminMobileAppNetworkWANIPv4Write calls /admin/mobile_app/network?form=wan_ipv4 with operation write.
func (c *Client) AdminMobileAppNetworkWANIPv4Write(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "wan_ipv4", "write", params)
}

// AdminMobileAppNetworkWANModeRead calls /admin/mobile_app/network?form=wan_mode with operation read.
func (c *Client) AdminMobileAppNetworkWANModeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "wan_mode", "read", params)
}

// AdminMobileAppNetworkWANModeWrite calls /admin/mobile_app/network?form=wan_mode with operation write.
func (c *Client) AdminMobileAppNetworkWANModeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network", "wan_mode", "write", params)
}

// AdminMobileAppNetworkOptimizeACSFilterMacsGet calls /admin/mobile_app/network_optimize?form=acs_filter_macs with operation get.
func (c *Client) AdminMobileAppNetworkOptimizeACSFilterMacsGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network_optimize", "acs_filter_macs", "get", params)
}

// AdminMobileAppNetworkOptimizeACSOptimizeRead calls /admin/mobile_app/network_optimize?form=acs_optimize with operation read.
func (c *Client) AdminMobileAppNetworkOptimizeACSOptimizeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network_optimize", "acs_optimize", "read", params)
}

// AdminMobileAppNetworkOptimizeACSOptimizeWrite calls /admin/mobile_app/network_optimize?form=acs_optimize with operation write.
func (c *Client) AdminMobileAppNetworkOptimizeACSOptimizeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/network_optimize", "acs_optimize", "write", params)
}

// AdminMobileAppNRDBlackListBlock calls /admin/mobile_app/nrd?form=black_list with operation block.
func (c *Client) AdminMobileAppNRDBlackListBlock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nrd", "black_list", "block", params)
}

// AdminMobileAppNRDBlackListList calls /admin/mobile_app/nrd?form=black_list with operation list.
func (c *Client) AdminMobileAppNRDBlackListList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nrd", "black_list", "list", params)
}

// AdminMobileAppNRDBlackListUnblock calls /admin/mobile_app/nrd?form=black_list with operation unblock.
func (c *Client) AdminMobileAppNRDBlackListUnblock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/nrd", "black_list", "unblock", params)
}

// AdminMobileAppPrivacyPolicyPrivacyPolicyRead calls /admin/mobile_app/privacy_policy?form=privacy_policy with operation read.
func (c *Client) AdminMobileAppPrivacyPolicyPrivacyPolicyRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/privacy_policy", "privacy_policy", "read", params)
}

// AdminMobileAppPrivacyPolicyPrivacyPolicyWrite calls /admin/mobile_app/privacy_policy?form=privacy_policy with operation write.
func (c *Client) AdminMobileAppPrivacyPolicyPrivacyPolicyWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/privacy_policy", "privacy_policy", "write", params)
}

// AdminMobileAppQuickSetupBatchdevicesRead calls /admin/mobile_app/quick_setup?form=batchdevices with operation read.
func (c *Client) AdminMobileAppQuickSetupBatchdevicesRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "batchdevices", "read", params)
}

// AdminMobileAppQuickSetupBatchdevicesWrite calls /admin/mobile_app/quick_setup?form=batchdevices with operation write.
func (c *Client) AdminMobileAppQuickSetupBatchdevicesWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "batchdevices", "write", params)
}

// AdminMobileAppQuickSetupBluetoothRead calls /admin/mobile_app/quick_setup?form=bluetooth with operation read.
func (c *Client) AdminMobileAppQuickSetupBluetoothRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "bluetooth", "read", params)
}

// AdminMobileAppQuickSetupBluetoothWrite calls /admin/mobile_app/quick_setup?form=bluetooth with operation write.
func (c *Client) AdminMobileAppQuickSetupBluetoothWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "bluetooth", "write", params)
}

// AdminMobileAppQuickSetupDcmpPreConfigPreConfigWrite calls /admin/mobile_app/quick_setup?form=dcmp_pre_config with operation pre_config_write.
func (c *Client) AdminMobileAppQuickSetupDcmpPreConfigPreConfigWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "dcmp_pre_config", "pre_config_write", params)
}

// AdminMobileAppQuickSetupDcmpPreConfigSavePreConfig calls /admin/mobile_app/quick_setup?form=dcmp_pre_config with operation save_pre_config.
func (c *Client) AdminMobileAppQuickSetupDcmpPreConfigSavePreConfig(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "dcmp_pre_config", "save_pre_config", params)
}

// AdminMobileAppQuickSetupEponymousDetectWrite calls /admin/mobile_app/quick_setup?form=eponymous_detect with operation write.
func (c *Client) AdminMobileAppQuickSetupEponymousDetectWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "eponymous_detect", "write", params)
}

// AdminMobileAppQuickSetupHeartbeatRead calls /admin/mobile_app/quick_setup?form=heartbeat with operation read.
func (c *Client) AdminMobileAppQuickSetupHeartbeatRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "heartbeat", "read", params)
}

// AdminMobileAppQuickSetupHeartbeatSync calls /admin/mobile_app/quick_setup?form=heartbeat with operation sync.
func (c *Client) AdminMobileAppQuickSetupHeartbeatSync(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "heartbeat", "sync", params)
}

// AdminMobileAppQuickSetupNewdeviceWrite calls /admin/mobile_app/quick_setup?form=newdevice with operation write.
func (c *Client) AdminMobileAppQuickSetupNewdeviceWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "newdevice", "write", params)
}

// AdminMobileAppQuickSetupNewgroupWrite calls /admin/mobile_app/quick_setup?form=newgroup with operation write.
func (c *Client) AdminMobileAppQuickSetupNewgroupWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "newgroup", "write", params)
}

// AdminMobileAppQuickSetupPreconfPreconfAdd calls /admin/mobile_app/quick_setup?form=preconf with operation preconf_add.
func (c *Client) AdminMobileAppQuickSetupPreconfPreconfAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "preconf", "preconf_add", params)
}

// AdminMobileAppQuickSetupPreconfPreconfCheck calls /admin/mobile_app/quick_setup?form=preconf with operation preconf_check.
func (c *Client) AdminMobileAppQuickSetupPreconfPreconfCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "preconf", "preconf_check", params)
}

// AdminMobileAppQuickSetupPreconfPreconfRead calls /admin/mobile_app/quick_setup?form=preconf with operation preconf_read.
func (c *Client) AdminMobileAppQuickSetupPreconfPreconfRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "preconf", "preconf_read", params)
}

// AdminMobileAppQuickSetupPreconfPreconfWrite calls /admin/mobile_app/quick_setup?form=preconf with operation preconf_write.
func (c *Client) AdminMobileAppQuickSetupPreconfPreconfWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/quick_setup", "preconf", "preconf_write", params)
}

// AdminMobileAppSecurityCategoryRead calls /admin/mobile_app/security?form=category with operation read.
func (c *Client) AdminMobileAppSecurityCategoryRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/security", "category", "read", params)
}

// AdminMobileAppSecurityHistoryClear calls /admin/mobile_app/security?form=history with operation clear.
func (c *Client) AdminMobileAppSecurityHistoryClear(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/security", "history", "clear", params)
}

// AdminMobileAppSecurityHistoryGet calls /admin/mobile_app/security?form=history with operation get.
func (c *Client) AdminMobileAppSecurityHistoryGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/security", "history", "get", params)
}

// AdminMobileAppSecurityHistoryRemove calls /admin/mobile_app/security?form=history with operation remove.
func (c *Client) AdminMobileAppSecurityHistoryRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/security", "history", "remove", params)
}

// AdminMobileAppSecurityInfoRead calls /admin/mobile_app/security?form=info with operation read.
func (c *Client) AdminMobileAppSecurityInfoRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/security", "info", "read", params)
}

// AdminMobileAppSecurityInfoWrite calls /admin/mobile_app/security?form=info with operation write.
func (c *Client) AdminMobileAppSecurityInfoWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/security", "info", "write", params)
}

// AdminMobileAppSecurityRuleRead calls /admin/mobile_app/security?form=rule with operation read.
func (c *Client) AdminMobileAppSecurityRuleRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/security", "rule", "read", params)
}

// AdminMobileAppSecurityRuleUpdate calls /admin/mobile_app/security?form=rule with operation update.
func (c *Client) AdminMobileAppSecurityRuleUpdate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/security", "rule", "update", params)
}

// AdminMobileAppSmartNetworkAppBlockListGet calls /admin/mobile_app/smart_network?form=app_block_list with operation get.
func (c *Client) AdminMobileAppSmartNetworkAppBlockListGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "app_block_list", "get", params)
}

// AdminMobileAppSmartNetworkAppDPITimeLimitAdd calls /admin/mobile_app/smart_network?form=app_dpi with operation time_limit_add.
func (c *Client) AdminMobileAppSmartNetworkAppDPITimeLimitAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "app_dpi", "time_limit_add", params)
}

// AdminMobileAppSmartNetworkAppDPITimeLimitModify calls /admin/mobile_app/smart_network?form=app_dpi with operation time_limit_modify.
func (c *Client) AdminMobileAppSmartNetworkAppDPITimeLimitModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "app_dpi", "time_limit_modify", params)
}

// AdminMobileAppSmartNetworkAppDPITimeLimitRemove calls /admin/mobile_app/smart_network?form=app_dpi with operation time_limit_remove.
func (c *Client) AdminMobileAppSmartNetworkAppDPITimeLimitRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "app_dpi", "time_limit_remove", params)
}

// AdminMobileAppSmartNetworkAppQoSGet calls /admin/mobile_app/smart_network?form=app_qos with operation get.
func (c *Client) AdminMobileAppSmartNetworkAppQoSGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "app_qos", "get", params)
}

// AdminMobileAppSmartNetworkAppQoSRemove calls /admin/mobile_app/smart_network?form=app_qos with operation remove.
func (c *Client) AdminMobileAppSmartNetworkAppQoSRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "app_qos", "remove", params)
}

// AdminMobileAppSmartNetworkAppQoSSet calls /admin/mobile_app/smart_network?form=app_qos with operation set.
func (c *Client) AdminMobileAppSmartNetworkAppQoSSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "app_qos", "set", params)
}

// AdminMobileAppSmartNetworkBandwidthGet calls /admin/mobile_app/smart_network?form=bandwidth with operation get.
func (c *Client) AdminMobileAppSmartNetworkBandwidthGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "bandwidth", "get", params)
}

// AdminMobileAppSmartNetworkBandwidthSet calls /admin/mobile_app/smart_network?form=bandwidth with operation set.
func (c *Client) AdminMobileAppSmartNetworkBandwidthSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "bandwidth", "set", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlAddAppTimeLimit calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation addAppTimeLimit.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlAddAppTimeLimit(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "addAppTimeLimit", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlAddOwnerInList calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation addOwnerInList.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlAddOwnerInList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "addOwnerInList", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlAddWhiteList calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation addWhiteList.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlAddWhiteList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "addWhiteList", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlDelActHistory calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation delActHistory.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlDelActHistory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "delActHistory", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlDelAppTimeLimit calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation delAppTimeLimit.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlDelAppTimeLimit(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "delAppTimeLimit", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlDelOwnerInList calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation delOwnerInList.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlDelOwnerInList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "delOwnerInList", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlDelWhiteList calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation delWhiteList.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlDelWhiteList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "delWhiteList", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlGetActHistory calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation getActHistory.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlGetActHistory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "getActHistory", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlGetAppBlockList calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation getAppBlockList.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlGetAppBlockList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "getAppBlockList", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlGetInsightData calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation getInsightData.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlGetInsightData(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "getInsightData", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlGetOwner calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation getOwner.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlGetOwner(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "getOwner", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlGetOwnerInList calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation getOwnerInList.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlGetOwnerInList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "getOwnerInList", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlIgnoreReq calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation ignoreReq.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlIgnoreReq(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "ignoreReq", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlModifyAppTimeLimit calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation modifyAppTimeLimit.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlModifyAppTimeLimit(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "modifyAppTimeLimit", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlModifyBaseInfo calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation modifyBaseInfo.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlModifyBaseInfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "modifyBaseInfo", params)
}

// AdminMobileAppSmartNetworkHCUpgPctlSetBonusTime calls /admin/mobile_app/smart_network?form=hc_upg_pctl with operation setBonusTime.
func (c *Client) AdminMobileAppSmartNetworkHCUpgPctlSetBonusTime(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "hc_upg_pctl", "setBonusTime", params)
}

// AdminMobileAppSmartNetworkPatrolCliAdd calls /admin/mobile_app/smart_network?form=patrol_cli with operation add.
func (c *Client) AdminMobileAppSmartNetworkPatrolCliAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_cli", "add", params)
}

// AdminMobileAppSmartNetworkPatrolCliDel calls /admin/mobile_app/smart_network?form=patrol_cli with operation del.
func (c *Client) AdminMobileAppSmartNetworkPatrolCliDel(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_cli", "del", params)
}

// AdminMobileAppSmartNetworkPatrolFilterBlock calls /admin/mobile_app/smart_network?form=patrol_filter with operation block.
func (c *Client) AdminMobileAppSmartNetworkPatrolFilterBlock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_filter", "block", params)
}

// AdminMobileAppSmartNetworkPatrolFilterList calls /admin/mobile_app/smart_network?form=patrol_filter with operation list.
func (c *Client) AdminMobileAppSmartNetworkPatrolFilterList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_filter", "list", params)
}

// AdminMobileAppSmartNetworkPatrolFilterTable calls /admin/mobile_app/smart_network?form=patrol_filter with operation table.
func (c *Client) AdminMobileAppSmartNetworkPatrolFilterTable(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_filter", "table", params)
}

// AdminMobileAppSmartNetworkPatrolFilterUnblock calls /admin/mobile_app/smart_network?form=patrol_filter with operation unblock.
func (c *Client) AdminMobileAppSmartNetworkPatrolFilterUnblock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_filter", "unblock", params)
}

// AdminMobileAppSmartNetworkPatrolInsightsGet calls /admin/mobile_app/smart_network?form=patrol_insights with operation get.
func (c *Client) AdminMobileAppSmartNetworkPatrolInsightsGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_insights", "get", params)
}

// AdminMobileAppSmartNetworkPatrolInsightsHistory calls /admin/mobile_app/smart_network?form=patrol_insights with operation history.
func (c *Client) AdminMobileAppSmartNetworkPatrolInsightsHistory(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_insights", "history", params)
}

// AdminMobileAppSmartNetworkPatrolInsightsRemove calls /admin/mobile_app/smart_network?form=patrol_insights with operation remove.
func (c *Client) AdminMobileAppSmartNetworkPatrolInsightsRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_insights", "remove", params)
}

// AdminMobileAppSmartNetworkPatrolOwnerAdd calls /admin/mobile_app/smart_network?form=patrol_owner with operation add.
func (c *Client) AdminMobileAppSmartNetworkPatrolOwnerAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_owner", "add", params)
}

// AdminMobileAppSmartNetworkPatrolOwnerBlock calls /admin/mobile_app/smart_network?form=patrol_owner with operation block.
func (c *Client) AdminMobileAppSmartNetworkPatrolOwnerBlock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_owner", "block", params)
}

// AdminMobileAppSmartNetworkPatrolOwnerDel calls /admin/mobile_app/smart_network?form=patrol_owner with operation del.
func (c *Client) AdminMobileAppSmartNetworkPatrolOwnerDel(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_owner", "del", params)
}

// AdminMobileAppSmartNetworkPatrolOwnerGet calls /admin/mobile_app/smart_network?form=patrol_owner with operation get.
func (c *Client) AdminMobileAppSmartNetworkPatrolOwnerGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_owner", "get", params)
}

// AdminMobileAppSmartNetworkPatrolOwnerList calls /admin/mobile_app/smart_network?form=patrol_owner with operation list.
func (c *Client) AdminMobileAppSmartNetworkPatrolOwnerList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_owner", "list", params)
}

// AdminMobileAppSmartNetworkPatrolOwnerSet calls /admin/mobile_app/smart_network?form=patrol_owner with operation set.
func (c *Client) AdminMobileAppSmartNetworkPatrolOwnerSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_owner", "set", params)
}

// AdminMobileAppSmartNetworkPatrolOwnerAvatarGet calls /admin/mobile_app/smart_network?form=patrol_owner_avatar with operation get.
func (c *Client) AdminMobileAppSmartNetworkPatrolOwnerAvatarGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_owner_avatar", "get", params)
}

// AdminMobileAppSmartNetworkPatrolOwnerAvatarSet calls /admin/mobile_app/smart_network?form=patrol_owner_avatar with operation set.
func (c *Client) AdminMobileAppSmartNetworkPatrolOwnerAvatarSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "patrol_owner_avatar", "set", params)
}

// AdminMobileAppSmartNetworkTMQoSRead calls /admin/mobile_app/smart_network?form=tm_qos with operation read.
func (c *Client) AdminMobileAppSmartNetworkTMQoSRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tm_qos", "read", params)
}

// AdminMobileAppSmartNetworkTMQoSWrite calls /admin/mobile_app/smart_network?form=tm_qos with operation write.
func (c *Client) AdminMobileAppSmartNetworkTMQoSWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tm_qos", "write", params)
}

// AdminMobileAppSmartNetworkTmpAviraAddOwnerInList calls /admin/mobile_app/smart_network?form=tmp_avira with operation addOwnerInList.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraAddOwnerInList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "addOwnerInList", params)
}

// AdminMobileAppSmartNetworkTmpAviraAddAllowedWebsites calls /admin/mobile_app/smart_network?form=tmp_avira with operation add_allowed_websites.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraAddAllowedWebsites(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "add_allowed_websites", params)
}

// AdminMobileAppSmartNetworkTmpAviraAddSecWhitelist calls /admin/mobile_app/smart_network?form=tmp_avira with operation add_sec_whitelist.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraAddSecWhitelist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "add_sec_whitelist", params)
}

// AdminMobileAppSmartNetworkTmpAviraBonusTimeSet calls /admin/mobile_app/smart_network?form=tmp_avira with operation bonusTimeSet.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraBonusTimeSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "bonusTimeSet", params)
}

// AdminMobileAppSmartNetworkTmpAviraCloudServiceStateCheck calls /admin/mobile_app/smart_network?form=tmp_avira with operation cloud_service_state_check.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraCloudServiceStateCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "cloud_service_state_check", params)
}

// AdminMobileAppSmartNetworkTmpAviraDelOwnerInList calls /admin/mobile_app/smart_network?form=tmp_avira with operation delOwnerInList.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraDelOwnerInList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "delOwnerInList", params)
}

// AdminMobileAppSmartNetworkTmpAviraGetFamilyTimeInfo calls /admin/mobile_app/smart_network?form=tmp_avira with operation getFamilyTimeInfo.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraGetFamilyTimeInfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "getFamilyTimeInfo", params)
}

// AdminMobileAppSmartNetworkTmpAviraGetInsightData calls /admin/mobile_app/smart_network?form=tmp_avira with operation getInsightData.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraGetInsightData(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "getInsightData", params)
}

// AdminMobileAppSmartNetworkTmpAviraGetOwnerInList calls /admin/mobile_app/smart_network?form=tmp_avira with operation getOwnerInList.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraGetOwnerInList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "getOwnerInList", params)
}

// AdminMobileAppSmartNetworkTmpAviraGetTodayInsightTimeUsagePro calls /admin/mobile_app/smart_network?form=tmp_avira with operation getTodayInsightTimeUsagePro.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraGetTodayInsightTimeUsagePro(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "getTodayInsightTimeUsagePro", params)
}

// AdminMobileAppSmartNetworkTmpAviraGetAllowedWebsites calls /admin/mobile_app/smart_network?form=tmp_avira with operation get_allowed_websites.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraGetAllowedWebsites(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "get_allowed_websites", params)
}

// AdminMobileAppSmartNetworkTmpAviraGetSecWhitelist calls /admin/mobile_app/smart_network?form=tmp_avira with operation get_sec_whitelist.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraGetSecWhitelist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "get_sec_whitelist", params)
}

// AdminMobileAppSmartNetworkTmpAviraIgnoreReq calls /admin/mobile_app/smart_network?form=tmp_avira with operation ignoreReq.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraIgnoreReq(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "ignoreReq", params)
}

// AdminMobileAppSmartNetworkTmpAviraISPGet calls /admin/mobile_app/smart_network?form=tmp_avira with operation ispGet.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraISPGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "ispGet", params)
}

// AdminMobileAppSmartNetworkTmpAviraModifyBaseInfo calls /admin/mobile_app/smart_network?form=tmp_avira with operation modifyBaseInfo.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraModifyBaseInfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "modifyBaseInfo", params)
}

// AdminMobileAppSmartNetworkTmpAviraNetworkQualityStartOptimize calls /admin/mobile_app/smart_network?form=tmp_avira with operation networkQualityStartOptimize.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraNetworkQualityStartOptimize(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "networkQualityStartOptimize", params)
}

// AdminMobileAppSmartNetworkTmpAviraOwnerClientListSet calls /admin/mobile_app/smart_network?form=tmp_avira with operation ownerClientListSet.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraOwnerClientListSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "ownerClientListSet", params)
}

// AdminMobileAppSmartNetworkTmpAviraRemoveAllowedWebsites calls /admin/mobile_app/smart_network?form=tmp_avira with operation remove_allowed_websites.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraRemoveAllowedWebsites(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "remove_allowed_websites", params)
}

// AdminMobileAppSmartNetworkTmpAviraRemoveSecWhitelist calls /admin/mobile_app/smart_network?form=tmp_avira with operation remove_sec_whitelist.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraRemoveSecWhitelist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "remove_sec_whitelist", params)
}

// AdminMobileAppSmartNetworkTmpAviraScanGet calls /admin/mobile_app/smart_network?form=tmp_avira with operation scanGet.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraScanGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "scanGet", params)
}

// AdminMobileAppSmartNetworkTmpAviraScanStart calls /admin/mobile_app/smart_network?form=tmp_avira with operation scanStart.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraScanStart(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "scanStart", params)
}

// AdminMobileAppSmartNetworkTmpAviraScanStop calls /admin/mobile_app/smart_network?form=tmp_avira with operation scanStop.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraScanStop(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "scanStop", params)
}

// AdminMobileAppSmartNetworkTmpAviraServiceStateCheck calls /admin/mobile_app/smart_network?form=tmp_avira with operation service_state_check.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraServiceStateCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "service_state_check", params)
}

// AdminMobileAppSmartNetworkTmpAviraServiceStatusCheck calls /admin/mobile_app/smart_network?form=tmp_avira with operation service_status_check.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraServiceStatusCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "service_status_check", params)
}

// AdminMobileAppSmartNetworkTmpAviraSetFamilyTimeInfo calls /admin/mobile_app/smart_network?form=tmp_avira with operation setFamilyTimeInfo.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraSetFamilyTimeInfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "setFamilyTimeInfo", params)
}

// AdminMobileAppSmartNetworkTmpAviraTmpGetSecinfo calls /admin/mobile_app/smart_network?form=tmp_avira with operation tmp_get_secinfo.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraTmpGetSecinfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "tmp_get_secinfo", params)
}

// AdminMobileAppSmartNetworkTmpAviraTmpGetSecv2info calls /admin/mobile_app/smart_network?form=tmp_avira with operation tmp_get_secv2info.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraTmpGetSecv2info(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "tmp_get_secv2info", params)
}

// AdminMobileAppSmartNetworkTmpAviraTmpSetSecinfo calls /admin/mobile_app/smart_network?form=tmp_avira with operation tmp_set_secinfo.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraTmpSetSecinfo(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "tmp_set_secinfo", params)
}

// AdminMobileAppSmartNetworkTmpAviraTmpSetSecv2info calls /admin/mobile_app/smart_network?form=tmp_avira with operation tmp_set_secv2info.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraTmpSetSecv2info(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "tmp_set_secv2info", params)
}

// AdminMobileAppSmartNetworkTmpAviraWhiteListAdd calls /admin/mobile_app/smart_network?form=tmp_avira with operation white_list_add.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraWhiteListAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "white_list_add", params)
}

// AdminMobileAppSmartNetworkTmpAviraWhiteListRemove calls /admin/mobile_app/smart_network?form=tmp_avira with operation white_list_remove.
func (c *Client) AdminMobileAppSmartNetworkTmpAviraWhiteListRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "tmp_avira", "white_list_remove", params)
}

// AdminMobileAppSmartNetworkWhiteListAdd calls /admin/mobile_app/smart_network?form=white_list with operation add.
func (c *Client) AdminMobileAppSmartNetworkWhiteListAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "white_list", "add", params)
}

// AdminMobileAppSmartNetworkWhiteListGet calls /admin/mobile_app/smart_network?form=white_list with operation get.
func (c *Client) AdminMobileAppSmartNetworkWhiteListGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "white_list", "get", params)
}

// AdminMobileAppSmartNetworkWhiteListRemove calls /admin/mobile_app/smart_network?form=white_list with operation remove.
func (c *Client) AdminMobileAppSmartNetworkWhiteListRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/smart_network", "white_list", "remove", params)
}

// AdminMobileAppVPNClientConfigRead calls /admin/mobile_app/vpn_client?form=config with operation read.
func (c *Client) AdminMobileAppVPNClientConfigRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_client", "config", "read", params)
}

// AdminMobileAppVPNClientConfigWrite calls /admin/mobile_app/vpn_client?form=config with operation write.
func (c *Client) AdminMobileAppVPNClientConfigWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_client", "config", "write", params)
}

// AdminMobileAppVPNClientServerInsert calls /admin/mobile_app/vpn_client?form=server with operation insert.
func (c *Client) AdminMobileAppVPNClientServerInsert(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_client", "server", "insert", params)
}

// AdminMobileAppVPNClientServerLoad calls /admin/mobile_app/vpn_client?form=server with operation load.
func (c *Client) AdminMobileAppVPNClientServerLoad(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_client", "server", "load", params)
}

// AdminMobileAppVPNClientServerRemove calls /admin/mobile_app/vpn_client?form=server with operation remove.
func (c *Client) AdminMobileAppVPNClientServerRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_client", "server", "remove", params)
}

// AdminMobileAppVPNClientServerUpdate calls /admin/mobile_app/vpn_client?form=server with operation update.
func (c *Client) AdminMobileAppVPNClientServerUpdate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_client", "server", "update", params)
}

// AdminMobileAppVPNServerAccountsInsert calls /admin/mobile_app/vpn_server?form=accounts with operation insert.
func (c *Client) AdminMobileAppVPNServerAccountsInsert(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "accounts", "insert", params)
}

// AdminMobileAppVPNServerAccountsRemove calls /admin/mobile_app/vpn_server?form=accounts with operation remove.
func (c *Client) AdminMobileAppVPNServerAccountsRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "accounts", "remove", params)
}

// AdminMobileAppVPNServerAccountsUpdate calls /admin/mobile_app/vpn_server?form=accounts with operation update.
func (c *Client) AdminMobileAppVPNServerAccountsUpdate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "accounts", "update", params)
}

// AdminMobileAppVPNServerCertCheck calls /admin/mobile_app/vpn_server?form=cert with operation check.
func (c *Client) AdminMobileAppVPNServerCertCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "cert", "check", params)
}

// AdminMobileAppVPNServerCertSet calls /admin/mobile_app/vpn_server?form=cert with operation set.
func (c *Client) AdminMobileAppVPNServerCertSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "cert", "set", params)
}

// AdminMobileAppVPNServerServerInsert calls /admin/mobile_app/vpn_server?form=server with operation insert.
func (c *Client) AdminMobileAppVPNServerServerInsert(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "server", "insert", params)
}

// AdminMobileAppVPNServerServerRead calls /admin/mobile_app/vpn_server?form=server with operation read.
func (c *Client) AdminMobileAppVPNServerServerRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "server", "read", params)
}

// AdminMobileAppVPNServerServerRemove calls /admin/mobile_app/vpn_server?form=server with operation remove.
func (c *Client) AdminMobileAppVPNServerServerRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "server", "remove", params)
}

// AdminMobileAppVPNServerServerUpdate calls /admin/mobile_app/vpn_server?form=server with operation update.
func (c *Client) AdminMobileAppVPNServerServerUpdate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "server", "update", params)
}

// AdminMobileAppVPNServerSingleRead calls /admin/mobile_app/vpn_server?form=single with operation read.
func (c *Client) AdminMobileAppVPNServerSingleRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpn_server", "single", "read", params)
}

// AdminMobileAppVpnconnCertSync calls /admin/mobile_app/vpnconn?form=cert with operation sync.
func (c *Client) AdminMobileAppVpnconnCertSync(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpnconn", "cert", "sync", params)
}

// AdminMobileAppVpnconnConnDisconnect calls /admin/mobile_app/vpnconn?form=conn with operation disconnect.
func (c *Client) AdminMobileAppVpnconnConnDisconnect(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpnconn", "conn", "disconnect", params)
}

// AdminMobileAppVpnconnConnList calls /admin/mobile_app/vpnconn?form=conn with operation list.
func (c *Client) AdminMobileAppVpnconnConnList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/vpnconn", "conn", "list", params)
}

// AdminMobileAppWirelessBandwidthEnhanceRead calls /admin/mobile_app/wireless?form=bandwidth_enhance with operation read.
func (c *Client) AdminMobileAppWirelessBandwidthEnhanceRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "bandwidth_enhance", "read", params)
}

// AdminMobileAppWirelessBandwidthEnhanceWrite calls /admin/mobile_app/wireless?form=bandwidth_enhance with operation write.
func (c *Client) AdminMobileAppWirelessBandwidthEnhanceWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "bandwidth_enhance", "write", params)
}

// AdminMobileAppWirelessBandwidthSwitchCheck calls /admin/mobile_app/wireless?form=bandwidth_switch with operation check.
func (c *Client) AdminMobileAppWirelessBandwidthSwitchCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "bandwidth_switch", "check", params)
}

// AdminMobileAppWirelessBandwidthSwitchRead calls /admin/mobile_app/wireless?form=bandwidth_switch with operation read.
func (c *Client) AdminMobileAppWirelessBandwidthSwitchRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "bandwidth_switch", "read", params)
}

// AdminMobileAppWirelessBandwidthSwitchWrite calls /admin/mobile_app/wireless?form=bandwidth_switch with operation write.
func (c *Client) AdminMobileAppWirelessBandwidthSwitchWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "bandwidth_switch", "write", params)
}

// AdminMobileAppWirelessBeamformingRead calls /admin/mobile_app/wireless?form=beamforming with operation read.
func (c *Client) AdminMobileAppWirelessBeamformingRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "beamforming", "read", params)
}

// AdminMobileAppWirelessBeamformingWrite calls /admin/mobile_app/wireless?form=beamforming with operation write.
func (c *Client) AdminMobileAppWirelessBeamformingWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "beamforming", "write", params)
}

// AdminMobileAppWirelessBridgeCheck calls /admin/mobile_app/wireless?form=bridge with operation check.
func (c *Client) AdminMobileAppWirelessBridgeCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "bridge", "check", params)
}

// AdminMobileAppWirelessBridgeRead calls /admin/mobile_app/wireless?form=bridge with operation read.
func (c *Client) AdminMobileAppWirelessBridgeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "bridge", "read", params)
}

// AdminMobileAppWirelessDedicatedBackhaulRead calls /admin/mobile_app/wireless?form=dedicated_backhaul with operation read.
func (c *Client) AdminMobileAppWirelessDedicatedBackhaulRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "dedicated_backhaul", "read", params)
}

// AdminMobileAppWirelessDedicatedBackhaulWrite calls /admin/mobile_app/wireless?form=dedicated_backhaul with operation write.
func (c *Client) AdminMobileAppWirelessDedicatedBackhaulWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "dedicated_backhaul", "write", params)
}

// AdminMobileAppWirelessIeee80211rRead calls /admin/mobile_app/wireless?form=ieee80211r with operation read.
func (c *Client) AdminMobileAppWirelessIeee80211rRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "ieee80211r", "read", params)
}

// AdminMobileAppWirelessIeee80211rWrite calls /admin/mobile_app/wireless?form=ieee80211r with operation write.
func (c *Client) AdminMobileAppWirelessIeee80211rWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "ieee80211r", "write", params)
}

// AdminMobileAppWirelessOperationModeRead calls /admin/mobile_app/wireless?form=operation_mode with operation read.
func (c *Client) AdminMobileAppWirelessOperationModeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "operation_mode", "read", params)
}

// AdminMobileAppWirelessOperationModeWrite calls /admin/mobile_app/wireless?form=operation_mode with operation write.
func (c *Client) AdminMobileAppWirelessOperationModeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "operation_mode", "write", params)
}

// AdminMobileAppWirelessPowerWrite calls /admin/mobile_app/wireless?form=power with operation write.
func (c *Client) AdminMobileAppWirelessPowerWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "power", "write", params)
}

// AdminMobileAppWirelessSmartAntennaRead calls /admin/mobile_app/wireless?form=smart_antenna with operation read.
func (c *Client) AdminMobileAppWirelessSmartAntennaRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "smart_antenna", "read", params)
}

// AdminMobileAppWirelessSmartAntennaWrite calls /admin/mobile_app/wireless?form=smart_antenna with operation write.
func (c *Client) AdminMobileAppWirelessSmartAntennaWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "smart_antenna", "write", params)
}

// AdminMobileAppWirelessWiFiScheduleRead calls /admin/mobile_app/wireless?form=wifi_schedule with operation read.
func (c *Client) AdminMobileAppWirelessWiFiScheduleRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "wifi_schedule", "read", params)
}

// AdminMobileAppWirelessWiFiScheduleWrite calls /admin/mobile_app/wireless?form=wifi_schedule with operation write.
func (c *Client) AdminMobileAppWirelessWiFiScheduleWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "wifi_schedule", "write", params)
}

// AdminMobileAppWirelessWLANRead calls /admin/mobile_app/wireless?form=wlan with operation read.
func (c *Client) AdminMobileAppWirelessWLANRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "wlan", "read", params)
}

// AdminMobileAppWirelessWLANWrite calls /admin/mobile_app/wireless?form=wlan with operation write.
func (c *Client) AdminMobileAppWirelessWLANWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wireless", "wlan", "write", params)
}

// AdminMobileAppWPSStateSet calls /admin/mobile_app/wps?form=state with operation set.
func (c *Client) AdminMobileAppWPSStateSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wps", "state", "set", params)
}

// AdminMobileAppWPSStatusGet calls /admin/mobile_app/wps?form=status with operation get.
func (c *Client) AdminMobileAppWPSStatusGet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/mobile_app/wps", "status", "get", params)
}
