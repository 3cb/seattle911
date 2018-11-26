<template>
    <div>
        <toolbar :map="map" :fireLayer="fireLayer"></toolbar>
        <div id='heatmap'></div>
    </div>
</template>

<script>
import Toolbar from "./Toolbar.vue";
var flatbuffers = require("../../node_modules/flatbuffers").flatbuffers;
var seattle = require("../seattle/schema_generated.js").seattle;
import axios from "axios";
import { DateTime } from "luxon";

export default {
  data() {
    return {
      map: null,

      fireLayer: {
        id: "fire",
        type: "heatmap",
        source: "fcalls",
        paint: {
          "heatmap-weight": {
            type: "exponential",
            stops: [[1, 0], [1, 1]]
          },
          "heatmap-intensity": {
            stops: [[1, 1], [13, 3]]
          },
          "heatmap-color": [
            "interpolate",
            ["linear"],
            ["heatmap-density"],
            0,
            "hsla(0, 100%, 80%, 0.2)",
            0.2,
            "hsla(0, 100%, 70%, 0.6)",
            0.4,
            "hsla(0, 100%, 60%, 0.8)",
            0.6,
            "hsla(0, 100%, 50%, 0.9)",
            0.8,
            "hsla(0, 100%, 40%, 0.95)",
            1,
            "hsla(0, 100%, 30%, 1)"
          ],
          "heatmap-radius": {
            base: 1.75,
            stops: [[10, 2.5], [22, 100]]
          }
        }
      }
    };
  },
  computed: {
    ffeatures() {
      if (this.$store.state.ui.showToday === true) {
        return this.$store.state.features.today.fire;
      } else {
        return this.$store.state.features.history.fire;
      }
    }
  },
  watch: {
    ffeatures(fft) {
      this.map.getSource("fcalls").setData({
        type: "FeatureCollection",
        features: fft
      });
    }
  },
  mounted() {
    this.$store.commit("resetFirePolice");

    mapboxgl.accessToken =
      "pk.eyJ1IjoibWFyY2NiIiwiYSI6ImNqYTR1enN2dGE0bWEyd3BhcTd6cnBzc3MifQ.Z4zYRzVCXv5zCqqdpgKZ-w";
    this.map = new mapboxgl.Map({
      container: "heatmap",
      style: "mapbox://styles/mapbox/dark-v9",
      center: [-122.335167, 47.608013],
      zoom: 11
    });
    // Add zoom and rotation controls to the map.
    this.map.addControl(new mapboxgl.NavigationControl());

    this.map.on("load", () => {
      axios(
        this.createDayRequest(
          DateTime.local()
            .setZone("America/Los_Angeles")
            .toISODate()
        )
      )
        .then(response => {
          let bytes = new Uint8Array(response.data);
          let buf = new flatbuffers.ByteBuffer(bytes);
          let message = seattle.Message.getRootAsMessage(buf);
          this.$store.commit("updateFeatures", {
            msg: message,
            type: "today"
          });

          this.map.addSource("fcalls", {
            type: "geojson",
            data: {
              type: "FeatureCollection",
              features: this.ffeatures
            }
          });
          this.map.addLayer(this.fireLayer);
        })
        .catch(error => {
          console.error(error);
        });
    });
  },
  methods: {
    // takes ISODate string parameter (YYYY-MM-DD)
    createDayRequest(date) {
      return {
        url: "/api/day/" + date,
        method: "get",
        responseType: "arraybuffer"
      };
    }
  },
  components: {
    Toolbar
  }
};
</script>

<style>
body {
  margin: 0;
  padding: 0;
}

#heatmap {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 100%;
}
</style>