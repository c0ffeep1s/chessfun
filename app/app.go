package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/make_move", makemove).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
