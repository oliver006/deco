package deco

import "testing"

func TestAdminDeviceMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminDeviceDeviceListRead",
			path:      "/admin/device",
			form:      "device_list",
			operation: "read",
			call:      (*Client).AdminDeviceDeviceListRead,
		},
		{
			name:      "AdminDeviceDeviceListRemove",
			path:      "/admin/device",
			form:      "device_list",
			operation: "remove",
			call:      (*Client).AdminDeviceDeviceListRemove,
		},
		{
			name:      "AdminDeviceMiniDeviceListRead",
			path:      "/admin/device",
			form:      "mini_device_list",
			operation: "read",
			call:      (*Client).AdminDeviceMiniDeviceListRead,
		},
		{
			name:      "AdminDeviceModeRead",
			path:      "/admin/device",
			form:      "mode",
			operation: "read",
			call:      (*Client).AdminDeviceModeRead,
		},
		{
			name:      "AdminDeviceSpeedtestClear",
			path:      "/admin/device",
			form:      "speedtest",
			operation: "clear",
			call:      (*Client).AdminDeviceSpeedtestClear,
		},
		{
			name:      "AdminDeviceSpeedtestGet",
			path:      "/admin/device",
			form:      "speedtest",
			operation: "get",
			call:      (*Client).AdminDeviceSpeedtestGet,
		},
		{
			name:      "AdminDeviceSpeedtestGetServer",
			path:      "/admin/device",
			form:      "speedtest",
			operation: "get_server",
			call:      (*Client).AdminDeviceSpeedtestGetServer,
		},
		{
			name:      "AdminDeviceSpeedtestRead",
			path:      "/admin/device",
			form:      "speedtest",
			operation: "read",
			call:      (*Client).AdminDeviceSpeedtestRead,
		},
		{
			name:      "AdminDeviceSpeedtestStop",
			path:      "/admin/device",
			form:      "speedtest",
			operation: "stop",
			call:      (*Client).AdminDeviceSpeedtestStop,
		},
		{
			name:      "AdminDeviceSpeedtestWrite",
			path:      "/admin/device",
			form:      "speedtest",
			operation: "write",
			call:      (*Client).AdminDeviceSpeedtestWrite,
		},
		{
			name:      "AdminDeviceSyncRead",
			path:      "/admin/device",
			form:      "sync",
			operation: "read",
			call:      (*Client).AdminDeviceSyncRead,
		},
		{
			name:      "AdminDeviceSyncWrite",
			path:      "/admin/device",
			form:      "sync",
			operation: "write",
			call:      (*Client).AdminDeviceSyncWrite,
		},
		{
			name:      "AdminDeviceSystemFactory",
			path:      "/admin/device",
			form:      "system",
			operation: "factory",
			call:      (*Client).AdminDeviceSystemFactory,
		},
		{
			name:      "AdminDeviceSystemGateway",
			path:      "/admin/device",
			form:      "system",
			operation: "gateway",
			call:      (*Client).AdminDeviceSystemGateway,
		},
		{
			name:      "AdminDeviceSystemReboot",
			path:      "/admin/device",
			form:      "system",
			operation: "reboot",
			call:      (*Client).AdminDeviceSystemReboot,
		},
		{
			name:      "AdminDeviceTimesettingRead",
			path:      "/admin/device",
			form:      "timesetting",
			operation: "read",
			call:      (*Client).AdminDeviceTimesettingRead,
		},
		{
			name:      "AdminDeviceTimesettingWrite",
			path:      "/admin/device",
			form:      "timesetting",
			operation: "write",
			call:      (*Client).AdminDeviceTimesettingWrite,
		},
	})
}
