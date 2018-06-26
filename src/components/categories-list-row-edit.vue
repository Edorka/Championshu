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
                <md-input type="text" v-model="edition.division" placeholder="Division"/>
                </md-input>
            </md-field>
        </div>
        <div class="md-layout-item md-size-15">
            <md-button v-on:click="save"> save </md-button>
        </div>
    </div>
</template>

<script>
import Vue from 'vue'
import { MdButton, MdField } from 'vue-material/dist/components'
import MdInput from 'vue-material/dist/components/MdChips'


export default {
  name: 'CategoriesListRowEdit',
  props: {
    target: Object
  },
  data() {
    return{
        edition: Object.assign({}, this.target)
    }
  },
  methods: {
    save: function() {
        //this.target.push(this.edition);
        Object.assign(this.target, this.edition)
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
