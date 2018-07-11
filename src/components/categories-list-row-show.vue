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
            <md-badge v-if="inscriptions.length > 0" 
                      class="md-primary md-square" 
                      :md-content="inscriptions.length">
                <router-link tag="md-button" class="md-icon-button"
                    :to="{ name: 'category-show', params: { id: target.id }}">
                    <md-icon>person</md-icon>
                </router-link>
            </md-badge>
            <router-link v-else tag="md-button" class="md-icon-button"
                :to="{ name: 'category-show', params: { id: target.id }}">
                <md-icon>person_add</md-icon>
            </router-link>
        </div>
    </div>
</template>

<script>
import Vue from 'vue'
import { mapActions, mapGetters } from 'vuex'
import { MdButton, MdIcon } from 'vue-material/dist/components'
import { MdContent } from 'vue-material/dist/components'
import { MdBadge } from 'vue-material/dist/components'
Vue.use(MdButton);
Vue.use(MdIcon);
Vue.use(MdContent);
Vue.use(MdBadge);

export default {
  name: 'CategoriesListRowShow',
  props: {
    target: Object
  },
  computed: {
    ...mapGetters(['getCompetitors']),
    inscriptions() {
      return this.getCompetitors(this.target.id)
    }
  },
  methods: {
    ...mapActions(['removeCategory']),
    edit() {
        this.$emit('edition');
    },
    remove() {
        this.removeCategory(this.target.id);
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
