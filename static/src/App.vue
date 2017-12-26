<template>
  <div id="app">
    <!-- <sidebar></sidebar> -->
    <streets></streets>
  </div>
</template>

<script>
import xs from 'xstream'
var flatbuffers = require("../../node_modules/flatbuffers").flatbuffers;
var seattle = require("../seattle/schema_generated.js").seattle;

import Side from './components/Sidebar.vue'
import Streets from './components/Mapbox.vue'

export default {
  name: 'app',
  data() {
    return {
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
    }
  },
  computed: {
    main$() {
      return xs.createWithMemory(this.producer);
    },
    update$() {
      return xs.from(this.main$);
    }
  },
  mounted() {
      this.main$.addListener(this.updateListener);
  },
  components: {
    Sidebar,
    Streets
  }
}
</script>

<style>

</style>
