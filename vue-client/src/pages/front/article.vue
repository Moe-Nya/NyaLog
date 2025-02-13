<script setup>
import { defineEmits, onBeforeMount, onMounted, onUpdated, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useRoute } from 'vue-router';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import client_id from '../../../config.json'
import { config } from 'md-editor-v3';
import ImageFiguresPlugin from '../../modules/markdown-it-image-figures.js';
import formatDate from '../../modules/Time';
import { useMeta } from 'vue-meta';
import { usePublicUserInfoStore } from '../../stores/userpublicinfo'

config({
  markdownItConfig: (mdit) => {
    mdit.use(ImageFiguresPlugin, { figcaption: true });
  },
});

const router = useRouter();
const userinfo = usePublicUserInfoStore();
const route = useRoute();

// 文章信息
const id = 'preview-only';
const articleid = ref(-1);
const title = ref('');
const text = ref('');
const shorttext = ref('');
const articleviews = ref('');
const articlelikes = ref('');
const aisummary = ref('');
const showcommentbox = ref(false);
const pushdate = ref(null);
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
async function loaduser() {
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
            window.localStorage.removeItem('usertoken');
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
    window.localStorage.removeItem('usertoken');
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
function senComment(item) {
    if (item != undefined) {
        if (item.replay === '') {
            window.$message.error('回复内容不能为空~');
            return
        }
        window.$loadingBar.start();
        axios.post('/comment/newcomment', { "articleid": Number(route.params.articleid), "comment": item.replay, "replayid": String(item.comid) }, { headers: { Authorization: window.localStorage.getItem('usertoken') } }).then(res => {
            if (res.data.code === 200) {
                window.$loadingBar.finish();
                window.$message.success('回复发送成功~');
                comments.value = [];
                pagenum.value = 1;
                queryComments();
                item.replay = '';
            } else {
                window.$loadingBar.error();
                errmsg(res.data.code);
            }
        })
        return
    }
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
async function queryArticle() {
    window.$loadingBar.start();
    axios.get(`/article/${route.params.articleid}`).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            articleid.value = res.data.article.articleid;
            title.value = res.data.article.articletitle;
            text.value = res.data.article.text;
            shorttext.value = res.data.article.shorttext;
            articleviews.value = res.data.article.articleviews;
            articlelikes.value = res.data.article.articlelikes;
            aisummary.value = res.data.article.aisummary;
            pushdate.value = res.data.article.CreatedAt;
            if (res.data.comswitch === 0) {
                showcommentbox.value = false;
            } else if (res.data.comswitch === 1) {
                showcommentbox.value = true;
            }
            return title.value;
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
const replaybox = ref(false);
const replayboxid = ref(-1);

async function queryComments() {
    axios.post('/comments', { "articleid": Number(route.params.articleid), "pagesize": pagesize.value, "pagenum": pagenum.value }).then(res => {
        if (res.data.code === 200) {
            res.data.comments.forEach(obj => {
                obj.replay = '';
            });
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
async function addMoreComment() {
    pagenum.value += 1;
    queryComments();
}

function clickReplay(key) {
    if (replayboxid.value === key && replaybox.value === true) {
        replaybox.value = false;
        replayboxid.value = -1;
        return;
    }
    replayboxid.value = key;
    if (!replaybox.value) {
        replaybox.value = true;
    }
}

function likeCom(item) {
    axios.post('/likecomment', { "comid": item.comid }).then(res => {
        if (res.data.code === 200) {
            item.likes = (BigInt(item.likes) + BigInt(1)).toString();
            window.$message.success('点赞成功~');
        } else {
            window.$message.error('点赞失败');
        }
    });
}

// github授权
function github() {
    window.open('https://github.com/login/oauth/authorize?client_id=' + client_id.client_id + '&redirect_uri=' + client_id.domain + '/callback' + '?article=' + route.params.articleid, "_self");
}

onMounted(() => {
    window.scrollTo(0, 0);
    queryArticle();
    loaduser();
    queryComments();
});

onUpdated(() => {
    let theTitle;
    let theShortText;
    while (theTitle = title.value, theShortText = shorttext.value) {
        if (theTitle !== '' && theShortText !== '' && userinfo.data.username !== '一个神秘用户') {
            useMeta({
                title: theTitle + " - " + userinfo.data.username,
                meta: [
                    { name: 'description', content: theShortText }
                ]
            });
            break;
        }
        setTimeout(loop, 200);
    }
});
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
        <div style="text-align: center; font-size: 16px; font-weight: 500; margin-bottom: 10px;">
            <i class="dateico"></i>
            <span style="color: #3d3d3d;">发布于 {{ formatDate(pushdate) }}</span>
        </div>
    </div>
    <div class="articlesummary" v-if="aisummary !== ''">
        <i class="aiicon"></i>
        <span style="font-size: 18px; font-weight: 700; color: #3d3d3d;"> AI文章摘要</span>
        <p style="margin-left: 20px; margin-right: 20px; font-size: 15px; font-weight: 500; color: #3d3d3d;">{{
            aisummary }}</p>
    </div>
    <div class="mdview">
        <MdPreview :editorId="id" :modelValue="text" previewTheme="default" style="color: #3d3d3d;" />
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
            placeholder="我要喵两句..."></textarea><br />
        <span class="comtextremain">还可输入{{ 500 - comtext.length }}个字符</span>
        <n-button @click="senComment(undefined)" :disabled="sendbtn" size="large" strong ghost class="sendcom">
            发送
        </n-button>
        <div class="commentsbox" style="margin-top: 50px;">
            <div v-for="item in comments" :key="item.comid">
                <div class="replaybackground">
                    <span class="replayinfo" :style="{ display: item.recomid == '' ? 'none' : 'block' }">
                        <span v-if="item.recominfo">
                            回复 {{ item.recominfo?.userid }}：
                            {{ item.recominfo?.commenttext }}
                        </span>
                        <span v-else>
                            评论不存在
                        </span>
                    </span>
                </div>
                <div class="commentuser">
                    <a :href="item.usersite" target="blank">
                        <img :src="item.profilephoto" class="commentuserprofile" />
                    </a>
                    <div class="commentsbox">
                        <a :href="item.usersite" style="color: inherit; text-decoration: none;" target="blank">
                            <span class="commentusername">{{ item.userid }}</span>
                        </a><br>
                        <div class="commentusercomment">
                            <span class="commentusercommentfont">{{ item.commenttext }}</span>
                        </div>
                        <div class="commentreply">
                            <span class="commentreplybutton" @click=clickReplay(item.comid)>回复</span>
                            <i class = "commentreplylike" @click="likeCom(item)"></i>
                            <span style="color:#676565cb; font-size: 14px;">{{ item.likes }}</span>
                        </div>
                        <div :style="{ display: replaybox && replayboxid === item.comid ? 'block' : 'none' }">
                            <textarea v-model=item.replay id="combox" class="comments" rows="4" cols="50" maxlength="500"
                                placeholder="我要喵两句..."></textarea><br />
                            <span class="comtextremain">还可输入{{ 500 - String(item.replay || '').length }}个字符</span>
                            <n-button @click="senComment(item)" :disabled="item.replay === '' || !loginstatus" size="large" strong ghost class="sendcom">
                                发送
                            </n-button>
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