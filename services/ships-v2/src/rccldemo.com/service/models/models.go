package models



import "time"

//
// Struct that maps to mysql table 'reservations'
//
type Reservation struct {
	Id int
	Shipcode string
	Vdsid string
	Saildate time.Time
}


//
// ship info from remote content service
// (e.g. calling http://localhost:8080/api/shipinfo/v1/AL )
//
type ShipInfo struct {
	ShipCode string `json:shipCode`
	ShipName string `json:shipName`
	ShipClass string `json:shipClass`
}

//
// Data this service sends back to callers.
//
//
type ReservationInfo struct {
	VdsId string
	Saildate time.Time
	ShipCode string
	ShipName string
	ShipClass string
}

// Builder for ReservationInfo from data we got from MySQL and remote content service.
func (this *ReservationInfo) Build(rez Reservation, shipInfo *ShipInfo) {
	this.VdsId = rez.Vdsid
	this.Saildate = rez.Saildate
	this.ShipClass = shipInfo.ShipClass
	this.ShipName = shipInfo.ShipName
	this.ShipCode = rez.Shipcode
}

