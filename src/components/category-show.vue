<template>
    <div class="md-gutter">
        <h2>{{ category.style }} -- {{ category.division }}</h2>
        <div class="md-gutter" v-if="competitors">
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
import { MdContent } from 'vue-material/dist/components'
import CategoryCompetitorAppend from '@/components/category-competitor-append.vue'
import CategoryCompetitorRow from '@/components/category-competitor-row.vue'
Vue.use(MdButton);
Vue.use(MdIcon);
Vue.use(MdContent);

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
    .md-layout-item {
        text-align: left;
        vertical-align: middle;
    }
    .md-layout-item .field{
    }
    .buttons{
        text-align: right;
    }
    .buttons button{
        display: inline-block;
    }
</style>
