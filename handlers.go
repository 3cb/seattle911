package main

import (
	"log"
	"net/http"

	"github.com/3cb/ssc"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// WS handlers request to upgrade to a websocket connection
// it adds the connection to the websocket Pool(ssc.Pool)
func WS(db *bolt.DB, pool *ssc.Pool, upgrader *websocket.Upgrader) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := pool.AddClientSocket("", upgrader, w, r)
		if err != nil {
			log.Printf("Unable to create client websocket connection: %v\n", err)
		} else {
			log.Printf("New websocket client connected.")
		}
	})
}

// SingleDate queries database for 911 calls on date in request
func SingleDate(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		date := vars["date"]
		buf := queryDB(db, date)
		w.Write(buf)
	})
}
