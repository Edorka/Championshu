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
    }
})
