package helpers

import (
	"rccl.com/structlog"
	"time"
)

var DefaultHost string = "unknown"

func DefaultSvcInfo() structlog.ServiceInfo {
	return structlog.ServiceInfo{
		Service: "RezService",
		Operation: "",
		Version: "1.0.0",
		Host: DefaultHost,
		DataCenter: "laptop",
		Cloud: "",
		Env: "dev",
	}
}


func LogConfig(port string, shipInfoUri string, shipClassUri string) {



	base := structlog.Base{
		EventType: structlog.Config,
		TraceId: "",
		VdsId: "",
		DateTime: time.Now(),
		Feature: "Reservations",
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

