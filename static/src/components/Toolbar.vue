<template>
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
          <date-picker></date-picker>
        </div>
      </div>
</template>

<script>
import DatePicker from "./DatePicker.vue";

export default {
  props: ["map", "fireLayer", "policeLayer"],
  methods: {
    toggleStyle() {
      this.$store.commit("toggleStyle");
    },
    toggleFire() {
      if (this.$store.state.ui.showFire === true) {
        this.map.removeLayer("fire");
      } else {
        this.map.addLayer(this.fireLayer);
      }
      this.$store.commit("toggleFire")
    },
    togglePolice() {
      if (this.$store.state.ui.showPolice === true) {
        this.map.removeLayer("police");
      } else {
        this.map.addLayer(this.policeLayer);
      }
      this.$store.commit("togglePolice")
    },
  },
  components: {
    DatePicker
  }
};
</script>

<style>
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
