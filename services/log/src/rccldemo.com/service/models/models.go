package models



import "time"

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

