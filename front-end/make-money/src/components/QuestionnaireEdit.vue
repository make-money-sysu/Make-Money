<template>
  <div class="edit-container">
    Edit Page
    <h2 @click="titleClick" v-if="!titleChange">{{qsItem.title}}</h2>
    <input type="text" name="qsTitle" v-if="titleChange"
    v-model="titleValue"
    @blur="onblur"
    @keyup.enter="onsubmit"
    ref="titleInput">
    <div class="content">
      <div class="questions" v-for="(qs, index) in qsItem.question">
        <div class="qs-left">
          <p class="qs-title">{{qs.num}}&nbsp;{{qs.title}}&nbsp;{{getMsg(qs)}}</p>
          <p v-for="option in qs.options" class="option">
            <label>
              <input type="radio" :name="`${qs.num}-${qs.title}`"
              v-if="qs.type === 'radio'">
              <input type="checkbox" :name="`${qs.num}-${qs.title}`"
              v-if="qs.type === 'checkbox'">{{option}}
            </label>
          </p>
          <textarea v-if="qs.type === 'textarea'"></textarea>
        </div>
        <div class="qs-right">
          <label>
            <input type="checkbox" :value="qs.isNeed" v-model="qs.isNeed">此题是否必填
          </label>
          <p>
            <span v-if="index !== 0" @click="goUp(index)">上移</span>
            <span v-if="index !== qsItem.question.length - 1" @click="goDown(index)">下移</span>
            <span @click="copy(index, qs)">复用</span>
            <span @click="del(index)">删除</span>
          </p>
        </div>
      </div>
      <div class="add">
        <transition name="slide">
          <div class="add-option" v-if="showBtn">
            <button @click="addRadio">单选</button>
            <button @click="addCheckbox">多选</button>
            <button @click="addTextarea">文本框</button>
          </div>
        </transition>
        <div class="add-item" @click="addItemClick">
          <span class="add-icon">+</span>
          <span>添加问题</span>
        </div>
      </div>
    </div>
    <div class="shadow" v-if="showAddQsDialog">
      <div class="add-qs-dialog">
        <header>
          <span>提示</span>
          <span class="close-btn" @click="showAddQsDialog = false">X</span>
        </header>
        <p>{{info}}</p>
        <label>输入题目标题<input type="text" v-model="qsInputTitle"></label>
        <label v-if="showAddOptionInput">输入选项<input type="text" v-model="qsInputOptions"></label>
        <div class="btn-box">
          <button class="yes" @click="validateAddQs">确定</button>
          <button @click="showAddQsDialog = false">取消</button>
        </div>
      </div>
    </div>
    <div class="shadow" v-if="showDialog">
      <div class="dialog">
        <header>
          <span>提示</span>
          <span class="close-btn" @click="cancel">X</span>
        </header>
        <p>{{info}}</p>
        <div class="btn-box">
          <button class="yes" @click="iterator.next()">确定</button>
          <button @click="cancel">取消</button>
        </div>
      </div>
    </div>
    <footer>
      <span>问卷截至日期</span>
      <calendar-input
        :showCalendar="false"
        :limit="limit"
        @getValue="getValue"></calendar-input>
        <div class="btn-box">
          <button class="save" @click="iterator = save(); iterator.next()">保存问卷</button>
          <button class="issue" @click="iterator = issueQs(); iterator.next()">发布问卷</button>
        </div>
    </footer>
  </div>
</template>

<script>
import calendaInput from './calendar-input'

export default {
  name: 'QuestionnaireEdit',
  components: {
    calendaInput
  },
  data () {
    return {
      qsItem: {},
      qslist: [],
      isError: false,
      showBtn: false,
      titleChange: false,
      titleValue: '',
      showAddQsDialog: false,
      showAddOptionInput: true,
      qsInputTitle: '',
      qsInputOptions: [],
      info: '',
      selectTime: '',
      addOptionType: '',
      limit: {},
      showDialog: false,
      iterator: {},
      isGoIndex: false
    }
  },
  beforeRouteEnter(to, from, next) {
    let num = to.params.num
    let theItem = {}
    if (num != 0) {
      let length = 1
      if (num < 0 || num > length) {
        alert('非法路由！')
        next('/')
      }
      else {
        for (let i = 0; i < length; i++) {

        }
      }
      if (theItem.state === 'noissue') {
        next()
      }
      else {
        alert('非法路由！')
        next('/')
      }
    }
    else {
      next()
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.limit = {
        minYear: new Date().getFullYear(),
        minMonth: new Date().getMonth() + 1,
        minDay: new Date().getDate(),
        maxYear: 2030,
        maxMonth: 3,
        maxDay: 20
      }
      if (this.$route.params.num == 0) {
        let item = {}
      }
    }
  }
}
</script>

<style></style>
