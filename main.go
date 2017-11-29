package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create new router instance
	r := mux.NewRouter()

	// serve static files
	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/dist/").Handler(http.FileServer(http.Dir("./static/")))

	// start server
	log.Fatal(http.ListenAndServe(":3030", r))
}
