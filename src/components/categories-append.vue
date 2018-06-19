<template>
    <div class="md-layout md-gutter">
      <div class="md-layout-item md-small-size-100">
            <md-field>
                <label>Style</label>
                <md-input v-model="style" placeholder="ChangQuan, NanQuan, ...">
                </md-input>
            </md-field>
      </div>
      <div class="md-layout-item md-small-size-100">
            <md-field>
                <label>Division</label>
                <md-input v-model="division" placeholder="Male, Female, Jr, ...">
                </md-input>
            </md-field>
      </div>
      <div class="md-layout-item md-small-size-100">
            <md-button class="md-primary" id="submit" v-on:click="append" :disabled=notValid> add </md-button>
      </div>
    </div>
</template>

<script>
import Vue from 'vue'
import { MdButton, MdField } from 'vue-material/dist/components'
import MdInput from 'vue-material/dist/components/MdChips'

Vue.use(MdButton)
Vue.use(MdField)
Vue.use(MdInput)
export default {
  name: 'CategoriesAppend',
  props: {
    target: Array
  },
  data() {
    return{
        style: '',
        division: '',
        errors: []
    }
  },
  methods: {
    append() {
        console.log('try', this.valid() );
        if ( this.valid() !== true ) { return false; }
        const id = this.target.length + 1;
        const style = this.style;
        const division = this.division;
        this.target.push({
            id,
            style,
            division
        });
        this.style = '';
        this.division = '';
    },      
    valid() {
      this.errors = [];
      if ( !this.style ) {
        this.errors.push('Style can\'t be empty');
      }
      if ( !this.division ) {
        this.errors.push('Division can\'t be empty');
      }
      return this.errors.length === 0;
    }
  },
  computed: {
    notValid() {
      return this.valid() !== true;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
</style>
