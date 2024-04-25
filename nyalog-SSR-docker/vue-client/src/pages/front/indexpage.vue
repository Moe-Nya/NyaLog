<script setup>
import { defineEmits } from 'vue';
import formatDate from '../../modules/Time';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { onMounted } from 'vue';

const router = useRouter();

// 页面信息
const pagesize = ref(9);
const pagenum = ref(1);
const totalarticle = ref(0);
const remainarticle = ref(0);
// 文章信息
const articles = ref([]);
const errimg = ref('/img/defaultarticleimg.svg');

// 请求文章
function queryArticle() {
    window.$loadingBar.start();
    axios.post("/articlelist", {"pagesize": pagesize.value, "pagenum": pagenum.value}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            totalarticle.value = res.data.total;
            articles.value = articles.value.concat(res.data.articles);
            remainarticle.value = totalarticle.value - articles.value.length;
            for (let i = 0; i < articles.value.length; i++) {
                let cname = '';
                axios.post('/categoryname', {"cid": articles.value[i].cid}).then(res => {
                    if (res.data.code === 200) {
                        cname = res.data.ctext;
                    }
                    articles.value[i].cidname = cname;
                });
            }
        } else {
            window.$loadingBar.error();
        }
    });
}
function articleImgError(item) {
    item.articleimg = errimg.value;
}
// 跳转文章浏览
function articleClick(item) {
    router.push(`/article/${item.articleid}`);
}

// 加载更多文章
function addMoreArticleBtn() {
    pagenum.value += 1;
    queryArticle();
}

onMounted(() => {
    queryArticle();
})
</script>
<template>
    <div class="articlesbox">
        <div @click="articleClick(item)" class="articles" v-for="item in articles" :key="item.articleid">
            <div class="articlespicbox">
                <img class="articlespic" :src="item.articleimg" @error="articleImgError(item)"/>
            </div>
            <div class="articlestitlebox">
                <div style="margin-left: 15px; margin-right: 12px; height: 115px;">
                    <span class="articlestitle">{{ item.articletitle }}</span>
                    <span class="articlesshorts">{{ item.shorttext }}</span>
                </div>
                <div class="articlesicons">
                    <div style="height: 25px; display: flex;">
                        <i v-show="(item.cidname !== '')" class="articlescategorysicon"></i>
                        <span v-show="(item.cidname !== '')" class="articlescategorys">{{ item.cidname }} |</span>
                        <i class="articlesviews"></i>
                        <span class="articlescategorys">{{ item.articleviews }}</span>
                        <i class="articleslikes"></i>
                        <span class="articlescategorys">{{ item.articlelikes }}</span>
                    </div>
                </div>
                <div class="articlestimesbox">
                    <div style="margin-left: 15px;">
                        <span class="articlestimes">{{ formatDate(item.CreatedAt) }}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="addmorearticleandcommentbox" style="margin-bottom: 40px;">
        <n-button strong secondary :disabled="(Number(remainarticle) === 0)" @click="addMoreArticleBtn" round size="large">
            加载更多！共{{ totalarticle }}篇文章 剩余{{ remainarticle }}篇
        </n-button>
    </div>
</template>