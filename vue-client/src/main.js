import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './modules/router'
import pinia from './modules/pinia'
import 'vfonts/FiraCode.css'
import naive from 'naive-ui/es/preset'
import axios from 'axios'

axios .defaults.baseURL = 'http://localhost:8080/api/v1'

const app = createApp(App)
app.use(router)
app.use(pinia)
app.use(naive)

app.mount('#app')
