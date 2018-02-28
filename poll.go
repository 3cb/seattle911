package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/3cb/ssc"
	"github.com/boltdb/bolt"
)

func poll(db *bolt.DB, pool *ssc.Pool) {
	ticker := time.NewTicker(time.Minute * 5)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	d := strings.Split(fmt.Sprint(time.Now().In(loc)), " ")[0]
	f := []FireCall{}
	p := []PoliceCall{}

	t1, t2 := getDate()
	f, err := updateFire(t1, t2)
	if err != nil {
		log.Printf("Unable to update fire calls: %v", err)
	} else {
		log.Printf("Fire updated")
	}
	p, err = updatePolice(t1, t2)
	if err != nil {
		log.Printf("Unable to update police calls: %v", err)
	} else {
		log.Printf("Police updated")
	}

	buf := serialize(f, p, d)
	err = updateDB(db, d, buf)
	if err != nil {
		log.Printf("Unable to save to database: %v\n", err)
	}
	pool.Inbound <- &ssc.Message{Type: 2, Payload: buf}

	for {
		<-ticker.C
		d := strings.Split(fmt.Sprint(time.Now().In(loc)), " ")[0]

		t1, t2 := getDate()
		f, err = updateFire(t1, t2)
		if err != nil {
			log.Printf("Unable to update fire calls: %v", err)
		} else {
			log.Printf("Fire updated")
		}
		p, err = updatePolice(t1, t2)
		if err != nil {
			log.Printf("Unable to update police calls: %v", err)
		} else {
			log.Printf("Police updated")
		}
		buf := serialize(f, p, d)
		err = updateDB(db, d, buf)
		if err != nil {
			log.Printf("Unable to save to database: %v\n", err)
		}
		pool.Inbound <- &ssc.Message{Type: 2, Payload: buf}
	}
}

func getDate() (string, string) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	t1 := strings.Split(fmt.Sprint(time.Now().In(loc)), " ")[0] + "T00:00:00"
	t2 := strings.Split(fmt.Sprint(time.Now().In(loc).Add(time.Hour*24)), " ")[0] + "T00:00:00"
	return t1, t2
}
