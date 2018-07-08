<template>
    <div class="md-layout md-gutter md-alignment-center-center">
        <div class="md-layout-item">
            <div class="field">{{item.name}}</div>
        </div>
        <div class="md-layout-item md-size-25 buttons">
            <md-button class="md-icon-button" >
                <md-icon>remove_circle</md-icon>
            </md-button>
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
  name: 'CategoryCompetitorRow',
  props: {
    item: Object
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
