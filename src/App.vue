<template>
  <div id="app">
      <div class="md-toolbar md-theme-default md-elevation-4 md-large md-dense">
        <div class="md-toolbar-row md-toolbar-offset">
          <h3 class="md-title">{{ encounter.name }}</h3>
        </div>
        <div class="md-toolbar-row md-toolbar-offset">
          <router-link to="/categories">categories</router-link>
        </div>
      </div>
      <router-view></router-view>
  </div>
</template>

<script>
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css' // This line here

import Vue from 'vue'
import { mapState } from 'vuex'
import { MdToolbar } from 'vue-material/dist/components'

Vue.use(MdToolbar);
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
}
</style>
