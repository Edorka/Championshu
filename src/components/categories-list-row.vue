<template >
    <div class="md-layout-item md-layout md-gutter">
      <CategoriesListRowShow v-if="!editing" :target="item" v-on:edition="editing=true" />
      <CategoriesListRowEdit v-else :target="item" v-on:saved="editing=false"/>
    </div>
</template>
<script>
import CategoriesListRowEdit from './categories-list-row-edit.vue'
import CategoriesListRowShow from './categories-list-row-show.vue'

export default {
  name: 'CategoriesListRow',
  props: {
    item: Object
  },
  components: {CategoriesListRowShow, CategoriesListRowEdit},
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
