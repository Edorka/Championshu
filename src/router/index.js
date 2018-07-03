import Vue from 'vue'
import Router from 'vue-router'
import Categories from '@/components/categories'
import CategoryShow from '@/components/category-show'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/categories',
            name: 'categories',
            component: Categories
        },
        {
            path: '/category/:id/show',
            name: 'category-show',
            component: CategoryShow
        },
        { path: '*', redirect: '/categories' }
    ]
})
