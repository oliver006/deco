package deco

import "testing"

func TestAdminWirelessMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminWirelessBandwidthEnhanceRead",
			path:      "/admin/wireless",
			form:      "bandwidth_enhance",
			operation: "read",
			call:      (*Client).AdminWirelessBandwidthEnhanceRead,
		},
		{
			name:      "AdminWirelessBandwidthEnhanceWrite",
			path:      "/admin/wireless",
			form:      "bandwidth_enhance",
			operation: "write",
			call:      (*Client).AdminWirelessBandwidthEnhanceWrite,
		},
		{
			name:      "AdminWirelessBeamformingRead",
			path:      "/admin/wireless",
			form:      "beamforming",
			operation: "read",
			call:      (*Client).AdminWirelessBeamformingRead,
		},
		{
			name:      "AdminWirelessBeamformingWrite",
			path:      "/admin/wireless",
			form:      "beamforming",
			operation: "write",
			call:      (*Client).AdminWirelessBeamformingWrite,
		},
		{
			name:      "AdminWirelessBridgeCheck",
			path:      "/admin/wireless",
			form:      "bridge",
			operation: "check",
			call:      (*Client).AdminWirelessBridgeCheck,
		},
		{
			name:      "AdminWirelessBridgeRead",
			path:      "/admin/wireless",
			form:      "bridge",
			operation: "read",
			call:      (*Client).AdminWirelessBridgeRead,
		},
		{
			name:      "AdminWirelessIeee80211rRead",
			path:      "/admin/wireless",
			form:      "ieee80211r",
			operation: "read",
			call:      (*Client).AdminWirelessIeee80211rRead,
		},
		{
			name:      "AdminWirelessIeee80211rWrite",
			path:      "/admin/wireless",
			form:      "ieee80211r",
			operation: "write",
			call:      (*Client).AdminWirelessIeee80211rWrite,
		},
		{
			name:      "AdminWirelessOperationModeRead",
			path:      "/admin/wireless",
			form:      "operation_mode",
			operation: "read",
			call:      (*Client).AdminWirelessOperationModeRead,
		},
		{
			name:      "AdminWirelessOperationModeWrite",
			path:      "/admin/wireless",
			form:      "operation_mode",
			operation: "write",
			call:      (*Client).AdminWirelessOperationModeWrite,
		},
		{
			name:      "AdminWirelessPowerRead",
			path:      "/admin/wireless",
			form:      "power",
			operation: "read",
			call:      (*Client).AdminWirelessPowerRead,
		},
		{
			name:      "AdminWirelessPowerWrite",
			path:      "/admin/wireless",
			form:      "power",
			operation: "write",
			call:      (*Client).AdminWirelessPowerWrite,
		},
		{
			name:      "AdminWirelessWLANRead",
			path:      "/admin/wireless",
			form:      "wlan",
			operation: "read",
			call:      (*Client).AdminWirelessWLANRead,
		},
		{
			name:      "AdminWirelessWLANWrite",
			path:      "/admin/wireless",
			form:      "wlan",
			operation: "write",
			call:      (*Client).AdminWirelessWLANWrite,
		},
	})
}
