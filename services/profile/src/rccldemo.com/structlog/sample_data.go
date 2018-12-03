package structlog

import (
	"errors"
	"log"
	"time"
)


func DemoLogError() {


	 traceId := "12345"
	 vdsId := "bjm100"

	 base := Base{
		EventType: Error,
		TraceId: traceId,
		VdsId: vdsId,
		DateTime: time.Now(),
		Feature: "Core-Content",
	}

	 svcInfo := ServiceInfo{
	 	Service: "ContentService",
	 	Operation: "GetShipInfo",
	 	Version: "1.0.0",
	 	Host: "localhost",
	 	DataCenter: "laptop",
	 	Cloud: "",
	 	Env: "dev",
	 }

	// throw an error
	err := errors.New("Simulated error occurred here.")

	context := make(map[string]interface{})
	context["shipcode"] = "AX"

	errInfo := ErrorInfo{
		ErrId: "ShipNoData",
		ErrMsg: "Ship not found",
		Blame: []string{"UserInput"},
		Context: context,
		Validations: []string{},
		Stack: ErrorToString(err),
		ErrRate: true,
		ExtErrId: "Core-13AF34",
	}


	errEvent := &ErrorEvent{
		Base: base,
		ServiceInfo: svcInfo,
		ErrorInfo: errInfo,
	}

	log.Println(errEvent.ToJson())


}


