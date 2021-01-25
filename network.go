package darktrace

import (
	"encoding/json"
)

func (c *Client) NetworkList(parameters ...RequestParameter) ([]interface{}, error) {
	// Setup params
	validParams := networkListParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/network", &params)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Conduct request
	responseData, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	// Unmarshal json body into structs and return
	var list []interface{}
	err = json.Unmarshal(responseData, &list)
	return list, err
}

func (c *Client) PcapList() ([]interface{}, error) {
	// Build request
	req, err := c.buildRequest("GET", "/pcaps", nil)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Conduct request
	responseData, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	// Unmarshal json body into structs and return
	var list []interface{}
	err = json.Unmarshal(responseData, &list)
	return list, err
}
