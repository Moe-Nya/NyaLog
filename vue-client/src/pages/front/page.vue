<script setup>
import { defineEmits } from 'vue';
import { useNavLocationStore } from '../../stores/navlocation'

const navloc = useNavLocationStore();

const pathname = window.location.pathname;
let suffix = pathname.substring(1); // 去除路径开头的斜杠
navloc.navloc = suffix;

</script>
<template>
    这是页面
</template>