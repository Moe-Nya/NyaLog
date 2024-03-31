<script setup>
import { onMounted, watch } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue'
import '../../../public/static/css/admin-login&registe.css'
import { useBlogSetStore } from '../../stores/blogset'
import errmsg from '../../modules/errmsg';
// 实例化路由管理，方便后面的路由跳转
const router = useRouter();

// 获取站点名称
const sitename = useBlogSetStore();
sitename.GetBlogInfo();

// 获取用户名、密码、2FA
const username = ref('');
const password = ref('');
const twoFA = ref('');

// 获取修改密码的uid、邮箱地址、2FA
const reuid = ref('')
const reemail = ref('')
const retwoFA = ref('')

// 获取修改密码的uid、新密码、邮箱验证码
const newuid = ref('')
const newpwd = ref('')
const emailcode = ref('')

const btnDown = ref(true);
const rePwdBtnDown = ref(true);
const newPwdBtnDown = ref(true);

// 忘记密码弹出框状态
const modifyPwdBox = {width: "550px"};
const showModifyPwdBox = ref(false);
const showModifyPwdvertifyBox = ref(false);

// 监听函数，监听这4个值
watch([username, password, twoFA], () => {
    const isPasswordValid = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]).*$/.test(password.value);
    btnDown.value = ![username.value, password.value, twoFA.value].every(value => value !== '') || !isPasswordValid;
});

// 监听修改密码信息值
watch([reuid, reemail, retwoFA], () => {
    rePwdBtnDown.value = ![reuid.value, reemail.value, retwoFA.value].every(value => value !== '');
});

// 监听新密码信息值
watch([newuid, newpwd, emailcode], () => {
    const isPasswordValid = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]).*$/.test(newpwd.value);
    newPwdBtnDown.value = ![newuid.value, newpwd.value, emailcode.value].every(value => value !== '') || !isPasswordValid;
});

function loginBtn() {
    axios.post("/login", {"uid": username.value, "password": password.value, "twofacode": twoFA.value}).then(response => {
        switch(response.data.code) {
            case 200:
                window.$message.success('登录成功');
                break;
            default:
                errmsg(response.data.code);
        }
        if (response.data.code == 200) {
            window.localStorage.setItem('token', response.data.token)
            router.push('/admin');
        }
    });
}

// 忘记密码按钮
function forGetPwd() {
    showModifyPwdBox.value = true;
}
// 忘记密码信息下一步按钮
function forgetPwdBtn() {
    showModifyPwdBox.value = false;
    axios.post("/resetpwd", {"uid":reuid.value, "email": reemail.value, "twofacode": retwoFA.value}).then(res => {
        if (res.data.code == 200) {
            showModifyPwdvertifyBox.value = true;
            window.localStorage.getItem('token', res.data.token);
            axios.post("/admin/sendemail", {"uid": reuid.value, "email": reemail.value}, {headers: {Authorization: window.localStorage.getItem("token")}}).then(emaires => {
                if (res.data.code == 200) {
                    window.$message.success('邮件发送成功');
                } else {
                    errmsg(res.data.code);
                }
            });
        } else {
            errmsg(res.data.code);
        }
    });
}
// 修改密码信息
// todo
function newPwdBtn() {

}

// 访问login页面时，清空token
function init() {
    window.localStorage.clear('token');
    // 若管理员用户不存在，跳转注册
    axios.get("/adminexist").then(response => {
        if (response.data.code == 200){
            router.push('/registe');
        } else {
            axios.get("/adminvalidate").then(res => {
                if (res.data.code == 200) {
                    router.push('/registe');
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
/*忘记密码*/
.forgetpwd{
    float: right;
}
</style>
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
                    <span class="forgetpwd" ><a href="#" @click="forGetPwd" style="color: gray;">忘记密码?</a></span>
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
    <n-modal
        v-model:show="showModifyPwdBox"
        class="custom-card vertifybox"
        preset="card"
        :style="modifyPwdBox"
        title="修改密码"
        size="huge"
        :bordered="false"
    >
            <div class="inputbox">
                <i class="userid"></i>
                <span class="input-text"> 用户UID</span>
                <n-input v-model:value="reuid" type="text" placeholder="用户UID"/>
            </div>
            <div class="inputbox">
                <i class="email"></i>
                <span class="input-text"> Email</span>
                <n-input v-model:value="reemail" type="text" placeholder="Email"/>
            </div>
            <div class="inputbox">
                <i class="vertify"></i>
                <span class="input-text"> 2FA验证码</span>
                <n-input v-model:value="retwoFA" type="text" placeholder="2FA验证码"/>
            </div>
        <template #footer>
            <div class="btn">
                <n-button :disabled="rePwdBtnDown" @click="forgetPwdBtn" strong round type="primary" size="large">
                    下一步
                </n-button>
            </div>
        </template>
  </n-modal>
  <n-modal
        v-model:show="showModifyPwdvertifyBox"
        class="custom-card vertifybox"
        preset="card"
        :style="modifyPwdBox"
        title="修改密码"
        size="huge"
        :bordered="false"
    >
            <div class="inputbox">
                <i class="userid"></i>
                <span class="input-text"> 用户UID</span>
                <n-input v-model:value="newuid" type="text" placeholder="用户UID"/>
            </div>
            <div class="inputbox">
                <i class="userpwd"></i>
                <span class="input-text"> 密码</span>
                <n-input
                v-model:value="newpwd"
                type="password"
                show-password-on="mousedown"
                placeholder="由数字、大小写字母、标点符号组成"
                :maxlength="20"
                />
            </div>
            <div class="inputbox">
                <i class="email"></i>
                <span class="input-text"> 邮箱验证码</span>
                <n-input v-model:value="emailcode" type="text" placeholder="邮箱验证码"/>
            </div>
        <template #footer>
            <div class="btn">
                <n-button :disabled="newPwdBtnDown" @click="newPwdBtn" strong round type="primary" size="large">
                    修改密码
                </n-button>
            </div>
        </template>
  </n-modal>
</template>