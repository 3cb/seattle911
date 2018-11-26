package main

// FireCall describes data structure for each 911 fire call
type FireCall struct {
	Address        string `json:"address"`
	DateTime       string `json:"datetime"`
	IncidentNumber string `json:"incident_number"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	Type           string `json:"type"`
}
