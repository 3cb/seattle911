<template>
    <div class="sb-picker">
        <div class="columns is-multiline is-centered">
            <div class="column is-10"></div>
            <div class="column is-10">
                <ul class="sb-date-input">
                    <li>
                        <el-radio-group v-model="radio">
                            <el-radio :label="1">Day</el-radio>
                            <el-radio :label="2">Month</el-radio>
                            <el-radio :label="3">Year</el-radio>
                        </el-radio-group>
                    </li>
                    <li v-show="radio == 1">
                        <el-date-picker
                            v-model="day"
                            type="date"
                            placeholder="Pick a Day"
                            :picker-options="pickerOptionsDay"
                        >
                        </el-date-picker>
                        <el-button
                          type="primary"
                          @click="submitDay"
                          plain
                          :disabled="!day">
                            Submit
                        </el-button>
                    </li>
                    <li v-show="radio == 2">
                        <el-date-picker
                            v-model="month"
                            type="month"
                            placeholder="Pick a Month"
                            :picker-options="pickerOptionsMonth"
                        >
                        </el-date-picker>
                        <el-button
                          type="primary"
                          @click="submitMonth"
                          plain
                          :disabled="!month">
                            Submit
                        </el-button>
                    </li>
                    <li v-show="radio == 3">
                        <el-date-picker
                            v-model="year"
                            type="year"
                            placeholder="Pick a Year"
                            :picker-options="pickerOptionsYear"
                        >
                        </el-date-picker>
                        <el-button
                          type="primary"
                          @click="submitYear"
                          plain
                          :disabled="!year">
                            Submit
                        </el-button>
                    </li>
                    <li><span>**Seattle is in PST</span></li>
                    <li>{{ month }}</li>
                </ul>
            </div>
            <div class="column is-10">
                <a class="button is-rounded sb-today-btn">Show Today's Calls</a>
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
      radio: 1,
      day: null,
      month: null,
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
    }
  },
  methods: {
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
        url: "/api/day/" + this.month.toISOString().split("T")[0]
      })
    },
    submitYear() {}
  },
  components: {}
};
</script>

<style>
.sb-picker {
  width: 100%;
  color: hsl(0, 0%, 21%);
  background-color: hsl(0, 0%, 96%);
}

.sb-date-input {
  width: 100%;
}

.date-submit {
  width: 50%;
}

.sb-today-btn {
  width: 100%;
}
</style>
 
