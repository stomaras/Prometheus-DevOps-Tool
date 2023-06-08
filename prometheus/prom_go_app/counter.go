package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
"github.com/prometheus/client_golang/prometheus/promhttp": help us with the creation of HTTP.handler instances to expose Prometheus Metrics via HTTP
with this package we will be able to expose the metrics on HTTP.
"github.com/prometheus/client_golang/prometheus/promauto": this sub package provide metrics constructors with automatic registration
*/

var REQUEST_COUNT = promauto.NewCounter(prometheus.CounterOpts{
	Name: "go_app_requests_count",
	Help: "Total App HTTP Requests count.",
})

func main() {
	// Start the application
	startMyApp()
}

func startMyApp() {
	router := mux.NewRouter()
	router.HandleFunc("/birthday/{name}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		greetings := fmt.Sprintf("Happy Birthday %s :)", name)
		rw.Write([]byte(greetings))

		REQUEST_COUNT.Inc()
	}).Methods("GET")

	log.Println("Starting the application server...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":8000", router)
}
