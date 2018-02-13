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

	wsp, err := ssc.NewSocketPool([]string{}, time.Second*0)
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
	r.Handle("/ws", wsHandler(db, wsp, upgrader))

	r.Handle("/api/day/{date}", queryHandler(db))

	// start server
	log.Fatal(http.ListenAndServe(":3030", r))
}

func wsHandler(db *bolt.DB, pool *ssc.SocketPool, upgrader *websocket.Upgrader) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := pool.AddClientSocket(upgrader, w, r)
		if err != nil {
			log.Printf("Unable to create client websocket connection: %v\n", err)
		} else {
			log.Printf("New websocket client connected.")
		}
	})
}

func queryHandler(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		date := vars["date"]
		buf := queryDB(db, date)
		w.Write(buf)
	})
}
