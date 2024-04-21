<script setup>
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { useRouter } from 'vue-router';
import { useRoute } from 'vue-router';
import MessageAPI from '../../components/message.vue';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import { onMounted, watch } from 'vue';
import { config } from 'md-editor-v3';
import ImageFiguresPlugin from '../../modules/markdown-it-image-figures.js';

config({
  markdownItConfig: (mdit) => {
    mdit.use(ImageFiguresPlugin, { figcaption: true });
  },
});

const router = useRouter();
const route = useRoute();

// 编辑器设置
// 编辑器设置-工具栏设置
const toolbars = [
  'bold',
  'underline',
  'italic',
  '-',
  'title',
  'strikeThrough',
  'sub',
  'sup',
  'quote',
  'unorderedList',
  'orderedList',
  'task',
  'image',
  '-',
  'codeRow',
  'code',
  'link',
  'table',
  'mermaid',
  'katex',
  '-',
  'revoke',
  'next',
  '=',
  'pageFullscreen',
  'fullscreen',
  'preview',
  'htmlPreview',
  'catalog',
];

// 编辑||新增切换
let edit = route.query.id || '';
function editorModel() {
    if (edit !== '') {
        axios.get(`/${edit}`).then(res => {
            if (res.data.code === 200) {
                pagetitle.value = res.data.page.ptitle;
                text.value = res.data.page.pcontent;
                pageurl.value = res.data.page.purl;
                pid.value = res.data.page.pid
            } else {
                window.$message.error(res.data.code);
            }
        });
    }
}

// 页面信息
const text = ref('# 说些什么吧~=w=');
const pagetitle = ref('');
const pageurl = ref('');
const pid = ref();
const formatCopiedText = (text) => {
  return `${text}  - from md-editor-v3`;
};
const mdHeadingId = (_text, _level, index) => `heading-${index}`;
const updatePageInfo = ref(true);
watch([text, pagetitle, pageurl], () => {
    updatePageInfo.value = ![text.value, pagetitle.value, pageurl.value].every(value => value !== '');
});

// 提交页面
function submitPage() {
    window.$loadingBar.start();
    if ((edit.toString()).length !== 0) {
        axios.post('/admin/editpage', {"pid": pid.value, "purl": pageurl.value, "ptitle": pagetitle.value, "pcontent": text.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
            if (res.data.code === 200) {
                window.$loadingBar.finish();
                window.$message.success('编辑成功');
            } else {
                window.$loadingBar.error();
                errmsg(res.data.code);
            }
        });
    } else {
        axios.post('/admin/newpage', {"purl": pageurl.value, "ptitle": pagetitle.value, "pcontent": text.value}, {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            window.$message.success('发布成功');
            router.push('/admin/pagelist');
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
        });
    }
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
    editorModel();
});
</script>
<template>
    <div class="infobox" style="margin-top: -20px;">
        <div class="inputbox">
            <i class="articletitleicon"></i>
            <span class="input-text"> 页面标题</span>
            <div style="display: flex;">
                <n-input v-model:value="pagetitle" type="text" placeholder="页面标题"/>
            </div>
            <span class="tips">好的标题会吸引读者点进去~</span>
        </div>
    </div>
    <div class="articlebox">
        <i class="articleicon"></i>
        <span class="input-text"> 正文</span>
        <MdEditor 
        style="height: 700px;"
        v-model="text"
        previewTheme="default"
        codeTheme="github"
        :formatCopiedText="formatCopiedText"
        :mdHeadingId="mdHeadingId"
        :toolbars
        noUploadImg/>
    </div>
    <div class="infobox">
        <div class="inputbox">
            <i class="url"></i>
            <span class="input-text"> 页面url</span>
            <div style="display: flex;">
                <n-input v-model:value="pageurl" type="text" placeholder="页面url"/>
            </div>
            <span class="tips">这里的字符将会作为页面索引的表示，例如https://example.com/索引值。</span>
        </div>
    </div>
    <div class="btn">
        <n-button style="margin-left: 5px;" :disabled="updatePageInfo" @click="submitPage" strong round type="primary" size="large">
            提交
        </n-button>
    </div>
</template>