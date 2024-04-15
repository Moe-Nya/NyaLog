import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import axios from 'axios'

export const usePublicUserInfoStore = defineStore('userinfo', () => {
    const data = reactive({
        username: "一个神秘用户",
        profilephoto: "/img/userprofile.png",
        email: "",
        slogan: ""
    })
    
    async function GetUserInfo() {
        const response = await axios.get('/queryuser');
        if (response.data.code === 200) {
            if (response.data.userinfo.username !== "") data.username = response.data.userinfo.username;
            if (response.data.userinfo.profilephoto !== "") data.profilephoto = response.data.userinfo.profilephoto;
            data.email = response.data.userinfo.email;
            data.slogan = response.data.userinfo.slogan;
        } else {
            console.log(response.data.message);
        }
        
    }
    return {GetUserInfo, data};
})