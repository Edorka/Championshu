<template>
    <ul>
        <li><input type="text" v-model="edition.style" placeholder="Style"/></li>
        <li><input type="text" v-model="edition.division" placeholder="Division"/></li>
        <li>
            <a v-on:click="save"> save </a>
        </li>
    </ul>
</template>

<script>
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
        this.editing = false;
        console.log('new', this.edition);

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
a {
  color: #42b983;
}
</style>
