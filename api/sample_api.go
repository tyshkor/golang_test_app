package api

import "net/http"

// NewHTTPHandler ...
func SampleAPIInit() {

	http.HandleFunc("/api1", HandlerAPI1)
	http.HandleFunc("/api2/json", HandlerAPI2JSON)
	http.HandleFunc("/api2/xml", HandlerAPI2XML)
}
