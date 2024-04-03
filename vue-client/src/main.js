import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './modules/router'
import pinia from './modules/pinia'
import 'vfonts/FiraCode.css'
import naive from 'naive-ui/es/preset'
import axios from 'axios'

axios.defaults.baseURL = 'http://192.168.31.145:8080/api/v1';
axios.defaults.timeout = 30 * 1000;

const app = createApp(App)
app.use(router)
app.use(pinia)
app.use(naive)
// 屏蔽控制台的警告信息
app.config.warnHandler = (msg, vm, trace) => {
};
app.mount('#app')
