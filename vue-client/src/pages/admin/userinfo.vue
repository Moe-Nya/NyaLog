<script setup>
import MessageAPI from '../../components/message.vue';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { usePublicUserInfoStore } from '../../stores/userpublicinfo'

const router = useRouter();
const userinfo = usePublicUserInfoStore();

// 用户信息动态
const username = ref('');
const email = ref('');
const profilephoto = ref('');
const slogan = ref('');

// 动态检查按钮是否生效
const updateUserinfo = ref(false);

// 监听用户信息值
watch([username, email], () => {
    const isEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value);
    updateUserinfo.value = ![username.value, email.value].every(value => value !== '') || !isEmail;
});

// 获取用户信息值
function init() {
    axios.get("/queryuser").then(res => {
        if (res.data.code === 200) {
            username.value = res.data.userinfo.username;
            email.value = res.data.userinfo.email;
            profilephoto.value = res.data.userinfo.profilephoto;
            slogan.value = res.data.userinfo.slogan;
        } else {
            errmsg(res.data.code);
        }
    });
}

// 更新用户信息按钮
function updateUserinfoBtn() {
    window.$loadingBar.start();
    axios.post("/admin/modifyuser", {"username": username.value, "email": email.value, "profilephoto": profilephoto.value, "slogan": slogan.value}, {headers: {Authorization: window.localStorage.getItem('token')}})
    .then(res => {
        if (res.data.code === 200) {
            userinfo.GetUserInfo();
            window.$loadingBar.finish();
            window.$message.success('用户信息更新成功');
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
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
    init();
});
</script>
<template>
    <div class="infotitlebox">
        <h2>用户信息</h2>
        <p style="color:gray; font-size: 14px;">在这里可以更新管理员用户的基础信息，部分信息将展示在你的站点主页。</p>
    </div>
    <div class="infobox">
        <div class="inputbox">
            <i class="userid"></i>
            <span class="input-text"> 用户名</span>
            <n-input v-model:value="username" type="text" placeholder="用户名"/>
            <span class="tips">用户名将会公开展示在站点header上</span>
        </div>
        <div class="inputbox">
            <i class="email"></i>
            <span class="input-text"> Email</span>
            <n-input v-model:value="email" type="text" placeholder="Email"/>
            <span class="tips">邮箱非常重要！它不仅是公开展示的信息，还能够辅助找回密码</span>
        </div>
        <div class="inputbox">
            <i class="profilephoto"></i>
            <span class="input-text"> 用户头像</span>
            <n-input v-model:value="profilephoto" type="text" placeholder="用户头像"/>
            <span class="tips">头像！就是门面！</span>
        </div>
        <div class="inputbox">
            <i class="slogan"></i>
            <span class="input-text"> Slogan</span>
            <n-input v-model:value="slogan" type="text" placeholder="Slogan"/>
            <span class="tips">好的slogan能够让他人更好理解自己</span>
        </div>
        <div class="btn">
            <n-button :disabled="updateUserinfo" @click="updateUserinfoBtn" strong round type="primary" size="large">
                更新
            </n-button>
        </div>
    </div>
</template>