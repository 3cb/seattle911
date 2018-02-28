package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// setQueryRange returns the query range t1 -> t2 and an error
// analyzes input to deterine if query is for single day, month, or year
func setQueryRange(t string, dur string) (string, string, error) {
	t1, t2 := "", ""
	switch dur {
	case "day":
		t1 = t + "T00:00:00"
		tp, err := time.Parse(time.RFC3339, t+"T00:00:00Z")
		if err != nil {
			return "", "", err
		}
		t2 = strings.Split(fmt.Sprint(tp.AddDate(0, 0, 1)), " ")[0] + "T00:00:00"
	case "month":
		t1 = t + "T00:00:00"
		tp, err := time.Parse(time.RFC3339, t1+"Z")
		if err != nil {
			return "", "", err
		}
		t2 = strings.Split(fmt.Sprint(tp.AddDate(0, 1, 0)), " ")[0] + "T00:00:00"
	case "year":
		t1 = t + "T00:00:00"
		tp, err := time.Parse(time.RFC3339, t1+"Z")
		if err != nil {
			return "", "", err
		}
		t2 = strings.Split(fmt.Sprint(tp.AddDate(1, 0, 0)), " ")[0] + "T00:00:00"
	}
	return t1, t2, nil
}

func updateFire(t1 string, t2 string) ([]FireCall, error) {
	f := []FireCall{}

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

func updatePolice(t1 string, t2 string) ([]PoliceCall, error) {
	p := []PoliceCall{}

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
