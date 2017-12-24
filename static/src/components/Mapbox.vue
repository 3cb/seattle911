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
    features() {
      return this.$store.state.features;
    }
  },
  watch: {
      features(ft) {
        this.map.getSource("calls").setData({
          "type": "FeatureCollection",
          "features": ft
        });
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
                this.map.addSource('calls', {
                    "type": "geojson",
                    "data": {
                        "type": "FeatureCollection",
                        "features": this.features
                    }
                })
                this.map.addLayer({
                    "id": "911",
                    "type": "symbol",
                    "source": 'calls',
                    "layout": {
                        "icon-image": "{icon}-15",
                        "text-field": "{title}",
                        "text-font": ["Open Sans Semibold", "Arial Unicode MS Bold"],
                        "text-offset": [0, 0.6],
                        "text-anchor": "top"
                    }
                });
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