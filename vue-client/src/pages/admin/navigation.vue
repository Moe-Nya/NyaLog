<script setup>
import axios from 'axios'
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import { useRouter } from 'vue-router';
import { onMounted, watch } from 'vue';

const router = useRouter();

// 导航栏信息
const navname = ref('');
const navurl = ref('');

// 新增导航栏
function addNavigation() {
    window.$loadingBar.start();
    axios.post('/admin/newnavigation', {"navtitle": navname.value, "navurl": navurl.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            windwo.$message.success('新增导航栏成功');
        } else {
            window.$loadingBar.finish();
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
});
</script>
<template>
    <div class="infotitlebox">
        <h2>导航栏管理</h2>
        <p style="color:gray; font-size: 14px;">在这里可以可以对导航栏进行设置，访问者可以通过导航栏到达其它页面。</p>
    </div>
    <div class="infobox">
        <div class="findmebox">
            <div>
                <i class="platformname"></i>
                <span class="input-text"> 导航栏名称</span><br />
                <n-input v-model:value="navname" type="text" placeholder="导航栏名称" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div>
                <i class="url"></i>
                <span class="input-text"> 导航栏Url</span><br />
                <n-input v-model:value="navurl" type="text" placeholder="导航栏Url tip:请加上http请求头" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div style="margin-top: 26px; display: flex;">
                <n-button :disabled="!(navname !== '' && navurl !== '')" @click="addNavigation" size="medium" strong round type="primary">新增</n-button>
            </div>
        </div>
        <hr class="categoryhr" /><br />
        <div class="findmebox">
            <div>
                <i class="platformname"></i>
                <span class="input-text"> 导航栏名称</span><br />
                <n-input  type="text" placeholder="导航栏名称" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div>
                <i class="url"></i>
                <span class="input-text"> 导航栏Url</span><br />
                <n-input type="text" placeholder="导航栏Url tip:请加上http请求头" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div style="margin-top: 26px; display: flex;">
                <n-button size="medium" style="margin-right: 5px;" strong round type="primary">新增</n-button>
                <n-button size="medium" strong round type="error">删除</n-button>
            </div>
        </div>
    </div>
</template>