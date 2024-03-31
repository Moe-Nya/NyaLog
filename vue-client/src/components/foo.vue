<script setup>
    import { ref, onMounted } from 'vue'
    import { usePublicUserInfoStore } from '../stores/userpublicinfo'
    import { useBlogSetStore } from '../stores/blogset'
    // 配置值，获取需要的footer数据
    const userinfo = usePublicUserInfoStore();
    const { GetBlogInfo, data } = useBlogSetStore();
    const createYear = ref(null);
    const startDate = ref(null)
    const timeDiff = ref(null);
    const days = ref(null);
    const hours = ref(null);
    const minutes = ref(null);
    const seconds = ref(null);
    userinfo.GetUserInfo();
    const currentYear = new Date().getFullYear();
    // 页面加载时挂载并且获取pinia中的值
    onMounted(() => {
        GetBlogInfo().then(() => {
            createYear.value = new Date(data.sitecreatetime).getFullYear();
            startDate.value = new Date(data.sitecreatetime);
            RunTime(startDate.value);
        })
    });
    // 用来实时计算博客上线的时长
    function RunTime(startDate) {
        setInterval(() => {
            const endDate = new Date();
            // 计算时间差（单位为毫秒）
            timeDiff.value = endDate - startDate;
            // 将时间差转换为天、小时、分钟和秒
            days.value = Math.floor(timeDiff.value / (1000 * 60 * 60 * 24));
            hours.value = Math.floor((timeDiff.value % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
            minutes.value = Math.floor((timeDiff.value % (1000 * 60 * 60)) / (1000 * 60));
            seconds.value = Math.floor((timeDiff.value % (1000 * 60)) / 1000);
        }, 1000);
    }
</script>

<style scoped>
    .footer{
        height:auto;
        font-weight: 500;
        font-size: 15px;
        text-align: center;
        position: absolute;
        bottom: 5px;
        left:50%;
        transform: translate(-50%);
    }

    .footer-heart{
        background: url("img/heart.svg") no-repeat;
        background-size: 15px 15px;
        background-position: center center;
        height: 20px;
        width: 20px;
        display: inline-block;
        animation: scaleAnimation 2s infinite;
        position: relative;
        top:3px;
    }

    .cuteword{
        animation: shakeAnimation 1s infinite;
        display: inline-block;
    }

    /* icon抖动 */
    @keyframes scaleAnimation {
        0% {
            transform: scale(1);
        }
        50% {
            transform: scale(1.5);
        }
        100% {
            transform: scale(1);
        }
    }

    /* 颜文字抖动 */
    @keyframes shakeAnimation {
        0% {
            transform: translateX(0);
        }
        10% {
            transform: translateX(-1px);
            transform: translateY(-px);
            transform: rotate(-2deg);
        }
        20% {
            transform: translateX(1px);
            transform: translateY(-1px);
            transform: rotate(2deg);
        }
        30% {
            transform: translateX(-1px);
            transform: translateY(1px);
            transform: rotate(-2deg);
        }
        40% {
            transform: translateX(1px);
            transform: translateY(-1px);
            transform: rotate(2deg);
        }
        50% {
            transform: translateX(-1px);
            transform: translateY(1px);
            transform: rotate(-2deg);
        }
        60% {
            transform: translateX(1px);
            transform: translateY(-1px);
            transform: rotate(2deg);
        }
        70% {
            transform: translateX(-1px);
            transform: translateY(1px);
            transform: rotate(-2deg);
        }
        80% {
            transform: translateX(1px);
            transform: translateY(-1px);
            transform: rotate(2deg);
        }
        90% {
            transform: translateX(-1px);
            transform: translateY(1px);
            transform: rotate(-2deg);
        }
        100% {
            transform: translateX(0);
        }
    }

    @media (max-width: 720px) {
        .border{
            width: 80%;
    }
}
</style>
<template>
    <footer class="footer">
        <p>
        版权所有 ©{{ createYear }}-{{ currentYear }} {{ userinfo.data.username }}<br>
        <a href="https://github.com/Moe-Nya/NyaLog" target="_blank" style="color: black; text-decoration: none;"><span>NyaLog</span></a> By <a href="https://github.com/Moe-Nya" target="_blank" style="color: black; text-decoration: none;"><span>Moe_Nya</span></a> With <i class="footer-heart"></i> | All Rights Reserved<br>
        <span class="cuteword">(●'◡'●)ﾉ</span>本博客可爱地运行了{{ days }}天{{ hours }}小时{{ minutes }}分{{ seconds }}秒<br>
        </p>
    </footer>
</template>