package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"rccldemo.com/service/controllers"
	"time"
)

func main() {



	port := os.Getenv("PORT")
	if (port == "") {
		fmt.Println("No PORT env set, using 8082");
		port = "8082"
	}

	// Important: Host must be localhost or this will never work in K8 with Istio proxy!!
	// It will be impossible to kubectl port-forward or use gateway ingress to access the service.
	host := "127.0.0.1"
	address := host + ":" + port
	fmt.Printf("Running on http://%v/royal/api/profile/bjm100\n", address)
	fmt.Printf("http://%v/royal/api/profile/mysql/vdsId - call mysql directly\n", address)

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
	router.HandleFunc("/royal/api/profile/mysql/{vdsId}", controllers.FindReservationMySqlHandler)
	router.HandleFunc("/royal/api/profile/{vdsId}", controllers.CallServiceHandler)
	router.HandleFunc("/royal/api/profile/health", controllers.HealthHandler)


	http.Handle("/", router)


	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, "Server error: %v", err)
	}
}
