import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
    routes: [
        {
            path:'/login',
            name: '登录',
            component : ()=> import('../pages/admin/login.vue')
        },
        {
            path:'/admin',
            name: '后台管理',
            component : ()=> import('../pages/admin/admin.vue')
        }
    ],
    history: createWebHistory()
})

export default router