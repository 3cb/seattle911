package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

func main() {
	// open database and create buckets of key/value pairs
	db, err := bolt.Open("fire.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = createBucket(db, "Fire")
	if err != nil {
		log.Fatal(err)
	}
	err = createBucket(db, "Police")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
