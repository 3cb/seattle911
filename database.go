package main

import (
	"fmt"
	"strings"

	"github.com/3cb/seattle911/seattle"
	flatbuffers "github.com/google/flatbuffers/go"
)

func serialize(fireCalls []FireCall, policeCalls []PoliceCall) []byte {
	builder := flatbuffers.NewBuilder(1024)
	seattle.MessageStartFireCallsVector(builder, len(fireCalls))
	for i := len(fireCalls) - 1; i >= 0; i-- {
		address := builder.CreateString(fireCalls[i].Address)
		datetime := builder.CreateString(fireCalls[i].DateTime)
		incidentNumber := builder.CreateString(fireCalls[i].IncidentNumber)
		latitude := builder.CreateString(fireCalls[i].Latitude)
		longitude := builder.CreateString(fireCalls[i].Longitude)
		tp := builder.CreateString(fireCalls[i].Type)

		seattle.FireCallStart(builder)
		seattle.FireCallAddAddress(builder, address)
		seattle.FireCallAddDatetime(builder, datetime)
		seattle.FireCallAddIncidentNumber(builder, incidentNumber)
		seattle.FireCallAddLatitude(builder, latitude)
		seattle.FireCallAddLongitude(builder, longitude)
		seattle.FireCallAddType(builder, tp)
		fire := seattle.FireCallEnd(builder)

		builder.PrependUOffsetT(fire)
	}
	fCalls := builder.EndVector(len(fireCalls))
	// builder.Reset()

	seattle.MessageStartPoliceCallsVector(builder, len(policeCalls))
	for j := len(policeCalls) - 1; j >= 0; j-- {
		atSceneTime := builder.CreateString(policeCalls[j].AtSceneTime)
		cadCdwID := builder.CreateString(policeCalls[j].CADCDWID)
		cadEventNumber := builder.CreateString(policeCalls[j].CADEventNumber)
		censusTract := builder.CreateString(policeCalls[j].CensusTract)
		districtSector := builder.CreateString(policeCalls[j].DistrictSector)
		eventClearanceCode := builder.CreateString(policeCalls[j].EventClearanceCode)
		eventClearanceDate := builder.CreateString(policeCalls[j].EventClearanceDate)
		eventClearanceDescription := builder.CreateString(policeCalls[j].EventClearanceDescription)
		eventClearanceGroup := builder.CreateString(policeCalls[j].EventClearanceGroup)
		eventClearanceSubgroup := builder.CreateString(policeCalls[j].EventClearanceSubgroup)
		generalOffenseNumber := builder.CreateString(policeCalls[j].GeneralOffenseNumber)
		hundredBlockLocation := builder.CreateString(policeCalls[j].HundredBlockLocation)
		initialTypeDescription := builder.CreateString(policeCalls[j].InitialTypeDescription)
		initialTypeGroup := builder.CreateString(policeCalls[j].InitialTypeGroup)
		initialTypeSubgroup := builder.CreateString(policeCalls[j].InitialTypeSubGroup)
		latitude := builder.CreateString(policeCalls[j].Latitude)
		longitude := builder.CreateString(policeCalls[j].Longitude)
		zoneBeat := builder.CreateString(policeCalls[j].ZoneBeat)

		seattle.PoliceCallStart(builder)
		seattle.PoliceCallAddAtSceneTime(builder, atSceneTime)
		seattle.PoliceCallAddCadCdwId(builder, cadCdwID)
		seattle.PoliceCallAddCadEventNumber(builder, cadEventNumber)
		seattle.PoliceCallAddCensusTract(builder, censusTract)
		seattle.PoliceCallAddDistrictSector(builder, districtSector)
		seattle.PoliceCallAddEventClearanceCode(builder, eventClearanceCode)
		seattle.PoliceCallAddEventClearanceDate(builder, eventClearanceDate)
		seattle.PoliceCallAddEventClearanceDescription(builder, eventClearanceDescription)
		seattle.PoliceCallAddEventClearanceGroup(builder, eventClearanceGroup)
		seattle.PoliceCallAddEventClearanceSubgroup(builder, eventClearanceSubgroup)
		seattle.PoliceCallAddGeneralOffenseNumber(builder, generalOffenseNumber)
		seattle.PoliceCallAddHundredBlockLocation(builder, hundredBlockLocation)
		seattle.PoliceCallAddInitialTypeDescription(builder, initialTypeDescription)
		seattle.PoliceCallAddInitialTypeGroup(builder, initialTypeGroup)
		seattle.PoliceCallAddInitialTypeSubgroup(builder, initialTypeSubgroup)
		seattle.PoliceCallAddLatitude(builder, latitude)
		seattle.PoliceCallAddLongitude(builder, longitude)
		seattle.PoliceCallAddZoneBeat(builder, zoneBeat)
		police := seattle.PoliceCallEnd(builder)

		builder.PrependUOffsetT(police)
	}
	pCalls := builder.EndVector(len(policeCalls))
	// builder.Reset()
	date := builder.CreateString(strings.Split(fireCalls[0].DateTime, "T")[0])

	seattle.MessageStart(builder)
	seattle.MessageAddDate(builder, date)
	seattle.MessageAddFireCalls(builder, fCalls)
	seattle.MessageAddFireCalls(builder, pCalls)
	msg := seattle.MessageEnd(builder)

	builder.Finish(msg)
	buf := builder.FinishedBytes()
	fmt.Print(buf)

	return buf
}
