<script setup>
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue'
import axios from 'axios'
import '../../../public/static/css/admin.css'
import '../../../public/static/css/admin-login&registe.css'
import errmsg from '../../modules/errmsg';
import { onMounted } from 'vue';
import { useBlogSetStore } from '../../stores/blogset'
import { usePublicUserInfoStore } from '../../stores/userpublicinfo'
import { RouterLink } from 'vue-router'
import { defineComponent, h } from "vue";
const router = useRouter();
const blogset = useBlogSetStore();
const userinfo = usePublicUserInfoStore();
// 定义后台首页
router.push('/admin/userinfo');

// 控制侧边栏弹出框的显示与否
const naviconshow = ref(true);

const menuOptions = [
    {
        type: "group",
            label: "用户",
            key: "user",
            children: [
            {
                label: () => h(
                    RouterLink,
                    {
                        to: {
                            name: '用户信息管理',
                        }
                    },
                    '用户信息管理'
                ),
                key: "userinfo"
            },
            {
                label: () => h(
                    RouterLink,
                    {
                        to: {
                            name: '修改密码',
                        }
                    },
                    '修改密码'
                ),
                key: "modifypwd"
            },
            ],
    },
    {
        type: "group",
            label: "系统设置",
            key: "system",
            children: [
            {
                label: () => h(
                    RouterLink,
                    {
                        to: {
                            name: '站点设置',
                        }
                    },
                    '站点设置'
                ),
                key: "siteset"
            },
            ],
    },
];

// 侧边菜单交换按钮
function navswitch() {
    const main = document.getElementById('main');
    main.classList.add('disable')
    const nav = document.getElementById('nav');
    nav.classList.add('visable');
    const userinfoDes = document.getElementById('userinfoDes');
    userinfoDes.classList.add('change');
    naviconshow.value = false;
}
function closeBtnBox() {
    naviconshow.value = true;
    const nav = document.getElementById('nav');
    nav.classList.remove('visable');
    const main = document.getElementById('main');
    main.classList.remove('disable')
}

// 登出
function loginOut() {
    window.$loadingBar.start();
    axios.post("/admin/loginout",{},{headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$message.success('注销登录成功');
            window.localStorage.clear('token');
            window.$loadingBar.finish();
            router.push('/login');
        } else {
            errmsg(res.data.code);
            window.$loadingBar.error();
        }
    });
}
// 返回主页
function reMain() {
    router.push('/')
}

// 进入页面权限验证
function Validate() {
    window.$loadingBar.start();
    axios.get("/admin/", {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code !== 200) {
            window.localStorage.clear('token');
            window.$loadingBar.finish();
            errmsg(res.data.code);
            router.push('/login');
        } else {
            window.$loadingBar.finish();
        }
    });
}

// 初始化
onMounted(() => {
    Validate();
});
</script>


<template>
    <div class="container">
        <div class="horizontalnav">
            <div class="naviconbox" v-show="naviconshow" @click="navswitch">
                <i class="navicon"></i>
            </div>
            <div class="navlogo">
                <span class="navlogotext">{{ blogset.data.sitename }}</span>
            </div>
        </div>
        <div class="nav" id="nav">
            <div class="closeBtnBox" @click="closeBtnBox">
                <i class="closeBtn"></i>
            </div>
            <n-space vertical class="navi">
                <n-layout has-sider>
                    <div>
                        <n-layout
                            bordered
                            collapse-mode="width"
                            :width="250"
                            show-trigger
                        >
                            <n-menu
                            @click="closeBtnBox"
                            default-value="userinfo"
                            style="min-width: 220px;"
                            :options="menuOptions"
                            />
                        </n-layout>
                    </div>
                </n-layout>
                <div class="userinfo">
                    <n-avatar
                    round
                    :size="60"
                    :src="userinfo.data.profilephoto"
                    />
                    <div class="userinfo-des" id="userinfoDes">
                        <span class="username">{{ userinfo.data.username }}</span><br>
                        <n-button @click="reMain" strong secondary type="primary" size="small">主页</n-button>
                        <n-button style="margin-left: 5px;" @click="loginOut" strong secondary type="error" size="small">登出</n-button>
                    </div>
                </div>
            </n-space>
        </div>
        <div class="main" id="main">
            <div class="main-container">
                <router-view />
            </div>
        </div>
        <foo />
    </div>
</template>
