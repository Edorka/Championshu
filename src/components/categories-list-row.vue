<template >
    <li>
      <template v-if="!editing">
        {{item.style}} {{item.division}}<a v-on:click="edit">edit </a>
      </template>
      <template v-else>
          <CategoriesListRowEdit :target="item"/>
      </template>
    </li>
</template>
<script>
import CategoriesListRowEdit from './categories-list-row-edit.vue'

export default {
  name: 'CategoriesListRow',
  props: {
    item: Object
  },
  components: {CategoriesListRowEdit},
  data(){ 
    return {
      editing: false
    };
  },
  methods: {
    edit() {
        this.editing = true;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

li {
  margin: 10px 15px;
}
</style>
