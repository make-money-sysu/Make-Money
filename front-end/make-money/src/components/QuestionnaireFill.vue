<template>
  <div class="fill-container">
    <div class="fill" v-if="!isError">
      <router-link to="/QuestionnaireList" tag="span" class="back">&lt; Back</router-link>
      <h2>{{qsItem.title}}</h2>
      <div class="content">
        <div class="content-item" v-for="item in qsItem.question">
          <p class="qs-title">
            {{item.num}}&nbsp;{{item.title}}&nbsp;{{getMsg(item)}}
          </p>
          <p v-for="option in item.options" class="option">
            <label>
              <input type="radio"
                :name="`${item.num}-${item.title}`"
                v-model="requiredItem[item.num]"
                v-if="item.type === 'radio"
                :value="option">
                <input type="checkbox"
                :name="`${item.num}-${item.title}`"
                v-model="requiredItem[item.num]"
                v-if="item.type === 'checkbox"
                :value="option">{{option}}
            </label>
          </p>
          <textarea
            v-if="item.type === 'textarea'"
            v-model="requiredItem[item.num]"></textarea>
        </div>
      </div>

      <transition name="fade">
        <div class="dialog" v-if="showDialog">
          <div class="submit-dialog" v-if="submitError">
            <header>
              <span>Tips</span>
              <span class="close-btn" @click="showDialog=false">X</span>
            </header>
            <p>{{info}}</p>
            <div class="btn-box">
              <button class="yes" @click="showDialog = false">Confirm</button>
              <button @click="showDialog = false">Cancel</button>
            </div>
          </div>
          <div class="submit-dialog" v-else>
            <header>
              <span>Tips</span>
              <span class="close-btn" @click="showDialog = false">X</span>
            </header>
            <p>{{info}}</p>
          </div>
        </div>
      </transition>

      <footer>
        <button class="submit" @click="submit">Submit</button>
      </footer>

    </div>
    <div class="error" v-else>404 Not Found
    </div>
  </div>
</template>

<script>
export default {
  name: 'QuestionnaireFill',
  data() {
    return {
      qsItem: {}, // 当前修改的问卷
      // 问卷列表
      qslist: [{ 'num': 1,
            'title': '第三份问卷',
            'time': '2017-3-27',
            'state': 'noissued',
            'stateTitle': '已发布',
            'checked': false,
            'question': [
              {'num': 'Q1', 'title': '单选题', 'type': 'radio', 'isNeed': true, 'options': ['选项一', '选项二']},
              {'num': 'Q2', 'title': '多选题', 'type': 'checkbox', 'isNeed': true, 'options': ['选项一', '选项二', '选项三', '选项四']},
              {'num': 'Q3', 'title': '文本题', 'type': 'textarea', 'isNeed': true}
            ]
          }],
      isError: false,
      showDialog: false,
      info: '',
      submitError: false,
      requiredItem: {}
    }
  },
  created() {
    this.fetchData()
  },
  mounted() {
    this.getRequiredItem()
  },
  methods: {
    fetchData() {

    },
    getMsg(item) {
      let msg = ''
      if (item.type === 'radio') {
        msg = '(单选题)'
      }
      else if (item.type === 'checkbox') {
        msg = '(多选题)'
      }
      else {
        msg = '(文本题)'
      }

      return item.isNeed ? `${msg} *` : msg
    },
    submit() {

    },
    getRequiredItem() {

    },
    validate() {

    }
  },
  watch: {
    '$route': 'fetchData'
  }
}
</script>

<style lang="scss" scoped>
@import '../style/QuestionnaireFill'
</style>
