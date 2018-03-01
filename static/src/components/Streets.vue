<template>
    <div>
      <div class="columns is-multiline is-gapless sb">
        <div class="column is-4">
          <a class="button is-rounded sb-style" @click="toggleStyle">Toggle Style</a>
        </div>
        <div class="column is-4">
          <a class="button is-rounded sb-fire" @click="toggleFire">Toggle Fire</a>
        </div>
        <div class="column is-4">
          <a class="button is-rounded sb-police" @click="togglePolice">Toggle Police</a>
        </div>
        <div class="column is-12">
           <a class="button is-rounded sb-date" @click="toggleDatePicker">
             <span v-show="!showDatePicker" class="icon is-small">
               <i class="fas fa-angle-double-right"></i>
             </span>
             <span v-show="showDatePicker" class="icon is-small">
               <i class="fas fa-angle-double-down"></i>
             </span>
             <span>{{ displayDates }}</span>
           </a>
        </div>
        <div class="column is-12">
          <date-picker v-show="showDatePicker"></date-picker>
        </div>
      </div>
      <div id='map'></div>
    </div>
</template>

<script>
import DatePicker from "./DatePicker.vue";
var flatbuffers = require("../../node_modules/flatbuffers").flatbuffers;
var seattle = require("../seattle/schema_generated.js").seattle;
import axios from "axios";
import { DateTime } from "luxon";

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
    showDatePicker() {
      return this.$store.state.ui.showDatePicker;
    },
    displayDates() {
      return this.$store.state.ui.showToday
        ? DateTime.local()
            .setZone("America/Los_Angeles")
            .toISODate()
        : this.$store.state.features.history.date.split("~").join(" to ");
    },
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
    toggleStyle() {
      this.$store.commit("toggleStyle");
    },
    // takes ISODate string parameter (YYYY-MM-DD)
    createDayRequest(date) {
      return {
        url: "/api/day/" + date,
        method: "get",
        responseType: "arraybuffer"
      };
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
    },
    toggleDatePicker() {
      this.$store.commit("toggleDatePicker");
    }
  },
  components: {
    DatePicker
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
  width: 360px;
  z-index: 10;
}

.sb-style {
  width: 100%;
  color: white;
  background-color: #3273dc;
}
.sb-style:hover {
  color: white;
  background-color: #5188e1;
}

.sb-fire {
  width: 100%;
  color: white;
  background-color: #b42222;
}
.sb-fire:hover {
  color: white;
  background-color: #de5454;
}

.sb-police {
  width: 100%;
  color: white;
  background-color: #034cc1;
}
.sb-police:hover {
  color: white;
  background-color: #3682fc;
}

.sb-date {
  width: 100%;
  color: hsl(0, 0%, 96%);
  background-color: hsl(0, 0%, 29%);
}
.sb-date:hover {
  color: hsl(0, 0%, 96%);
  background-color: hsl(0, 0%, 42%);
}
</style>