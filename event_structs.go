package darktrace

// EventType type of events
type EventType string

// Types of events that can be received.
const (
	EventTypePolicyBreach = EventType("policybreach")
	EventTypeConnection   = EventType("connection")
	EventTypeNotice       = EventType("notice")
)

type Event struct {
	Time      string    `json:"time"`
	Timems    int64     `json:"timems"`
	Action    string    `json:"action"`
	EventType EventType `json:"eventType"`
	PolicyBreachEvent
	NoticeEvent
	ConnectionEvent
}

type PolicyBreachEvent struct {
	Pbid                 int64         `json:"pbid"`
	Pid                  int64         `json:"pid"`
	Phid                 int64         `json:"phid"`
	CreationTime         int64         `json:"creationTime"`
	CreationTimestamp    string        `json:"creationTimestamp"`
	Name                 string        `json:"name"`
	Components           []int64       `json:"components"`
	DidRestrictions      []interface{} `json:"didRestrictions"`
	DidExclusions        []interface{} `json:"didExclusions"`
	Throttle             int64         `json:"throttle"`
	SharedEndpoints      bool          `json:"sharedEndpoints"`
	Interval             int64         `json:"interval"`
	Sequenced            bool          `json:"sequenced"`
	Active               bool          `json:"active"`
	Retired              bool          `json:"retired"`
	InstanceID           int64         `json:"instanceID"`
	Acknowledged         bool          `json:"acknowledged"`
	State                string        `json:"state"`
	Score                float64       `json:"score"`
	CommentCount         int64         `json:"commentCount"`
	ComponentBreaches    []int64       `json:"componentBreaches"`
	ComponentBreachTimes []int64       `json:"componentBreachTimes"`
	Devices              []int64       `json:"devices"`
	DeviceLabels         []string      `json:"deviceLabels"`
}

type NoticeEvent struct {
	Nid             int64  `json:"nid"`
	Info            string `json:"info"`
	UID             string `json:"uid"`
	Direction       string `json:"direction"`
	Mlid            int64  `json:"mlid"`
	Type            string `json:"type"`
	Dpcode          string `json:"dpcode"`
	Dpcodetitle     string `json:"dpcodetitle"`
	Msg             string `json:"msg"`
	DestinationPort int64  `json:"destinationPort"`
	Size            int64  `json:"size"`
	Details         string `json:"details"`
	SourceDevice    struct {
		ID         int64  `json:"id"`
		Did        int64  `json:"did"`
		Objecttype string `json:"objecttype"`
		Macaddress string `json:"macaddress"`
		IP         string `json:"ip"`
		Ips        []struct {
			IP     string `json:"ip"`
			Timems int64  `json:"timems"`
			Time   string `json:"time"`
			Sid    int64  `json:"sid"`
		} `json:"ips"`
		Sid             int64   `json:"sid"`
		Hostname        string  `json:"hostname"`
		Time            string  `json:"time"`
		Typename        string  `json:"typename"`
		Typelabel       string  `json:"typelabel"`
		DeduplicatedIds []int64 `json:"deduplicated_ids"`
	} `json:"sourceDevice"`
	DestinationDevice struct {
		ID         int64  `json:"id"`
		Did        int64  `json:"did"`
		Objecttype string `json:"objecttype"`
		Macaddress string `json:"macaddress"`
		IP         string `json:"ip"`
		Ips        []struct {
			IP     string `json:"ip"`
			Timems int64  `json:"timems"`
			Time   string `json:"time"`
			Sid    int64  `json:"sid"`
		} `json:"ips"`
		Sid             int64   `json:"sid"`
		Hostname        string  `json:"hostname"`
		Time            string  `json:"time"`
		Os              string  `json:"os"`
		Typename        string  `json:"typename"`
		Typelabel       string  `json:"typelabel"`
		DeduplicatedIds []int64 `json:"deduplicated_ids"`
	} `json:"destinationDevice"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type ConnectionEvent struct {
	Recentbytessincems  int64  `json:"recentbytessincems"`
	Recentbytessince    string `json:"recentbytessince"`
	Recentbytesout      int64  `json:"recentbytesout"`
	Recentbytesin       int64  `json:"recentbytesin"`
	Totalbytessincems   int64  `json:"totalbytessincems"`
	Totalbytessince     string `json:"totalbytessince"`
	Totalbytesout       int64  `json:"totalbytesout"`
	Totalbytesin        int64  `json:"totalbytesin"`
	UID                 string `json:"uid"`
	Status              string `json:"status"`
	Sdid                int64  `json:"sdid"`
	Ddid                int64  `json:"ddid"`
	Port                int64  `json:"port"`
	SourcePort          int64  `json:"sourcePort"`
	DestinationPort     int64  `json:"destinationPort"`
	Info                string `json:"info"`
	Direction           string `json:"direction"`
	Applicationprotocol string `json:"applicationprotocol"`
	Protocol            string `json:"protocol"`
	SourceDevice        struct {
		ID         int64  `json:"id"`
		Did        int64  `json:"did"`
		Objecttype string `json:"objecttype"`
		Macaddress string `json:"macaddress"`
		IP         string `json:"ip"`
		Ips        []struct {
			IP     string `json:"ip"`
			Timems int64  `json:"timems"`
			Time   string `json:"time"`
			Sid    int64  `json:"sid"`
		} `json:"ips"`
		Sid             int64   `json:"sid"`
		Hostname        string  `json:"hostname"`
		Time            string  `json:"time"`
		Typename        string  `json:"typename"`
		Typelabel       string  `json:"typelabel"`
		DeduplicatedIds []int64 `json:"deduplicated_ids"`
	} `json:"sourceDevice"`
	DestinationDevice struct {
		ID         int64  `json:"id"`
		Did        int64  `json:"did"`
		Objecttype string `json:"objecttype"`
		Macaddress string `json:"macaddress"`
		IP         string `json:"ip"`
		Ips        []struct {
			IP     string `json:"ip"`
			Timems int64  `json:"timems"`
			Time   string `json:"time"`
			Sid    int64  `json:"sid"`
		} `json:"ips"`
		Sid             int64   `json:"sid"`
		Hostname        string  `json:"hostname"`
		Time            string  `json:"time"`
		Os              string  `json:"os"`
		Typename        string  `json:"typename"`
		Typelabel       string  `json:"typelabel"`
		DeduplicatedIds []int64 `json:"deduplicated_ids"`
	} `json:"destinationDevice"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}
