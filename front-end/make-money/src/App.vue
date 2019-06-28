<template>
  <div id="app">
    <v-header></v-header>
    <router-view></router-view>
  </div>
</template>

<script>
import vHeader from './components/public/header'

export default {
  name: 'App',
  components: {
    vHeader
  },
  created() {
    console.log("Start Main")
  }
}
</script>

<style lang="scss">
@import 'style/app'
</style>
