package main

import (
	"fmt"

	"github.com/3cb/seattle911/seattle"
	"github.com/boltdb/bolt"
)

func createBucket(db *bolt.DB, name string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err2 := tx.CreateBucketIfNotExists([]byte(name))
		if err2 != nil {
			return fmt.Errorf("Error creating bucket: %s", err2)
		}
		return nil
	})
}

func updateDB(date string, msg seattle.Message) error {

	return nil
}
