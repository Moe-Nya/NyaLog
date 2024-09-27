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
import { onBeforeUnmount } from 'vue';

const handleBeforeUnload = (event) => {
    pageCache();
    // 显示提示信息
    event.preventDefault();
    event.returnValue = ''; // 某些浏览器需要这个
};

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
                DeleteCache();
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
            DeleteCache();
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

// 页面缓存
function pageCache() {
    // 创建页面缓存
    if (edit === '') {
        window.localStorage.setItem('page', JSON.stringify({pagetitle: pagetitle.value, 
            text: text.value, pageurl: pageurl.value}));
    } else {
        // 编辑页面 根据页面名称存储
        window.localStorage.setItem(edit.toString(), JSON.stringify({pagetitle: pagetitle.value, 
            text: text.value, pageurl: pageurl.value}));
    }
}

onBeforeUnmount(() => {
    window.removeEventListener('beforeunload', handleBeforeUnload);
    pageCache();
});

// 页面加载时读取缓存
function readCache() {
    if (edit.toString().length == 0 && window.localStorage.getItem('page')) {
        const page = JSON.parse(window.localStorage.getItem('page'));
        if (page) {
            pagetitle.value = page.pagetitle;
            text.value = page.text;
            pageurl.value = page.pageurl;
        }
    } else if (window.localStorage.getItem(edit.toString())) {
        const page = JSON.parse(window.localStorage.getItem(edit.toString()));
        if (page) {
            pagetitle.value = page.pagetitle;
            text.value = page.text;
            pageurl.value = page.pageurl;
        }
    } else {
        editorModel();
    }
}

// 清除缓存
function DeleteCache() {
    if (edit.toString().length == 0) {
        window.localStorage.removeItem('page');
        text.value = '# 说些什么吧~=w=';
        pagetitle.value = '';
        pageurl.value = '';
    } else {
        window.localStorage.removeItem(edit.toString());
        editorModel();
    }
}

// 初始化
onMounted(() => {
    window.addEventListener('beforeunload', handleBeforeUnload);
    Validate();
    readCache();
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
    <div class="inputbox">
        <i class="clearcache"></i>
        <span class="input-text"> 清空本地缓存</span>
        <n-button style="margin-left: 20px;" @click="DeleteCache" strong round type="error" size="small">
            清除
        </n-button><br/>
        <span class="tips">清除掉保存在浏览器中的本页页面信息~</span>
    </div>
    <div class="btn">
        <n-button style="margin-left: 5px;" :disabled="updatePageInfo" @click="submitPage" strong round type="primary" size="large">
            提交
        </n-button>
    </div>
</template>