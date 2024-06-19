package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api/time", getCurrentTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
