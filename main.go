package main

import (
	"fmt"
	"log"
	"net/http"

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
	wsp, err := ssc.NewSocketPool(nil, config)
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

	// start server
	log.Fatal(http.ListenAndServe(":3030", r))
}

func createBucket(db *bolt.DB, name string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err2 := tx.CreateBucketIfNotExists([]byte(name))
		if err2 != nil {
			return fmt.Errorf("Error creating bucket: %s", err2)
		}
		return nil
	})
}
