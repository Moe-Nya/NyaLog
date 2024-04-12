<script setup>
import { defineEmits } from 'vue';
import { useNavLocationStore } from '../../stores/navlocation'

const navloc = useNavLocationStore();

// 路由定位
function loadRouter() {
    const pathname = window.location.pathname;
    let suffix = pathname.substring(1); // 去除路径开头的斜杠
    navloc.navloc = suffix;
}

onMounted(() => {
    loadRouter();
})
</script>
<template>
    这是归档
</template>