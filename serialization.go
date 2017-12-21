package main

import (
	"github.com/3cb/seattle911/seattle"
	flatbuffers "github.com/google/flatbuffers/go"
)

func serialize(fireCalls []FireCall, policeCalls []PoliceCall, date string) []byte {
	builder := flatbuffers.NewBuilder(1024)

	fire := []flatbuffers.UOffsetT{}
	for _, call := range fireCalls {
		address := builder.CreateString(call.Address)
		datetime := builder.CreateString(call.DateTime)
		incidentNumber := builder.CreateString(call.IncidentNumber)
		latitude := builder.CreateString(call.Latitude)
		longitude := builder.CreateString(call.Longitude)
		tp := builder.CreateString(call.Type)

		seattle.FireCallStart(builder)
		seattle.FireCallAddAddress(builder, address)
		seattle.FireCallAddDatetime(builder, datetime)
		seattle.FireCallAddIncidentNumber(builder, incidentNumber)
		seattle.FireCallAddLatitude(builder, latitude)
		seattle.FireCallAddLongitude(builder, longitude)
		seattle.FireCallAddType(builder, tp)
		fire = append(fire, seattle.FireCallEnd(builder))
	}
	seattle.MessageStartFireCallsVector(builder, len(fireCalls))
	for i := len(fire) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(fire[i])
	}
	fCalls := builder.EndVector(len(fireCalls))

	police := []flatbuffers.UOffsetT{}
	for _, call := range policeCalls {
		atSceneTime := builder.CreateString(call.AtSceneTime)
		cadCdwID := builder.CreateString(call.CADCDWID)
		cadEventNumber := builder.CreateString(call.CADEventNumber)
		censusTract := builder.CreateString(call.CensusTract)
		districtSector := builder.CreateString(call.DistrictSector)
		eventClearanceCode := builder.CreateString(call.EventClearanceCode)
		eventClearanceDate := builder.CreateString(call.EventClearanceDate)
		eventClearanceDescription := builder.CreateString(call.EventClearanceDescription)
		eventClearanceGroup := builder.CreateString(call.EventClearanceGroup)
		eventClearanceSubgroup := builder.CreateString(call.EventClearanceSubgroup)
		generalOffenseNumber := builder.CreateString(call.GeneralOffenseNumber)
		hundredBlockLocation := builder.CreateString(call.HundredBlockLocation)
		initialTypeDescription := builder.CreateString(call.InitialTypeDescription)
		initialTypeGroup := builder.CreateString(call.InitialTypeGroup)
		initialTypeSubgroup := builder.CreateString(call.InitialTypeSubGroup)
		latitude := builder.CreateString(call.Latitude)
		longitude := builder.CreateString(call.Longitude)
		zoneBeat := builder.CreateString(call.ZoneBeat)

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
		police = append(police, seattle.PoliceCallEnd(builder))
	}
	seattle.MessageStartPoliceCallsVector(builder, len(policeCalls))
	for i := len(police) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(police[i])
	}
	pCalls := builder.EndVector(len(policeCalls))

	d := builder.CreateString(date)

	seattle.MessageStart(builder)
	seattle.MessageAddDate(builder, d)
	seattle.MessageAddFireCalls(builder, fCalls)
	seattle.MessageAddPoliceCalls(builder, pCalls)
	msg := seattle.MessageEnd(builder)

	builder.Finish(msg)
	buf := builder.FinishedBytes()

	return buf
}
