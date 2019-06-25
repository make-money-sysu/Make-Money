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
import calendarInput from './calendar-input'
import global_ from './Global'
import axios from 'axios'
export default {
  name: 'QuestionnaireEdit',
  components: {
    calendarInput
  },
  data () {
    return {
      currentId: "", // 当前用户
      qsItem: {}, // 当前修改的问卷
      // 问卷列表
      /*
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
          */
      qslist: [],
      isError: false,
      showBtn: false,
      titleChange: false, // 标题是否修改
      titleValue: '', // 标题值
      showAddQsDialog: false, // 是否显示添加问卷对话框
      showAddOptionInput: true, // 是否显示添加选项对话框
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
    // alert('here is beforeRouteEnter')
    const qsGetUrl = global_.url + 'survey'
    var temp;
    // 获取问卷列表
    axios.get(qsGetUrl)
      .then(response => {
        temp = response.data.data // 问卷list

        let num = to.params.num // 从URL中获取要编辑的问卷序号
        let theItem = {}
        // 编辑问卷逻辑
        if (num != 0) {
          let length = temp.length // 获取总问卷列表长度
            alert(num)
              /*
              if (num < 0 || num > length) {
                // 序号越界
                alert('非法路由！')
                next('/')
              }
              else {
                // 获取要编辑的问卷到theItem
                for (let i = 0; i < length; i++) {
                    if (temp[i].id == num) {
                      theItem = temp[i]
                      break
                    }
                }
              }
              */
            // 获取要编辑的问卷到theItem
              for (let i = 0; i < length; i++) {
                if (temp[i].id == num) {
                  let questionnaire = {
                      "num": temp[i].id,
                      "publisher_id": temp[i].publisher_id,
                      "title": temp[i].title,
                      "stateTitle": temp[i].state == 0 ? '未发布' : '已发布',
                      "time": temp[i].create_time.substr(0, 10),
                      "state": temp[i].state == 0 ? false: true,
                      "checked": temp[i].checked == 0 ? false: true,
                      "question": JSON.parse(temp[i].content)
                      }

                    theItem = questionnaire
                    alert("Get!")
                    break
                  }
                }
              console.log(theItem)
              if (theItem.state === false) {
                // 只能编辑未发布的问卷
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
        })
        .catch(error => {
          console.log(error)
        })
  },
  created() {
    alert('here is created')
    this.fetchData()
  },
  methods: {
    fetchData() {
      // alert('in fetchData')
      const userGetUrl = global_.url + 'user'
      console.log(userGetUrl)
      axios.get(userGetUrl)
        .then(response => {
          console.log(response.data.data.id)
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
                  "publisher_id": item.publisher_id,
                  "title": item.title,
                  "stateTitle": item.state == 0 ? '未发布' : '已发布',
                  "time": item.create_time.substr(0, 10),
                  "state": item.state == 0 ? false: true,
                  "checked": item.checked == 0 ? false: true,
                  "question": JSON.parse(item.content)
                }
                this.qslist.push(questionnaire)
              })

              this.limit = {
                minYear: new Date().getFullYear(),
                minMonth: new Date().getMonth() + 1,
                minDay: new Date().getDate(),
                maxYear: 2030,
                maxMonth: 3,
                maxDay: 20
              }
              // 新建问卷逻辑
              if (this.$route.params.num == 0) {
                // Initialize Questionnaire content
                let item = {}
                item.num = this.qslist.length + 1
                item.title = 'Default Title'
                item.time = ''
                item.state = false
                item.question = []
                item.stateTitle = '未发布'
                item.checked = false
                this.qsItem = item
                this.qslist.push(this.qsItem)
              }
              // 编辑问卷逻辑
              else {
                let i = 0
                for (let length = this.qslist.length; i < length; i++) {
                  // alert('listNum ' + this.qslist[i].num)
                  // alert('routeNum ' + this.$route.params.num)
                  if (this.qslist[i].num == this.$route.params.num) {
                    console.log(temp[i].content)
                    let questionnaire = {
                      "num": temp[i].id,
                      "publisher_id": temp[i].publisher_id,
                      "title": temp[i].title,
                      "stateTitle": temp[i].state == 0 ? '未发布' : '已发布',
                      "time": temp[i].create_time.substr(0, 10),
                      "state": temp[i].state == 0 ? false: true,
                      "checked": temp[i].checked == 0 ? false: true,
                      "question": JSON.parse(temp[i].content)
                    }
                    this.qsItem = questionnaire
                    console.log("Match!")
                    // alert("Match")
                    break
                  }
                }

                // console.log(this.currentId)
                // console.log(this.qsItem.publisher_id)
                if (this.currentId != this.qsItem.publisher_id) {
                  alert("不是当前用户！")
                }
                if (i == this.qslist.length) {
                  this.isError = true
                }
              }
            console.log(this.qsItem)
            })
            .catch(error => {
              alert("Cannot Find the Questionnaire!")
              //console.log(error)
            })
        })
        .catch(error => {
          alert("Login Expire!")
          // console.log(error)
        })
    },
    // 获取问题类型
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
    // 标题失去焦点时，修改完成
    onblur() {
      // alert('in onblur')
      this.titleValue = this.titleValue.trim()
      this.qsItem.title = this.titleValue === '' ? this.qsItem.title : this.titleValue
      this.titleChange = false
    },
    // 用户enter时，修改完成
    onsubmit() {
      this.titleValue = this.titleValue.trim()
      this.qsItem.title = this.titleValue === '' ? this.qsItem.title: this.titleValue
    },
    // 修改标题内容
    titleClick() {
      this.titleChange = !this.titleChange
      setTimeout(() => {
        this.$refs.titleInput.focus()
      }, 150)
    },
    // 交换两个问题
    swapItems(oldIndex, newIndex) {
      let [tmpVal] = this.qsItem.question.splice(newIndex, 1, this.qsItem.question[oldIndex])
      this.qsItem.question.splice(oldIndex, 1, tmpVal)
    },
    goUp(index) {
      this.swapItems(index, index - 1)
    },
    goDown(index) {
      this.swapItems(index, index + 1)
    },
    copy (index, qs) {
      if (this.questionLength === 10)
        return alert('Questionnaire is full!')
      // shallow copy
      qs = Object.assign({}, qs)
      // insert to the new position
      this.qsItem.question.splice(index, 0, qs)
    },
    del(index) {
      this.qsItem.question.splice(index, 1)
    },
    addItemClick() {
      // 若选择类型框没有弹出，则弹出
      if (this.showBtn === false) {
        // check the question list size
        this.questionLength === 10 ? alert('Questionnaire is full!') : this.showBtn = !this.showBtn
      }
      // 否则，收起选择类型框
      else {
        this.showBtn = !this.showBtn
      }
    },
    showAddDialog(msg, showOption) {
      this.showAddQsDialog = true
      // 该题目是否有选项
      this.showAddOptionInput = showOption
      // 提示信息
      this.info = msg
      // 清空输入框
      this.qsInputTitle = ''
      this.qsInputOptions = ''
    },
    // Add Radio Question
    addRadio() {
      if (this.questionLength === 10)
        return alert('Questionnaire is full!')
      this.showAddDialog('分别在下面的输入框中输入问题的名称以及选项，选项用半角逗号“,”分隔开', true)
      this.addOptionType = 'radio'
    },
    // Add Checkbox Question
    addCheckbox() {
      if (this.questionLength === 10)
        return alert('Questionnaire is full!')
      this.showAddDialog('分别在下面的输入框中输入问题的名称以及选项，选项用半角逗号“,”分隔开', true)
      this.addOptionType = 'checkbox'
    },
    // Add Textarea Question
    addTextarea() {
      if (this.questionLength === 10)
        return alert('Questionnaire is full!')
      this.showAddDialog('请输入问题名称', false)
      this.addOptionType = 'textarea'
    },
    // 检查每个题目的合法性
    validateAddQs() {
      // Check title
      let qsTitle = this.qsInputTitle.trim()
      if (qsTitle === '') {
        alert('Title Cannot be empty!')
        return
      }
      // Check option if existed
      if (this.showAddOptionInput) {
        let qsOptions = this.qsInputOptions.trim()
        if (qsOptions === '') {
          alert('Option Cannot be empty!')
          return
        }
        qsOptions = qsOptions.split(',')
        for (let i =0, length = qsOptions.length; i < length; i++) {
          if (qsOptions[i].trim() === '') {
            alert('Some Option is empty!')
            return
          }
        }
        this.qsItem.question.push({
          'num': this.qsItem.question.length,
          'title': qsTitle,
          'type': this.addOptionType,
          'isNeed': true,
          'options': qsOptions
        })
        this.showAddQsDialog = false
      }
      // Textarea Question
      else {
        this.qsItem.question.push({
          'num': this.qsItem.question.length - 1,
          'title': qsTitle,
          'type': 'textarea',
          'isNeed': true
        })
        this.showAddQsDialog = false
      }
    },
    // 获取截止时间
    getValue(selectTime) {
      this.selectTime = selectTime
      this.qsItem.time = this.selectTime
    },
    // 保存问卷
    *save() {
      this.showDialog = true
      this.info = 'Confirm Save?'
      yield
      if (this.qsItem.question.length === 0) {
        this.showDialog = false
        alert('The Questionnaire is empty!')
      }
      else {
        // Save in the database
        const qsPostUrl = global_.url + 'survey'
        // console.log(qsItem)
        axios.post(qsPostUrl, JSON.stringify({
          title: qsItem.title,
          content: JSON.stringify(qsItem.question)
        }))
          .then(response => {
            // console.log(response.data)
          })
          .catch(error => {
            alert('Post QS Fail!')
            console.log(error)
          })
        this.info = 'Issue it Now?'
        this.isGoIndex = true
      }

      yield
      this.qsItem.state = 'inissue'
      this.qsItem.stateTitle = '发布中'
      // Save in the database
      this.showDialog = false
      this.$route.push({path: '/QuestionnaireList'})
    },
    *issueQs() {
      this.showDialog = true
      this.info = '确认发布？'
      yield
      if (this.qsItem.question.length === 0) {
        this.showDialog = false
        alert('问卷为空，无法保存！')
      }
      else {
        this.qsItem.state = 'inissue'
        this.qsItem.stateTitle = '发布中'
        // save in DB
        // 新建问卷，使用post
        if (this.$route.params.num == 0) {
          console.log('use post')
          const qsPostUrl = global_.url + 'survey'
          // console.log(this.qsItem)
          axios.post(qsPostUrl, JSON.stringify({
            title: this.qsItem.title,
            content: JSON.stringify(this.qsItem.question)
          }))
            .then(response => {
              // console.log(response.data)
            })
            .catch(error => {
              alert('Post QS Fail!')
              console.log(error)
            })
        }
        // 旧问卷编辑，使用PUT
        else {
          console.log('use put')
          console.log(this.qsItem.title)
          const qsPutUrl = global_.url + 'survey/' + this.qsItem.num
          axios.put(qsPutUrl, JSON.stringify({
            title: this.qsItem.title,
            content: JSON.stringify(this.qsItem.question)
          }))
            .then(response => {
              // console.log(response.data)
            })
            .catch(error => {
              alert('Put QS Fail!')
              console.log(error)
            })
        }

        this.showDialog = false
        this.$router.push({path: '/QuestionnaireList'})
      }
    },
    // 取消对话框
    cancel() {
      this.showDialog = false;
      // 若已经保存但不发布，直接跳回
      if (this.isGoIndex === true) {
        this.$router.push({path: 'QuestionnaireList'})
      }
    }
  },
  computed: {
    questionLength() {
      return this.qsItem.question.length
    }
  },
  watch: {
    // 路由发生变化时重新抓取数据
    '$route': 'fetchData',
    // 重新安排问题列表的num
    qsItem: {
      handler(newVal) {
        /*
        newVal.question.forEach((item, index) => {
          item.num = `Q${index + 1}`
        })
        */
      },
      // deep observation
      deep: true
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../style/QuestionnaireEdit';
</style>
