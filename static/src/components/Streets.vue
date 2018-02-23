<template>
    <div>
      <div class="sb">
        <a class="button is-rounded sb-style" @click="toggleStyle">Toggle Style</a>
        <a class="button is-rounded sb-fire" @click="toggleFire">Toggle Fire</a>
        <a class="button is-rounded sb-police" @click="togglePolice">Toggle Police</a>
      </div>
      <div id='map'></div>
    </div>
</template>

<script>
var flatbuffers = require("../../node_modules/flatbuffers").flatbuffers;
var seattle = require("../seattle/schema_generated.js").seattle;
import axios from "axios";
// import moment from "moment";
import { DateTime } from 'luxon'

export default {
  data() {
    return {
      map: null,
      showPolice: true,
      showFire: true,

      fireLayer: {
        id: "fire",
        type: "circle",
        source: "fcalls",
        paint: {
          "circle-radius": 6,
          "circle-color": "#B42222"
        }
      },
      policeLayer: {
        id: "police",
        type: "circle",
        source: "pcalls",
        paint: {
          "circle-radius": 6,
          "circle-color": "#034cc1"
        }
      }
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
      axios(this.createRequest())
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
    toggleStyle() {
      this.$store.commit("toggleStyle");
    },
    createRequest() {
      return {
        url: "/api/day/" + DateTime.local().setZone("America/Los_Angeles").toISODate(),
        method: "get",
        responseType: "arraybuffer"
      }
    },
    toggleFire() {
      if (this.showFire === true) {
        this.map.removeLayer("fire");
      } else {
        this.map.addLayer(this.fireLayer);
      }
      this.showFire = !this.showFire;
    },
    togglePolice() {
      if (this.showPolice === true) {
        this.map.removeLayer("police");
      } else {
        this.map.addLayer(this.policeLayer);
      }
      this.showPolice = !this.showPolice;
    }
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

.sb {
  position: absolute;
  top: 0;
  bottom: 0;
  height: 50px;
  width: 400px;
  z-index: 10;
}

.sb-style {
  color: white;
  background-color: #3273dc;
}
.sb-style:hover {
  color: white;
  background-color: #5188e1;
}

.sb-fire {
  color: white;
  background-color: #b42222;
}
.sb-fire:hover {
  color: white;
  background-color: #de5454;
}

.sb-police {
  color: white;
  background-color: #034cc1;
}
.sb-police:hover {
  color: white;
  background-color: #3682fc;
}
</style>