<template>
    <md-content class="md-elevation-2 md-layout-item md-layout md-gutter md-alignment-center-center">
        <div class="md-layout-item">
            <md-field>
                <label>Name</label>
                <md-input type="text" v-model="name" placeholder="Competitor's name">
                </md-input>
            </md-field>
        </div>
        <div class="md-layout-item md-size-15">
            <md-button class="md-button" id="submit" 
                v-on:click="append" :disabled=notValid >
                <md-icon>add_circle</md-icon>
                APPEND
            </md-button>
        </div>
    </md-content>
</template>

<script>
import Vue from "vue";
import { MdButton, MdField } from "vue-material/dist/components";
import MdInput from "vue-material/dist/components/MdChips";
import { mapActions } from "vuex";

Vue.use(MdButton);
Vue.use(MdField);
Vue.use(MdInput);

export default {
  name: "CategoryCompetitorAppend",
  props: ["categoryId"],
  data() {
    return {
      name: "",
      errors: []
    };
  },
  methods: {
    ...mapActions(["addCategoryCompetitor"]),
    append() {
      if (this.valid() !== true) {
        return false;
      }
      this.addCategoryCompetitor({
        categoryId: this.categoryId,
        name: this.name
      });
      this.name = "";
    },
    valid() {
      this.errors = [];
      if (!this.name) {
        this.errors.push("Competitor's name can't be empty");
      }
      return this.errors.length === 0;
    }
  },
  computed: {
    notValid() {
      return this.valid() !== true;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.md-layout-item {
  text-align: left;
  vertical-align: middle;
}

.md-content {
  margin: 10px 15px;
  padding: 5px 15px;
}
</style>
