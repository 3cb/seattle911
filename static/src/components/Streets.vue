<template>
    <div id='map'></div>
</template>

<script>
import xs from "xstream";
var flatbuffers = require("../../node_modules/flatbuffers").flatbuffers;
var seattle = require("../seattle/schema_generated.js").seattle;
import axios from 'axios'
import moment from 'moment'

export default {
  data() {
    return {
      map: null,
      producer: {
        start: listener => {
          this.$store.commit("startWS");
          this.$store.state.ws.onmessage = event => {
            let bytes = new Uint8Array(event.data);
            let buf = new flatbuffers.ByteBuffer(bytes);
            let message = seattle.Message.getRootAsMessage(buf);
            listener.next(message);
          };
        },
        stop: () => {
          console.log("No longer listening to websocket.");
        }
      },
      updateListener: {
        next: message => {
          this.$store.commit("updateCalls", message);
          this.$store.commit("updateFeatures");
        },
        error: err => {
          console.error("Error on main stream: ", err);
        },
        close: () => {
          console.log("No longer listening to main stream.");
        }
      }
    };
  },
  computed: {
    main$() {
      return xs.createWithMemory(this.producer);
    },
    update$() {
      return xs.from(this.main$);
    },
    calls() {
      return this.$store.state.calls;
    },
    ffeatures() {
      return this.$store.state.features.fire;
    },
    pfeatures() {
        return this.$store.state.features.police;
    }
  },
  watch: {
      ffeatures(fft) {
        this.map.getSource("fcalls").setData({
          "type": "FeatureCollection",
          "features": fft
        });
      },
      pfeatures(pft) {
          this.map.getSource("pcalls").setData({
              "type": "FeatureCollection",
              "features": pft
          })
      }
  },
  mounted() {
    this.main$.addListener(this.updateListener);

    mapboxgl.accessToken =
      "pk.eyJ1IjoibWFyY2NiIiwiYSI6ImNqYTR1enN2dGE0bWEyd3BhcTd6cnBzc3MifQ.Z4zYRzVCXv5zCqqdpgKZ-w";
    this.map = new mapboxgl.Map({
      container: "map",
      style: "mapbox://styles/mapbox/streets-v10",
      center: [-122.335167, 47.608013],
      zoom: 11
    });

    this.map.on('load', () => {
        axios({
            url: '/day/' + moment().format().split('T')[0],
                method: 'get',
                responseType: 'arraybuffer'
            })
            .then(response => {
                let bytes = new Uint8Array(response.data);
                let buf = new flatbuffers.ByteBuffer(bytes);
                let message = seattle.Message.getRootAsMessage(buf);
                this.$store.commit('updateCalls', message)
                this.$store.commit("updateFeatures");
                this.map.addSource('fcalls', {
                    "type": "geojson",
                    "data": {
                        "type": "FeatureCollection",
                        "features": this.ffeatures
                    }
                })
                this.map.addSource('pcalls', {
                    "type": "geojson",
                    "data": {
                        "type": "FeatureCollection",
                        "features": this.pfeatures
                    }
                })
                this.map.addLayer({
                    "id": "fire",
                    "type": "circle",
                    "source": "fcalls",
                    "paint": {
                        "circle-radius": 6,
                        "circle-color": "#B42222"
                    }
                });
                this.map.addLayer({
                    "id": "police",
                    "type": "circle",
                    "source": "pcalls",
                    "paint": {
                        "circle-radius": 6,
                        "circle-color": "#034cc1"
                    }
                    
                })
            })
            .catch(error => {
                console.error(error)
            })
    })
  }
};
</script>

<style>
#map {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 100%;
}
</style>