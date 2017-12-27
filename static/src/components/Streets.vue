<template>
    <div id='map'></div>
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
      container: "map",
      style: "mapbox://styles/mapbox/streets-v10",
      center: [-122.335167, 47.608013],
      zoom: 10.75
    });
    // Add zoom and rotation controls to the map.
    this.map.addControl(new mapboxgl.NavigationControl());

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
            id: "fire",
            type: "circle",
            source: "fcalls",
            paint: {
              "circle-radius": 6,
              "circle-color": "#B42222"
            }
          });
          this.map.addLayer({
            id: "police",
            type: "circle",
            source: "pcalls",
            paint: {
              "circle-radius": 6,
              "circle-color": "#034cc1"
            }
          });

          var popupFire = new mapboxgl.Popup({
            closeButton: false,
            closeOnClick: false
          });

          this.map.on("mouseenter", "fire", (e) => {
            // Change the cursor style as a UI indicator.
            this.map.getCanvas().style.cursor = "pointer";

            // Populate the popup and set its coordinates
            // based on the feature found.
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

          this.map.on("mouseenter", "police", (e) => {
            // Change the cursor style as a UI indicator.
            this.map.getCanvas().style.cursor = "pointer";

            // Populate the popup and set its coordinates
            // based on the feature found.
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

.mapboxgl-popup {
  max-width: 400px;
  font: 12px/20px "Helvetica Neue", Arial, Helvetica, sans-serif;
}
</style>