package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rccldemo.com/service/helpers"
	"rccldemo.com/service/models"
	"rccldemo.com/service/remote"
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

	data, err := remote.CallRemoteBookingService(vdsId)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("<html><body>Error: %v</body></html>", err)
		w.Write([]byte(msg))

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
		msg := fmt.Sprintf("Error parsing remote service call results: %v", err)
		w.Write([]byte(msg))
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



