<template>
  <div id="app">
    mainpage
    <router-view></router-view>
  </div>
</template>

<script>
import vHeader from './public/header'

export default {
  name: 'App',
  components: {
    vHeader
  }
}
</script>

<style>
</style>
