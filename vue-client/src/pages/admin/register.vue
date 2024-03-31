<script setup>
import axios from 'axios';
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue'
import '../../../public/static/css/admin-login&registe.css'
// 实例化路由管理，方便后面的路由跳转
const router = useRouter();
// 获取UID、用户名、密码、邮箱地址
const uid = ref('');
const username = ref('');
const password = ref('');
const email = ref('');

// 按钮状态
const btnDown = ref(true);
const vertifyBtnDown = ref(true);

// 二维码信息
const qrInfo = ref('');

// 验证弹出框宽度
const vertifybox = {width: "550px"};

// 验证码信息
const emailcode = ref('');
const twoFAcode = ref('');

// 监听函数，监听注册信息值
watch([uid, username, password, email], () => {
    const isEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value);
    const isPasswordValid = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]).*$/.test(password.value);
    btnDown.value = ![uid.value, username.value, password.value, email.value].every(value => value !== '') || !isPasswordValid || !isEmail;
});

// 监听验证信息值
watch([emailcode, twoFAcode], () => {
    vertifyBtnDown.value = ![emailcode.value, twoFAcode.value].every(value => value !== '');
});

// 弹出框状态
const showVertifyBox = ref(false);

// 点击下一步按钮
function nextBtn() {
    axios.post("/adminregistration", {"uid": uid.value, "username": username.value, "password": password.value, "email": email.value}).then(response => {
        if (response.data.code == 200) {
            qrInfo.value = "data:image/png;base64,"+response.data.qrcode;
            window.localStorage.setItem('temptoken', response.data.temptoken);
            showVertifyBox.value = true;
            const params = {
                "uid": uid.value,
                "email": email.value
            }
            axios.post("/admin/sendemail", params, {headers: {Authorization: window.localStorage.getItem("temptoken")}}).then(res => {
                if (res.data.code == 200) {
                    window.$message.success('邮件发送成功');
                } else {
                    window.$message.error('邮件发送失败,请检查邮箱地址');
                }
            });
        }
    });
}
//注册按钮
function registeBtn() {
    axios.post("/admin/adminvalidate", {"emailcode": emailcode.value, "twofacode": twoFAcode.value}, {headers: {Authorization: window.localStorage.getItem("temptoken")}}).then(res => {
        if (res.data.code == 200) {
            window.$message.success('用户注册成功');
            window.localStorage.clear("temptoken");
            router.push('/login');
        } else {
            window.$message.error('用户注册失败');
            router.push('/registe');
        }
    });
}

// 打开页面检查若用户存在则跳转登录
function init() {
    axios.get("/adminexist").then(response => {
        if (response.data.code != 200){
            axios.get("/adminvalidate").then(res => {
                if (res.data.code != 200) {
                    router.push('/login');
                }
            });
        }
    });
}
onMounted(() => {
    init();
});
</script>

<style scoped>
.vertifybox{
    display: flex;
}
</style>
<template>
    <div class="pageback">
        <div class="loginbox">
            <div class="loginbox-logo">
                <p class="registelogo">注册</p>
            </div>
            <div class="login-card">
                <div class="inputbox">
                    <i class="userid"></i>
                    <span class="input-text"> 用户UID</span>
                    <n-input v-model:value="uid" type="text" placeholder="用户UID,用户信息唯一标识"/>
                </div>
                <div class="inputbox">
                    <i class="userid"></i>
                    <span class="input-text"> 用户名</span>
                    <n-input v-model:value="username" type="text" placeholder="用户名"/>
                </div>
                <div class="inputbox">
                    <i class="userpwd"></i>
                    <span class="input-text"> 密码</span>
                    <n-input
                    v-model:value="password"
                    type="password"
                    show-password-on="mousedown"
                    placeholder="由数字、大小写字母、标点符号组成"
                    :maxlength="20"
                    />
                </div>
                <div class="inputbox">
                    <i class="email"></i>
                    <span class="input-text"> Email</span>
                    <n-input v-model:value="email" type="text" placeholder="Email"/>
                </div>
                <div class="btn">
                    <n-button :disabled="btnDown" @click="nextBtn" strong round type="primary" size="large">
                        下一步
                    </n-button>
                </div>
            </div>
        </div>
    </div>
    <n-modal
        v-model:show="showVertifyBox"
        class="custom-card vertifybox"
        preset="card"
        :style="vertifybox"
        title="注册信息验证"
        size="huge"
        :bordered="false"
    >
        <p>扫码二维码或者输入key绑定2FA</p><p style="color: red;">注意!请务必确保绑定成功</p>
        <img :src="qrInfo" alt="Base64 Image" :style="{ width: '200px', height: '200px' }"/>
        <template #footer>
            <div class="inputbox">
                <i class="email"></i>
                <span class="input-text"> 邮箱验证码</span>
                <n-input v-model:value="emailcode" type="text" placeholder="邮箱验证码"/>
            </div>
            <div class="inputbox">
                <i class="vertify"></i>
                <span class="input-text"> 2FA验证码</span>
                <n-input v-model:value="twoFAcode" type="text" placeholder="2FA验证码"/>
            </div>
            <div class="btn">
                    <n-button :disabled="vertifyBtnDown" @click="registeBtn" strong round type="primary" size="large">
                        注册
                    </n-button>
                </div>
        </template>
  </n-modal>
</template>