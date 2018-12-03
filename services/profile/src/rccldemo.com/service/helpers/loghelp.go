package helpers

import (
	"fmt"
	"rccldemo.com/structlog"
	"time"
)

var DefaultHost string = "unknown"

func DefaultSvcInfo() structlog.ServiceInfo {
	return structlog.ServiceInfo{
		Service: "profile",
		Operation: "",
		Version: "1.10",
		Host: "127.0.0.1",
		DataCenter: "Cloud",
		Cloud: "Google",
		Env: "dev",
	}
}


func LogConfig(port string, shipInfoUri string, shipClassUri string) {



	base := structlog.Base{
		EventType: structlog.Config,
		TraceId: "",
		VdsId: "",
		DateTime: time.Now(),
		Feature: "Profile",
	}

	svcInfo := DefaultSvcInfo()

	// Setup to log config properties
	properties := make(map[string]interface{})
	properties["PORT"] = port;
	properties["CONTENT_SERVICE_SHIP_CLASS"] = shipClassUri
	properties["CONTENT_SERVICE_SHIP_INFO"] = shipInfoUri



	configInfo := structlog.ConfigInfo{
		Properties: properties,
	}

	configEvt := &structlog.ConfigEvent{
		Base: base,
		ServiceInfo: svcInfo,
		ConfigInfo: configInfo,
	}


	structlog.GetLogger().Println(configEvt.ToJson())


}


func LogError(vdsId string, traceId string, err error, errId string, errmsg string, cxt map[string]interface{} ) {

	base := structlog.Base{
		EventType: structlog.Error,
		TraceId: traceId,
		VdsId: vdsId,
		DateTime: time.Now(),
		Feature: "Reservations",
	}

	svcInfo := DefaultSvcInfo()



	errInfo :=structlog.ErrorInfo{
		ErrId: errId,
		ErrMsg: errmsg,
		Blame: []string{},
		Context: cxt,
		Validations: nil,
		Stack: structlog.ErrorToString(err),
		ErrRate: true,
		ExtErrId: "",
	}

	errEvent := &structlog.ErrorEvent{
		Base: base,
		ServiceInfo: svcInfo,
		ErrorInfo: errInfo,
	}

	structlog.GetLogger().Println(errEvent.ToJson())

}

//
// Get time elapsed in milliseconds since start time.
//
func GetElapsed(start time.Time) int64 {
	// Get elapsed time in millseconds
	return int64(time.Since(start) / time.Millisecond)
}

//
// Includes timer wrapper
// Example use:
//    start := time.Now()
//    ... do things that take time...
//    LogServiceMetric(start, GetElapsed(start), vdsId ...)
//

func LogServiceMetric(start time.Time, elapsedInMillis int64, vdsId string, traceId string, service string,
	operation string,
	method string, status int, context map[string]interface{}) {





	base := structlog.Base{
		EventType: structlog.ServiceMetric,
		TraceId: traceId,
		VdsId: vdsId,
		DateTime: start,
		Feature: "Profile",
	}

	svcInfo := DefaultSvcInfo()

	metric := structlog.ServiceMetricInfo{
		Service: service,
		Operation: operation,
		Method: method,
		Latency: elapsedInMillis,
		Status: status,
		Tags: nil,
		Context: context,
	}


	metricEvent := &structlog.ServiceMetricEvent{
		Base: base,
		ServiceInfo: svcInfo,
		Metric: metric,
	}

	//structlog.GetLogger().Println(errEvent.ToJson())
	// Log to stdout
	fmt.Println(metricEvent.ToJson())

}

