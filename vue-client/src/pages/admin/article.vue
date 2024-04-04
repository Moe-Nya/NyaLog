<script setup>
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { useRouter } from 'vue-router';
import MessageAPI from '../../components/message.vue'
import axios from 'axios'
import errmsg from '../../modules/errmsg';

const router = useRouter();

// 编辑器设置
const formatCopiedText = (text) => {
  return `${text}  - from md-editor-v3`;
};
const mdHeadingId = (_text, _level, index) => `heading-${index}`;

// 进入页面权限验证
function Validate() {
    window.$loadingBar.start();
    axios.get("/admin/", {headers: {Authorization: window.localStorage.getItem('token')}}).then(res => {
        if (res.data.code !== 200) {
            window.localStorage.clear('token');
            window.$loadingBar.finish();
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
});
const text = ref('# Hello Editor');
</script>
<template>
    <div class="articlebox">
        <MdEditor 
        style="height: 700px;"
        v-model="text"
        previewTheme="default"
        codeTheme="github"
        :formatCopiedText="formatCopiedText"
        :mdHeadingId="mdHeadingId"/>
    </div>
</template>