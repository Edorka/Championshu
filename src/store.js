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
        ],
        "inscriptions": []
    },
    getters: {
        getCategory: (state) => (id) => {
            id = parseInt(id);
            console.log('searching id', id, state.categories );
            return state.categories.find( category => category.id === id)
        },
        getCompetitors: (state) => (id) => {
            id = parseInt(id);
            const byCategoryId =  insc => insc.categoryId === id;
            let inscriptions = state.inscriptions.find(byCategoryId);
            return inscriptions ? inscriptions.competitors : [];
            const segment = state.inscriptions.find( insc => insc.categoryId === id);
            return segment.competitors || [];
        }
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
        },
        ADD_CATEGORY_COMPETITOR: (state, competitor) => {
            const categoryId = competitor.categoryId;
            const byCategoryId =  insc => insc.categoryId === categoryId;
            let position = state.inscriptions.findIndex(byCategoryId) 
            if ( position === -1 ) {
                const newOne = { categoryId, competitors: [ competitor ] };
                Vue.set(state, 'inscriptions', [...state.inscriptions, newOne]);
            } else {
                const competitors = state.inscriptions[position].competitors;
                const updated = {
                    categoryId, 
                    competitors: [...competitors, competitor]
                }
                Vue.set(state.inscriptions, position, updated);
            }
            console.log('done', state.inscriptions);
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
        },
        addCategoryCompetitor: (context, inscription) => {
            context.commit('ADD_CATEGORY_COMPETITOR', inscription);
        }
    }
})
