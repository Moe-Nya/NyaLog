<script setup>
import axios from 'axios';
import { onMounted } from 'vue';
import { useMeta } from 'vue-meta';
import { usePublicUserInfoStore } from '../../stores/userpublicinfo'

const userinfo = usePublicUserInfoStore();

// 页面信息
const pagesize = ref(12);
const pagenum = ref(1);
const totalfriendlink = ref(0);
const remainfriendlink = ref(0);
// 文章信息
const friendlinks = ref([]);
const errimg = ref('/img/userprofile.png');

// 加载友链
function queryFriendLink() {
    window.$loadingBar.start();
    axios.post('/friendlinks', {"pagesize": pagesize.value, "pagenum": pagenum.value}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            totalfriendlink.value = res.data.total;
            friendlinks.value = friendlinks.value.concat(res.data.friendlinks);
            remainfriendlink.value = totalfriendlink.value - friendlinks.value.length;
        } else {
            window.$loadingBar.error();
        }
    });
}
function friendlinkImgError(item) {
    item.friendprofile = errimg.value;
}

// 加载更多文章
function addMoreFriendlinkBtn() {
    pagenum.value += 1;
    queryFriendLink();
}

onMounted(() => {
    queryFriendLink();
})


onUpdated(() => {
    useMeta({
        title: '友情链接' + " - " + userinfo.data.username,
    });
});
</script>
<template>
    <div class="articlesbox">
        <div class="friendlinksbox" v-for="item in friendlinks" :key="item.friendlinkid">
            <a :href="item.friendsite" style="text-decoration: none;" target="_blank">
                <div class="friendlinksprofilebox">
                    <img class="friendlinksprofile" :src="item.friendprofile" @error="friendlinkImgError(item)" />
                </div>
                <div class="sitesnamebox">
                    <div style="margin-left: 10px; margin-right: 10px;">
                        <span class="sitesname">{{ item.friendname }}</span>
                    </div>
                </div>
                <div class="sitesdescriptionbox">
                    <div style="margin-left: 10px; margin-right: 10px;">
                        <span
                            class="sitesdescription">{{ item.frienddescription }}</span>
                    </div>
                </div>
            </a>
        </div>
    </div>
    <div class="addmorearticleandcommentbox" style="margin-bottom: 40px;">
        <n-button strong secondary :disabled="(Number(remainfriendlink) === 0)" @click="addMoreFriendlinkBtn" round size="large">
            加载更多！共{{ totalfriendlink }}个友链 剩余{{ remainfriendlink }}个
        </n-button>
    </div>
</template>