<script setup>
import axios from 'axios'
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import { useRouter } from 'vue-router';
import { onMounted, watch } from 'vue';

const router = useRouter();

// 评论页面信息
const pagesize = ref(7);
const pagenum = ref(1);
const totalfriendlink = ref(0);

// 友链编辑器信息
const friendid = ref(-1);
const editorswitch = ref(false); // false是新增 true是编辑
const friendlinkeditorbtn = ref('');
const showfriendlinkeditor = ref(false);
const friendlinkeditortitle = ref('');
const friendname = ref('');
const friendsite = ref('');
const friendprofile = ref('');
const frienddescription = ref('');

watch([showfriendlinkeditor], () => {
    if (!showfriendlinkeditor.value) {
        friendid.value = -1;
        friendname.value = '';
        friendsite.value = '';
        friendprofile.value = '';
        frienddescription.value = '';
    }
});

// 编辑友链
function friendLinkEditor(item) {
    if (!item.friendname) {
        friendlinkeditortitle.value = '新增友链';
        friendlinkeditorbtn.value = '新增';
        editorswitch.value = false;
    } else {
        friendlinkeditortitle.value = '编辑友链';
        friendlinkeditorbtn.value = '更新';
        friendid.value = item.friendlinkid;
        friendname.value = item.friendname;
        friendsite.value = item.friendsite;
        friendprofile.value = item.friendprofile;
        if (friendprofile.value === '/img/userprofile.png') friendprofile.value = '';
        frienddescription.value = item.frienddescription;
        editorswitch.value = true;
    }
    showfriendlinkeditor.value = true;
}
function friendLinkEdit() {
    window.$loadingBar.start();
    if (!editorswitch.value) {
        axios.post('/admin/newfriendlink', {"friendname": friendname.value, "friendsite": friendsite.value, "friendprofile": friendprofile.value, "frienddescription": frienddescription.value}, 
        {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
            if (res.data.code === 200) {
                window.$loadingBar.finish();
                showfriendlinkeditor.value = false;
                window.$message.success('新增友链成功');
                friendname.value = '';
                friendsite.value = '';
                friendprofile.value = '';
                frienddescription.value = '';
                queryFriendLink();
            } else {
                friendname.value = '';
                friendsite.value = '';
                friendprofile.value = '';
                frienddescription.value = '';
                window.$loadingBar.error();
                errmsg(res.data.code);
            }
        });
    } else {
        axios.post('/admin/modifyfriendlink', {"friendlinkid": friendid.value, "friendname": friendname.value, "friendsite": friendsite.value, "friendprofile": friendprofile.value, "frienddescription": frienddescription.value},
        {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
            if (res.data.code === 200) {
                window.$loadingBar.finish();
                showfriendlinkeditor.value = false;
                window.$message.success('更新友链成功');
                friendid.value = -1;
                friendname.value = '';
                friendsite.value = '';
                friendprofile.value = '';
                frienddescription.value = '';
                queryFriendLink();
            } else {
                friendid.value = -1;
                friendname.value = '';
                friendsite.value = '';
                friendprofile.value = '';
                frienddescription.value = '';
                window.$loadingBar.error();
                errmsg(res.data.code);
            }
        });
    }
}

// 加载友链列表
const friendlinks = ref([]);
function queryFriendLink() {
    axios.post('/friendlinks', {"pagesize": pagesize.value, "pagenum": pagenum.value}).then(res => {
        if (res.data.code === 200) {
            console.log(res.data);
            totalfriendlink.value = res.data.total;
            friendlinks.value = res.data.friendlinks;
        }
    });
}
const friendlinkeditorbtndisabled = ref(true);
watch([friendname, friendsite, friendprofile, frienddescription], () => {
    friendlinkeditorbtndisabled.value = ![friendname.value, friendsite.value, friendprofile.value, frienddescription.value].every(value => value !== '');
});

// 错误头像
function errorFirendProfile(item) {
    item.friendprofile = '/img/userprofile.png'
}

// 加载更多友链
const previousPage = ref(true);
const nextPage = ref(false);
watch([pagenum, totalfriendlink, friendlinks], () => {
    if (pagenum.value <= 1) {
        previousPage.value = true;
    } else {
        previousPage.value = false;
    }
    if (pagenum.value >= Math.ceil(totalfriendlink.value / pagesize.value)) {
        nextPage.value = true;
    } else {
        nextPage.value = false;
    }
})

function previousPageBtn() {
    pagenum.value -= 1;
    queryFriendLink();
}
function nextPageBtn() {
    pagenum.value += 1;
    queryFriendLink();
}

// 删除友链
const deletefriendlinkid = ref(-1);
const showfriendlinkdeletor = ref(false);
function deleteFriendLink(item) {
    showfriendlinkdeletor.value = true;
    deletefriendlinkid.value = item.friendlinkid;
}
function onPositiveClick() {
    window.$loadingBar.start();
    axios.post('/admin/deletefriendlink', {"friendlinkid": deletefriendlinkid.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            deletefriendlinkid.value = -1;
            window.$message.success('删除成功');
            queryFriendLink();
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    })
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
    queryFriendLink();
});
</script>
<template>
    <div class="infotitlebox">
        <h2>友链管理</h2>
        <p style="color:gray; font-size: 14px;">在这里可以设置各种站外链接，方便自己和访问者能够访问到其它自己感兴趣的站点~</p>
        <p style="color:gray; font-size: 14px;">tips:不要轻易给他人设置友情链接哦！</p>
    </div>
    <div class="infobox" style="display: flex; flex-wrap: wrap; min-height: 485px;">
        <div class="friendlinkbox" v-for="item in friendlinks" :key="item.friendlinkid">
            <div class="friendlinkcard">
                <a :href="item.friendsite" target="blank" style="text-decoration: none; color: black">
                    <div class="friendlinkinfo">
                        <div class="friendlinkimg">
                            <img :src="item.friendprofile" @error="errorFirendProfile(item)" style="height: 100px; width: 100px; border-radius: 50%;"/>
                        </div>
                        <div class="friendlinknamebox">
                            <span class="friendlinkname">{{ item.friendname }}</span>
                        </div>
                        <div class="friendlinkdescriptionbox">
                            <span class="friendlinkdescription">{{ item.frienddescription }}</span>
                        </div>
                    </div>
                </a>
                <div class="friendlinkbtn">
                    <n-button @click="friendLinkEditor(item)" strong type="primary" size="small" ghost style="margin-right: 10px; margin-left: 17px;">
                        更新
                    </n-button>
                    <n-button @click="deleteFriendLink(item)" strong type="error" size="small" ghost>
                        删除
                    </n-button>
                </div>
            </div>
        </div>
        <div class="friendlinkbox" @click="friendLinkEditor" style="cursor: pointer;">
            <div class="friendlinkcard">
                <div style="margin-top: 40px;">
                    <img src="/img/add.svg" style="height:140px; width: 140px;" />
                </div>
            </div>
        </div>
    </div>
    <div class="addmorebtnbox">
        <n-button strong secondary :disabled="previousPage" @click="previousPageBtn" round size="large" class="addmorebtn" style="background-color: rgb(186, 193, 197);">
            <
        </n-button>&nbsp&nbsp&nbsp
        <n-button strong secondary :disabled="nextPage" @click="nextPageBtn" round size="large" class="addmorebtn" style="background-color: rgb(186, 193, 197);">
            >
        </n-button>
    </div>
    <n-modal
        v-model:show="showfriendlinkeditor"
        preset="card"
        style="width: 550px;"
        :title="friendlinkeditortitle"
        size="huge"
        :bordered="false"
    >
        <div class="inputbox">
            <i class="userid"></i>
            <span class="input-text"> 友链名称</span>
            <n-input v-model:value="friendname" type="text" placeholder="友链名称"/>
        </div>
        <div class="inputbox">
            <i class="url"></i>
            <span class="input-text"> 友链地址</span>
            <n-input v-model:value="friendsite" type="text" placeholder="友链地址"/>
        </div>
        <div class="inputbox">
            <i class="profilephoto"></i>
            <span class="input-text"> 友链图标</span>
            <n-input v-model:value="friendprofile" type="text" placeholder="友链图标"/>
        </div>
        <div class="inputbox">
            <i class="slogan"></i>
            <span class="input-text"> 描述</span>
            <n-input v-model:value="frienddescription" type="text" placeholder="描述"/>
        </div>
        <template #footer>
            <div class="btn">
                <n-button :disabled="friendlinkeditorbtndisabled" @click="friendLinkEdit" strong round type="primary" size="large">
                    {{ friendlinkeditorbtn }}
                </n-button>
            </div>
        </template>
  </n-modal>
  <n-modal
        v-model:show="showfriendlinkdeletor"
        v-model:value="delepid"
        :mask-closable="false"
        preset="dialog"
        type="error"
        title="确认删除该友链吗？"
        content="删除后不可恢复！"
        positive-text="确认"
        negative-text="取消"
        @positive-click="onPositiveClick"
        @negative-click="onNegativeClick"
    />
</template>