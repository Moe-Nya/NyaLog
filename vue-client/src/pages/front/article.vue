<script setup>
import { defineEmits, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useRoute } from 'vue-router';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import client_id from '../../../config.json'

const router = useRouter();
const route = useRoute();

// 文章信息
const id = 'preview-only';
const articleid = ref(-1);
const title = ref('');
const text = ref('');
const articleviews = ref('');
const articlelikes = ref('');
const aisummary = ref('');
const showcommentbox = ref(false);
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
    axios.get('/comment/comuserinfo', { headers: { Authorization: window.localStorage.getItem('usertoken') } }).then(res => {
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
        await navigator.clipboard.writeText('我在NyaLog发布了文章:《' + title.value + '》，快来看看吧~  ' + pageUrl);
        window.$message.success('复制文章地址成功');
    } catch (err) {
        window.$message.error('复制文章地址失败');
    }
}

// 喜欢文章
function likeMe() {
    axios.post('/articlelike', { "articleid": Number(route.params.articleid) }).then(res => {
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
watch([comtext, loginstatus], () => {
    if (comtext.value !== '' && loginstatus.value) {
        sendbtn.value = false;
    } else {
        sendbtn.value = true;
    }
});
function senComment() {
    window.$loadingBar.start();
    axios.post('/comment/newcomment', { "articleid": Number(route.params.articleid), "comment": comtext.value }, { headers: { Authorization: window.localStorage.getItem('usertoken') } }).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('评论发送成功~');
            comments.value = [];
            pagenum.value = 1;
            queryComments();
            comtext.value = '';
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    })
}

// 页面内容加载
function queryArticle() {
    window.$loadingBar.start();
    axios.get(`/article/${route.params.articleid}`).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            articleid.value = res.data.article.articleid;
            title.value = res.data.article.articletitle;
            text.value = res.data.article.text;
            articleviews.value = res.data.article.articleviews;
            articlelikes.value = res.data.article.articlelikes;
            aisummary.value = res.data.article.aisummary;
            if (res.data.comswitch === 0) {
                showcommentbox.value = false;
            } else if (res.data.comswitch === 1) {
                showcommentbox.value = true;
            }
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 评论内容
const pagesize = ref(10);
const pagenum = ref(1);
const comments = ref([]);
const showaddmorebtn = ref(false);

function queryComments() {
    axios.post('/comments', { "articleid": Number(route.params.articleid), "pagesize": pagesize.value, "pagenum": pagenum.value }).then(res => {
        if (res.data.code === 200) {
            comments.value = comments.value.concat(res.data.comments);
            if (res.data.comments.length === 0) {
                showaddmorebtn.value = true;
            } else {
                showaddmorebtn.value = false;
            }
        } else {
            showaddmorebtn.value = true;
        }
    })
}
function addMoreComment() {
    pagenum.value += 1;
    queryComments();
}

// github授权
function github() {
    window.open('https://github.com/login/oauth/authorize?client_id=' + client_id.client_id + '&redirect_uri=' + client_id.domain + '/callback' + '?article=' + route.params.articleid, "_self");
}

// 路由定位
function loadRouter() {
    const pathname = window.location.pathname;
    let suffix = pathname.substring(1); // 去除路径开头的斜杠
    sessionStorage.setItem('nav', suffix);
}

onMounted(() => {
    loadRouter();
    queryArticle();
    loaduser();
    queryComments();
})
</script>
<template>
    <div class="articleandpagetitle">
        <div class="titlestylebox">
            <span class="titlestyle">{{ title }}</span>
        </div>
        <div style="text-align: center; font-size: 16px; font-weight: 500; margin-bottom: 10px;">
            <i class="views"></i>
            <span style="color: #3d3d3d;"> {{ articleviews }} | </span>
            <i class="likes"></i>
            <span style="color: #3d3d3d;">{{ articlelikes }}</span>
        </div>
    </div>
    <div class="articlesummary" v-if="aisummary !== ''">
        <i class="aiicon"></i>
        <span style="font-size: 18px; font-weight: 700; color: #3d3d3d;"> AI文章摘要</span>
        <p style="margin-left: 20px; margin-right: 20px; font-size: 15px; font-weight: 500; color: #3d3d3d;">{{
            aisummary }}</p>
    </div>
    <div class="mdview">
        <MdPreview :editorId="id" :modelValue="text" previewTheme="mdeditor" style="color: #3d3d3d;" />
    </div>
    <div class="headview">
        <MdCatalog :editorId="id" :scrollElement="scrollElement" />
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
    <div class="commentbox" v-show="showcommentbox">
        <img v-show="showuserprofile" class="comuserprofile" :src="userprofile" />&nbsp
        <img v-show="showgithub" @click="github" class="comuserprofile" src="/img/github.svg" />&nbsp
        <span v-show="showusername" class="comusername">{{ username }}</span>
        <n-button v-show="showloginout" @click="loginout" size="small" strong ghost
            style="margin-top: 2px; margin-left: 5px;">
            登出
        </n-button>
    </div>
    <div style="display: block; width: 100%;" v-show="showcommentbox">
        <textarea v-model="comtext" id="combox" class="comments" rows="4" cols="50" maxlength="500"
            placeholder="请输入评论..."></textarea>
        <span class="comtextremain">还可输入{{ 500 - comtext.length }}个字符</span>
        <n-button @click="senComment" :disabled="sendbtn" size="large" strong ghost class="sendcom">
            发送
        </n-button>
        <div class="commentsbox">
            <div v-for="item in comments" :key="item.comid">
                <div class="commentuser">
                    <a :href="item.usersite" target="blank">
                        <img :src="item.profilephoto" class="commentuserprofile" />
                    </a>
                    <div>
                        <a :href="item.usersite" style="color: inherit; text-decoration: none;" target="blank">
                            <span class="commentusername">{{ item.userid }}</span>
                        </a><br>
                        <div class="commentusercomment">
                            <span class="commentusercommentfont">{{ item.commenttext }}</span>
                        </div>
                    </div>
                </div>
                <hr class="dashed-border" />
            </div>
        </div>
        <div class="addmorearticleandcommentbox">
            <n-button strong secondary :disabled="showaddmorebtn" @click="addMoreComment" round size="large">
                点我查看更多~
            </n-button>
        </div>
    </div>
</template>