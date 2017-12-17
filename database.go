package main

import (
	"fmt"

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

func updateDB(db *bolt.DB, date string, msg []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Messages"))
		err := b.Put([]byte(date), msg)
		if err != nil {
			return err
		}
		return nil
	})
}
