<script setup>
import { onMounted } from 'vue';
import axios from 'axios';
import { useNavLocationStore } from '../stores/navlocation'
import MessageAPI from '../components/message.vue'
import errmsg from '../modules/errmsg';
import { useRouter } from 'vue-router';
import { useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();

const navloc = useNavLocationStore();

// 路由定位
function loadRouter() {
    let pathname = window.location.pathname;
    let suffix = pathname.substring(1); // 去除路径开头的斜杠
    navloc.navloc = suffix;
}
// 确认code是否存在并且执行后续操作
function checkCode() {
    let code = route.query.code;
    if (code === undefined) {
        router.push(`/article/${route.query.article}?position=0`);
    }
    axios.get('/githuboauthcode', {params: {"code": code}}).then(res => {
        if (res.data.code === 200) {
            window.localStorage.setItem('usertoken', res.data.token);
            router.push(`/article/${route.query.article}?position=1`);
        } else {
            router.push(`/article/${route.query.article}?position=0`);
        }
    })
}

onMounted(() => {
    loadRouter();
    checkCode();
})
</script>
<template>
    <div class="callbackbox">
        <img src="../../public/img/oauth.svg" style="height: 30%; width: 30%;"/>
        <span style="font-size: 20px; font-weight: bold;">三方登录授权中...</span>
    </div>
</template>