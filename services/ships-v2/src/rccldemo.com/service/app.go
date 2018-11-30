package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"rccldemo.com/service/controllers"
	"time"
)

func main() {

	port := os.Getenv("PORT")
	if (port == "") {
		log.Println("No PORT env set, using 8070");
		port = "8070"
	}

	// Important: Host must be localhost or this will never work in K8 with Istio proxy!!
	// It will be impossible to kubectl port-forward or use gateway ingress to access the service.
	host := "127.0.0.1"
	address := host + ":" + port
	fmt.Printf("Running on http://%v/royal/api/ships/AL \n", address)


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
	router.HandleFunc("/royal/api/ships/{shipcode:[A-Z][A-Z]}", controllers.FindShipHandler)
	router.HandleFunc("/royal/api/ships/health", controllers.HealthHandler)

	/* These will not be accessible in the
	router.HandleFunc("/health", controllers.HealthHandler)
	router.HandleFunc("/getprof/{vdsId:[0-9]+}", controllers.FindReservationHandler)
	*/
	http.Handle("/", router)


	log.Println(server.ListenAndServe())
}
