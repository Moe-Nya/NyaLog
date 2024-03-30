import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import axios from 'axios'

export const usePublicUserInfoStore = defineStore('userinfo', () => {
    const data = reactive({
        username: "一个神秘站点",
        profilephoto: "",
        email: "",
        slogan: ""
    })
    
    async function GetUserInfo() {
        const response = await axios.get('/queryuser');
        if (response.data.code === 200) {
            data.username = response.data.userinfo.username;
            data.profilephoto = response.data.userinfo.profilephoto;
            data.email = response.data.userinfo.email;
            data.slogan = response.data.userinfo.slogan;
        } else {
            console.log(response.data.message);
        }
        
    }
    return {GetUserInfo, data};
})