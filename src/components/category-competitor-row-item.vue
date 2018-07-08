<template>
    <div class="md-layout md-gutter md-alignment-center-center">
        <div class="md-layout-item">
            <div class="field">{{target.style}}</div>
        </div>
        <div class="md-layout-item">
            <div class="field">{{target.division}}</div>
        </div>
        <div class="md-layout-item md-size-25 buttons">
            <md-button class="md-icon-button" v-on:click="edit">
                <md-icon>create</md-icon>
            </md-button>
            <md-button class="md-icon-button" v-on:click="remove">
                <md-icon>remove_circle</md-icon>
            </md-button>
            <router-link tag="md-button" class="md-icon-button"
                :to="{ name: 'category-show', params: { id: 123 }}">
                <md-icon>remove_circle</md-icon>
            </router-link>
        </div>
    </div>
</template>

<script>
import Vue from 'vue'
import { mapActions } from 'vuex'
import { MdButton, MdIcon } from 'vue-material/dist/components'
import { MdContent } from 'vue-material/dist/components'
Vue.use(MdButton);
Vue.use(MdIcon);
Vue.use(MdContent);

export default {
  name: 'CategoriesListRowShow',
  props: {
    target: Object
  },
  methods: {
    ...mapActions(['removeCompetitor']),
    edit() {
        this.$emit('edition');
    },
    remove() {
        this.removeCompetitor(this.target.id);
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
    .buttons{
        text-align: right;
    }
    .buttons button{
        display: inline-block;
    }
</style>
