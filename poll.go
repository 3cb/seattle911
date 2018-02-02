package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/3cb/ssc"
	"github.com/boltdb/bolt"
)

func poll(db *bolt.DB, pool *ssc.SocketPool) {
	ticker := time.NewTicker(time.Minute * 5)
	loc, _ := time.LoadLocation("America/Los_Angeles")
	d := strings.Split(fmt.Sprint(time.Now().In(loc)), " ")[0]
	f := []FireCall{}
	p := []PoliceCall{}

	f, err := updateFire()
	if err != nil {
		log.Printf("Unable to update fire calls: %v", err)
	} else {
		log.Printf("Fire updated")
	}
	p, err = updatePolice()
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
	pool.Pipes.Inbound <- ssc.Message{Type: 2, Payload: buf}

	for {
		<-ticker.C
		d := strings.Split(fmt.Sprint(time.Now().In(loc)), " ")[0]
		f, err = updateFire()
		if err != nil {
			log.Printf("Unable to update fire calls: %v", err)
		} else {
			log.Printf("Fire updated")
		}
		p, err = updatePolice()
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
		pool.Pipes.Inbound <- ssc.Message{Type: 2, Payload: buf}
	}
}

func updateFire() ([]FireCall, error) {
	f := []FireCall{}

	t1, t2 := getDate()
	api := "https://data.seattle.gov/resource/grwu-wqtk.json?$where="
	queryString := url.QueryEscape("datetime between '" + t1 + "' and '" + t2 + "'")
	resp, err := http.Get(api + queryString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func updatePolice() ([]PoliceCall, error) {
	p := []PoliceCall{}

	t1, t2 := getDate()
	api := "https://data.seattle.gov/resource/pu5n-trf4.json?$where="
	queryString := url.QueryEscape("at_scene_time between '" + t1 + "' and '" + t2 + "'")
	resp, err := http.Get(api + queryString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func getDate() (string, string) {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	t1 := strings.Split(fmt.Sprint(time.Now().In(loc)), " ")[0] + "T00:00:00"
	t2 := strings.Split(fmt.Sprint(time.Now().In(loc).Add(time.Hour*24)), " ")[0] + "T00:00:00"
	return t1, t2
}
