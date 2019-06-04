<template>
  <div class="Myhead">
    <router-link class="link" to="/HomePage" ><img id='logo' src="../../assets/mainpage.png"></router-link>
    <router-link class="link" to="/personalInfo"><img id="personInfo" src="../../assets/personInfo.png"></router-link>
    <router-link class="link" to="/order"><img id="order" src="../../assets/order.png"></router-link>
    <router-link class="link" to="/Login"><img id="login" src="../../assets/login.png"></router-link>
    <router-link class="link" to="/register"><img id="register" src="../../assets/register.png"></router-link>
  </div>
</template>

<script>
  export default {
    name: 'vHeader'
  }
</script>

<style lang="scss" scoped>
  @import '../../style/header'
</style>
