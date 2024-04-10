<script setup>
import { useRouter } from 'vue-router';
import axios from 'axios'
import MessageAPI from '../../components/message.vue'
import errmsg from '../../modules/errmsg';
import { useBlogSetStore } from '../../stores/blogset'
import { usePublicUserInfoStore } from '../../stores/userpublicinfo'
import '../../../public/static/css/main.css'
import { onMounted } from 'vue';

const router = useRouter();
const blogset = useBlogSetStore();
const userinfo = usePublicUserInfoStore();

// 加载findme
const findmes = ref([]);
function queryFindMe() {
    axios.get('/findmes').then(res => {
        if (res.data.code === 200) {
            findmes.value = res.data.findme;
        } else {
            window.$message.success('网络连接似乎不顺畅哦！');
        }
    });
}

onMounted(() => {
    queryFindMe();
});
</script>
<template>
    <div class="box">
        <div class="headerbox">
            <div class="headerimg">
                <img :src="blogset.data.sitebackground" class="headerimgstyle" />
            </div>
            <div class="header">
                <div class="headerinfo">
                    <div class="userprofile center">
                        <img :src="userinfo.data.profilephoto" class="userprofileimgstyle" />
                    </div>
                    <div class="userinfo">
                        <div class="usernamebox">
                            <div>
                                <span class="username">{{ userinfo.data.username }}</span>
                            </div>
                            <div class="search">
                                <img src="../../../public/img/search.svg" style="height: 25px; width: 25px;"/>
                            </div>
                        </div>
                        <div class="userdescriptionbox">
                            <span class="userdescription">{{ userinfo.data.slogan }}</span>
                        </div>
                        <div class="findmebox">
                            <!-- 这里的图片需要替换动态哦！ -->
                            <div class="findmeiconbox" v-for="item in findmes" :key="item.findid">
                                <a :href="item.href" target="blank">
                                    <n-tooltip trigger="hover">
                                        <template #trigger>
                                            <div class="findme">
                                                <img :src="item.icon" class="findmeicon"/>
                                            </div>
                                        </template>
                                            {{ item.description }}
                                    </n-tooltip>
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="navigation">
                这是导航栏
            </div>
        </div>
        <div class="mainbox">
            <p>这是main</p>
        </div>
    </div>
    <foo />
</template>