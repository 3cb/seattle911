<template>
<div class="sb-picker">
    <div class="columns is-multiline is-centered">
      <div class ="column is-10"></div>      
      <div class ="column is-8">
        <div class="tabs is-centered">
          <ul>
            <li :class="tabs.day"><a @click="setTab(1)">Day</a></li>
            <li :class="tabs.month"><a @click="setTab(2)">Month</a></li>
            <li :class="tabs.year"><a @click="setTab(3)">Year</a></li>
          </ul>
        </div>

        <div v-show="tabs.is == 1" class="field">
          <div class="control has-text-centered">
            <el-date-picker
                v-model="day"
                type="date"
                placeholder="Pick a Day"
                :picker-options="pickerOptionsDay"
            >
            </el-date-picker>
          </div>
        </div>
        <div v-show="tabs.is == 1" class="field">
          <div class="control has-text-centered">
            <button
              class="button is-outlined is-link"
              @click="submitDay"
              :disabled="!day">
                Submit
            </button>
          </div>
        </div>
        <div v-show="tabs.is == 2" class="field">
          <div class="control has-text-centered">
            <el-date-picker
                v-model="month"
                type="month"
                placeholder="Pick a Month"
                :picker-options="pickerOptionsMonth"
            >
            </el-date-picker>
          </div>
        </div>
        <div v-show="tabs.is == 2" class="field">
          <div class="control has-text-centered">
            <button
              class="button is-outlined is-link"
              @click="submitMonth"
              :disabled="!month">
                Submit
            </button>
          </div>
        </div>
        <div v-show="tabs.is == 3" class="field">
          <div class="control has-text-centered">
            <el-date-picker
                v-model="year"
                type="year"
                placeholder="Pick a Year"
                :picker-options="pickerOptionsYear"
            >
            </el-date-picker>
          </div>
        </div>
        <div v-show="tabs.is == 3" class="field">
          <div class="control has-text-centered">
            <button
              class="button is-outlined is-link"
              @click="submitYear"
              :disabled="!year">
                Submit
            </button>
          </div>
        </div>
      </div>
      <div class ="column is-10">
        <button
          @click="toggleDisplay"
          class="button is-link sb-today-btn"
          :disabled="!historyDate">
            {{ displayButtonLabel }}
        </button>
      </div>
    </div>
</div>
</template>

<script>
import { DateTime } from "luxon";
var flatbuffers = require("../../node_modules/flatbuffers").flatbuffers;
var seattle = require("../seattle/schema_generated.js").seattle;
import axios from "axios";

export default {
  data() {
    return {
      tabs: {
        is: 1,
        day: "is-active",
        month: "",
        year: ""
      },
      day: null,
      // first day of selected month: "YYYY-MM-DD"
      month: null,
      // first day of selected year: "YYYY-MM-DD"
      year: null,
      pickerOptionsDay: {
        disabledDate(time) {
          let date = DateTime.fromISO(
            DateTime.fromISO(time.toISOString())
              .setZone("America/Los_Angeles")
              .toISODate()
          );
          let now = DateTime.fromISO(
            DateTime.local()
              .setZone("America/Los_Angeles")
              .toISODate()
          );
          let endDiff = now.diff(date, "days");
          if (endDiff.days > 0) {
            return false;
          } else {
            return true;
          }
        }
      },
      pickerOptionsMonth: {
        disabledDate(time) {
          let date = DateTime.fromISO(
            DateTime.fromISO(time.toISOString())
              .setZone("America/Los_Angeles")
              .toISODate()
          );
          let now = DateTime.fromISO(
            DateTime.local()
              .setZone("America/Los_Angeles")
              .toISODate()
          );
          let diff = now.diff(date, "months");
          if (diff.months > 0) {
            return false;
          } else {
            return true;
          }
        }
      },
      pickerOptionsYear: {
        disabledDate(time) {
          let date = DateTime.fromISO(
            DateTime.fromISO(time.toISOString())
              .setZone("America/Los_Angeles")
              .toISODate()
          );
          let now = DateTime.fromISO(
            DateTime.local()
              .setZone("America/Los_Angeles")
              .toISODate()
          );
          let diff = now.diff(date, "years");
          if (diff.years > 0) {
            return false;
          } else {
            return true;
          }
        }
      }
    };
  },
  computed: {
    showToday() {
      return this.$store.state.ui.showToday;
    },
    displayButtonLabel() {
      return this.showToday ? "Show Historical Calls" : "Show Today's Calls";
    },
    historyDate() {
      return this.$store.state.features.history.date;
    }
  },
  methods: {
    setTab(number) {
      switch (number) {
        case 1:
          this.tabs.is = 1;
          this.tabs.day = "is-active";
          this.tabs.month = "";
          this.tabs.year = "";
          break;
        case 2:
          this.tabs.is = 2;
          this.tabs.day = "";
          this.tabs.month = "is-active";
          this.tabs.year = "";
          break;
        case 3:
          this.tabs.is = 3;
          this.tabs.day = "";
          this.tabs.month = "";
          this.tabs.year = "is-active";
          break;
      }
    },
    submitDay() {
      axios({
        url: "/api/day/" + this.day.toISOString().split("T")[0],
        method: "get",
        responseType: "arraybuffer"
      })
        .then(response => {
          let bytes = new Uint8Array(response.data);
          let buf = new flatbuffers.ByteBuffer(bytes);
          let message = seattle.Message.getRootAsMessage(buf);
          this.$store.commit("updateFeatures", {
            msg: message,
            type: "history"
          });
          this.$store.commit("showHistory");
        })
        .catch(err => {
          console.error("error getting historical 911 call data", err);
        });
    },
    submitMonth() {
      axios({
        url: "/api/month/" + this.month.toISOString().split("T")[0],
        method: "get",
        responseType: "arraybuffer"
      })
        .then(response => {
          let bytes = new Uint8Array(response.data);
          let buf = new flatbuffers.ByteBuffer(bytes);
          let message = seattle.Message.getRootAsMessage(buf);
          this.$store.commit("updateFeatures", {
            msg: message,
            type: "history"
          });
          this.$store.commit("showHistory");
        })
        .catch(err => {
          console.error("error getting historical 911 call data", err);
        });
    },
    submitYear() {
      axios({
        url: "/api/year/" + this.year.toISOString().split("T")[0],
        method: "get",
        responseType: "arraybuffer"
      })
        .then(response => {
          let bytes = new Uint8Array(response.data);
          let buf = new flatbuffers.ByteBuffer(bytes);
          let message = seattle.Message.getRootAsMessage(buf);
          this.$store.commit("updateFeatures", {
            msg: message,
            type: "history"
          });
          this.$store.commit("showHistory");
        })
        .catch(err => {
          console.error("error getting historical 911 call data", err);
        });
    },
    toggleDisplay() {
      this.$store.state.ui.showToday
        ? this.$store.commit("showHistory")
        : this.$store.commit("showToday");
    }
  }
};
</script>

<style>
.sb-picker {
  width: 100%;
  color: hsl(0, 0%, 21%);
  background-color: hsl(0, 0%, 96%);
}

.sb-today-btn {
  width: 100%;
}
</style>
 
