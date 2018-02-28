package main

import (
	"log"
	"net/http"
	"strings"

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
// Date should be in format of "YYYY-MM-DD"
func SingleDate(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		date := vars["date"]
		buf := queryDB(db, date)
		if len(buf) == 0 {
			t1, t2, err := setQueryRange(date, "day")
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("invalid date string"))
				return
			}
			f, err := updateFire(t1, t2)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("unable to get data"))
				return
			}
			p, err := updatePolice(t1, t2)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("unable to get data"))
				return
			}
			buf = serialize(f, p, date)
			err = updateDB(db, date, buf)
			if err != nil {
				log.Printf("Unable to save to database: %v\n", err)
			}
		}
		w.Write(buf)
	})
}

// Month queries Socrata API for call data for requested date
func Month(db *bolt.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		date := vars["month"]
		t1, t2, err := setQueryRange(date, "month")
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("invalid date string"))
			return
		}
		f, err := updateFire(t1, t2)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("unable to get data"))
			return
		}
		p, err := updatePolice(t1, t2)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("unable to get data"))
			return
		}

		dateRange := date + "~" + strings.Split(t2, "T")[0]
		buf := serialize(f, p, dateRange)
		w.Write(buf)
	})
}
