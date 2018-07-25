<template >
    <md-content class="md-elevation-7 md-layout-item md-layout md-gutter">
      <CategoriesListRowShow v-if="!editing" :target="item" v-on:edition="editing=true" />
      <CategoriesListRowEdit v-else :target="item" v-on:saved="editing=false"/>
    </md-content>
</template>
<script>
import CategoriesListRowEdit from "./list-row-edit.vue";
import CategoriesListRowShow from "./list-row-show.vue";

export default {
  name: "CategoriesListRow",
  props: {
    item: Object
  },
  components: { CategoriesListRowShow, CategoriesListRowEdit },
  data() {
    return {
      editing: false
    };
  },
  methods: {
    edit() {
      this.editing = true;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.md-content {
  margin: 10px 15px;
  padding: 5px 15px;
}
</style>
