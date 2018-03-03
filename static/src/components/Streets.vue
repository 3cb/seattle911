<template>
    <div>
      <toolbar :map="map" :fireLayer="fireLayer" :policeLayer="policeLayer"></toolbar>
      <div id='map'></div>
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
        type: "circle",
        source: "fcalls",
        paint: {
          "circle-radius": {
            base: 1.75,
            stops: [[10, 2], [22, 100]]
          },
          "circle-color": "#B42222"
        }
      },
      policeLayer: {
        id: "police",
        type: "circle",
        source: "pcalls",
        paint: {
          "circle-radius": {
            base: 1.75,
            stops: [[10, 2], [22, 100]]
          },
          "circle-color": "#034cc1"
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
    },
    pfeatures() {
      if (this.$store.state.ui.showToday === true) {
        return this.$store.state.features.today.police;
      } else {
        return this.$store.state.features.history.police;
      }
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
    this.$store.commit("resetFirePolice")

    mapboxgl.accessToken =
      "pk.eyJ1IjoibWFyY2NiIiwiYSI6ImNqYTR1enN2dGE0bWEyd3BhcTd6cnBzc3MifQ.Z4zYRzVCXv5zCqqdpgKZ-w";
    this.map = new mapboxgl.Map({
      container: "map",
      style: 'mapbox://styles/mapbox/light-v9',
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
          this.map.addSource("pcalls", {
            type: "geojson",
            data: {
              type: "FeatureCollection",
              features: this.pfeatures
            }
          });
          this.map.addLayer(this.fireLayer);
          this.map.addLayer(this.policeLayer);

          var popupFire = new mapboxgl.Popup({
            closeButton: false,
            closeOnClick: false
          });

          this.map.on("mouseenter", "fire", e => {
            this.map.getCanvas().style.cursor = "pointer";

            popupFire
              .setLngLat(e.features[0].geometry.coordinates)
              .setHTML(e.features[0].properties.description)
              .addTo(this.map);
          });

          this.map.on("mouseleave", "fire", () => {
            this.map.getCanvas().style.cursor = "";
            popupFire.remove();
          });

          var popupPolice = new mapboxgl.Popup({
            closeButton: false,
            closeOnClick: false
          });

          this.map.on("mouseenter", "police", e => {
            this.map.getCanvas().style.cursor = "pointer";

            popupPolice
              .setLngLat(e.features[0].geometry.coordinates)
              .setHTML(e.features[0].properties.description)
              .addTo(this.map);
          });

          this.map.on("mouseleave", "police", () => {
            this.map.getCanvas().style.cursor = "";
            popupPolice.remove();
          });
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

#map {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 100%;
  z-index: 1;
}

.mapboxgl-popup {
  max-width: 400px;
  font: 12px/20px "Helvetica Neue", Arial, Helvetica, sans-serif;
}
</style>