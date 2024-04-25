import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
    history: createWebHistory(),
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
            component : ()=> import('../pages/admin/admin.vue'),
            children : [
                {path : 'userinfo', name: '用户信息管理', component: ()=> import('../pages/admin/userinfo.vue')},
                {path : 'modifypwd', name: '修改密码', component: ()=> import('../pages/admin/modifypwd.vue')},
                {path : 'siteset', name: '站点设置', component: ()=> import('../pages/admin/siteset.vue')},
                {path : 'articlelist', name: '文章管理', component: ()=> import('../pages/admin/articlelist.vue')},
                {path : 'articleedit', name: '文章编辑', component: ()=> import('../pages/admin/article.vue')},
                {path : 'commentlist', name: '评论管理', component: ()=> import('../pages/admin/comment.vue')},
                {path : 'categorylist', name: '文章分类管理', component: ()=> import('../pages/admin/category.vue')},
                {path : 'pagelist', name: '页面管理', component: ()=> import('../pages/admin/pagelist.vue')},
                {path : 'pageedit', name: '页面编辑', component: ()=> import('../pages/admin/page.vue')},
                {path : 'findme', name: 'findme管理', component: ()=> import('../pages/admin/findme.vue')},
                {path : 'navigation', name: '导航栏管理', component: ()=> import('../pages/admin/navigation.vue')},
                {path : 'friendlink', name: '友链管理', component: ()=> import('../pages/admin/friendlink.vue')},
            ]
        },
        {
            path:'/',
            name: '主页',
            component : ()=> import('../pages/front/main.vue'),
            children : [
                {path : '', name: '主页', component: ()=> import('../pages/front/indexpage.vue')},
                {path : '/archive', name: '归档', component: ()=> import('../pages/front/archive.vue')},
                {path : '/:pageid', name: '页面', component: ()=> import('../pages/front/page.vue')},
                {path : '/article/:articleid', name: '文章', component: ()=> import('../pages/front/article.vue')},
                {path : '/friendlink', name: '友链', component: ()=> import('../pages/front/friendlink.vue')},
                {path : '/callback', name: 'GitHub登录', component: ()=> import('../components/callback.vue')},
            ]
        }
    ],
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