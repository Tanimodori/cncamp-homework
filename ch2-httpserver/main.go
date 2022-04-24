package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var port = ":8888"
	http.HandleFunc("/healthz", healthzHandler)
	fmt.Println("Starting http server listening at", port)
	http.ListenAndServe(port, nil)
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
