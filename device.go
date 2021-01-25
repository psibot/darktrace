package darktrace

import (
	"encoding/json"
)

// DeviceList returns an array of devices from the submitted parameters
// Supported params: ip, iptime, mac, seensince
func (c *Client) DeviceList(parameters ...RequestParameter) ([]Device, error) {
	// Setup params
	validParams := deviceListParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/devices", &params)

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
	var devices []Device
	err = json.Unmarshal(responseData, &devices)
	return devices, err
}

// DeviceInfo returns an array of devices from the submitted parameters
// Supported params: datatype, did, externaldomain, fulldevicedetails, odid, showallgraphdata, similardevices
func (c *Client) DeviceInfo(parameters ...RequestParameter) (*DeviceInfo, error) {
	// Setup params
	validParams := deviceInfoParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/deviceinfo", &params)

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
	var device DeviceInfo
	err = json.Unmarshal(responseData, &device)
	return &device, err
}

// SimilarDevices returns an array of devices with similar characteristics
// Supported params: did, count
func (c *Client) SimilarDevices(parameters ...RequestParameter) ([]SimilarDevice, error) {
	// Setup params
	validParams := similarDevicesParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/similardevices", &params)

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
	var devices []SimilarDevice
	err = json.Unmarshal(responseData, &devices)
	return devices, err
}

// DeviceMetrics returns an array of device metrics
// Supported params: applicationprotocol, breachtimes, ddid, destinationport, did, endtime, from, fulldevicedetails,
//					 interval, metric, odid, port, protocol, sourceport, starttime, to
func (c *Client) DeviceMetrics(parameters ...RequestParameter) ([]DeviceMetric, error) {
	// Setup params
	validParams := deviceMetricsParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/metricdata", &params)

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
	var deviceMetrics []DeviceMetric
	err = json.Unmarshal(responseData, &deviceMetrics)
	return deviceMetrics, err
}

// EndpointDetails returns an array of device metrics
// Supported params: additionalinfo, devices, hostname, ip, score
func (c *Client) EndpointDetails(parameters ...RequestParameter) (*EndpointDetails, error) {
	// Setup params
	validParams := endpointDetailsParams()
	params, err := setupParams(parameters, validParams)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Build request
	req, err := c.buildRequest("GET", "/endpointdetails", &params)

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
	var endpointDetails EndpointDetails
	err = json.Unmarshal(responseData, &endpointDetails)
	return &endpointDetails, err
}
