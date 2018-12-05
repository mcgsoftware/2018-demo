package models

import (
	"encoding/json"
	"time"
)

//
// Struct that maps to mysql table 'reservations'
//

type Reservation struct {
	VdsId     string    `json:"vdsId"`
	ShipCode  string    `json:"shipCode"`
	ShipName  string    `json:"shipName"`
	ShipClass string    `json:"shipClass"`
	SailDate  time.Time `json:"sailDate"`
}

type ErrorResponse struct {
	ErrorMsg     string    `json:"errMsg"`
}
func (evt *ErrorResponse) ToJson() string {
	jsonBytes, _ := json.Marshal(&evt)
	return (string(jsonBytes))
}
func (evt *ErrorResponse) ToJsonBytes() []byte {
	jsonBytes, _ := json.Marshal(&evt)
	return (jsonBytes)
}