package aluminum

import (
	"net/http"
	"time"
)

//Aluminum is the default log struct if you don't want to make your own
type Aluminum struct {
	Env        string         `json:"env" required:"false"`            //All
	Time       time.Time      `json:"timestamp" required:"true"`       //All
	Debug      bool           `json:"debugging" required:"true"`       //All
	Method     string         `json:"requestMethod" required:"false"`  //All
	URL        string         `json:"requestURL" required:"false"`     //All
	StatusCode string         `json:"statusCode" required:"false"`     //Info/Error
	Latency    string         `json:"requestLatency" required:"false"` //Info
	ClientIP   string         `json:"clientIP" required:"false"`       //All
	Response   *http.Response `json:"response" required:"false"`       //Fatal
	User       string         `json:"user" required:"true"`            //All
	Message    string         `json:"message" required:"true"`         //Fatal
	LogType    string         `json:"logType" required:"true"`         //All
}

var (
	start  = []byte{byte('{'), byte('"')}
	end    = []byte{byte('"'), byte('}')}
	br     = []byte{byte('"'), byte(':'), byte('"')}
	se     = []byte{byte('"'), byte(','), byte('"')}
	unkown = []byte("Unkown")
)
