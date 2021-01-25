package darktrace

import (
	"fmt"
	"net/url"
)

func setupParams(reqParams []RequestParameter, valid []string) (url.Values, error) {
	params := url.Values{}

	// Set request parameters
	for _, op := range reqParams {
		pKey, pVal := op()
		if !contains(pKey, valid) {
			return nil, fmt.Errorf("invalid parameter: %s", pKey)
		}
		params.Set(pKey, pVal)
	}

	return params, nil
}

func eventAcknowledgeParams() []string {
	return []string{"pbid"}
}

func eventUnAcknowledgeParams() []string {
	return []string{"pbid"}
}

func eventDetailsParams() []string {
	return []string{
		"applicationprotocol",
		"count",
		"ddid",
		"deduplicate",
		"destinationport",
		"did",
		"endtime",
		"eventtype",
		"externalhostname",
		"from",
		"fulldevicedetails",
		"intext",
		"msg",
		"odid",
		"pbid",
		"port",
		"protocol",
		"sourceport",
		"starttime",
		"to",
		"uid",
	}
}

func modelComponentParams() []string {
	return []string{"cid"}
}

func modelBreachesParams() []string {
	return []string{
		"deviceattop",
		"did",
		"endtime",
		"expandenums",
		"from",
		"historicmodelonly",
		"includeacknowledged",
		"includebreachurl",
		"minimal",
		"minscore",
		"pbid",
		"pid",
		"starttime",
		"to",
		"uuid",
	}
}

func deviceListParams() []string {
	return []string{
		"ip",
		"iptime",
		"mac",
		"seensince",
	}
}

func similarDevicesParams() []string {
	return []string{
		"did",
		"count",
	}
}

func deviceMetricsParams() []string {
	return []string{
		"applicationprotocol",
		"breachtimes",
		"ddid",
		"destinationport",
		"did",
		"endtime",
		"from",
		"fulldevicedetails",
		"interval",
		"metric",
		"odid",
		"port",
		"protocol",
		"sourceport",
		"starttime",
		"to",
	}
}

func deviceInfoParams() []string {
	return []string{
		"datatype",
		"did",
		"externaldomain",
		"fulldevicedetails",
		"odid",
		"showallgraphdata",
		"similardevices",
	}
}

func endpointDetailsParams() []string {
	return []string{
		"additionalinfo",
		"devices",
		"hostname",
		"ip",
		"score",
	}
}

func networkListParams() []string {
	return []string{
		"applicationprotocol",
		"destinationport",
		"did",
		"endtime",
		"from",
		"fulldevicedetails",
		"intext",
		"ip",
		"metric",
		"port",
		"protocol",
		"sourceport",
		"starttime",
		"to",
	}
}

func contains(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
