package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rccldemo.com/service/models"
)



//
//  Handler for GET /reservations/{vdsId}
//
func FindShipHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	shipCode := params["shipcode"]

	w.Header().Set("Content-Type", "application/json")

	resultSlice := []models.ShipInfo{}


	shipInfo := models.ShipInfo{}
	shipInfo.ShipCode = shipCode
	shipInfo.ShipName = "Allure"
	shipInfo.ShipClass = "Monster Class"


	// Append results to slice
	resultSlice = append(resultSlice, shipInfo)



	// Return results to caller
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(resultSlice)
	w.Write(json)

}




