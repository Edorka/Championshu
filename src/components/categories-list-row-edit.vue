<template>
    <div class="md-layout md-gutter">
        <div class="md-layout-item">
            <md-field>
                <label>Division</label>
                <md-input type="text" v-model="edition.style" placeholder="Style">
                </md-input>
            </md-field>
        </div>
        <div class="md-layout-item">
            <md-field>
                <label>Division</label>
                <md-input type="text" v-model="edition.division" placeholder="Division" />
            </md-field>
        </div>
        <div class="md-layout-item md-size-15">
            <md-button class="md-icon-button" v-on:click="save">
                <md-icon>save_alt</md-icon>
            </md-button>
        </div>
    </div>
</template>

<script>
import { mapActions } from 'vuex'


export default {
  name: 'CategoriesListRowEdit',
  props: {
    target: Object
  },
  data() {
    return {
        edition: Object.assign({}, this.target)
    }
  },
  methods: {
    ...mapActions(['updateCategory']),
    save() {
        this.updateCategory(this.edition);
        this.$emit('saved');
    }      
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .md-layout-item {
        text-align: left;
        vertical-align: middle;
    }
</style>
