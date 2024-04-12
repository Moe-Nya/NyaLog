<script setup>
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import formatDate from '../../modules/Time';

const router = useRouter();

// 分页信息
const pagesize = ref(4);
const pagenum = ref(1);
const totalpages = ref(0);
const remainpages = ref(0);
// 文章信息
const articles = ref([]);
const errimg = '../../../public/img/defaultarticleimg.svg';

// 请求文章
function queryArticle() {
    window.$loadingBar.start();
    axios.post("/articlelist", {"pagesize": pagesize.value, "pagenum": pagenum.value}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            totalpages.value = res.data.total;
            articles.value = res.data.articles;
            remainpages.value = totalpages.value - articles.value.length;
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 文章图片错误
function articleImgError(item) {
    item.articleimg = errimg;
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
// 删除文章
const delearticleid = ref();
const showConfirm = ref(false);
function deleArticleBtn(articleid) {
    delearticleid.value = articleid;
    showConfirm.value = true;
}
function onNegativeClick() {
    showConfirm.value = false;
}
function onPositiveClick() {
    window.$loadingBar.start();
    axios.post("/admin/deletearticle", {"articleid": delearticleid.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('文章删除成功');
            let totalpage = Math.ceil((totalpages.value - 1) / pagesize.value);
            if (pagenum.value > totalpage) {
                pagenum.value = totalpage;
            }
            queryArticle();
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    })
}

// 新增文章
function addArticle() {
    router.push({name: '文章编辑'});
}

// 加载更多文章
const previousPage = ref(true);
const nextPage = ref(false);
watch([pagenum, totalpages, articles], () => {
    if (pagenum.value <= 1) {
        previousPage.value = true;
    } else {
        previousPage.value = false;
    }
    if (pagenum.value >= Math.ceil(totalpages.value / pagesize.value)) {
        nextPage.value = true;
    } else {
        nextPage.value = false;
    }
})

function previousPageBtn() {
    pagenum.value -= 1;
    queryArticle();
}
function nextPageBtn() {
    pagenum.value += 1;
    queryArticle();
}

// 编辑文章
function articleEditor(articleid) {
    router.push({name: '文章编辑', query: {id: articleid}});
}

// 初始化
onMounted(() => {
    Validate();
    queryArticle();
});
</script>
<template>
    <div class="infotitlebox">
        <h2>文章管理</h2>
        <n-button @click="addArticle" strong round type="info" size="large">+新增文章</n-button>
        <p style="color:gray; font-size: 14px;">文章是网站的公开内容，也是博客系统最重要的部分，漂亮的样式可没用哦~更重要的是文章的内容质量！</p>
        <p style="color:gray; font-size: 14px;">在这里可以管理文章内容，包括新增、编辑、删除。</p>
    </div>
    <div class="infobox">
        <div class="articlelistbox" v-for="item in articles" :key="item.articleid">
            <div class="articlelistboxinfo">
                    <div class="piccontainer" @click="articleEditor(item.articleid)">
                        <img class="pic" :src="item.articleimg" @error="articleImgError(item)" alt="Image">
                    </div>
                <div class="articleinfo" @click="articleEditor(item.articleid)">
                        <span class="articletitle commentinfo">{{ item.articletitle }}</span>
                        <div class="articleshortcontainer">
                            <p class="articleshort">{{ item.shorttext }}</p><br />
                        </div>
                    <div>
                        <span class="articletime commentinfo">创建时间{{ formatDate(item.CreatedAt) }}</span>
                    </div>
                </div>
                <div class="deletebtn">
                    <n-button @click="deleArticleBtn(item.articleid)" strong round type="error" size="large">
                        删除
                    </n-button>
                </div>
            </div>
        </div>
    </div>
    <div class="addmorebtnbox">
        <n-button strong secondary :disabled="previousPage" @click="previousPageBtn" round size="large" class="addmorebtn" style="background-color: rgb(186, 193, 197);">
            <
        </n-button>&nbsp&nbsp&nbsp
        <n-button strong secondary :disabled="nextPage" @click="nextPageBtn" round size="large" class="addmorebtn" style="background-color: rgb(186, 193, 197);">
            >
        </n-button>
    </div>
    <n-modal
        v-model:show="showConfirm"
        :mask-closable="false"
        preset="dialog"
        type="error"
        title="确认删除该文章吗？"
        content="删除后不可恢复！"
        positive-text="确认"
        negative-text="取消"
        @positive-click="onPositiveClick"
        @negative-click="onNegativeClick"
        />
</template>