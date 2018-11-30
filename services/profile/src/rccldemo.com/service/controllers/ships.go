package controllers

import (
	"fmt"
	"net/http"
	"rccldemo.com/service/remote"
)

func CallServiceHandler(w http.ResponseWriter, r *http.Request) {

	data, err := remote.CallRemoteShips()
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("<html><body>Error: %v</body></html>", err)
		w.Write([]byte(msg))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}



