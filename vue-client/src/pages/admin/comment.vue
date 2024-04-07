<script setup>
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import formatDate from '../../modules/Time';
import { onMounted } from 'vue';

const router = useRouter();

// 评论页面信息
const pagesize = ref(6);
const pagenum = ref(1);
const totalcomments = ref(0);
// 评论信息
const comments = ref([]);
// 删除按钮逻辑
const showDeleteComConfirm = ref(false);
function deleCommentBtn() {
    showDeleteComConfirm.value = true;
}
function onNegativeClick() {
    showDeleteComConfirm.value = false;
}
function onPositiveClick(value) {
    window.$loadingBar.start();
    axios.post('/admin/deletecomments', {"comid": value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('评论删除成功');
            let totalcoms = Math.ceil((totalcomments.value - 1) / pagesize.value);
            if (pagenum.value > totalcoms) {
                pagenum.value = totalcoms;
            }
            queryComment();
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 加载评论
function queryComment() {
    window.$loadingBar.start();
    axios.post('/admin/allcomments', {"pagesize": pagesize.value, "pagenum": pagenum.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).
    then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            totalcomments.value = res.data.total;
            comments.value = res.data.comments;
        }
    });
}

// 加载更多评论
const previousPage = ref(true);
const nextPage = ref(false);
watch([pagenum, totalcomments, comments], () => {
    if (pagenum.value <= 1) {
        previousPage.value = true;
    } else {
        previousPage.value = false;
    }
    if (pagenum.value >= Math.ceil(totalcomments.value / pagesize.value)) {
        nextPage.value = true;
    } else {
        nextPage.value = false;
    }
})

function previousPageBtn() {
    pagenum.value -= 1;
    queryComment();
}
function nextPageBtn() {
    pagenum.value += 1;
    queryComment();
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

onMounted(() => {
    Validate();
    queryComment();
});
</script>
<template>
    <div class="infotitlebox">
        <h2>评论管理</h2>
        <p style="color:gray; font-size: 14px;">来了解一下访问者们都说了写什么吧~ 不同的观点也许会带来一些思想上的启发，请谨慎使用删除按钮！</p>
    </div>
    <div class="infobox">
        <div class="commentbox" v-for="item in comments" :key="item.comid">
            <div class="commentboxinfo">
                <div style="display: flex; align-items: center;">
                    <img :src="item.profilephoto" style="height:40px; width: 40px; border-radius: 50%;">
                    <a :href="item.usersite" target="_blank" style="color: black; text-decoration: none;"><span style="margin-left: 10px; margin-bottom: 10px; font-weight: bold; font-size: 16px;" class="commentinfo">{{ item.userid }}</span></a>
                    <a :href="'/article/'+item.articleid" target="_blank" style="margin-left: 10px; margin-bottom: 8px; color: gray; font-weight: bold; font-size: 14px;" class="commentinfo">来自文章id: {{ item.articleid }}</a>
                    <span style="margin-left: 10px; margin-bottom: 8px; color: gray; font-weight: bold; font-size: 14px;" class="commentinfo">发布于: {{ formatDate(item.CreatedAt) }}</span>
                </div>
                <div style="margin-left: 50px; margin-top: -10px;">
                    <span class="articleshort">{{ item.commenttext }}</span>
                </div>
            </div>
            <div class="deletecomment">
                <n-button @click="deleCommentBtn" strong round type="error" size="medium">
                    删除
                    <n-modal
                    v-model:show="showDeleteComConfirm"
                    :mask-closable="false"
                    preset="dialog"
                    type="error"
                    title="确认删除该评论吗？"
                    content="删除后不可恢复！"
                    positive-text="确认"
                    negative-text="取消"
                    @positive-click="onPositiveClick(item.comid)"
                    @negative-click="onNegativeClick"
                    />
                </n-button>
            </div>
        </div>
    </div>
    <div class="addmorebtnbox">
        <n-button strong secondary :disabled="previousPage" @click="previousPageBtn" round size="large" class="addmorebtn">
            <
        </n-button>&nbsp&nbsp&nbsp
        <n-button strong secondary :disabled="nextPage" @click="nextPageBtn" round size="large" class="addmorebtn">
            >
        </n-button>
    </div>
</template>