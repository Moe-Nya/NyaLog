<script setup>
import axios from 'axios'
import errmsg from '../../modules/errmsg';
import MessageAPI from '../../components/message.vue'
import { useRouter } from 'vue-router';
import { watch } from 'vue';

const router = useRouter();

// 更新按钮状态
const updateSiteInfo = ref(true);

// 站点信息
const sitename = ref('');
const sitebackground = ref('');
const createdate = ref('');
const commentEnable = ref();
const aiEnable = ref();
const aiCategory = ref();

// AI选择器
const aiSwitchBtn = ref();
const aiSummarySwitchBtn = ref(false);
const aiswitchCanDown = ref(true);
watch([aiSwitchBtn], () => {
    if (aiSwitchBtn.value) {
        aiEnable.value = 1;
        aiswitchCanDown.value = false;
    } else {
        aiEnable.value = 0;
        aiswitchCanDown.value = true;
    }
});
const aiSummarySwitch = ({ focused, checked }) => {
    const style = {};
    if (checked) {
        aiCategory.value = 1;
        style.background = "#18A058";
    if (focused) {
      style.boxShadow = "0 0 0 2px #18A058";
    }
    } else {
        aiCategory.value = 0;
        style.background = "#18A058";
        if (focused) {
        style.boxShadow = "0 0 0 2px #18A058";
        }
    }
    return style;
};

// 监测评论是否开启
const commentSwitch = ref();
watch([commentSwitch], () => {
    if (commentSwitch.value) {
        commentEnable.value = 1;
    } else {
        commentEnable.value = 0;
    }
});

// 站点创建时间不可是未来
const timestamp = ref();
function disablePreviousDate(ts) {
    return ts > Date.now();
}

// 监听timestamp的变化
watch([timestamp], () => {
    createdate.value = new Date(timestamp.value).toISOString();
});

// 监听输入框内容决定按钮是否能按下
watch([sitename, sitebackground, createdate], () => {
    updateSiteInfo.value = ![sitename.value, sitebackground.value, createdate.value].every(value => value !== '');
});

// 更新按钮
function updateSiteInfoBtn() {
    window.$loadingBar.start();
    axios.post("/admin/blogset", {"sitename": sitename.value, "sitebackground": sitebackground.value, "sitecreatetime": createdate.value, "commentswitch": commentEnable.value, "aiswitch": aiEnable.value, "aicategory": aiCategory.value}, 
    {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('博客信息更新成功');
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

// 获取初始信息
function init() {
    axios.get("/admin/blogsetlist", {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => { 
        if (res.data.code === 200) {
            sitename.value = res.data.blogsetinfo.sitename;
            sitebackground.value = res.data.blogsetinfo.sitebackground;
            timestamp.value = new Date(res.data.blogsetinfo.sitecreatetime).getTime();
            commentEnable.value = res.data.blogsetinfo.commentswitch;
            if (commentEnable.value === 1) {
                commentSwitch.value = true;
            } else {
                commentSwitch.value = false;
            } 
            aiEnable.value = res.data.blogsetinfo.aiswitch;
            if (aiEnable.value === 1) {
                aiSwitchBtn.value = true;
            } else {
                aiSwitchBtn.value = false;
            }
            aiCategory.value = res.data.blogsetinfo.aicategory;
            if (aiCategory.value === 1) {
                aiSummarySwitchBtn.value = true;
            } else {
                aiSummarySwitchBtn.value = false;
            }
        } else {
            errmsg(res.data.code);
        }
    })
}

// 初始化
onMounted(() => {
    Validate();
    init();
});
</script>
<template>
    <div class="userinfotitlebox">
        <h2>站点设置</h2>
        <p style="color:gray; font-size: 14px;">这里可以对站点信息进行设置更新，这些信息都是会公开展示的哦~</p>
    </div>
    <div class="infobox">
        <div class="inputbox">
            <i class="sitename"></i>
            <span class="input-text"> 站点名称</span>
            <n-input v-model:value="sitename" type="text" placeholder="站点名称"/>
            <span class="tips">叫什么好呢？~</span>
        </div>
        <div class="inputbox">
            <i class="profilephoto"></i>
            <span class="input-text"> 首页背景图</span>
            <n-input v-model:value="sitebackground" type="text" placeholder="首页背景图"/>
            <span class="tips">选择一张好看的背景图片吧~</span>
        </div>
        <div class="inputbox">
            <i class="time"></i>
            <span class="input-text"> 创建时间</span>
            <div class="inputbox">
                <n-date-picker style="margin-top: -10px;" v-model:value="timestamp" type="date" :is-date-disabled="disablePreviousDate" placeholder="站点创建时间"/>
            </div>
            <span class="tips">选择一个时间吧！未来不可选~ tip:该选项会影响到footer的运行时长</span>
        </div>
        <div class="inputbox">
            <i class="comment"></i>
            <span class="input-text"> 评论启用</span>&nbsp&nbsp&nbsp
            <n-switch size="large" 
            style="margin-left: 20px; margin-bottom: 5px;"
            v-model:value="commentSwitch"
            /><br />
            <span class="tips">启用评论的开关，畅所欲言吧！~</span>
        </div>
        <div class="inputbox">
            <i class="aisummary"></i>
            <span class="input-text"> AI摘要启用</span>
            <n-switch size="large" 
            style="margin-left: 20px; margin-bottom: 5px;"
            v-model:value="aiSwitchBtn"
            /><br />
            <span class="tips">人工智能的未来会是什么样呢？~</span>
        </div>
        <div class="inputbox">
            &nbsp&nbsp&nbsp&nbsp
            <n-switch :rail-style="aiSummarySwitch" v-model:value="aiSummarySwitchBtn" :disabled="aiswitchCanDown">
                <template #checked>
                    通仪千问
                </template>
                <template #unchecked>
                    GPT3.5
                </template>
            </n-switch>
        </div>
        <div class="btn">
            <n-button style="margin-left: 5px;" :disabled="updateSiteInfo" @click="updateSiteInfoBtn" strong round type="primary" size="large">
                更新信息
            </n-button>
        </div>
    </div>
</template>