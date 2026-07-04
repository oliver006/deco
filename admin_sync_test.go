package deco

import "testing"

func TestAdminSyncMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminSyncCheckFirmware",
			path:      "/admin/sync",
			form:      "",
			operation: "check_firmware",
			call:      (*Client).AdminSyncCheckFirmware,
		},
		{
			name:      "AdminSyncForceUpgrade",
			path:      "/admin/sync",
			form:      "",
			operation: "force_upgrade",
			call:      (*Client).AdminSyncForceUpgrade,
		},
		{
			name:      "AdminSyncForceUpgradeLTE",
			path:      "/admin/sync",
			form:      "",
			operation: "force_upgrade_lte",
			call:      (*Client).AdminSyncForceUpgradeLTE,
		},
		{
			name:      "AdminSyncSyncCheck",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_check",
			call:      (*Client).AdminSyncSyncCheck,
		},
		{
			name:      "AdminSyncSyncConfig",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_config",
			call:      (*Client).AdminSyncSyncConfig,
		},
		{
			name:      "AdminSyncSyncDetectSlave",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_detect_slave",
			call:      (*Client).AdminSyncSyncDetectSlave,
		},
		{
			name:      "AdminSyncSyncDownloadBigfirm",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_download_bigfirm",
			call:      (*Client).AdminSyncSyncDownloadBigfirm,
		},
		{
			name:      "AdminSyncSyncDownloadStatusBigfirm",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_download_status_bigfirm",
			call:      (*Client).AdminSyncSyncDownloadStatusBigfirm,
		},
		{
			name:      "AdminSyncSyncEMMCCheck",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_emmc_check",
			call:      (*Client).AdminSyncSyncEMMCCheck,
		},
		{
			name:      "AdminSyncSyncEMMCConfig",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_emmc_config",
			call:      (*Client).AdminSyncSyncEMMCConfig,
		},
		{
			name:      "AdminSyncSyncFirmware",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_firmware",
			call:      (*Client).AdminSyncSyncFirmware,
		},
		{
			name:      "AdminSyncSyncGetCfg",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_get_cfg",
			call:      (*Client).AdminSyncSyncGetCfg,
		},
		{
			name:      "AdminSyncSyncGetInfo",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_get_info",
			call:      (*Client).AdminSyncSyncGetInfo,
		},
		{
			name:      "AdminSyncSyncISPProfile",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_isp_profile",
			call:      (*Client).AdminSyncSyncISPProfile,
		},
		{
			name:      "AdminSyncSyncSubconfig",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_subconfig",
			call:      (*Client).AdminSyncSyncSubconfig,
		},
		{
			name:      "AdminSyncSyncUpdateDevList",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_update_dev_list",
			call:      (*Client).AdminSyncSyncUpdateDevList,
		},
		{
			name:      "AdminSyncSyncUpgrade",
			path:      "/admin/sync",
			form:      "",
			operation: "sync_upgrade",
			call:      (*Client).AdminSyncSyncUpgrade,
		},
	})
}
