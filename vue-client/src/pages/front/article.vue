<script setup>
import { defineEmits, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useRoute } from 'vue-router';
import { useNavLocationStore } from '../../stores/navlocation'
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import client_id from '../../../config.json'

const navloc = useNavLocationStore();
const router = useRouter();
const route = useRoute();
// 文章信息
const id = 'preview-only';
const articleid = ref();
const title = ref('');
const text = ref('');
const articleviews = ref('');
const articlelikes = ref('');
const aisummary = ref('');
const scrollElement = document.documentElement;

// 评论用户信息
const username = ref('');
const userprofile = ref('');
const usersite = ref('');

// 用户状态
const showuserprofile = ref(false);
const showgithub = ref(true);
const showusername = ref(false)
const showloginout = ref(false);

let pageUrl = window.location.href;

// 加载用户信息
function loaduser() {
    axios.get('/comment/comuserinfo', {headers: {Authorization: window.localStorage.getItem('usertoken')}}).then(res => {
        if (res.data.code === 200) {
            username.value = res.data.userinfo.username;
            userprofile.value = res.data.userinfo.userprofile;
            usersite.value = res.data.userinfo.usersite;
            showuserprofile.value = true;
            showusername.value = true;
            showgithub.value = false;
            showloginout.value = true;
            loginstatus.value = true;
        } else {
            window.localStorage.clear('usertoken');
            showuserprofile.value = false;
            showgithub.value = true;
            showusername.value = false;
            showloginout.value = false;
            loginstatus.value = false;
        }
        if (route.query.position === '1') {
            window.location.href = '#combox';
        }
    });
}

// 登出用户
function loginout() {
    window.localStorage.clear('usertoken');
    loaduser();
}

// 分享文章
async function shareArticle() {
    try {
        await navigator.clipboard.writeText('我在NyaLog发布了文章:《'+ title.value + '》，快来看看吧~  ' + pageUrl);
        window.$message.success('复制文章地址成功');
    } catch (err) {
        window.$message.error('复制文章地址失败');
    }
}

// 喜欢文章
function likeMe() {
    axios.post('/articlelike', {"articleid": route.params.articleid}).then(res => {
        if (res.data.code === 200) {
            window.$message.success('=W=');
        } else {
            window.$message.error('点赞失败');
        }
    });
}

// 评论逻辑
const comtext = ref('');
const loginstatus = ref(false);
const sendbtn = ref(true);
watch([comtext,loginstatus], () => {
    if (comtext.value !== '' && loginstatus.value) {
        sendbtn.value = false;
    } else {
        sendbtn.value = true;
    }
});
function senComment() {
    
}


// 页面内容加载
function queryArticle() {
    window.$loadingBar.start();
    axios.get(`${navloc.navloc}`).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            articleid.value = res.data.article.articleid;
            title.value = res.data.article.articletitle;
            text.value = res.data.article.text;
            articleviews.value = res.data.article.articleviews;
            articlelikes.value = res.data.article.articlelikes;
            aisummary.value = res.data.article.aisummary;
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// github授权
function github() {
    window.open('https://github.com/login/oauth/authorize?client_id=' + client_id.client_id +'&redirect_uri=http://192.168.31.145:5173/callback' + '?article=' + articleid.value, "_self");
}

// 路由定位
function loadRouter() {
    const pathname = window.location.pathname;
    let suffix = pathname.substring(1); // 去除路径开头的斜杠
    navloc.navloc = suffix;
}

onMounted(() => {
    loadRouter();
    queryArticle();
    loaduser();
})
</script>
<template>
        <div class="articleandpagetitle">
            <div>
                <span class="titlestyle">{{ title }}</span>
            </div>
            <div style="text-align: center; font-size: 16px; font-weight: 500; margin-bottom: 10px;">
                <i class="views"></i>
                <span style="color: #3d3d3d;"> {{ articleviews }} | </span>
                <i class="likes"></i>
                <span style="color: #3d3d3d;">likes: {{ articlelikes }}</span>
            </div>
        </div>
        <div class="articlesummary" v-if="aisummary !== ''">
            <i class="aiicon"></i>
            <span style="font-size: 18px; font-weight: 700; color: #3d3d3d;"> AI文章摘要</span>
            <p style="margin-left: 20px; margin-right: 20px; font-size: 15px; font-weight: 500; color: #3d3d3d;">{{ aisummary }}</p>
        </div>
        <div class="mdview">
            <MdPreview :editorId="id" :modelValue="text" previewTheme="mdeditor" style="color: #3d3d3d;"/>
        </div>
        <div class="headview">
            <MdCatalog :editorId="id" :scrollElement="scrollElement"/>
        </div>
        <div class="articleinfoicon">
            <div class="likearticlebtnbox" @click="likeMe" style="margin-right: 10px;">
                <i class="likearticlebtn"></i>
            </div>
            <div class="likearticlebtnbox" @click="shareArticle">
                <i class="share"></i>
            </div>
        </div>
        <hr class="hr" />
        <div class="commentbox">
            <img v-show="showuserprofile" class="comuserprofile" :src="userprofile" />&nbsp
            <img v-show="showgithub" @click="github" class="comuserprofile" src="../../../public/img/github.svg" />&nbsp
            <span v-show="showusername" class="comusername">{{ username }}</span>
            <n-button v-show="showloginout" @click="loginout" size="small" strong ghost style="margin-top: 2px; margin-left: 5px;">
            登出
            </n-button>
        </div>
        <div style="width: 100%;">
        <textarea v-model="comtext" id="combox" class="comments" rows="4" cols="50" maxlength="500" placeholder="请输入评论..."></textarea>
        <span class="comtextremain">还可输入{{ 500 - comtext.length }}个字符</span>
        <n-button @click="senComment" :disabled="sendbtn" size="large" strong ghost class="sendcom">
            发送
        </n-button>
    </div>
</template>