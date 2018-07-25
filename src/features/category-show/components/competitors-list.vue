<template>
    <div id="competitors-list" 
        class="md-gutter md-content" v-if="items.length > 0">
        <CompetitorsListRow  v-for="item in items" :key="item.id" :item="item" />
    </div>
    <p v-else>Sorry, no items found.</p>
</template>

<script>
import CompetitorsListRow from "./competitors-list-row.vue";
import { mapActions } from "vuex";

export default {
  name: "CompetitorsList",
  components: {
    CompetitorsListRow
  },
  props: {
    items: Array
  },
  methods: {
    ...mapActions(["retrieveData"]),
  },
  data() {
    return {
    };
  },
  created() {
    this.retrieveData();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  color: #42b983;
}
#competitors-list {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  padding: 0.5em 2.5em;
}
</style>
