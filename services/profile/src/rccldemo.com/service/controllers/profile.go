package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rccldemo.com/service/helpers"
	"rccldemo.com/service/models"
	"rccldemo.com/service/remote"
	"rccldemo.com/structlog"
	"time"
)

func CallServiceHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	vdsId := params["vdsId"]

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")  // Hack for web demo!



	start := time.Now()
	logCxt := map[string]interface{}{ "vdsId" : vdsId }
	const traceId = "abcd1234"
	const service = "booking"
	const operation = "bookings"
	const method = "GET"

	// Get tracing headers for Jaeger to ses spans without breaking between services.
	var traceHdrInfo *helpers.TraceHeaders = helpers.BuildFromRequestHeader(r)

	data, err := remote.CallRemoteBookingService(traceHdrInfo, vdsId)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		helpers.LogError(vdsId, traceId, "Failed calling bookings.",
			err, "ProfileErrBookingFetch", structlog.GetSrcLocation() , nil)

		errResp := models.ErrorResponse{
			ErrorMsg: "Error fetching profile booking information.",
		}


		w.Write(errResp.ToJsonBytes())

		// Log the latency, even though an error calling remote service
		defer helpers.LogServiceMetric(start, helpers.GetElapsed(start), vdsId,
			traceId , service, operation, method, http.StatusInternalServerError, logCxt )
		return
	}

	// Log the latency, with success from calling remote service
	defer helpers.LogServiceMetric(start, helpers.GetElapsed(start), vdsId,
		traceId , service, operation, method, http.StatusOK, logCxt )

	// Convert back into data structure
	var res []models.Reservation
	if err := json.Unmarshal(data, &res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		helpers.LogError(vdsId, traceId, "Error parsing booking service call results.",
			err, "ProfileErrBookingFetchParse", structlog.GetSrcLocation() , nil)

		errResp := models.ErrorResponse{
			ErrorMsg: "Error fetching profile booking information.",
		}

		w.Write(errResp.ToJsonBytes())
		return
	}


	profile := models.Profile{}
	profile.Reservations = res
	profile.FirstName = "Brian"
	profile.LastName = "McGinnis"

	json, _ := json.Marshal(profile)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}



