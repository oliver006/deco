package deco

import "testing"

func TestAdminClientMethods(t *testing.T) {
	testAdminMethodCases(t, []adminMethodCase{
		{
			name:      "AdminClientAddrReservationAdd",
			path:      "/admin/client",
			form:      "addr_reservation",
			operation: "add",
			call:      (*Client).AdminClientAddrReservationAdd,
		},
		{
			name:      "AdminClientAddrReservationGetlist",
			path:      "/admin/client",
			form:      "addr_reservation",
			operation: "getlist",
			call:      (*Client).AdminClientAddrReservationGetlist,
		},
		{
			name:      "AdminClientAddrReservationModify",
			path:      "/admin/client",
			form:      "addr_reservation",
			operation: "modify",
			call:      (*Client).AdminClientAddrReservationModify,
		},
		{
			name:      "AdminClientAddrReservationRemove",
			path:      "/admin/client",
			form:      "addr_reservation",
			operation: "remove",
			call:      (*Client).AdminClientAddrReservationRemove,
		},
		{
			name:      "AdminClientBlackListBlock",
			path:      "/admin/client",
			form:      "black_list",
			operation: "block",
			call:      (*Client).AdminClientBlackListBlock,
		},
		{
			name:      "AdminClientBlackListList",
			path:      "/admin/client",
			form:      "black_list",
			operation: "list",
			call:      (*Client).AdminClientBlackListList,
		},
		{
			name:      "AdminClientBlackListUnblock",
			path:      "/admin/client",
			form:      "black_list",
			operation: "unblock",
			call:      (*Client).AdminClientBlackListUnblock,
		},
		{
			name:      "AdminClientClientAccessRead",
			path:      "/admin/client",
			form:      "client_access",
			operation: "read",
			call:      (*Client).AdminClientClientAccessRead,
		},
		{
			name:      "AdminClientClientAccessWrite",
			path:      "/admin/client",
			form:      "client_access",
			operation: "write",
			call:      (*Client).AdminClientClientAccessWrite,
		},
		{
			name:      "AdminClientClientListRead",
			path:      "/admin/client",
			form:      "client_list",
			operation: "read",
			call:      (*Client).AdminClientClientListRead,
		},
		{
			name:      "AdminClientClientListRemove",
			path:      "/admin/client",
			form:      "client_list",
			operation: "remove",
			call:      (*Client).AdminClientClientListRemove,
		},
		{
			name:      "AdminClientClientListWrite",
			path:      "/admin/client",
			form:      "client_list",
			operation: "write",
			call:      (*Client).AdminClientClientListWrite,
		},
		{
			name:      "AdminClientTrafficStatClient",
			path:      "/admin/client",
			form:      "traffic_stat",
			operation: "client",
			call:      (*Client).AdminClientTrafficStatClient,
		},
		{
			name:      "AdminClientTrafficStatList",
			path:      "/admin/client",
			form:      "traffic_stat",
			operation: "list",
			call:      (*Client).AdminClientTrafficStatList,
		},
	})
}
