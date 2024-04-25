import { createApp } from 'vue';
import './style.css';
import '../public/static/css/font.css';
import App from './App.vue';
import router from './modules/router';
import pinia from './modules/pinia';
import 'vfonts/FiraCode.css';
import naive from 'naive-ui/es/preset';
import axios from 'axios';
import url from '../config.json';
import { createMetaManager} from 'vue-meta';

axios.defaults.baseURL = url.api_url;
axios.defaults.timeout = 30 * 1000;

const app = createApp(App);
app.use(router);
app.use(pinia);
app.use(naive);
app.use(createMetaManager(false, {
    meta: { tag: 'meta', nameless: true }
}));
// 屏蔽控制台的警告信息
app.config.warnHandler = (msg, vm, trace) => {
};
app.mount('#app');
