package deco

import (
	"encoding/json"
	"fmt"
)

func (c *Client) adminPost(path, form, operation string, params map[string]interface{}) (map[string]interface{}, error) {
	req := request{
		Operation: operation,
		Params:    params,
	}
	jsonRequest, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = c.doEncryptedPost(fmt.Sprintf(";stok=%s%s", c.stok, path), EndpointArgs{form: form}, jsonRequest, false, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
