import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        ws: null,
        wsConnected: false
    },
    mutations: {
        startWS(state) {
            state.ws = new WebSocket("ws://localhost:3030/ws")
            state.ws.binaryType = 'arraybuffer'
            state.ws.onopen = event => {
                console.log(event)
                state.wsConnected = true
            }
        }
    }
})