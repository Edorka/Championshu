<template>
  <div class="categories-list">
    <h1>Categories</h1>
    items:{{items.length}}
    <ul v-if="items.length > 0">
        <CategoriesListRow  v-for="item in items" :key="item.id" :item="item" />
    </ul>
    <p v-else>Sorry, no items found.</p>
  </div>
</template>

<script>
import CategoriesListRow from './categories-list-row.vue' 

export default {
  name: 'CategoriesList',
  components: {
    CategoriesListRow
  },
  props: {
    items: Array
  },
  data ()  { return {
    item: {'name': 'test'}
  }}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
a {
  color: #42b983;
}
</style>
