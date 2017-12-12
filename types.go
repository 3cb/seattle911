package main

// FireMessage describes data structure for a day of 911 fire calls as it is stored in databased and sent over websocket connection
type FireMessage struct {
	Type  string     `json:"type"`
	Date  string     `json:"date"`
	Calls []FireCall `json:"calls"`
}

// FireCall describes data structure for each 911 fire call
type FireCall struct {
	Address        string `json:"address"`
	DateTime       string `json:"date_time"`
	IncidentNumber string `json:"incident_number"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	Type           string `json:"type"`
}

// PoliceMessage describes data structure for a day of 911 police calls as it is stored in database and sent over websocket connection
type PoliceMessage struct {
	Type  string       `json:"type"`
	Date  string       `json:"date"`
	Calls []PoliceCall `json:"calls"`
}

// PoliceCall describes data structure for individual 911 police calls
type PoliceCall struct {
	AtSceneTime               string `json:"at_scene_time"`
	CADCDWID                  string `json:"cad_cdw_id"`
	CADEventNumber            string `json:"cad_event_number"`
	CensusTract               string `json:"census_tract"`
	DistrictSector            string `json:"district_sector"`
	EventClearanceCode        string `json:"event_clearance_code"`
	EventClearanceDate        string `json:"event_clearance_date"`
	EventClearanceDescription string `json:"event_clearance_description"`
	EventClearanceGroup       string `json:"event_clearance_group"`
	EventClearanceSubgroup    string `json:"event_clearance_subgroup"`
	GeneralOffenseNumber      string `json:"general_offense_number"`
	HundredBlockLocation      string `json:"hundred_block_location"`
	InitialTypeDescription    string `json:"initial_type_description"`
	InitialTypeGroup          string `json:"initial_type_group"`
	InitialTypeSubGroup       string `json:"initial_type_subgroup"`
	Latitude                  string `json:"latitude"`
	Longitude                 string `json:"longitude"`
	ZoneBeat                  string `json:"zone_beat"`
}
