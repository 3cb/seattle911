package main

import (
	"github.com/3cb/seattle911/seattle"
	flatbuffers "github.com/google/flatbuffers/go"
)

func serialize(fireCalls []FireCall, date string) []byte {
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

	d := builder.CreateString(date)

	seattle.MessageStart(builder)
	seattle.MessageAddDate(builder, d)
	seattle.MessageAddFireCalls(builder, fCalls)
	msg := seattle.MessageEnd(builder)

	builder.Finish(msg)
	buf := builder.FinishedBytes()

	return buf
}
