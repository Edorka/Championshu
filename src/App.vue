<template>
  <div id="app">
      <h1>{{ encounter.name }}</h1>
      <nav>
         <router-link to="/categories">categories</router-link>
      </nav>
      <router-view></router-view>
  </div>
</template>

<script>
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css' // This line here

import { mapState } from 'vuex'

export default {
  name: 'app',
  computed: mapState(['encounter']),
}
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  padding: .5em 2.5em;
}
</style>
