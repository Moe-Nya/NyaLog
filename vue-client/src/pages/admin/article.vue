<script setup>
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { useRouter } from 'vue-router';
import { useRoute } from 'vue-router';
import MessageAPI from '../../components/message.vue'
import axios from 'axios'
import errmsg from '../../modules/errmsg';
import { watch } from 'vue';
import { config } from 'md-editor-v3';
import ImageFiguresPlugin from '../../modules/markdown-it-image-figures.js';
import { onBeforeUnmount } from 'vue';

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

const handleBeforeUnload = (event) => {
    articleCache();
    // 显示提示信息
    event.preventDefault();
    event.returnValue = ''; // 某些浏览器需要这个
};

// 编辑||新增切换
let edit = route.query.id || '';

// 文章信息
const articletitle = ref('');
const text = ref('# 说些什么吧~=w=');
const articlePhoto = ref('');
const category = ref(-1);
const aiswitch = ref(0);
const aiSwitchBtn = ref(false);
watch([aiSwitchBtn], () => {
    if (aiSwitchBtn.value) {
        aiswitch.value = 1;
    } else {
        aiswitch.value = 0;
    }
});
const updateArticleInfo = ref(true);
watch([articletitle, text], () => {
    updateArticleInfo.value = ![articletitle.value, text.value].every(value => value !== '');
});
// 编辑||新增切换
function editorModel() {
    if ((edit.toString()).length !== 0) {
        axios.get(`/article/${edit}`).then(res => {
            if (res.data.code === 200) {
                articletitle.value = res.data.article.articletitle;
                text.value = res.data.article.text;
                articlePhoto.value = res.data.article.articleimg;
                category.value = res.data.article.cid;
                aiswitch.value = res.data.article.aiswitch;
                if (aiswitch.value === 1) {
                    aiSwitchBtn.value = true;
                } else {
                    aiSwitchBtn.value = false;
                }
            } else {
                window.$message.error(res.data.code);
            }
        });
    }
}

const formatCopiedText = (text) => {
  return `${text}  - from md-editor-v3`;
};
const mdHeadingId = (_text, _level, index) => `heading-${index}`;

// 类型选择器
const categorysOptions = ref([]);

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

// 提交文章
function submitArticle() {
    window.$loadingBar.start();
    if (edit != '') {
        axios.post('/admin/editarticle', {"articleid": Number(edit), "articleimg": articlePhoto.value, "articletitle": articletitle.value, "cid": category.value, "aiswitch": aiswitch.value, "text": text.value, "articlecategory": 0}, {headers: {Authorization: window.localStorage.getItem('token')}}).
        then(res => {
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
        axios.post("/admin/newarticle", {"articletitle": articletitle.value, "text": text.value, "articleimg": articlePhoto.value, "cid": category.value, "aiswitch": aiswitch.value, "articlecategory": 0}, {headers: {Authorization: window.localStorage.getItem('token')}}).
        then(res => {
            if (res.data.code === 200) {
                window.$loadingBar.finish();
                window.$message.success('发布成功');
                DeleteCache();
                router.push('/admin/articlelist');
            } else {
                window.$loadingBar.error();
                errmsg(res.data.code);
            }
        });
    }
}

// 初始化
function init() {
    axios.get("/category").then(res=> {
        if (res.data.code === 200) {
            for (let i = 0; i < res.data.category.length; i++) {
                categorysOptions.value = categorysOptions.value.concat({label: res.data.category[i].categorytext, value: res.data.category[i].cid});
            }
        }
    });
    categorysOptions.value = categorysOptions.value.concat({label : '不设置分组', value: -1});
}

// 文章缓存
function articleCache() {
// 创建文章
if (edit === '') {
    window.localStorage.setItem('article', JSON.stringify({articletitle: articletitle.value, 
        text: text.value, articlePhoto: articlePhoto.value, category: category.value, aiswitch: aiswitch.value}));
    } else {
        // 编辑文章 根据文章id缓存
        window.localStorage.setItem(edit.toString(), JSON.stringify({articletitle: articletitle.value, 
            text: text.value, articlePhoto: articlePhoto.value, category: category.value, aiswitch: aiswitch.value}));
    }
}

onBeforeUnmount(() => {
    window.removeEventListener('beforeunload', handleBeforeUnload);
    articleCache();
});

// 页面加载时读取缓存
function readCache() {
    if (edit.toString().length == 0 && window.localStorage.getItem('article')) {
        const article = JSON.parse(window.localStorage.getItem('article'));
        if (article) {
            articletitle.value = article.articletitle;
            text.value = article.text;
            articlePhoto.value = article.articlePhoto;
            category.value = article.category;
            aiswitch.value = article.aiswitch;
            if (aiswitch.value === 1) {
                aiSwitchBtn.value = true;
            } else {
                aiSwitchBtn.value = false;
            }
        }
    } else if (window.localStorage.getItem(edit.toString())) {
        const article = JSON.parse(window.localStorage.getItem(edit.toString()));
        if (article) {
            articletitle.value = article.articletitle;
            text.value = article.text;
            articlePhoto.value = article.articlePhoto;
            category.value = article.category;
            aiswitch.value = article.aiswitch;
            if (aiswitch.value === 1) {
                aiSwitchBtn.value = true;
            } else {
                aiSwitchBtn.value = false;
            }
        }
    } else {
        editorModel();
    }
}

// 清除缓存
function DeleteCache() {
    if (edit.toString().length == 0) {
        window.localStorage.removeItem('article');
        text.value = '# 说些什么吧~=w=';
        articletitle.value = '';
        articlePhoto.value = '';
        category.value = -1;
        aiswitch.value = 0;
        aiSwitchBtn.value = false;
    } else {
        window.localStorage.removeItem(edit.toString());
        editorModel();
    }
}

// 初始化
onMounted(() => {
    window.addEventListener('beforeunload', handleBeforeUnload);
    Validate();
    init();
    readCache();
});
</script>
<template>
    <div class="infobox" style="margin-top: -20px;">
        <div class="inputbox">
            <i class="articletitleicon"></i>
            <span class="input-text"> 文章标题</span>
            <div style="display: flex;">
                <n-input v-model:value="articletitle" type="text" placeholder="文章标题"/>
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
            <i class="profilephoto"></i>
            <span class="input-text"> 文章封面图片</span>
            <n-input v-model:value="articlePhoto" type="text" placeholder="文章封面图片"/>
            <span class="tips">一定要设置哦，否则只能显示默认图片作为封面啦~</span>
        </div>
    </div>
    <div class="infobox" style="margin-bottom: -10px;">
        <div class="inputbox">
            <i class="articlecategory"></i>
            <span class="input-text"> 分类选择</span>
            <div style="display: flex;">
                <n-select v-model:value="category" :options="categorysOptions" placeholder="分类选择"/>
            </div>
            <span class="tips">选择文章分类~</span>
        </div>
    </div>
    <div class="inputbox">
        <i class="aisummary"></i>
        <span class="input-text"> AI摘要启用</span>
        <n-switch size="large" 
        style="margin-left: 20px;"
        v-model:value="aiSwitchBtn"
        /><br />
        <span class="tips">人工智能的未来会是什么样呢？~</span>
    </div>
    <div class="inputbox">
        <i class="clearcache"></i>
        <span class="input-text"> 清空本地缓存</span>
        <n-button style="margin-left: 20px;" @click="DeleteCache" strong round type="error" size="small">
            清除
        </n-button><br/>
        <span class="tips">清除掉保存在浏览器中的本页文章信息~</span>
    </div>
    <div class="btn">
        <n-button style="margin-left: 5px;" :disabled="updateArticleInfo" @click="submitArticle" strong round type="primary" size="large">
            提交
        </n-button>
    </div>
</template>