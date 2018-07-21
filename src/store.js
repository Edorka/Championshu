import Vuex from 'vuex'
import Vue from 'vue'
import axios from 'axios'

Vue.use(Vuex)

const headers = new Headers({
    'Content-Type': 'application/json',
    'Accept': 'application/json',
});

export default new Vuex.Store({
    state: {
        "encounter": {
            "name": "First championship"
        },
        "categories": [],
        "inscriptions": []
    },
    getters: {
        getCategory: (state) => (id) => {
            id = parseInt(id);
            return state.categories.find( category => category.id === id)
        },
        getCompetitors: (state) => (id) => {
            id = parseInt(id);
            const byCategoryId =  insc => insc.categoryId === id;
            let inscriptions = state.inscriptions.find(byCategoryId);
            return inscriptions ? inscriptions.competitors : [];
        }
    },
    mutations: {
        RETRIEVE_DATA: (state, data) => {
            Vue.set(state, 'categories', data.categories);
        },
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
        },
        NEW_CATEGORY_COMPETITOR: (state, competitor) => {
            const categoryId = competitor.categoryId;
            const newOne = { categoryId, competitors: [ competitor ] };
            Vue.set(state, 'inscriptions', [...state.inscriptions, newOne]);
        },
        APPEND_TO_CATEGORY_COMPETITORS: (state, {position, competitor}) => {
            const categoryId = competitor.categoryId;
            const competitors = state.inscriptions[position].competitors;
            const updated = {
                categoryId, 
                competitors: [...competitors, competitor]
            }
            Vue.set(state.inscriptions, position, updated);
        }

    },
    actions: {
        retrieveData: (context) => {
            const params = { headers };
            fetch(process.env.VUE_APP_API_HOST + '/categories', params)
                .catch( error => console.error(error))
                .then( response => response.json() )
                .then( data => context.commit('RETRIEVE_DATA', data) );
        },
        addCategory: (context, category) => {
            const params = { method: 'POST', headers, body: JSON.stringify(category) };
            fetch(process.env.VUE_APP_API_HOST + '/categories', params)
                .catch( error => console.error(error))
                .then( response => response.ok ? response.json() : response.text() )
                .then( data => context.commit('ADD_CATEGORY', data) );
        },
        removeCategory: (context, targetId) => {
            context.commit('REMOVE_CATEGORY', targetId);
        },
        updateCategory: (context, updated) => {
            context.commit('UPDATE_CATEGORY', updated);
        },
        addCategoryCompetitor: (context, competitor) => {
            const categoryId = competitor.categoryId;
            const byCategoryId =  insc => insc.categoryId === categoryId;
            let position = context.state.inscriptions.findIndex(byCategoryId);
            if ( position === -1 ) {
                context.commit('NEW_CATEGORY_COMPETITOR', competitor);
            } else {
                context.commit('APPEND_TO_CATEGORY_COMPETITORS', {position, competitor});
            }
        }
    }
})
