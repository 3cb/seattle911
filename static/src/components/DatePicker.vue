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
                            v-model="today"
                            type="date"
                            placeholder="Pick a Day"
                            :picker-options="pickerOptionsDay"
                        >
                        </el-date-picker>
                        <el-button type="primary" plain>Submit</el-button>
                    </li>
                    <li v-show="radio == 2">
                        <el-date-picker
                            v-model="month"
                            type="month"
                            placeholder="Pick a Month"
                            :picker-options="pickerOptionsMonth"
                        >
                        </el-date-picker>
                        <el-button type="primary" plain>Submit</el-button>
                    </li>
                    <li v-show="radio == 3">
                        <el-date-picker
                            v-model="year"
                            type="year"
                            placeholder="Pick a Year"
                            :picker-options="pickerOptionsYear"
                        >
                        </el-date-picker>
                        <el-button type="primary" plain>Submit</el-button>
                    </li>
                    <li><span>**Seattle is in PST</span></li>
                    {{ today }}
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

export default {
  data() {
    return {
      radio: 1,
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
      },
      local: {
        dow: 0, // Sunday is the first day of the week
        hourTip: "Select Hour", // tip of select hour
        minuteTip: "Select Minute", // tip of select minute
        secondTip: "Select Second", // tip of select second
        yearSuffix: "", // suffix of head year
        monthsHead: "January_February_March_April_May_June_July_August_September_October_November_December".split(
          "_"
        ), // months of head
        months: "Jan_Feb_Mar_Apr_May_Jun_Jul_Aug_Sep_Oct_Nov_Dec".split("_"), // months of panel
        weeks: "Su_Mo_Tu_We_Th_Fr_Sa".split("_"), // weeks,
        cancelTip: "cancel",
        submitTip: "confirm"
      },
      //   today: DateTime.local().setZone("America/Los_Angeles"),
      //   day: DateTime.local()
      //     .setZone("America/Los_Angeles")
      //     .toISO(),
      today: null,
      month: null,
      year: null
    };
  },
  methods: {},
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
 
