package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//
//  Handler for post log
//
func LogHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")  // Hack for web demo!

	posted, err := ioutil.ReadAll(r.Body)
	if (err != nil) || (len(posted) < 10) {
		var nuErr error = fmt.Errorf("Error reading the body", err)
		fmt.Println(nuErr)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the error we just got to stderr
	fmt.Fprintln(os.Stderr, string(posted))


	// Return log event we got to caller so they know what we logged.
	w.WriteHeader(http.StatusOK)
	w.Write(posted)

}

//
//  Handler for post log
//

/* NOT_USED parses input. Better for tightly checked logging.
func ErrorLogHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")  // Hack for web demo!

	decoder := json.NewDecoder(r.Body)

	var data structlog.ErrorEvent
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Failed to parse log event.", err)
		w.Write([]byte(msg))
		return
	}





	// Return results to caller
	w.WriteHeader(http.StatusOK)

}
*/

