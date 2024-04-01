<script setup>
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue'
import axios from 'axios'
import '../../../public/static/css/admin.css'
import errmsg from '../../modules/errmsg';
import { usePublicUserInfoStore } from '../../stores/userpublicinfo'
import { useBlogSetStore } from '../../stores/blogset'
import { onMounted } from 'vue';
const router = useRouter();
const userinfo = usePublicUserInfoStore();
const blogset = useBlogSetStore();

// 控制侧边栏弹出框的显示与否
const naviconshow = ref(true);

const menuOptions = [
    {
        type: "group",
            label: "用户",
            key: "user",
            children: [
            {
                label: "更改用户信息",
                key: "narrator"
            },
            {
                label: "更改密码",
                key: "sheep-man"
            }
            ]
    },
];

// 侧边菜单交换按钮
function navswitch() {
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
}

// 登出
function loginOut() {
    window.$loadingBar.start();
    axios.post("/admin/loginout",{},{headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code == 200) {
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
                            :collapsed-width="64"
                            :collapsed-icon-size="22"
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
                        <span class="username">11</span><br>
                        <n-button @click="reMain" strong secondary type="primary" size="small">主页</n-button>
                        <n-button style="margin-left: 5px;" @click="loginOut" strong secondary type="error" size="small">登出</n-button>
                    </div>
                </div>
            </n-space>
        </div>
        <div class="main">欸嘿嘿嘿{{ userinfo.data.profilephoto }}</div>
        
        <foo />
    </div>
</template>
