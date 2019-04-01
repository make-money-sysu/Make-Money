<template>
  <div class="header">
    <router-link class="link" to="/"><img id='logo' src="../../assets/mainpage.png"></router-link>
    <router-link class="link" to="/personalInfo"><img id="icon" src="../../assets/icon.png"></router-link>
    <router-link class="link" to="/order"><img id="order" src="../../assets/order.png"></router-link>
  </div>
</template>

<script>
  export default {
    name: 'header'
  }
</script>

<style lang="scss" scoped>
  @import '../../style/header'
</style>
