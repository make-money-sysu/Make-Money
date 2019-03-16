<template>
  <div class="qs-list">
    QuestionnaireList
    <ul v-if="qslist.length == 0 ? false : true">
      <li></li>
      <li>标题</li>
      <li>截至时间</li>
      <li>状态</li>
      <li>操作</li>
      <li>操作<span @click="$router.push({name: 'qsEdit', param: {num: 0}})">+新建问卷</span></li>
    </ul>
    <template v-for="item in qslist">
      <ul>
        <li>
          <input type="checkbox" v-model="item.checked">
        </li>
        <li>{{item.title}}</li>
        <li>{{item.time}}</li>
        <li :class="item.state === 'inissue' ? 'inissue' : ''">{{item.stateTitle}}</li>
        <li>
          <button @click="iterator = edit(item); iterator.next()">编辑</button>
          <button @click="iterator = delItem(item.num); iterator.next()">删除</button>
          <router-link :to="`/fill/${item.num}`" tag="button">查看问卷</router-link>
          <button @click="iterator = watchData(item); iterator.next()">查看数据</button>
        </li>
      </ul>
    </template>
    <router-view></router-view>
  </div>
</template>

<script>
  export default {
    name: 'qsList',
    data() {
      return {
        qslist: [{ 'num': 3,
            'title': '第三份问卷',
            'time': '2017-3-27',
            'state': 'issueed',
            'stateTitle': '已发布',
            'checked': false,
            'question': [
              {'num': 'Q1', 'title': '单选题', 'type': 'radio', 'isNeed': true, 'options': ['选项一', '选项二']},
              {'num': 'Q2', 'title': '多选题', 'type': 'checkbox', 'isNeed': true, 'options': ['选项一', '选项二', '选项三', '选项四']},
              {'num': 'Q3', 'title': '文本题', 'type': 'textarea', 'isNeed': true}
            ]
          }]
        }
    }
  }
</script>

<style>
</style>
