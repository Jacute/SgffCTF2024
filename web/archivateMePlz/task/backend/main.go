package main

import (
	"api/api"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Mkdir("uploads", 0755)
	os.Mkdir("archives", 0755)

	go runCleaner("archives")

	router := api.InitHandlers()

	log.Fatal(http.ListenAndServe(":5555", router))
}
