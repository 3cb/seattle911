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
		seattle.FireCallStart(builder)
		seattle.FireCallAddAddress(builder, builder.CreateString(fireCalls[i].Address))
		seattle.FireCallAddDatetime(builder, builder.CreateString(fireCalls[i].DateTime))
		seattle.FireCallAddIncidentNumber(builder, builder.CreateString(fireCalls[i].IncidentNumber))
		seattle.FireCallAddLatitude(builder, builder.CreateString(fireCalls[i].Latitude))
		seattle.FireCallAddLongitude(builder, builder.CreateString(fireCalls[i].Longitude))
		seattle.FireCallAddType(builder, builder.CreateString(fireCalls[i].Type))
		fire := seattle.FireCallEnd(builder)
		builder.PrependUOffsetT(fire)
	}
	fCalls := builder.EndVector(len(fireCalls))
	// builder.Reset()

	seattle.MessageStartPoliceCallsVector(builder, len(policeCalls))
	for j := len(policeCalls) - 1; j >= 0; j-- {
		seattle.PoliceCallStart(builder)
		seattle.PoliceCallAddAtSceneTime(builder, builder.CreateString(policeCalls[j].AtSceneTime))
		seattle.PoliceCallAddCadCdwId(builder, builder.CreateString(policeCalls[j].CADCDWID))
		seattle.PoliceCallAddCadEventNumber(builder, builder.CreateString(policeCalls[j].CADEventNumber))
		seattle.PoliceCallAddCensusTract(builder, builder.CreateString(policeCalls[j].CensusTract))
		seattle.PoliceCallAddDistrictSector(builder, builder.CreateString(policeCalls[j].DistrictSector))
		seattle.PoliceCallAddEventClearanceCode(builder, builder.CreateString(policeCalls[j].EventClearanceCode))
		seattle.PoliceCallAddEventClearanceDate(builder, builder.CreateString(policeCalls[j].EventClearanceDate))
		seattle.PoliceCallAddEventClearanceDescription(builder, builder.CreateString(policeCalls[j].EventClearanceDescription))
		seattle.PoliceCallAddEventClearanceGroup(builder, builder.CreateString(policeCalls[j].EventClearanceGroup))
		seattle.PoliceCallAddEventClearanceSubgroup(builder, builder.CreateString(policeCalls[j].EventClearanceSubgroup))
		seattle.PoliceCallAddGeneralOffenseNumber(builder, builder.CreateString(policeCalls[j].GeneralOffenseNumber))
		seattle.PoliceCallAddHundredBlockLocation(builder, builder.CreateString(policeCalls[j].HundredBlockLocation))
		seattle.PoliceCallAddInitialTypeDescription(builder, builder.CreateString(policeCalls[j].InitialTypeDescription))
		seattle.PoliceCallAddInitialTypeGroup(builder, builder.CreateString(policeCalls[j].InitialTypeGroup))
		seattle.PoliceCallAddInitialTypeSubgroup(builder, builder.CreateString(policeCalls[j].InitialTypeSubGroup))
		seattle.PoliceCallAddLatitude(builder, builder.CreateString(policeCalls[j].Latitude))
		seattle.PoliceCallAddLongitude(builder, builder.CreateString(policeCalls[j].Longitude))
		seattle.PoliceCallAddZoneBeat(builder, builder.CreateString(policeCalls[j].ZoneBeat))
		police := seattle.PoliceCallEnd(builder)
		builder.PrependUOffsetT(police)
	}
	pCalls := builder.EndVector(len(policeCalls))
	// builder.Reset()
	date := strings.Split(fireCalls[0].DateTime, "T")[0]
	seattle.MessageStart(builder)
	seattle.MessageAddDate(builder, builder.CreateString(date))
	seattle.MessageAddFireCalls(builder, fCalls)
	seattle.MessageAddFireCalls(builder, pCalls)
	msg := seattle.MessageEnd(builder)

	builder.Finish(msg)
	buf := builder.FinishedBytes()
	fmt.Print(buf)

	return buf
}
