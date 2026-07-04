package deco

// AdminAdministrationAccountAppget calls /admin/administration?form=account with operation appget.
func (c *Client) AdminAdministrationAccountAppget(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "account", "appget", params)
}

// AdminAdministrationAccountAppset calls /admin/administration?form=account with operation appset.
func (c *Client) AdminAdministrationAccountAppset(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "account", "appset", params)
}

// AdminAdministrationAccountMCUCheck calls /admin/administration?form=account with operation mcu_check.
func (c *Client) AdminAdministrationAccountMCUCheck(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "account", "mcu_check", params)
}

// AdminAdministrationAccountMCURead calls /admin/administration?form=account with operation mcu_read.
func (c *Client) AdminAdministrationAccountMCURead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "account", "mcu_read", params)
}

// AdminAdministrationAccountMCUWrite calls /admin/administration?form=account with operation mcu_write.
func (c *Client) AdminAdministrationAccountMCUWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "account", "mcu_write", params)
}

// AdminAdministrationAccountRead calls /admin/administration?form=account with operation read.
func (c *Client) AdminAdministrationAccountRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "account", "read", params)
}

// AdminAdministrationAccountSet calls /admin/administration?form=account with operation set.
func (c *Client) AdminAdministrationAccountSet(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "account", "set", params)
}

// AdminAdministrationAccountWrite calls /admin/administration?form=account with operation write.
func (c *Client) AdminAdministrationAccountWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "account", "write", params)
}

// AdminAdministrationLocalInsert calls /admin/administration?form=local with operation insert.
func (c *Client) AdminAdministrationLocalInsert(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "local", "insert", params)
}

// AdminAdministrationLocalLoad calls /admin/administration?form=local with operation load.
func (c *Client) AdminAdministrationLocalLoad(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "local", "load", params)
}

// AdminAdministrationLocalRemove calls /admin/administration?form=local with operation remove.
func (c *Client) AdminAdministrationLocalRemove(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "local", "remove", params)
}

// AdminAdministrationLocalUpdate calls /admin/administration?form=local with operation update.
func (c *Client) AdminAdministrationLocalUpdate(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "local", "update", params)
}

// AdminAdministrationLoginRead calls /admin/administration?form=login with operation read.
func (c *Client) AdminAdministrationLoginRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "login", "read", params)
}

// AdminAdministrationLoginWrite calls /admin/administration?form=login with operation write.
func (c *Client) AdminAdministrationLoginWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "login", "write", params)
}

// AdminAdministrationModeRead calls /admin/administration?form=mode with operation read.
func (c *Client) AdminAdministrationModeRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "mode", "read", params)
}

// AdminAdministrationModeWrite calls /admin/administration?form=mode with operation write.
func (c *Client) AdminAdministrationModeWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "mode", "write", params)
}

// AdminAdministrationRecoveryRead calls /admin/administration?form=recovery with operation read.
func (c *Client) AdminAdministrationRecoveryRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "recovery", "read", params)
}

// AdminAdministrationRecoveryWrite calls /admin/administration?form=recovery with operation write.
func (c *Client) AdminAdministrationRecoveryWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "recovery", "write", params)
}

// AdminAdministrationRemoteRead calls /admin/administration?form=remote with operation read.
func (c *Client) AdminAdministrationRemoteRead(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "remote", "read", params)
}

// AdminAdministrationRemoteWrite calls /admin/administration?form=remote with operation write.
func (c *Client) AdminAdministrationRemoteWrite(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "remote", "write", params)
}

// AdminAdministrationViewLoad calls /admin/administration?form=view with operation load.
func (c *Client) AdminAdministrationViewLoad(params map[string]interface{}) (map[string]interface{}, error) {
	return c.adminPost("/admin/administration", "view", "load", params)
}
