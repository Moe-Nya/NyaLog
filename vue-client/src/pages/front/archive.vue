<script setup>
import { defineEmits } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import formatDate from '../../modules/Time';
import { useMeta } from 'vue-meta';
import { usePublicUserInfoStore } from '../../stores/userpublicinfo'

const router = useRouter();
const userinfo = usePublicUserInfoStore();

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

onMounted(() => {
    queryArchive();
})

onUpdated(() => {
    useMeta({
        title: '归档' + " - " + userinfo.data.username,
    });
});
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