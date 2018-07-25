<template>
    <div id="category-detail">
        <h2>{{ category.style }} -- {{ category.division }}</h2>
        <CompetitorsList :items="competitors" />
        <CompetitorsAppend :category-id="category.id" />
    </div>
</template>

<script>
import Vue from "vue";
import { mapGetters } from "vuex";

import { MdButton, MdIcon } from "vue-material/dist/components";
import CompetitorsAppend from "./components/competitors-append.vue";
import CompetitorsList from "./components/competitors-list.vue";
Vue.use(MdButton);
Vue.use(MdIcon);

export default {
  name: "CategoryShow",
  components: { CompetitorsAppend, CompetitorsList },
  props: {},
  data() {
    return {};
  },
  computed: {
    ...mapGetters(["getCategory", "getCompetitors"]),
    category() {
      const categoryId = this.$route.params.id;
      return this.getCategory(categoryId);
    },
    competitors() {
      const categoryId = this.$route.params.id;
      return this.getCompetitors(categoryId);
    }
  }
};
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
#competitors-list {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  padding: 0.5em 2.5em;
}
</style>
