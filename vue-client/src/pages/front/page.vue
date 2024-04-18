<script setup>
import { defineEmits, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/preview.css';
import axios from 'axios';
import errmsg from '../../modules/errmsg';
import { config } from 'md-editor-v3';
import ImageFiguresPlugin from '../../modules/markdown-it-image-figures.js';

config({
  markdownItConfig: (mdit) => {
    mdit.use(ImageFiguresPlugin, { figcaption: true });
  },
});

const route = useRoute();

const id = 'preview-only';
const title = ref('')
const text = ref('');
const scrollElement = document.documentElement;

// 页面内容加载
function queryPage() {
    window.$loadingBar.start();
    axios.get(`${route.params.pageid}`).then(res => {
        if (res.data.code === 200) {
            window.$loadingBar.finish();
            text.value = res.data.page.pcontent;
            title.value = res.data.page.ptitle;
        } else {
            window.$loadingBar.error();
            errmsg(res.data.code);
        }
    });
}

onMounted(() => {
    queryPage();
})
</script>
<template>
    <div class="articleandpagetitle">
        <span class="titlestyle" style="margin-bottom: 15px; color: #3d3d3d;">{{ title }}</span>
    </div>
    <div class="mdview">
        <MdPreview :editorId="id" :modelValue="text" previewTheme="mdeditor" style="color: #3d3d3d;"/>
    </div>
    <div class="headview">
        <MdCatalog :editorId="id" :scrollElement="scrollElement"/>
    </div>
</template>