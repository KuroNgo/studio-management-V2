
import { defineStore } from 'pinia'

export const useUserStore = defineStore('AccountStores', {
    state: () => ({
        data: null,
        RoleRegister:{
            error:false,
            success:false,
        },
        host: 'http://localhost:8000'
    }),
    actions: {
        async GetUser() {
            try {
                const res = await fetch(`${this.host}/api/client/v1/category/get-all`, {
                    method: "GET",
                    headers: {
                        "Content-Type": 'application/json'
                    },
                })
                const data = await res.json();
                // console.log(data)

            } catch (error) {
                console.error("Lỗi:", error);
            }
        },
        async Register(fullName, userName, email, passWord, phone, role, provider, verified, enable) {
            try {
                const res = await fetch(`${this.host}/api/client/v1/register`, {
                    method: "POST",
                    headers: {
                        "Content-Type": 'application/json'
                    },
                    // mode: 'cors',
                    body: JSON.stringify({fullName, userName, email, passWord, phone, role, provider, verified, enable}),
                })
                if (!res.ok) {
                    const errorText = await res.text(); // Đọc nội dung lỗi
                    this.RoleRegister.error=true;
                    this.RoleRegister.success=false;
                    //throw new Error(`Lỗi yêu cầu API: ${res.status} - ${errorText}`);
                }else{
                    this.RoleRegister.success=true;
                    this.RoleRegister.error=false;
                }
                //const user = await res.json();
                //console.log(user);
            } catch (error) {
                console.log(error)
            }

        }
    },
})