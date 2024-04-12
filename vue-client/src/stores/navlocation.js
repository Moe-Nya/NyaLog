import { defineStore } from 'pinia'

export const useNavLocationStore = defineStore('navloc', () => {
    const navloc = ref('/');
    return {navloc}
})