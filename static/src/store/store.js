import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        ws: null,
        wsConnected: false,

        calls: {
            today: {
                fire: [], // { address, datetime, incident_number, latitude, longitude, type }
                police: [] // { at_scene_time, cad_cdw_id, cad_event_number, census_tract, district_sector, event_clearance_coe, event_clearance_date, event_clearance_description, event_clearance_group, event_clearance_subgroup, general_offense_number, hundred_block_location, initial_type_description, initial_type_group, initial_type_subgroup, latitude, longitude, zone_beat }
            },
            history: {
                fire: [],
                police: []
            }
        }
    },
    mutations: {
        startWS(state) {
            state.ws = new WebSocket("ws://localhost:3030/ws")
            state.ws.binaryType = 'arraybuffer'
            state.ws.onopen = event => {
                console.log(event)
                state.wsConnected = true
            }
        },
        updateCalls(state, msg) {
            state.calls.today.fire = []
            var lenFire = msg.fireCallsLength()
            for (let i = 0; i < lenFire; i++) {
                state.calls.today.fire.push({
                    address: msg.fireCalls(i).address(),
                    datetime: msg.fireCalls(i).datetime(),
                    incidentNumber: msg.fireCalls(i).incidentNumber(),
                    latitude: msg.fireCalls(i).latitude(),
                    longitude: msg.fireCalls(i).longitude(),
                    type: msg.fireCalls(i).type()
                })
            }
            console.log(state.calls.today.fire)

            state.calls.today.police = []
            var lenPolice = msg.policeCallsLength()
            for (let i = 0; i < lenPolice; i++) {
                state.calls.today.police.push({
                    atSceneTime: msg.policeCalls(i).atSceneTime(),
                    cadCdwId: msg.policeCalls(i).cadCdwId(),
                    cadEventNumber: msg.policeCalls(i).cadEventNumber(),
                    censusTract: msg.policeCalls(i).censusTract(),
                    districtSector: msg.policeCalls(i).districtSector(),
                    eventClearanceCode: msg.policeCalls(i).eventClearanceCode(),
                    eventClearanceDate: msg.policeCalls(i).eventClearanceDate(),
                    eventClearanceDescription: msg.policeCalls(i).eventClearanceDescription(),
                    eventClearanceGroup: msg.policeCalls(i).eventClearanceGroup(),
                    eventClearanceSubgroup: msg.policeCalls(i).eventClearanceSubgroup(),
                    generalOffenseNumber: msg.policeCalls(i).generalOffenseNumber(),
                    hundredBlockLocation: msg.policeCalls(i).hundredBlockLocation(),
                    initialTypeDescription: msg.policeCalls(i).initialTypeDescription(),
                    initialTypeGroup: msg.policeCalls(i).initialTypeGroup(),
                    initialTypeSubgroup: msg.policeCalls(i).initialTypeSubgroup(),
                    latitude: msg.policeCalls(i).latitude(),
                    longitude: msg.policeCalls(i).longitude(),
                    zoneBeat: msg.policeCalls(i).zoneBeat()
                })
            }
            console.log(state.calls.today.police)
        }
    }
})