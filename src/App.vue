<template>
  <div id="app">
      <h1>{{ encounter.name }}</h1>
      <div class="categories-list">
        <h2>Categories</h2>
        <CategoriesList :items="categories"/>
        <CategoriesAppend :target="categories"/>
      </div>
  </div>
</template>

<script>
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css' // This line here

import { mapState } from 'vuex'
import CategoriesList from './components/categories-list.vue'
import CategoriesAppend from './components/categories-append.vue'


export default {
  name: 'app',
  components: {
    CategoriesList,
    CategoriesAppend
  },
  computed: mapState(['encounter', 'categories']),
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
