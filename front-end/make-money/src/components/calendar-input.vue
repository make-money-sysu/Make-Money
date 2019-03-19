<template>
  <div class="calendar-input-container blue-theme" v-show="show" :style="containerStyle">
    <label for="calendar-input">
      <input type="text" ref="mainInput" class="calendar-input" readonly="readonly"
      :class="{focus: isfocus}"
      @click="focus"
      v-model="selectValue">
    </label>

    <transition name="fade">
      <div id="calendar" v-show="changeShowCalendar">
        <div id="calendar-header">
          <span class="arrow" @click="subMonth">&lt;</span>
          <span id="date-box">
            {{trueSelectYear}}年{{trueSelectMonth}}月
          </span>
          <span class="arrow" @click="addMonth">&gt;</span>
        </div>
        <div class="week">
          <span v-for="(item, index) in week" :class="{weekend: index === 0 || index === 6}">{{item}}</span>
        </div>
        <div class="days">
          <span v-for="(item, index) in renderData"
            :class="{
              weekend: index % 7 === 0 || index % 7 === 6,
              unselect: unselectArr.includes(index),
              select: index === firstDayInWeek + trueSelectDay - 1
            }"
            @click="changeSelectDay(index)">{{item}}</span>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  name: 'calendar-input',
  props: {
    // 组件通信
    // 是否显示整个组件
    show: {
      type: Boolean,
      default: true
    },
    // 是否显示日历
    showCalendar: {
      type: Boolean,
      default: true
    },
    // 日期可选范围
    limit: {
      type: Object,
      default() {
        return {
          minYear: 1900,
          minMonth: 1,
          minDay: 1,
          maxYear: 2030,
          maxMonth: 3,
          maxDay: 20
        }
      }
    },
    // 组件容器样式
    containerStyle: {
      type: Object
    }
  },
  data() {
    return {
      isfocus: false,
      changeShowCalendar: this.showCalendar,
      week: ["日", "一", "二", "三", "四", "五", "六"],
      date: new Date(),
      selectYear: new Date().getFullYear(),
      selectMonth: new Date().getMonth() + 1,
      selectDay: new Date().getDate()
    }
  },
  methods: {
    focus() {
      this.isfocus = !this.isfocus;
      this.changeShowCalendar = !this.changeShowCalendar
    },
    subMonth() {
      if (this.selectMonth === 1) {
        this.selectYear -= 1;
        this.selectMonth = 12;
      }
      else {
        this.selectMonth -= 1;
      }

      if (this.selectYear < this.limit.minYear) {
        this.selectYear = this.limit.minYear
      }
      if (this.trueSelectYear === this.limit.minYear) {
        if (this.selectMonth <= this.limit.minMonth) {
          this.selectMonth = this.limit.minMonth;
        }
      }
    },
    addMonth() {

    }
  }
}
</script>

<style>

</style>
