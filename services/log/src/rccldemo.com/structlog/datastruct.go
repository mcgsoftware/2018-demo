package structlog

import (
	"bytes"
	"encoding/json"
	"time"
)

//
// Define event type enum.
//
type EventType int

const (
	Error EventType = iota
	Metric
	ServiceMetric
	Config
)


func (etype EventType) String() string {
	switch etype {
	case Error:
		return "Error"
	case Metric:
		return "Metric"
	case ServiceMetric:
		return "ServiceMetric"
	case Config:
		return "Config"
	default:
		return "Unknown"
	}
}

var fromString = map[string]EventType{
	"Error":  			Error,
	"ServiceMetric":  	ServiceMetric,
	"Metric": 			Metric,
	"Config": 			Config,
}

// MarshalJSON marshals the enum as a quoted json string
func (s *EventType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *EventType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value,
	// 'Error' in this case.
	*s = fromString[j]
	return nil
}

//
// Base struct for log events
//
type Base struct {
	EventType EventType `json:"eventType"`
	TraceId   string    `json:"traceId,omitempty"`
	VdsId     string    `json:"userId,omitempty"`
	DateTime  time.Time `json:"datetime"`
	Feature   string    `json:"feature,omitempty"`
}
func (evt *Base) ToJson() string {
	jsonBytes, _ := json.Marshal(&evt)
	return (string(jsonBytes))
}


//
// ServiceInfo struct for log events
//
type ServiceInfo struct {
	Service     string `json:"service"`
	Operation   string `json:"operation"`
	Version     string `json:"version"`
	Host        string `json:"host"`
	DataCenter  string `json:"datacenter"`
	Cloud       string `json:"cloud,omitempty"`
	Env         string `json:"env"`
}
func (evt *ServiceInfo) ToJson() string {
	jsonBytes, _ := json.Marshal(&evt)
	return (string(jsonBytes))
}


//
// ErrorInfo struct for log events
//
type ErrorInfo struct {
	ErrId       string                   `json:"errId"`
	ErrMsg      string                   `json:"errMsg"`
	Blame       []string                 `json:"blame,omitempty"`
	Context     map[string]interface{}   `json:"context,omitempty"`
	Validations []string                 `json:"validations,omitempty"`
	Stack       string                   `json:"stack,omitempty"`
	ErrRate     bool                     `json:"errRate,omitempty"`
	ExtErrId	string                   `json:"extErrId,omitempty"`
}
func (evt *ErrorInfo) ToJson() string {
	jsonBytes, _ := json.Marshal(&evt)
	return (string(jsonBytes))
}

//
// Config struct for log events
//
type ConfigInfo struct {
	Properties     map[string]interface{}   `json:"properties"`
}
func (evt *ConfigInfo) ToJson() string {
	jsonBytes, _ := json.Marshal(&evt)
	return (string(jsonBytes))
}

//=========================================
// Event types
//=========================================


//
// Error log event type
//
type ErrorEvent struct {
	Base         Base         `json:"event"`
	ServiceInfo  ServiceInfo  `json:"serviceInfo"`
	ErrorInfo    ErrorInfo    `json:"error"`
}
func (evt *ErrorEvent) ToJson() string {
	jsonBytes, _ := json.Marshal(&evt)
	return (string(jsonBytes))
}


//
// Config log event type
//
type ConfigEvent struct {
	Base         Base         `json:"event"`
	ServiceInfo  ServiceInfo  `json:"serviceInfo"`
	ConfigInfo   ConfigInfo   `json:"config"`
}
func (evt *ConfigEvent) ToJson() string {
	jsonBytes, _ := json.Marshal(&evt)
	return (string(jsonBytes))
}

