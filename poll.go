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
	d := strings.Split(fmt.Sprint(time.Now()), " ")[0]
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
	pool.Pipes.InboundBytes <- ssc.Data{Type: 2, Payload: buf}

	for {
		date := <-ticker.C
		d := strings.Split(fmt.Sprint(date), " ")[0]
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
		pool.Pipes.InboundBytes <- ssc.Data{Type: 2, Payload: buf}
	}
}

func updateFire() ([]FireCall, error) {
	f := []FireCall{}

	t1, t2 := getDate(time.Now())
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

	t1, t2 := getDate(time.Now())
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

func getDate(day time.Time) (string, string) {
	t := strings.Split(fmt.Sprint(day), " ")
	t1 := t[0] + "T00:00:00"
	t = strings.Split(fmt.Sprint(day.Add(time.Hour*24)), " ")
	t2 := t[0] + "T00:00:00"
	return t1, t2
}
