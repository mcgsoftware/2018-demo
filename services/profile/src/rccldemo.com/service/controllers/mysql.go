package controllers



import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rccldemo.com/service/models"
	_ "github.com/go-sql-driver/mysql"   // must do or will get a runtime error when connecting to mysql!
	"time"
)



//
// Gets a database connection from pool.
// Note: We tell mysql driver to convert Date and Datetime columns into time.Time
//
func connect() (db *sql.DB) {
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

	fmt.Println("connecting to mysql: ", connStr)

	db, err := sql.Open(dbDriver, connStr)

	if (err != nil) {
		panic(err.Error())
	}

	return db
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
	db := connect()
	defer db.Close()

	const sql = "select r.vdsid, r.shipcode, s.name, r.saildate, s.class from reservations r, ships s where r.shipcode = s.shipcode AND vdsid = ?"

	//const sql = "SELECT resid, vdsid, shipcode, saildate FROM reservations where vdsid = ?"
    fmt.Println("SQL: ", sql, "vdsId = ", vdsId)

	query, err := db.Query(sql, vdsId)
	if err != nil {
		// make sure to return the sql statement to caller via error stack so it's logged.
		return nil, fmt.Errorf("MySQL query failed: %s. Error: ", sql, err)
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
			return nil, err
		}

		// load fields into struct
		rez.VdsId = dvdsid
		rez.ShipCode = shipcode
		rez.SailDate = saildate
		rez.ShipName = shipname
		rez.ShipClass = shipclass

		// save struct to slice
		reservations = append(reservations, rez)

		//fmt.Println("Reservation: ", rez)
	}


	return reservations, nil

}


//
//  Handler for GET /reservations/{vdsId}
//
func FindReservationMySqlHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	vdsId := params["vdsId"]

	fmt.Println("VDSID => vdsId")

	w.Header().Set("Content-Type", "application/json")


	reservations, err := getReservations(vdsId)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("<html><body>Error: %v</body></html>", err)
		w.Write([]byte(msg))
		return
	}


	// Return results to caller
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(reservations)
	w.Write(json)


}





