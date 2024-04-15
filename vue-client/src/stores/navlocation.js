import { defineStore } from 'pinia'

export const useNavLocationStore = defineStore('navloc', () => {
    const navloc = ref('/');

    function changeStatus(newnavloc) {
        navloc.value = newnavloc
    }

    return {navloc, changeStatus}
})