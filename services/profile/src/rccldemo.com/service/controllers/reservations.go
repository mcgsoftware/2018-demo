package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rccldemo.com/service/models"
	"time"
)


//
//  Handler for GET /reservations/{vdsId}
//
func FindReservationHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	vdsId := params["vdsId"]


	w.Header().Set("Content-Type", "application/json")



	resultSlice := []models.ReservationInfo{}


	rez := models.Reservation{}
	rez.Shipcode = "AL"
	rez.Saildate = time.Now()
	rez.Vdsid = vdsId
	rez.Id = 100

	shipInfo := &models.ShipInfo{}
	shipInfo.ShipName = "Allure"
	shipInfo.ShipClass = "Monster Class"

	result := models.ReservationInfo{}
	result.Build(rez, shipInfo)

	// Append results to slice
	resultSlice = append(resultSlice, result)



	// Return results to caller
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(resultSlice)
	w.Write(json)

}




