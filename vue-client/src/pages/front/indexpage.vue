<script setup>
import { defineEmits } from 'vue';
import { useNavLocationStore } from '../../stores/navlocation'

const navloc = useNavLocationStore();


function loadRouter() {
    const pathname = window.location.pathname;
    let suffix = pathname.substring(1);
    if (suffix === '') {
        suffix = '/';
    }
    navloc.navloc = suffix;
}

onMounted(() => {
    loadRouter();
})
</script>
<template>
    这是主页
</template>