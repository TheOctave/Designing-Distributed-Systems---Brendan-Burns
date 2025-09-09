package main

import (
	"log"
	"net/http"
	"topz/pkg/topz"
)

func main() {
	http.HandleFunc("/tops", topz.HandleRequest)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8088", nil))
}
