
import { defineStore } from 'pinia'

export const useUserStore = defineStore('myStoreId',{
    state: () => ({
        data: null,
    }),
    actions: {
        async GetUser() {
            try {
                const res = await fetch("http://localhost:8000/api/client/v1/category/get-all",{
                    method:"GET",
                    headers:{
                        "Content-Type": 'application/json'
                    },
                })
                const data = await res.json();
                console.log(data)
                
            } catch (error) {
                console.error("Lỗi:", error);
            }
        }
    }
})