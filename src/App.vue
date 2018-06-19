<template>
  <div id="app">
      <CategoriesList :items="categories"/>
      <CategoriesAppend :target="categories"/>
  </div>
</template>

<script>
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css' // This line here

import CategoriesList from './components/categories-list.vue'
import CategoriesAppend from './components/categories-append.vue'

export default {
  name: 'app',
  components: {
    CategoriesList,
    CategoriesAppend
  },
  data() {return {
    categories: [
        {'id': 1,'style':'Chang Quan','division':'Sr. Male'},
        {'id': 2,'style':'Chang Quan','division':'Sr. Female'},
        {'id': 3,'style':'Chang Quan','division':'Jr. Mixed'},
    ]
  }
  }
}
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  padding: 1.5em .5em;
}
</style>
