package main

import (
	"log"
	"net/http"

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

	// launch empty websocket pool
	config := ssc.PoolConfig{
		IsReadable: true,
		IsWritable: true,
		IsJSON:     false,
	}
	wsp, err := ssc.NewSocketPool(config)
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
	r.Handle("/ws", wsHandler(wsp, upgrader))

	// start server
	log.Fatal(http.ListenAndServe(":3030", r))
}

func wsHandler(pool *ssc.SocketPool, upgrader *websocket.Upgrader) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Unable to connect to new client via websocket.")
			return
		}
		s := ssc.NewSocketInstance("", pool.Config)
		s.Connection = conn
		pool.ReadStack[s] = true
		pool.WriteStack[s] = true
		log.Printf("New client connected via websocket.")
	})
}
