<template>
    <div class="sb-picker">
        <div class="columns is-multiline is-centered">
            <div class="column is-10"></div>
            <div class="column is-10">
                <ul class="sb-date-input">
                    <li><span>Select Date:</span></li>
                    <li>
                        <vue-datepicker-local
                            v-model="day"
                            :local="local"
                            placeholder="Pick a Day"
                            :disabled-date="disabledDate"
                            show-buttons
                        >
                        </vue-datepicker-local>
                        <!-- <a class="button is-rounded date-submit" @click="submitDay">Submit</a> -->
                    </li>
                    <li><span>**Seattle is in PST</span></li>
                    {{ day }}
                </ul>
            </div>
            <div class="column is-10">
                <a class="button is-rounded sb-today-btn">Show Today's Calls</a>
            </div>
        </div>
    </div>
</template>

<script>
import VueDatepickerLocal from "vue-datepicker-local";
import { DateTime } from 'luxon'

export default {
  data() {
    return {
      local: {
        dow: 0, // Sunday is the first day of the week
        hourTip: "Select Hour", // tip of select hour
        minuteTip: "Select Minute", // tip of select minute
        secondTip: "Select Second", // tip of select second
        yearSuffix: "", // suffix of head year
        monthsHead: "January_February_March_April_May_June_July_August_September_October_November_December".split("_"), // months of head
        months: "Jan_Feb_Mar_Apr_May_Jun_Jul_Aug_Sep_Oct_Nov_Dec".split("_"), // months of panel
        weeks: "Su_Mo_Tu_We_Th_Fr_Sa".split("_"), // weeks,
        cancelTip: "cancel",
        submitTip: "confirm"
      },
      today: DateTime.local().setZone("America/Los_Angeles"),
      day: DateTime.local().setZone("America/Los_Angeles").toISO()
    };
  },
  methods: {
    //   setDay(date) {
    //       this.day = date
    //   },
      disabledDate(date) {
          var dayISO = date.toISOString().split("T")[0]
          var todayISO = this.today.toISODate()
          diff = todayISO.diff(dayISO, 'days')
          if (diff.days > 0) {
              return true
          }

      }
  },
  components: {
    VueDatepickerLocal
  }
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
 
