<template>
  <div id="categories">
    <CategoriesList :items="categories"/>
    <CategoriesAppend :target="categories"/>
  </div>
</template>

<script>
import { mapState } from "vuex";
import CategoriesList from "./components/list.vue";
import CategoriesAppend from "./components/append.vue";

export default {
  name: "Categories",
  components: {
    CategoriesList,
    CategoriesAppend
  },
  computed: mapState(["encounter", "categories"])
};
</script>

<style>
#categories {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  padding: 0.5em 2.5em;
}
</style>
