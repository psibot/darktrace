package darktrace

type Device struct {
	ID  int64  `json:"id"`
	IP  string `json:"ip,omitempty"`
	Ips []struct {
		IP     string `json:"ip"`
		Timems int64  `json:"timems"`
		Time   string `json:"time"`
		Sid    int64  `json:"sid"`
	} `json:"ips,omitempty"`
	Did             int     `json:"did,omitempty"`
	Sid             int64   `json:"sid,omitempty"`
	Hostname        string  `json:"hostname,omitempty"`
	Time            int64   `json:"time,omitempty"`
	Endtime         int64   `json:"endtime,omitempty"`
	Os              string  `json:"os,omitempty"`
	Typename        string  `json:"typename,omitempty"`
	Typelabel       string  `json:"typelabel,omitempty"`
	DeduplicatedIds []int64 `json:"deduplicated_ids,omitempty"`
}

type DeviceInfo struct {
	DeviceInfo []struct {
		Did             int `json:"did"`
		SimilarityScore int `json:"similarityScore"`
		GraphData       []struct {
			Time  int64 `json:"time"`
			Count int   `json:"count"`
		} `json:"graphData"`
		Info struct {
			TotalUsed            int `json:"totalUsed"`
			TotalServed          int `json:"totalServed"`
			TotalDevicesAndPorts int `json:"totalDevicesAndPorts"`
			DevicesAndPorts      []struct {
				DeviceAndPort string `json:"deviceAndPort"`
				Size          int    `json:"size"`
			} `json:"devicesAndPorts"`
			PortsUsed []struct {
				Port      interface{} `json:"port"`
				Size      int         `json:"size"`
				FirstTime int64       `json:"firstTime"`
			} `json:"portsUsed"`
			PortsServed []struct {
				Port interface{} `json:"port"`
				Size int         `json:"size"`
			} `json:"portsServed"`
			DevicesUsed []struct {
				Did       interface{} `json:"did"`
				Size      int         `json:"size"`
				FirstTime int64       `json:"firstTime,omitempty"`
			} `json:"devicesUsed"`
			DevicesServed []struct {
				Did  interface{} `json:"did"`
				Size int         `json:"size"`
			} `json:"devicesServed"`
		} `json:"info"`
	} `json:"deviceInfo"`
}

type SimilarDevice struct {
	Did        int    `json:"did"`
	Objecttype string `json:"objecttype"`
	Score      int    `json:"score"`
	IP         string `json:"ip"`
	Ips        []struct {
		IP     string `json:"ip"`
		Timems int64  `json:"timems"`
		Time   string `json:"time"`
		Sid    int64  `json:"sid"`
	} `json:"ips"`
	Sid       int64  `json:"sid"`
	Hostname  string `json:"hostname,omitempty"`
	FirstSeen int64  `json:"firstSeen"`
	LastSeen  int64  `json:"lastSeen"`
	Os        string `json:"os,omitempty"`
	Typename  string `json:"typename"`
	Typelabel string `json:"typelabel"`
	Priority  int    `json:"priority,omitempty"`
}

type DeviceMetric struct {
	Time   string `json:"time"`
	TimeMS int64  `json:"timems"`
	Size   int    `json:"size"`
	In     int    `json:"in"`
	Out    int    `json:"out"`
}

type Model struct {
	Name  string `json:"name"`
	Pid   int    `json:"pid"`
	Phid  int    `json:"phid"`
	UUID  string `json:"uuid"`
	Logic struct {
		Data    []interface{} `json:"data"`
		Type    string        `json:"type"`
		Version int           `json:"version"`
	} `json:"logic"`
	Throttle        int  `json:"throttle"`
	SharedEndpoints bool `json:"sharedEndpoints"`
	Actions         struct {
		Alert    bool `json:"alert"`
		Antigena struct {
		} `json:"antigena"`
		Breach      bool        `json:"breach"`
		Model       bool        `json:"model"`
		SetPriority interface{} `json:"setPriority"`
		SetTag      interface{} `json:"setTag"`
		SetType     interface{} `json:"setType"`
	} `json:"actions"`
	Tags        []string `json:"tags"`
	Interval    int      `json:"interval"`
	Sequenced   bool     `json:"sequenced"`
	Active      bool     `json:"active"`
	Modified    string   `json:"modified"`
	ActiveTimes struct {
		Devices struct {
		} `json:"devices"`
		Tags struct {
		} `json:"tags"`
		Type    string `json:"type"`
		Version int    `json:"version"`
	} `json:"activeTimes"`
	Priority      int           `json:"priority"`
	AutoUpdatable bool          `json:"autoUpdatable"`
	AutoUpdate    bool          `json:"autoUpdate"`
	AutoSuppress  bool          `json:"autoSuppress"`
	Description   string        `json:"description"`
	Behaviour     string        `json:"behaviour"`
	Defeats       []interface{} `json:"defeats"`
	Created       struct {
		By string `json:"by"`
	} `json:"created"`
	Edited struct {
		By string `json:"by"`
	} `json:"edited"`
	History []struct {
		Modified string `json:"modified"`
		Active   bool   `json:"active"`
		Message  string `json:"message"`
		By       string `json:"by"`
		Phid     int    `json:"phid"`
	} `json:"history"`
	Message string `json:"message"`
	Version int    `json:"version"`
}

type ModelComponent struct {
	Cid       int `json:"cid"`
	Chid      int `json:"chid"`
	Mlid      int `json:"mlid"`
	Threshold int `json:"threshold"`
	Interval  int `json:"interval"`
	Logic     struct {
		Data    interface{} `json:"data"`
		Version string      `json:"version"`
	} `json:"logic"`
	Filters []interface{} `json:"filters"`
	Active  bool          `json:"active"`
}

type ModelBreach struct {
	CreationTime int64 `json:"creationTime"`
	CommentCount int   `json:"commentCount"`
	Pbid         int64 `json:"pbid"`
	Time         int64 `json:"time"`
	Model        struct {
		Then struct {
			Name  string `json:"name"`
			Pid   int    `json:"pid"`
			Phid  int    `json:"phid"`
			UUID  string `json:"uuid"`
			Logic struct {
				Data    interface{} `json:"data"`
				Type    string      `json:"type"`
				Version int         `json:"version"`
			} `json:"logic"`
			Throttle        int  `json:"throttle"`
			SharedEndpoints bool `json:"sharedEndpoints"`
			Actions         struct {
				Alert    bool `json:"alert"`
				Antigena struct {
				} `json:"antigena"`
				Breach      bool `json:"breach"`
				Model       bool `json:"model"`
				SetPriority bool `json:"setPriority"`
				SetTag      bool `json:"setTag"`
				SetType     bool `json:"setType"`
			} `json:"actions"`
			Tags        []string `json:"tags"`
			Interval    int      `json:"interval"`
			Sequenced   bool     `json:"sequenced"`
			Active      bool     `json:"active"`
			ReadOnly    bool     `json:"readOnly"`
			Modified    string   `json:"modified"`
			ActiveTimes struct {
				Devices struct {
				} `json:"devices"`
				Tags struct {
				} `json:"tags"`
				Type    string `json:"type"`
				Version int    `json:"version"`
			} `json:"activeTimes"`
			Priority      int           `json:"priority"`
			AutoUpdatable bool          `json:"autoUpdatable"`
			AutoUpdate    bool          `json:"autoUpdate"`
			AutoSuppress  bool          `json:"autoSuppress"`
			Description   string        `json:"description"`
			Behaviour     string        `json:"behaviour"`
			Defeats       []interface{} `json:"defeats"`
			Created       struct {
			} `json:"created"`
			Edited struct {
			} `json:"edited"`
			Version int `json:"version"`
		} `json:"then"`
		Now struct {
			Name  string `json:"name"`
			Pid   int    `json:"pid"`
			Phid  int    `json:"phid"`
			UUID  string `json:"uuid"`
			Logic struct {
				Data    interface{} `json:"data"`
				Type    string      `json:"type"`
				Version int         `json:"version"`
			} `json:"logic"`
			Throttle        int  `json:"throttle"`
			SharedEndpoints bool `json:"sharedEndpoints"`
			Actions         struct {
				Alert    bool `json:"alert"`
				Antigena struct {
				} `json:"antigena"`
				Breach      bool `json:"breach"`
				Model       bool `json:"model"`
				SetPriority bool `json:"setPriority"`
				SetTag      bool `json:"setTag"`
				SetType     bool `json:"setType"`
			} `json:"actions"`
			Tags        []string `json:"tags"`
			Interval    int      `json:"interval"`
			Sequenced   bool     `json:"sequenced"`
			Active      bool     `json:"active"`
			ReadOnly    bool     `json:"readOnly"`
			Modified    string   `json:"modified"`
			ActiveTimes struct {
				Devices struct {
				} `json:"devices"`
				Tags struct {
				} `json:"tags"`
				Type    string `json:"type"`
				Version int    `json:"version"`
			} `json:"activeTimes"`
			Priority      int           `json:"priority"`
			AutoUpdatable bool          `json:"autoUpdatable"`
			AutoUpdate    bool          `json:"autoUpdate"`
			AutoSuppress  bool          `json:"autoSuppress"`
			Description   string        `json:"description"`
			Behaviour     string        `json:"behaviour"`
			Defeats       []interface{} `json:"defeats"`
			Created       struct {
			} `json:"created"`
			Edited struct {
			} `json:"edited"`
			Version int `json:"version"`
		} `json:"now"`
	} `json:"model"`
	TriggeredComponents []struct {
		Time      int64 `json:"time"`
		Cbid      int64 `json:"cbid"`
		Cid       int   `json:"cid"`
		Chid      int   `json:"chid"`
		Size      int   `json:"size"`
		Threshold int   `json:"threshold"`
		Interval  int   `json:"interval"`
		Metric    struct {
			Mlid  int    `json:"mlid"`
			Name  string `json:"name"`
			Label string `json:"label"`
		} `json:"metric"`
		TriggeredFilters []struct {
			Cfid       int    `json:"cfid"`
			ID         string `json:"id"`
			FilterType string `json:"filterType"`
			Arguments  struct {
				Value interface{} `json:"value"`
			} `json:"arguments,omitempty"`
			ComparatorType string `json:"comparatorType"`
			Trigger        struct {
				Value interface{} `json:"value"`
			} `json:"trigger"`
		} `json:"triggeredFilters"`
	} `json:"triggeredComponents"`
	Score  float64 `json:"score"`
	Device struct {
		Did             int     `json:"did"`
		Objecttype      string  `json:"objecttype"`
		Hostname        string  `json:"hostname"`
		IP              string  `json:"ip"`
		Sid             int64   `json:"sid"`
		Typelabel       string  `json:"typelabel"`
		Typename        string  `json:"typename"`
		DeduplicatedIds []int64 `json:"deduplicated_ids"`
	} `json:"device"`
}

type EnumResponse struct {
	Results map[string][]Enum
}

type Enum struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden,omitempty"`
}

type EndpointDetails struct {
	Hostname   string `json:"hostname"`
	Dgascore   int    `json:"dgascore"`
	Firsttime  int64  `json:"firsttime"`
	Popularity int    `json:"popularity"`
	Devices    []struct {
		Did        int    `json:"did"`
		Objecttype string `json:"objecttype"`
		Macaddress string `json:"macaddress,omitempty"`
		Vendor     string `json:"vendor,omitempty"`
		IP         string `json:"ip"`
		Ips        []struct {
			IP     string `json:"ip"`
			Timems int64  `json:"timems"`
			Time   string `json:"time"`
			Sid    int64  `json:"sid"`
		} `json:"ips"`
		Sid             int64   `json:"sid"`
		Hostname        string  `json:"hostname"`
		FirstSeen       int64   `json:"firstSeen"`
		LastSeen        int64   `json:"lastSeen"`
		Os              string  `json:"os"`
		Typename        string  `json:"typename"`
		Typelabel       string  `json:"typelabel"`
		DeduplicatedIds []int64 `json:"deduplicated_ids"`
	} `json:"devices"`
	Ips []struct {
		IP        string `json:"ip"`
		Firsttime int64  `json:"firsttime"`
		Lasttime  int64  `json:"lasttime"`
	} `json:"ips"`
	Locations []struct {
		Latitude  int    `json:"latitude"`
		Longitude int    `json:"longitude"`
		Country   string `json:"country"`
		City      string `json:"city"`
	} `json:"locations"`
}
