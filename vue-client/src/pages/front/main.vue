<script setup>
import { useRouter } from 'vue-router';
import { useRoute } from 'vue-router';
import axios from 'axios';
import MessageAPI from '../../components/message.vue';
import errmsg from '../../modules/errmsg';
import { useBlogSetStore } from '../../stores/blogset';
import { usePublicUserInfoStore } from '../../stores/userpublicinfo';
import { useNavLocationStore } from '../../stores/navlocation';
import '../../../public/static/css/main.css';
import { inject, onMounted } from 'vue';

const router = useRouter();
const route = useRoute();
const blogset = useBlogSetStore();
const userinfo = usePublicUserInfoStore();
const navloc = useNavLocationStore();

// 加载findme
const findmes = ref([]);
function queryFindMe() {
    axios.get('/findmes').then(res => {
        if (res.data.code === 200) {
            findmes.value = res.data.findme;
        } else {
            errmsg(res.data.code);
        }
    });
}
// findme图标错误
function findmeIconErr(item) {
    item.icon = '/img/nullicon.svg';
}

// 加载导航栏
const navs = ref([]);
function queryNavigation() {
    axios.get('/navigations').then(res => {
        if (res.data.code === 200) {
            navs.value = res.data.navigations;
        } else {
            errmsg(res.data.code);
        }
    });
}
// 导航栏逻辑
const activeItemId = ref('');
function selectNavItem(item) {
    const route = router.resolve({ path: item.navurl }).href;
    const absolutePath = route.replace(/\/article/, '');
    router.push(absolutePath);
    activeItemId.value = item.navurl;
}
function loadNavLocation() {
    activeItemId.value = navloc.navloc;
}

// 搜索框
const showsearchbox = ref(false);
const searchinfo = ref('');
function onPositiveClick() {
    let searchUrl = 'https://www.google.com/search?q=site:' + window.location.hostname + ' ' + encodeURIComponent(searchinfo.value);
    window.open(searchUrl, '_blank');
}
watch([showsearchbox], () => {
   if (!showsearchbox.value) {
        searchinfo.value = '';
   }
});
// 搜索按钮
function search() {
    showsearchbox.value = true
}

onMounted(() => {
    queryFindMe();
    queryNavigation();
    loadNavLocation();
});
</script>
<template>
    <n-back-top :right="100" />
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
                                <span class="usernameinfo">{{ userinfo.data.username }}</span>
                            </div>
                            <div class="search" @click="search">
                                <img src="/img/search.svg" style="height: 25px; width: 25px;"/>
                            </div>
                        </div>
                        <div class="userdescriptionbox">
                            <span class="userdescription">{{ userinfo.data.slogan }}</span>
                        </div>
                        <div class="findmebox">
                            <div class="findmeiconbox" v-for="item in findmes" :key="item.findid">
                                <a :href="item.href" target="blank">
                                    <n-tooltip trigger="hover">
                                        <template #trigger>
                                            <div class="findme">
                                                <img :src="item.icon" @error="findmeIconErr(item)" class="findmeicon"/>
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
            <div class="navigation1">
                <div class="nav1">
                    <div
                    v-for="item in navs"
                    :key="item.navid"
                    style="margin-right: 15px;"
                    :class="{ 'active': activeItemId === item.navurl }"
                    @click="selectNavItem(item)"
                    >
                    {{ item.navtitle }}
                    </div>
                </div>
            </div>
        </div>
        <div class="mainbox">
            <div class="maininfo">
                <router-view />
            </div>
        </div>
    </div>
    <foo />
    <n-modal
        v-model:show="showsearchbox"
        :mask-closable="false"
        preset="dialog"
        title="请输入搜索内容"
        positive-text="确认"
        negative-text="取消"
        @positive-click="onPositiveClick"
        @negative-click="onNegativeClick"
    >
        <n-input v-model:value="searchinfo" type="text" placeholder="请输入搜索内容"/>
    </n-modal>
</template>