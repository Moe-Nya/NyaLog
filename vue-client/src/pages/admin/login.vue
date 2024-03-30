<script setup>
import { onMounted, watch } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue'
import '../../../public/static/css/admin-login&registe.css'
import { useBlogSetStore } from '../../stores/blogset'
// 实例化路由管理，方便后面的路由跳转
const router = useRouter();

// 获取站点名称
const sitename = useBlogSetStore();
sitename.GetBlogInfo();

// 获取用户名、密码、2FA
const username = ref('');
const password = ref('');
const twoFA = ref('');

const btnDown = ref(true);

// 监听函数，监听这三个值
watch([username, password, twoFA], () => {
    const isPasswordValid = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]).*$/.test(password.value)
    btnDown.value = ![username.value, password.value, twoFA.value].every(value => value !== '') || !isPasswordValid;
});

function loginBtn() {
    axios.post("/login", {"uid": username.value, "password": password.value, "twofacode": twoFA.value}).then(response => {
        switch(response.data.code) {
            case 3002: 
                window.$message.error('2FA验证码错误');
                break;
            case 2008: 
                window.$message.error('密码错误');
                break;
            case 2007: 
                window.$message.error('UID错误');
                break;
            case 2009: 
                window.$message.error('IP登录受限');
                break;
            case 200:
                window.$message.success('登录成功');
                break;
        }
        if (response.data.code == 200) {
            window.localStorage.setItem('token', response.data.token)
            router.push('/admin');
        }
    });
}

// 访问login页面时，清空token
function init() {
    window.localStorage.clear('token')
}
onMounted(() => {
    init();
});
</script>

<template>
    <div class="pageback">
        <div class="loginbox">
            <div class="loginbox-logo">
                <p class="loginlogo">登录到 <span class="loginlogo-sitename">{{ sitename.data.sitename }}</span></p>
            </div>
            <div class="login-card">
                <div class="inputbox">
                    <i class="userid"></i>
                    <span class="input-text"> 用户UID</span>
                    <n-input v-model:value="username" type="text" placeholder="用户UID"/>
                </div>
                <div class="inputbox">
                    <i class="userpwd"></i>
                    <span class="input-text"> 密码</span>
                    <span class="forgetpwd"><a href="#" style="color: gray; text-decoration: none;">忘记密码?</a></span>
                    <n-input
                    v-model:value="password"
                    type="password"
                    show-password-on="mousedown"
                    placeholder="由数字、大小写字母、标点符号组成"
                    :maxlength="20"
                    />
                </div>
                <div class="inputbox">
                    <i class="vertify"></i>
                    <span class="input-text"> 2FA验证码</span>
                    <n-input v-model:value="twoFA" type="text" placeholder="2FA验证码"/>
                </div>
                <div class="btn">
                    <n-button :disabled="btnDown" @click="loginBtn" strong round type="primary" size="large">
                        登录
                    </n-button>
                </div>
            </div>
        </div>
        <n-message-provider />
    </div>
</template>