package darktrace

import (
	"encoding/json"
)

func (c *Client) ListModels() ([]Model, error) {
	// Build request
	req, err := c.buildRequest("GET", "/models", nil)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Do request
	responseData, err := c.doRequest(req)

	// Unmarshal json body into structs and return
	var models []Model
	err = json.Unmarshal(responseData, &models)
	return models, err
}

func (c *Client) ModelComponents(parameters ...RequestParameter) ([]ModelComponent, error) {
	// Setup params
	validParams := modelComponentParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/components", &params)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Do request
	responseData, err := c.doRequest(req)

	// Unmarshal json body into structs and return
	var components []ModelComponent
	err = json.Unmarshal(responseData, &components)
	return components, err
}

func (c *Client) ModelBreaches(parameters ...RequestParameter) ([]ModelBreach, error) {
	// Setup params
	validParams := modelBreachesParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/modelbreaches", &params)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Do request
	responseData, err := c.doRequest(req)

	// Unmarshal json body into structs and return
	var breaches []ModelBreach
	err = json.Unmarshal(responseData, &breaches)
	return breaches, err
}
