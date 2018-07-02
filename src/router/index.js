import Vue from 'vue'
import Router from 'vue-router'
import Categories from '@/components/categories'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/categories',
            name: 'categories',
            component: Categories
        },
        { path: '*', redirect: '/categories' }
    ]
})
