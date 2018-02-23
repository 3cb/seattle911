<template>
  <div id="app">
    <streets v-if="showStreets"></streets>
    <heatmap v-else></heatmap>
  </div>
</template>

<script>
import xs from 'xstream'
var flatbuffers = require("../node_modules/flatbuffers").flatbuffers;
var seattle = require("./seattle/schema_generated.js").seattle;
import Streets from './components/Streets.vue'
import Heatmap from './components/Heatmap.vue'

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
          this.$store.commit("updateFeatures",{
            msg: message,
            type: 'today'
          });
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
    showStreets() {
      return this.$store.state.ui.showStreets
    },
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
    Streets,
    Heatmap
  }
}
</script>

<style>

</style>
