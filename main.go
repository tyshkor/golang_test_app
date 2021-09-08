package main

import (
	"log"
	"net/http"

	"github.com/tyshkor/golang_test_app/api"
)

func main() {
	api.SampleAPIInit()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
