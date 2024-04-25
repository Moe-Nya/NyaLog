<script setup>
import axios from 'axios'
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import { useRouter } from 'vue-router';
import { onMounted, watch } from 'vue';

const router = useRouter();

// 新增findme
const iconurl = ref('');
const offsiteurl = ref('');
const offsitename = ref('');
const description = ref('');
const previewicon = ref('/img/nullicon.svg');
const addfindmebtn = ref(true);
watch([iconurl, offsiteurl], () => {
    addfindmebtn.value = ![iconurl.value, offsiteurl.value].every(value => value !== '');
});
watch([iconurl], () => {
    if (iconurl.value !== '') {
        previewicon.value = iconurl.value;
    } else {
        previewicon.value = '/img/nullicon.svg';
    }
});
function previewiconerr() {
    previewicon.value = '/img/nullicon.svg';
}
function addFindMe() {
    window.$loadingBar.start();
    axios.post('/admin/newfindme', {"icon": iconurl.value, "href": offsiteurl.value, "text": offsitename.value, "description": description.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).
    then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('新增findme成功');
            iconurl.value = '';
            offsiteurl.value = '';
            offsitename.value = '';
            description.value = '';
            queryFindMe();
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 更新findme
function updateFindMe(item) {
    window.$loadingBar.start();
    axios.post('/admin/modifyfindme', {"findid": item.findid, "icon": item.icon, "href": item.href, "text": item.text, "description": item.description}, {headers: {Authorization: window.localStorage.getItem('token')}}).
    then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('更新成功');
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 删除findme
function deleteFindme(item) {
    window.$loadingBar.start();
    axios.post('/admin/deletefindme', {"findid": item.findid}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('删除成功');
            queryFindMe();
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

// 查询findme
const findmes = ref([]);
function queryFindMe() {
    axios.get('/findmes').then(res => {
        if (res.data.code === 200) {
            findmes.value = res.data.findme;
            for (let i = 0; i < findmes.value.length; i++) {
                if (typeof findmes.value[i] === 'object' && findmes.value[i] !== null) {
                    findmes.value[i].tmp = findmes.value[i].icon;
                }
            }
        } else {
            errmsg(res.data.code);
        }
    });
}
function previewIconEditErr(item) {
    item.icon = '/img/nullicon.svg';
}
function updateSrc(item) {
    item.icon = item.tmp;
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
    queryFindMe();
});
</script>
<template>
    <div class="infotitlebox">
        <h2>Findme管理</h2>
        <p style="color:gray; font-size: 14px;">在这里可以设置一些站外链接，让大家可以在其它地方找到你哦！</p>
        <p style="color:gray; font-size: 14px;">tips:在这里可以输入站外链接的icon、链接、站点名字和描述，需要自己去找一个icon哦~请注意icon的质量~</p>
    </div>
    <div class="infobox">
        <div class="findmebox">
            <div>
                <i class="url"></i>
                <span class="input-text"> iconUrl</span><br />
                <n-input v-model:value="iconurl" type="text" placeholder="iconUrl" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div>
                <i class="url"></i>
                <span class="input-text"> 外链Url</span><br />
                <n-input v-model:value="offsiteurl" type="text" placeholder="外链Url tip:请加上http请求头" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div>
                <i class="platformname"></i>
                <span class="input-text"> 外链名称</span><br />
                <n-input v-model:value="offsitename" type="text" placeholder="外链名称" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div>
                <i class="slogan"></i>
                <span class="input-text"> 描述</span><br />
                <n-input v-model:value="description" type="text" placeholder="描述" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div style="margin-top: 26px; display: flex;">
                <n-tooltip trigger="hover">
                    <template #trigger>
                        <div class="previewiconbox">
                            <a :href="offsiteurl" target="blank">
                                <img :src="previewicon" @error="previewiconerr" class="previewicon"/>
                            </a>
                        </div>
                    </template>
                        {{ description }}
                </n-tooltip>
                <n-button :disabled="addfindmebtn" @click="addFindMe" size="medium" strong round type="primary">新增</n-button>
            </div>
        </div>
        <hr class="categoryhr" /><br />
        <div class="findmebox" v-for="item in findmes" :key="item.findid">
            <div>
                <i class="url"></i>
                <span class="input-text"> iconUrl</span><br />
                <n-input v-model:value="item.tmp" @blur="updateSrc(item)" type="text" placeholder="iconUrl" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div>
                <i class="url"></i>
                <span class="input-text"> 外链Url</span><br />
                <n-input v-model:value="item.href" type="text" placeholder="外链Url" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div>
                <i class="platformname"></i>
                <span class="input-text"> 外链名称</span><br />
                <n-input v-model:value="item.text" type="text" placeholder="外链名称" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div>
                <i class="slogan"></i>
                <span class="input-text"> 描述</span><br />
                <n-input v-model:value="item.description" type="text" placeholder="描述" style="min-width:230px; width: 230px; border-radius: 5px; margin-right: 10px; margin-bottom: 8px;"/>
            </div>
            <div style="margin-top: 27px; display: flex;">
                <n-tooltip trigger="hover">
                    <template #trigger>
                        <div class="previewiconbox">
                            <a :href="item.href" target="blank">
                                <img :src="item.icon" @error="previewIconEditErr(item)" class="previewicon"/>
                            </a>
                        </div>
                    </template>
                        {{ item.description }}
                </n-tooltip>
                <n-button size="medium" :disabled="!(item.tmp !== '' && item.href !== '')" @click="updateFindMe(item)" style="margin-right: 5px;" strong round type="primary">更新</n-button>
                <n-button size="medium" @click="deleteFindme(item)" strong round type="error">删除</n-button>
            </div>
        </div>
    </div>
</template>