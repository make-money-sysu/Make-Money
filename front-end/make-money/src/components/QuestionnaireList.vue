<template>
  <div class="qs-list">
    QuestionnaireList
    <ul v-if="qslist.length == 0 ? false : true">
      <li></li>
      <li>标题</li>
      <li>截止时间</li>
      <li>状态</li>
      <li>操作<span @click="$router.push({name: 'qsEdit', param: {num: 0}})">+新建问卷</span></li>
    </ul>
    <template v-for="item in qslist">
      <ul>
        <li>
          <input type="checkbox" v-model="item.checked">
        </li>
        <li>{{item.title}}</li>
        <li>{{item.time}}</li>
        <li :class="item.state === 'issued' ? 'issued' : ''">{{item.stateTitle}}</li>
        <li>
          <button @click="iterator = edit(item); iterator.next()">编辑</button>
          <button @click="iterator = delItem(item.num); iterator.next()">删除</button>
          <router-link :to="`/QuestionnaireFill/${item.num}`" tag="button">查看问卷</router-link>
          <button @click="iterator = watchData(item); iterator.next()">查看数据</button>
        </li>
      </ul>
    </template>

    <div class="list-bottom" v-if="qslist.length == 0 ? false : true">
      <label><input type="checkbox" id="all-check" v-model="selectAll">全选</label>
      <button @click="iterator = delItems(); iterator.next()">删除</button>
    </div>
    <div class="add-qs" v-if="qslist.length === 0">
      <button class="add-btn" @click="$router.push({name: 'qsEdit', params: {num: 0}})">+&nbsp;&nbsp;新建问卷</button>
    </div>
    <div class="shadow" v-if="showDialog">
      <div class="del-dialog">
        <header>
          <span>提示</span>
          <span class="close-btn" @click="showDialog = false">X</span>
        </header>
        <p>{{info}}</p>
        <div class="btn-box">
          <button class="yes" @click="iterator.next();">确定</button>
          <button @click="showDialog = false">取消</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  export default {
    name: 'qsList',
    data() {
      return {
        showDialog: false,
        info: '',
        iterator: {},
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
          }]
        }
    },
    mounted() {
      // LHW Server: 182.254.206.244
      // LT Server: 139.199.166.124
      axios.defaults.withCredentials=true;

      const url = 'http://139.199.166.124:8080/survey'
      /*
      axios.get(url)
        .then(response => {
          console.log(response)
        })
        */
      var temp;
      /*
      axios.get(url)
        .then(response => {
          temp = response.data.data
          // Handle list
          temp.forEach(item => {
            console.log(item.id) // 获取num
            console.log(item.name) // 获取title
            console.log(item.content) // 获取questions
          })
          console.log(response.data.data)
        })
      */

      const url_login = 'http://139.199.166.124:8080/login'

      axios.post(url_login, JSON.stringify({
          id: 666,
          password: "123456"
        }))
        .then(response => {
          console.log(response.data)

          axios.post(url, JSON.stringify({
            id: 1,
            publisher_id: 666,
            name: "testing post2",
            content: "[Test POST 最新v2]"
          }))
          .then(response => {
            console.log(response.data)
          })

        })




      this.qslist.forEach(item => {
        let [year, month, day] = item.time.split('-')
        if (year < new Date().getFullYear()) {
          item.state = 'issued'
          item.stateTitle = '已发布'
        }
        else if (year == new Date().getFullYear() && month < new Date().getMonth() + 1) {
          item.state = 'issued'
          item.stateTitle = '已发布'
        }
        else if (year == new Date().getFullYear() && month == new Date().getMonth() + 1 && day < new Date().getDate()) {
          item.state = 'issued'
          item.stateTitle = '已发布'
        }
        else {
          item.state = ''
          item.stateTitle = '未发布'
        }
      })
    },
    methods: {
      showDialogMsg(info) {
        this.showDialog = true;
        this.info = info;
      },
      *delItem(num) {
        yield this.showDialogMsg('确认要删除此问卷')

        yield (() => {
          let index = 0;
          for (let length = this.qslist.length; index < length; index++) {
            if (this.qslist[index].num == num) {
              break;
            }
          }
          this.qslist.splice(index, 1)
          this.showDialog = false;
        })();
      },
      *delItems() {
        yield this.showDialogMsg('确认要删除选中的问卷？')

        yield(() => {
          this.showDialog = false;
          if (this.selectAll) {
            this.qslist = [];
            return;
          }

          if (this.selectGroup == [])
            return;

          this.selectGroup.forEach(item => {
            if (this.qslist.indexOf(item) > -1) {
              this.qslist.splice(this.qslist.indexOf(item), 1);
            }
          })
        })();
      },
      *edit(item) {
        yield (() => {
          if (item.state == 'noissue') {
            this.showDialogMsg('确认要编辑？')
          }
          else {
            this.showDialogMsg('只有未发布的问卷才能编辑');
          }
        })();

        yield (() => {
          if (item.state !== 'noissue') {
            this.showDialog = false;
          }
          else {
            this.showDialog = false;
            this.$router.push({name: 'qsEdit', params: {num: item.num}})
          }
        })();
      },
      *watchData(item) {
        yield (() => {
          if (item.state == 'noissue') {
            this.showDialogMsg('未发布的问卷无数据查阅')
          }
          else {
            this.$router.push({name: 'qsData', params: {num: item.num}})
          }
        })();

        yield this.showDialog = false;
      }
    },
    computed: {
      selectAll: {
        get() {
          return this.selectCount === this.qslist.length && this.selectCount !== 0;
        },
        set(value) {
          this.qslist.forEach(item => {
            item.checked = value;
          });
          return value;
        }
      },
      selectCount() {
        let i = 0;
        this.qslist.forEach(item => {
          if (item.checked)
            i++;
        });
        return i
      },
      selectGroup() {
        let group = [];
        this.qslist.forEach(item => {
          if (item.checked)
            group.push(item)
        });
        return group;
      }
    },
    watch: {
      qslist: {
        handler(val) {
          val.forEach( (item, index) => {
            item.num = index + 1
          })
          // Save in database
        },
        deep: true
      }
    }
  }
</script>

<style lang="scss" scoped>
@import '../style/QuestionnaireList';
</style>
