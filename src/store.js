import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        "encounter": {
            "name": "First championship"
        },
        "categories": [
            {'id': 1,'style':'Chang Quan','division':'Sr. Male'},
            {'id': 2,'style':'Chang Quan','division':'Sr. Female'},
            {'id': 3,'style':'Chang Quan','division':'Jr. Mixed'},
        ]
    },
    mutations: {
        ADD_CATEGORY: (state, category) => {
            category.id = state.categories.length + 1;
            category.competitors = [];
            state.categories.push(category);
        },
        REMOVE_CATEGORY: (state, targetId) => {
            state.categories = state.categories.filter( current => current.id != targetId );
        },
        UPDATE_CATEGORY: (state, updated) => {
            const index = state.categories.findIndex((current) => current.id === updated.id);
            Vue.set(state.categories, index, updated);
        }
    },
    actions: {
        addCategory: (context, category) => {
            context.commit('ADD_CATEGORY', category);
        },
        removeCategory: (context, targetId) => {
            context.commit('REMOVE_CATEGORY', targetId);
        },
        updateCategory: (context, updated) => {
            context.commit('UPDATE_CATEGORY', updated);
        }
    }
})
