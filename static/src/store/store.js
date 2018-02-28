import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    ws: null,
    wsConnected: false,
    wsProtocol: location.hostname === "localhost" ? "ws://" : "wss://",

    ui: {
      showToday: true,
      showStreets: true,
      showDatePicker: false,
      pickerText: 'Show Date Picker'
    },

    features: {
      today: {
        fire: [],
        police: []
      },
      history: {
        type: '', // 'single' or 'range'
        fire: [],
        police: []
      }
    }
  },
  
  mutations: {
    startWS(state) {
      state.ws = new WebSocket(state.wsProtocol + location.host + "/ws")
      state.ws.binaryType = 'arraybuffer'
      state.ws.onopen = event => {
        state.wsConnected = true
      }
    },

    toggleStyle(state) {
      state.ui.showStreets = !state.ui.showStreets
    },

    toggleDatePicker(state) {
      if (state.ui.showDatePicker === false) {
        state.ui.showDatePicker = true
        state.ui.pickerText = 'Hide Date Picker'
      } else {
        state.ui.showDatePicker = false
        state.ui.pickerText = 'Show Date Picker'
      }
    },

    showToday(state) {
      state.ui.showToday = true
    },
    showHistory(state) {
      state.ui.showToday = false
    },

    updateFeatures(state, {
      msg,
      type
    }) {
      var fire = []
      for (let i = 0; i < msg.fireCallsLength(); i++) {
        fire.push({
          "type": "Feature",
          "properties": {
            "description": "<strong>" + msg.fireCalls(i).type().toUpperCase() + "</strong><p>Time: " + msg.fireCalls(i).datetime().split("T")[1].split(".")[0] + "</p><p>Date: " + msg.fireCalls(i).datetime().split("T")[0] + "</p><p>Location: " + msg.fireCalls(i).address() + "</p><p>Incident #: " + msg.fireCalls(i).incidentNumber() + "</p>"
          },
          "geometry": {
            "type": "Point",
            "coordinates": [parseFloat(msg.fireCalls(i).longitude()), parseFloat(msg.fireCalls(i).latitude())]
          }
        })
      }

      var police = []
      for (let i = 0; i < msg.policeCallsLength(); i++) {
        police.push({
          "type": "Feature",
          "properties": {
            "description": "<strong>" + msg.policeCalls(i).initialTypeDescription().toUpperCase() + "</strong><p>At Scene Time: " + msg.policeCalls(i).atSceneTime().split("T")[1].split(".")[0] + "</p><p>Date: " + msg.policeCalls(i).atSceneTime().split("T")[0] + "</p><p>Location: " + msg.policeCalls(i).hundredBlockLocation() + "</p><p>Event #: " + msg.policeCalls(i).cadEventNumber() + "</p>"
          },
          "geometry": {
            "type": "Point",
            "coordinates": [parseFloat(msg.policeCalls(i).longitude()), parseFloat(msg.policeCalls(i).latitude())]
          }
        })
      }
      console.log(fire)
      console.log(police)

      if (type === 'today') {
        state.features.today.fire = fire
        state.features.today.police = police
      } else if (type === 'history') {
        state.features.history.fire = fire
        state.features.history.police = police
      }
    }
  }
})
