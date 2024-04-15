<script setup>
import { defineEmits } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import formatDate from '../../modules/Time';

const router = useRouter();

// 查询归档
const archives = ref([]);
function queryArchive() {
    window.$loadingBar.start();
    axios.get('/archive').then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            archives.value = res.data.archives;
        } else {
            window.$loadingBar.error();
        }
    });
}

// 浏览文章
function archiveClick(aid) {
    router.push(`/article/${aid}`);
}

// 路由定位
function loadRouter() {
    const pathname = window.location.pathname;
    let suffix = pathname.substring(1); // 去除路径开头的斜杠
    sessionStorage.setItem('nav', suffix);
}

onMounted(() => {
    loadRouter();
    queryArchive();
})
</script>
<template>
    <div class="archivesbox">
        <div class="archiveslistbox" v-for="item in archives" :key="item.year">
            <div class="archiveyear">
                <span class="archiveyeartitle">{{ item.year }}</span>
            </div>
            <div class="archivesarticlebox" @click="archiveClick(i.articleid)" v-for="i in item.articles" :key="i.articleid">
                <i class="dot"></i>
                <div class="archivearticletitlewidth">
                    <span class="archivesarticlename">{{ i.articletitle }}</span>
                </div>
                <div class="archivearticledatewidth">
                    <span class="archivesarticledate">{{ formatDate(i.CreatedAt) }}</span>
                </div>
            </div>
        </div>
    </div>
</template>