package main

import (
	"fmt"
	"httpserver/metrics"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var port = ":8080"
	metrics.Register()
	mux := http.NewServeMux()

	mux.HandleFunc("/delay", delay)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", healthzHandler)
	fmt.Println("Starting http server listening at", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// set http code 200
	var code = 200
	// copy header to response
	for key := range r.Header {
		w.Header().Set(key, r.Header.Get(key))
	}
	// read Version and send to response
	var VERSION = "VERSION"
	w.Header().Set(VERSION, os.Getenv(VERSION))
	// print client ip and code
	var ROUTE = "GET /healthz"
	fmt.Println(ROUTE, code, "from", r.RemoteAddr)
	// flush result
	w.WriteHeader(code)
	fmt.Fprintln(w, ROUTE)
}

func delay(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("%d", randInt)))
}
