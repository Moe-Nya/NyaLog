<script setup>
import axios from 'axios'
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import { useRouter } from 'vue-router';
import { onMounted } from 'vue';

const router = useRouter();

// 分类信息
const createcategoryinput = ref('');
const categorylist = ref([]);

// 新增分类
function createcategory() {
    window.$loadingBar.start();
    axios.post('/admin/newcategory', {"categorytext": createcategoryinput.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('新增分类成功');
            queryCategory();
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 请求分类列表
function queryCategory() {
    axios.get('/category').then(res => {
        if (res.data.code === 200) {
            categorylist.value = res.data.category;
        } else {
            errmsg(res.data.code);
        }
    });
}

// 分类操作
function editCategory(cid, categorytext) {
    window.$loadingBar.start();
    axios.post('admin/editcategory', {"cid": cid, "categorytext": categorytext}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('更改分类成功');
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
            categorylist.value = [];
        }
    });
}
function deleteCategory(cid) {
    window.$loadingBar.start();
    axios.post('/admin/deletecategory', {"cid": cid}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('删除分类成功');
            queryCategory();
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
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
    queryCategory();
});
</script>
<template>
    <div class="userinfotitlebox">
        <h2>文章分类管理</h2>
        <p style="color:gray; font-size: 14px;">这里可以对文章分类进行设置更新，给文章分类，有助于让读者有目的地选择性阅读~</p>
    </div>
    <div class="infobox">
        <div class="inputbox" style="display: flex; flex-wrap: wrap;">
            <div style="margin-right: 10px; margin-top: 8px;">
                <i class="articlecategory"></i>
                <span class="input-text"> 文章分类</span>
                <n-input v-model:value="createcategoryinput" type="text" placeholder="文章分类"/>
            </div>
            <div class="deletecomment">
                <n-button :disabled="!(createcategoryinput !== '')" @click="createcategory" strong round type="primary" size="medium">
                    新增
                </n-button>
            </div>
        </div>
        <hr class="categoryhr" /><br />
        <div class="inputbox" style="display: flex; flex-wrap: wrap;" v-for="item in categorylist" :key="item.cid">
            <div style="margin-right: 10px; margin-top: 8px;">
                <i class="articlecategory"></i>
                <span class="input-text"> 文章分类 cid:{{ item.cid }}</span>
                <n-input v-model:value="item.categorytext" type="text" placeholder="文章分类"/>
            </div>
            <div class="deletecomment">
                <n-button :disabled="!(item.categorytext !== '')" @click="editCategory(item.cid ,item.categorytext)" strong round type="primary" size="medium" style="margin-right: 10px; float: left">
                    更改
                </n-button>
                <n-button @click="deleteCategory(item.cid)" strong round type="error" size="medium" style="float: left;">
                    删除
                </n-button>
            </div>
        </div>
    </div>
    
</template>