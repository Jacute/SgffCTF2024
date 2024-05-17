package api

import "github.com/gorilla/mux"

func InitHandlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/archivator", getArchive).Methods("GET")
	router.HandleFunc("/api/archivator", createArchive).Methods("POST")

	return router
}
