package controllers

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql" // must do or will get a runtime error when connecting to mysql!
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"rccldemo.com/service/helpers"
	"rccldemo.com/service/models"
	"rccldemo.com/structlog"
	"time"
)


const sql_query = "select r.vdsid, r.shipcode, s.name, r.saildate, s.class from reservations r, ships s where r.shipcode = s.shipcode AND vdsid = ?"


//
// Gets a database connection from pool.
// Note: We tell mysql driver to convert Date and Datetime columns into time.Time
//
func connect() (db *sql.DB, err error) {
	/* local conn
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Alamo2009"
	dbName := "demo"
	*/

	/*
	To connect to your database use these details
Server: sql9.freemysqlhosting.net
Database : sql9267914
Username: sql9267914
Password: zbhMi8NH7x
Port number: 3306
	 */
	const dbDriver = "mysql"
	const dbHost = "sql9.freemysqlhosting.net"
	const dbUser = "sql9267914"
	const dbPass = "zbhMi8NH7x"
	const dbName = "sql9267914"

	const connStr = dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)"+ "/" + dbName +"?parseTime=true"

	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")

	//user:password@tcp(localhost:5555)/dbname?charset=utf8

	// fmt.Println("connecting to mysql: ", connStr)

	db, err = sql.Open(dbDriver, connStr)

	if (err != nil) {
		return nil, errors.Wrapf(err, "MySQL connect failed: %s", connStr)
	}

	return db, nil
}

/*
mysql> select r.vdsid, r.shipcode, s.name, r.saildate, s.class from reservations r, ships s where r.shipcode = s.shipcode AND vdsid = 'bjm100';
+--------+----------+----------+------------+-------+
| vdsid  | shipcode | name     | saildate   | class |
+--------+----------+----------+------------+-------+
| bjm100 | AL       | Allure   | 2018-12-30 | Oasis |
| bjm100 | HM       | Harmony  | 2018-12-30 | Oasis |
| bjm100 | SY       | Symphony | 2019-03-14 | Oasis |
+--------+----------+----------+------------+-------+
 */

//
// Exec query to fetch reservations from MySQL by vdsId
//
func getReservations(vdsId string) ([]models.Reservation, error) {

	// open mysql conn
	db, err := connect();
	if err != nil {
		return nil, errors.Wrapf(err, "MySQL failed to get connection")
	}
	defer db.Close()

	query, err := db.Query(sql_query, vdsId)
	if err != nil {
		// make sure to return the sql statement to caller via error stack so it's logged.
		return nil, errors.Wrapf(err, "MySQL query failed: %s", sql_query)
	}



	rez := models.Reservation{}
	reservations := []models.Reservation{}
	for query.Next() {

		var dvdsid string
		var shipcode string
		var shipname string
		var shipclass string
		var saildate time.Time
		err := query.Scan(&dvdsid, &shipcode, &shipname, &saildate, &shipclass)
		if (err != nil) {
			return nil, errors.Wrapf(err,"MySQL query results scan error: %s.", sql_query )
		}

		// load fields into struct
		rez.VdsId = dvdsid
		rez.ShipCode = shipcode
		rez.SailDate = saildate
		rez.ShipName = shipname
		rez.ShipClass = shipclass

		// save struct to slice
		reservations = append(reservations, rez)

	}


	return reservations, nil

}


//
//  Handler for GET reservations/{vdsId}
//
func FindReservationMySqlHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	vdsId := params["vdsId"]


	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")  // Hack for web demo!

	start := time.Now()
	logCxt := map[string]interface{}{ "sql" : sql_query, "vdsId" : vdsId }
	const traceId = "abcd1234"
	const service = "MySQL"
	const operation = "query"
	const method = "Sql"


	reservations, err := getReservations(vdsId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		helpers.LogError(vdsId, traceId, "Failed to query reservations.",
			err, "BookingErrFetchQuery", structlog.GetSrcLocation() , nil)

		errResp := models.ErrorResponse{
			ErrorMsg: "Error querying profile",
		}

		msg := errResp.ToJsonBytes()
		w.Write(msg)

		// Log the latency, even though an error calling mysql
		defer helpers.LogServiceMetric(start, helpers.GetElapsed(start), vdsId,
			traceId , service, operation, method, http.StatusInternalServerError, logCxt )

		return
	} else {
		// Log the latency, note success calling mysql
		defer helpers.LogServiceMetric(start, helpers.GetElapsed(start), vdsId,
			traceId , service, operation, method, http.StatusOK, logCxt )
	}



	if (len(reservations) < 1) {
		w.WriteHeader(http.StatusNotFound)
		msg := "{ \"error\" : \"No bookings found for: " + vdsId + "\"}"
		w.Write([]byte(msg))
		return
	}



	// Return results to caller
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(reservations)
	w.Write(json)


}





