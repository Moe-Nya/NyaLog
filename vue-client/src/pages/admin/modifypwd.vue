<script setup>
import MessageAPI from '../../components/message.vue'
import axios from 'axios'
import errmsg from '../../modules/errmsg';
import { useRouter } from 'vue-router';
import { usePublicUserInfoStore } from '../../stores/userpublicinfo'

const router = useRouter();
const userinfo = usePublicUserInfoStore();

// 静态用户信息
const email = userinfo.data.email;

// 动态监测输入框
const uid = ref('');
const password = ref('');
const emailcode = ref('');

// 监测两个框是否被正确填入
const updatePwd = ref(true);
watch([uid, password, emailcode], () => {
    const isPassword = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]).*$/.test(password.value);
    updatePwd.value = ![uid.value, password.value, emailcode.value].every(value => value !== '') || !isPassword;
});

// 发送邮件验证码
function sendEmail() {
    window.$loadingBar.start();
    axios.post("/admin/sendemail", {"uid": uid.value, "email": email}, {headers: {Authorization: window.localStorage.getItem("token")}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('邮件发送成功');
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 修改密码
function updateUserinfoBtn() {
    window.$loadingBar.start();
    axios.post("/admin/resetpwdcode", {"uid": uid.value, "password": password.value, "code": emailcode.value}, {headers: {Authorization: window.localStorage.getItem("token")}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('密码修改成功');
            router.push('/login');
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
            window.$loadingBar.error();
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
    <div class="userinfotitlebox">
        <h2>修改密码</h2>
        <p style="color:gray; font-size: 14px;">请好好记住自己的密码哦~ 修改密码需要用到个人邮箱账户，记得点击发送邮件验证码。</p>
    </div>
    <div class="infobox">
        <div class="inputbox">
            <i class="userid"></i>
            <span class="input-text"> 用户UID</span>
            <n-input v-model:value="uid" type="text" placeholder="用户UID"/>
            <span class="tips">注意不要和用户名弄混了哦~</span>
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
            <span class="tips">输入想要修改的新密码</span>
        </div>
        <div class="inputbox">
            <i class="email"></i>
            <span class="input-text"> 邮箱验证码</span>
            <n-input v-model:value="emailcode" type="text" placeholder="邮箱验证码"/>
            <span class="tips">去看看邮箱哦~</span>
        </div>
        <div class="btn">
            <n-button @click="sendEmail" strong round type="info" size="large">
                发送验证码
            </n-button>
            <n-button style="margin-left: 5px;" :disabled="updatePwd" @click="updateUserinfoBtn" strong round type="primary" size="large">
                修改密码
            </n-button>
        </div>
    </div>
</template>