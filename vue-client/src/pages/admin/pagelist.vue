<script setup>
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import formatDate from '../../modules/Time';

const router = useRouter();

// 页面信息
const pageinfo = ref([]);

// 分页信息
const pagesize = ref(6);
const pagenum = ref(1);
const totalpages = ref(0);

// 添加页面
function addPage() {
    router.push({name: '页面编辑'});
}

// 请求页面列表
function queryPage() {
    window.$loadingBar.start();
    axios.post('/admin/pages',{"pagesize": pagesize.value, "pagenum": pagenum.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            totalpages.value = res.data.total;
            pageinfo.value = res.data.pages;
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 加载更多页面
const previousPage = ref(true);
const nextPage = ref(false);
watch([pagenum, totalpages, pageinfo], () => {
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
    queryPage();
}
function nextPageBtn() {
    pagenum.value += 1;
    queryPage();
}

// 删除页面
const delepid = ref();
const showConfirm = ref(false);
function delePageBtn(pid) {
    delepid.value = pid;
    showConfirm.value = true;
}
function onPositiveClick() {
    window.$loadingBar.start();
    axios.post('/admin/deletepage', {"pid": delepid.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('页面删除成功');
            let totalpage = Math.ceil((totalpages.value - 1) / pagesize.value);
            if (pagenum.value > totalpage) {
                pagenum.value = totalpage;
            }
            queryPage();
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}
function onNegativeClick() {
    showConfirm.value = false;
}

// 编辑页面
function pageEditor(pid) {
    router.push({name: '页面编辑', query: {id: pid}});
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
// 初始化
onMounted(() => {
    Validate();
    queryPage();
});
</script>
<template>
    <div class="infotitlebox">
        <h2>页面管理</h2>
        <n-button @click="addPage" strong round type="info" size="large">+新增页面</n-button>
        <p style="color:gray; font-size: 14px;">页面是网站的公开内容，你可以用markdown做出各种好看的页面，然后使用导航栏系统将页面展示在主页。</p>
        <p style="color:gray; font-size: 14px;">在这里可以管理页面内容，包括新增、编辑、删除。</p>
    </div>
    <div class="infobox">
        <div class="articlelistbox" style="height:90px" v-for="item in pageinfo" :key="item.pid">
            <div class="articlelistboxinfo">
                <div class="articleinfo" @click="pageEditor(item.purl)">
                        <span class="articletitle commentinfo">{{ item.ptitle }}</span>
                    <div>
                        <span class="articletime commentinfo">创建时间:{{ formatDate(item.CreatedAt) }}</span>
                    </div>
                </div>
                <div class="deletebtn" style="margin-top: 15px;">
                    <n-button @click="delePageBtn(item.pid)" strong round type="error" size="large">
                        删除
                    </n-button>
                </div>
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
    <n-modal
        v-model:show="showConfirm"
        v-model:value="delepid"
        :mask-closable="false"
        preset="dialog"
        type="error"
        title="确认删除该页面吗？"
        content="删除后不可恢复！"
        positive-text="确认"
        negative-text="取消"
        @positive-click="onPositiveClick"
        @negative-click="onNegativeClick"
        />
</template>