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
                v-if="item.type === 'radio'"
                :value="option">
                <input type="checkbox"
                :name="`${item.num}-${item.title}`"
                v-model="requiredItem[item.num]"
                v-if="item.type === 'checkbox'"
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
import axios from 'axios';
import global_ from './Global'

export default {
  name: 'QuestionnaireFill',
  data() {
    return {
      currentId: "",
      qsItem: {}, // 当前修改的问卷
      // 问卷列表
      /*
      qslist: [{ 'num': 3,
            'title': '第三份问卷',
            'time': '2017-3-27',
            'state': 'issued',
            'stateTitle': '已发布',
            'checked': false,
            'question': [
              {'num': 'Q1', 'title': '单选题', 'type': 'radio', 'isNeed': true, 'options': ['选项一', '选项二']},
              {'num': 'Q2', 'title': '多选题', 'type': 'checkbox', 'isNeed': true, 'options': ['选项一', '选项二', '选项三', '选项四']},
              {'num': 'Q3', 'title': '文本题', 'type': 'textarea', 'isNeed': true}
            ]
          }],
          */
      qslist: [],
      isError: false, // 是否出错
      showDialog: false, // 是否显示对话框
      info: '', // 对话框内容
      submitError: false, // 提交是否出错
      requiredItem: {}
    }
  },
  // 页面生成时自动抓取数据
  created() {
    // alert("created")
    this.fetchData()
  },
  mounted() {
  },
  methods: {
    fetchData() {
      // alert('fetchData')
      // 获取数据库内容
      axios.defaults.withCredentials = true;

      // 获取用户信息
      const userGetUrl = global_.url + 'user'
      console.log(userGetUrl)

      axios.get(userGetUrl)
        .then(response => {
          this.currentId = response.data.data.id

          // 获取问卷列表
          const qsGetUrl = global_.url + 'survey'
          var temp;

          axios.get(qsGetUrl)
            .then(response => {
              temp = response.data.data
              // Handle list
              temp.forEach(item => {
                /*
                console.log(item.id) // 获取问卷id
                console.log(item.publisher_id) // 获取问卷发起者
                console.log(item.checked) // 获取checked属性
                console.log(item.content) // 获取questions
                console.log(item.create_time.substr(0, 10)) // 获取问卷发布时间
                console.log(item.state) // 获取问卷状态
                console.log(item.title) // 获取问卷标题
                */
                let questionnaire = {
                  "num": item.id,
                  "title": item.title,
                  "stateTitle": item.state == 0 ? '未发布' : '已发布',
                  "time": item.create_time.substr(0, 10),
                  "state": item.state == 0 ? false: true, //"state": true
                  // "state": true,
                  "checked": item.checked == 0 ? false: true,
                  "question": JSON.parse(item.content)
                  }
                this.qslist.push(questionnaire)
              })

              // 找到要查看的问卷
              let i = 0;
              for (let length = this.qslist.length; i < length; i++) {
                // alert(this.qslist[i].num)
                if (this.qslist[i].num == this.$route.params.num) {
                  this.qsItem = this.qslist[i]
                  // alert("Match")
                  break;
                }
              }
              if (this.qsItem.state == false) {
                alert("当前问卷未发布，不能填写！")
                this.$router.push({path: '/QuestionnaireList'})
              }
              this.getRequiredItem()
            })
            .catch(error => {
              alert('Get Quesionnaire Error!')
              console.log(error)
              this.$router.push({path: '/QuestionnaireList'})
            })
        })
        .catch(error => {
          alert("Login Expire!")
          this.$router.push({path: '/QuestionnaireList'})
        })
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
      // alert(this.qsItem.state)
      if (this.qsItem.state === true) {
        // 校验
        let result = this.validate()
        if (result) {
          // 填写问卷
          const dosurveyUrl = global_.url + '/do_survey'

          axios.post(dosurveyUrl, JSON.stringify({
            survey_id: this.qsItem.num,
            recipient_id: this.currentId,
            content: "nothing"
          }))
            .then(response => {
              // console.log(response)
            })
            .catch(error => {
              alert("Submit Error!")
              this.$router.push({path: '/QuestionnaireList'})
            })

          this.showDialog = true;
          this.submitError = false;
          this.info = 'Submit Successfully!'
          console.log('Submit Successfully')
          setTimeout(() => {
            this.showDialog = false
          }, 500)
          setTimeout(() => {
            this.$router.push({path: '/QuestionnaireList'})
          }, 1500)
        }
        else {
          this.showDialog = true;
          this.submitError = true;
          this.info = 'Fail to submit! Format invalid!'
        }
      }
      else {
        this.showDialog = true;
        this.submitError = true;
        this.info = 'Fail to submit! Only Inissue Questionnaire Can be Submitted!'
      }
    },
    getRequiredItem() {
      console.log("in getRequiredItem")
      console.log(this.qsItem)
      this.qsItem.question.forEach(item => {
        if (item.isNeed) {
          if (item.type === 'checkbox') {
            this.requiredItem[item.num] = []
          }
          else {
            this.requiredItem[item.num] = ''
          }
        }
      })
    },
    validate() {
      for (let i in this.requiredItem) {
        if (this.requiredItem[i].length === 0)
          return false
      }
      return true
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
