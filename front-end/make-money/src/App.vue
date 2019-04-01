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
  }
}
</script>

<style>
</style>
