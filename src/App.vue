<template>
  <div id="app">
    <CategoriesList :items="categories"/>
    <CategoriesAppend :target="categories"/>
  </div>
</template>

<script>
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
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
