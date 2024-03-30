import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
    routes: [
        {
            path:'/registe',
            name: '注册',
            component : ()=> import('../pages/admin/register.vue')
        },
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

// 三个传入参数，你去哪里，你从哪来，next
router.beforeEach((to, from, next) => {
	// 先获取token
    const token = window.localStorage.getItem('token')
    // 如果访问的是登录页面，则直接放行
    if (to.path === '/login') return next()
    // 如果访问的是用户管理页面，则判断有没有token，没有的话就跳转登录页面。
    if (!token && to.path=='/admin') {
        next('/login')
    } else {
        next()
    }
})

export default router