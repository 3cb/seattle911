package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/3cb/ssc"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

func main() {
	// open database and create buckets of key/value pairs
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = createBucket(db, "Messages")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	wsp := ssc.NewPool([]string{}, time.Second*0)
	err = wsp.Start()
	if err != nil {
		log.Fatal("Unable to start pool to serve websocket connections.")
	}

	// start 5 minute polling goroutine
	go poll(db, wsp)

	// create new router instance
	r := mux.NewRouter()

	// serve static files
	r.Handle("/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/dist/").Handler(http.FileServer(http.Dir("./static/")))

	// handle websocket requests
	upgrader := &websocket.Upgrader{}
	r.Handle("/ws", WS(db, wsp, upgrader))
	r.Handle("/api/day/{date}", SingleDate(db))

	// start server
	log.Fatal(http.ListenAndServe(":3030", r))
}
