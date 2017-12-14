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
