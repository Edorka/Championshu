<template>
    <div class="md-gutter" v-if="items.length > 0">
        <CategoriesListRow  v-for="item in items" :key="item.id" :item="item" />
    </div>
    <p v-else>Sorry, no items found.</p>
</template>

<script>
import CategoriesListRow from "./categories-list-row.vue";
import { mapActions } from "vuex";

export default {
  name: "CategoriesList",
  components: {
    CategoriesListRow
  },
  props: {
    items: Array
  },
  methods: {
    ...mapActions(["retrieveData"]),
  },
  data() {
    return {
    };
  },
  created() {
    this.retrieveData();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  color: #42b983;
}
</style>
