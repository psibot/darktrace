package darktrace

import (
	"encoding/json"
)

func (c *Client) AcknowledgeEvent(parameters ...RequestParameter) error {
	// Setup params
	validParams := eventAcknowledgeParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return err
	}

	// Build request
	req, err := c.buildRequest("POST", "/acknowledgeevent", &params)

	// Handle errors
	if err != nil {
		return err
	}

	// Conduct request
	_, err = c.doRequest(req)

	return err
}

func (c *Client) UnAcknowledgeEvent(parameters ...RequestParameter) error {
	// Setup params
	validParams := eventUnAcknowledgeParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return err
	}

	// Build request
	req, err := c.buildRequest("POST", "/unacknowledgeevent", &params)

	// Handle errors
	if err != nil {
		return err
	}

	// Conduct request
	_, err = c.doRequest(req)

	return err
}

func (c *Client) EventList(parameters ...RequestParameter) ([]Event, error) {
	// Setup params
	validParams := eventDetailsParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/details", &params)

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
	var events []Event
	err = json.Unmarshal(responseData, &events)
	return events, err
}
