package deco

// AdminSystemEnvarRead calls /admin/system?form=envar with operation read.
func (c *Client) AdminSystemEnvarRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/system", "envar", "read", params)
}

// AdminSystemEnvarWrite calls /admin/system?form=envar with operation write.
func (c *Client) AdminSystemEnvarWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/system", "envar", "write", params)
}

// AdminSystemLogoutLogout calls /admin/system?form=logout with operation logout.
func (c *Client) AdminSystemLogoutLogout(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/system", "logout", "logout", params)
}
