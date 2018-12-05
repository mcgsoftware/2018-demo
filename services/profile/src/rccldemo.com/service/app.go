package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"rccldemo.com/service/controllers"
	"rccldemo.com/service/helpers"
	"time"
)

func main() {

	host := os.Getenv("HOST")
	if (host == "") {
		host = "127.0.0.1"   // Must be this or Istio wont work
	}

	port := os.Getenv("PORT")
	if (port == "") {
		port = "8082"
	}
	address := host + ":" + port

	sampleUrl := "http://" + address + "/royal/api/profile/bjm100"
	helpers.LogConfig(host, port, sampleUrl)


	server := &http.Server{
		Addr:         address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// setup router
	// HACK: Due to Istio's ingress gateway, we have to prepend a path that will match in the
	// bookinfo gateway config file: 'bookinfo-gateway-mcg.yaml'.
	router := mux.NewRouter()
	router.HandleFunc("/royal/api/profile/{vdsId}", controllers.CallServiceHandler)
	router.HandleFunc("/royal/api/profile/health", controllers.HealthHandler)


	http.Handle("/", router)


	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, "Server error: %v", err)
	}
}
