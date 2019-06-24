<template>
  <div id="head-container">
    <div class="Myhead">
      <span id="title">{{title}}</span>
      <router-link class="link" to="/HomePage" ><img id='logo' src="../../assets/mainpage.png"></router-link>
      <router-link class="link" to="/personalInfo"><img id="personInfo" src="../../assets/personInfo.png"></router-link>
      <router-link class="link" to="/Login"><img id="login" src="../../assets/login.png"></router-link>
      <router-link class="link" to="/register"><img id="register" src="../../assets/register.png"></router-link>
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
