package deco

// AdminClientAddrReservationAdd calls /admin/client?form=addr_reservation with operation add.
func (c *Client) AdminClientAddrReservationAdd(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "addr_reservation", "add", params)
}

// AdminClientAddrReservationGetlist calls /admin/client?form=addr_reservation with operation getlist.
func (c *Client) AdminClientAddrReservationGetlist(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "addr_reservation", "getlist", params)
}

// AdminClientAddrReservationModify calls /admin/client?form=addr_reservation with operation modify.
func (c *Client) AdminClientAddrReservationModify(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "addr_reservation", "modify", params)
}

// AdminClientAddrReservationRemove calls /admin/client?form=addr_reservation with operation remove.
func (c *Client) AdminClientAddrReservationRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "addr_reservation", "remove", params)
}

// AdminClientBlackListBlock calls /admin/client?form=black_list with operation block.
func (c *Client) AdminClientBlackListBlock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "black_list", "block", params)
}

// AdminClientBlackListList calls /admin/client?form=black_list with operation list.
func (c *Client) AdminClientBlackListList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "black_list", "list", params)
}

// AdminClientBlackListUnblock calls /admin/client?form=black_list with operation unblock.
func (c *Client) AdminClientBlackListUnblock(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "black_list", "unblock", params)
}

// AdminClientClientAccessRead calls /admin/client?form=client_access with operation read.
func (c *Client) AdminClientClientAccessRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "client_access", "read", params)
}

// AdminClientClientAccessWrite calls /admin/client?form=client_access with operation write.
func (c *Client) AdminClientClientAccessWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "client_access", "write", params)
}

// AdminClientClientListRead calls /admin/client?form=client_list with operation read.
func (c *Client) AdminClientClientListRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "client_list", "read", params)
}

// AdminClientClientListRemove calls /admin/client?form=client_list with operation remove.
func (c *Client) AdminClientClientListRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "client_list", "remove", params)
}

// AdminClientClientListWrite calls /admin/client?form=client_list with operation write.
func (c *Client) AdminClientClientListWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "client_list", "write", params)
}

// AdminClientTrafficStatClient calls /admin/client?form=traffic_stat with operation client.
func (c *Client) AdminClientTrafficStatClient(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "traffic_stat", "client", params)
}

// AdminClientTrafficStatList calls /admin/client?form=traffic_stat with operation list.
func (c *Client) AdminClientTrafficStatList(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/client", "traffic_stat", "list", params)
}
