<template>
    <div class="md-gutter" v-if="items.length > 0">
        <CategoriesListRow  v-for="item in items" :key="item.id" :item="item" />
    </div>
    <p v-else>Sorry, no items found.</p>
</template>

<script>
import Vue from 'vue'
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
a {
  color: #42b983;
}
</style>
