<template>
  <div id="head-container">
    <div class="Myhead">
      <span id="title">{{title}}</span>
      <router-link class="link1" to="/HomePage" >
        <img id='logo' src="../../assets/mainpage.png">
        <span class="tooltiptext">主页面</span>
      </router-link>


      <router-link class="link2" to="/personalInfo">
        <img id="personInfo" src="../../assets/personInfo.png">
        <span class="tooltiptext2">个人信息</span>
      </router-link>

      <router-link class="link3" to="/Login">
        <img id="login" src="../../assets/login.png">
        <span class="tooltiptext3">登陆</span>
      </router-link>

      <router-link class="link4" to="/register">
        <img id="register" src="../../assets/register.png">
        <span class="tooltiptext4">注册</span>
      </router-link>
      <img src="../../assets/sysu.png" id="school">
    </div>
  </div>
</template>

<script>
  export default {
    name: 'vHeader',
    data() {
      return {
        title: "中山大学挣闲钱系统"
      }
    }
  }
</script>

<style lang="scss" scoped>
  @import '../../style/header.scss'
</style>
