<template>
    <div class="columns is-multiline is-gapless sb">
        <div class="column is-12">
          <a class="button is-rounded sb-style" @click="toggleStyle">Toggle Style</a>
        </div>
        <div class="column is-12">
          <date-picker></date-picker>
        </div>
      </div>
</template>

<script>
import DatePicker from "./DatePicker.vue";

export default {
  props: ["map", "fireLayer"],
  methods: {
    toggleStyle() {
      this.$store.commit("toggleStyle");
    }
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
