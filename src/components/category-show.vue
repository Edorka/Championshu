<template>
    <div id="category-detail">
        <h2>{{ category.style }} -- {{ category.division }}</h2>
        <div id="competitors-list" 
            class="md-gutter md-content" v-if="competitors.length > 0">
          <CategoryCompetitorRow  v-for="competitor in competitors" 
              :key="competitor.id" :item="competitor" />
        </div>
        <p v-else>No competitors on this category.</p>
        <CategoryCompetitorAppend :category-id="category.id" />
    </div>
</template>

<script>
import Vue from 'vue'
import { mapGetters } from 'vuex'

import { MdButton, MdIcon } from 'vue-material/dist/components'
import CategoryCompetitorAppend from '@/components/category-competitor-append.vue'
import CategoryCompetitorRow from '@/components/category-competitor-row.vue'
Vue.use(MdButton);
Vue.use(MdIcon);

export default {
  name: 'CategoryShow',
  components: { CategoryCompetitorAppend, CategoryCompetitorRow },
  props: {
  },
  data() {
      return {}
  },
  computed: {
    ...mapGetters(['getCategory', 'getCompetitors']),
    category() {
      const categoryId = this.$route.params.id;
      return this.getCategory(categoryId)
    },
    competitors() {
      const categoryId = this.$route.params.id;
      return this.getCompetitors(categoryId);
    }
  }
}</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
    #competitors-list {
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
      text-align: center;
      padding: .5em 2.5em;
    }
</style>
