package darktrace

import (
	"encoding/json"
)

// EnumsList returns a full list of enums
func (c *Client) EnumsList() (*EnumResponse, error) {
	// Build request
	req, err := c.buildRequest("GET", "/enums", nil)

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
	var enums EnumResponse
	err = json.Unmarshal(responseData, &enums)
	return &enums, err
}

// EnumsApplicationProtocolsList returns the list of application protocol enums
func (c *Client) EnumsApplicationProtocolsList() ([]Enum, error) {
	// Build request
	req, err := c.buildRequest("GET", "/enums/applicationprotocols", nil)

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
	var enums []Enum
	err = json.Unmarshal(responseData, &enums)
	return enums, err
}

// EnumsCountriesList returns the list of country enums
func (c *Client) EnumsCountriesList() ([]Enum, error) {
	// Build request
	req, err := c.buildRequest("GET", "/enums/countries", nil)

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
	var enums []Enum
	err = json.Unmarshal(responseData, &enums)
	return enums, err
}

// EnumsDestinationDeviceTypesList returns the list of destination device type enums
func (c *Client) EnumsDestinationDeviceTypesList() ([]Enum, error) {
	// Build request
	req, err := c.buildRequest("GET", "/enums/destinationdevicetypes", nil)

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
	var enums []Enum
	err = json.Unmarshal(responseData, &enums)
	return enums, err
}

// EnumsProtocolsList returns the list of protocol enums
func (c *Client) EnumsProtocolsList() ([]Enum, error) {
	// Build request
	req, err := c.buildRequest("GET", "/enums/protocols", nil)

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
	var enums []Enum
	err = json.Unmarshal(responseData, &enums)
	return enums, err
}

// EnumsSourceDeviceTypesList returns the list of source device type enums
func (c *Client) EnumsSourceDeviceTypesList() ([]Enum, error) {
	// Build request
	req, err := c.buildRequest("GET", "/enums/sourcedevicetypes", nil)

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
	var enums []Enum
	err = json.Unmarshal(responseData, &enums)
	return enums, err
}
