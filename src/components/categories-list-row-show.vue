<template>
    <div class="md-layout md-gutter">
        <div class="md-layout-item md-alignment-center-center">
            <md-content>{{target.style}}</md-content>
        </div>
        <div class="md-layout-item md-alignment-center-center">
            <md-content>{{target.division}}</md-content>
        </div>
        <div class="md-layout-item">
            <md-button v-on:click="edit">edit </md-button>
        </div>
    </div>
</template>

<script>
import Vue from 'vue'
import { MdButton } from 'vue-material/dist/components'
import { MdContent } from 'vue-material/dist/components'
Vue.use(MdButton);
Vue.use(MdContent);

export default {
  name: 'CategoriesListRowShow',
  props: {
    target: Object
  },
  methods: {
    edit() {
        this.$emit('edition');
        console.log('edition target', this.target);
    }      
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .md-content {
        display: inline-flex;
        justify-content: center;
        align-items: center;
    }
</style>
