<template>
    <div class="md-layout md-gutter">
        <div class="md-layout-item">
            <span>{{target.style}}</span>
        </div>
        <div class="md-layout-item">
            <span>{{target.division}}</span>
        </div>
        <div class="md-layout-item md-size-15">
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
<style lang="scss" scoped>
    .md-layout-item {
        text-align: left;
        vertical-align: middle;
    }
    .md-layout-item span{
        margin-top: 1em;
    }
    
</style>
