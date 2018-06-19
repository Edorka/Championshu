<template>
    <ul>
        <li><input type="text" v-model="style" placeholder="Style"/></li>
        <li><input type="text" v-model="division" placeholder="Division"/></li>
        <li>
            <button id="submit" v-on:click="append" :disabled=notValid> add </button>
        </li>
    </ul>
</template>

<script>
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
