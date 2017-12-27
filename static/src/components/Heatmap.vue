<template>
    <div id='heatmap'></div>
</template>

<script>
var flatbuffers = require("../../node_modules/flatbuffers").flatbuffers;
var seattle = require("../seattle/schema_generated.js").seattle;
import axios from "axios";
import moment from "moment";

export default {
  data() {
    return {
      map: null
    };
  },
  computed: {
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
        type: "FeatureCollection",
        features: fft
      });
    },
    pfeatures(pft) {
      this.map.getSource("pcalls").setData({
        type: "FeatureCollection",
        features: pft
      });
    }
  },
  mounted() {
    mapboxgl.accessToken =
      "pk.eyJ1IjoibWFyY2NiIiwiYSI6ImNqYTR1enN2dGE0bWEyd3BhcTd6cnBzc3MifQ.Z4zYRzVCXv5zCqqdpgKZ-w";
    this.map = new mapboxgl.Map({
      container: "heatmap",
      style: "mapbox://styles/mapbox/dark-v9",
      center: [-122.335167, 47.608013],
      zoom: 10.75
    });

    this.map.on("load", () => {
      axios({
        url:
          "/api/day/" +
          moment()
            .format()
            .split("T")[0],
        method: "get",
        responseType: "arraybuffer"
      })
        .then(response => {
          let bytes = new Uint8Array(response.data);
          let buf = new flatbuffers.ByteBuffer(bytes);
          let message = seattle.Message.getRootAsMessage(buf);
          this.$store.commit("updateCalls", message);
          this.$store.commit("updateFeatures");
          this.map.addSource("fcalls", {
            type: "geojson",
            data: {
              type: "FeatureCollection",
              features: this.ffeatures
            }
          });
          this.map.addSource("pcalls", {
            type: "geojson",
            data: {
              type: "FeatureCollection",
              features: this.pfeatures
            }
          });
          this.map.addLayer({
            id: "fire-heat",
            type: "heatmap",
            source: "fcalls",
            paint: {
                   "heatmap-weight": {
                      "type": "exponential",
                      "stops": [
                          [0, 0],
                          [6, 1]
                      ]
                  },
                  "heatmap-intensity": {
                      "stops": [
                          [0, 1],
                          [9, 3]
                      ]
                  },
                  "heatmap-color": [
                      "interpolate",
                      ["linear"],
                      ["heatmap-density"],
                      0, "rgba(33,102,172,0)",
                      0.2, "rgb(103,169,207)",
                      0.4, "rgb(209,229,240)",
                      0.6, "rgb(253,219,199)",
                      0.8, "rgb(239,138,98)",
                      1, "rgb(178,24,43)"
                  ],
                   "heatmap-radius": {
                      "stops": [
                          [0, 2],
                          [9, 20]
                      ]
                  },
            }
          });
          this.map.addLayer({
              "id": "police-heat",
              "type":"heatmap",
              "source": "pcalls",
              "paint": {
                  "heatmap-weight": {
                      "type": "exponential",
                      "stops": [
                          [0, 0],
                          [6, 1]
                      ]
                  },
                  "heatmap-intensity": {
                      "stops": [
                          [0, 1],
                          [9, 3]
                      ]
                  },
                "heatmap-color": [
                      "interpolate",
                      ["linear"],
                      ["heatmap-density"],
                      0, "#eff3ff",
                      0.2, "#c6dbef",
                      0.4, "#9ecae1",
                      0.6, "#6baed6",
                      0.8, "#3182bd",
                      1, "#08519c"
                  ],
                   "heatmap-radius": {
                      "stops": [
                          [0, 2],
                          [9, 20]
                      ]
                  }
              }
          })
        })
        .catch(error => {
          console.error(error);
        });
    });
  }
};
</script>

<style>
body {
    margin:0;
    padding:0;
}

#heatmap {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 100%;
}
</style>