<script lang="js">
import { createDiscreteApi } from 'naive-ui'
import { defineComponent } from 'vue'

const {message, notification, dialog, loadingBar} = createDiscreteApi(
    ['message', 'dialog', 'notification', 'loadingBar']
)
// 这是naiveUI的'message', 'dialog', 'notification', 'loadingBar'几个UI组件的全局定义
// 在使用的时候需要在目标vue页面中引入：import MessageAPI from '../../components/message.vue'
// 调用使用：window.$message.success('登录成功');等等
window.$message = message;
window.$notification = notification;
window.$dialog = dialog;
window.$loadingBar = loadingBar;

export default defineComponent({});
</script>

<template>
    <div></div>
</template>