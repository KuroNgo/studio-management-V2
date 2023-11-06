<template>
  <main class="py-20 h-screen">
    <div class="flex h-full w-3/4 mx-auto border border-solid border-gray-500">
      <div class="w-2/4 mt-auto mb-auto px-32">
        <h4 class="Heading4 uppercase text-center mb-11">ĐĂNG KÝ</h4>
        <form method="get" onsubmit="return false;">
          <div class="relative mb-8">
            <input
              type="text"
              v-model="inputFullName.value"
              placeholder="Hãy điền họ và tên của bạn"
              class="form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full" />
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2 left-3">
              Họ và tên
            </p>
          </div>

          <div class="relative mb-3">
            <input
              type="text"
              v-model="inputUserName.value"
              placeholder="Không chứa các ký tự *,#,&"
              pattern="^[^*#&]+$"
              :class="{
                ' outline-red-700 outline-2': inputUserName.error,
              }"
              class="form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full" />
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2 left-3">
              Tên tài khoản
            </p>
            <p class="invisible peer-invalid:visible text-pink-600 Caption pl-3">
              Tên tài khoảng không chứa các ký tự *,#,&
            </p>
          </div>

          <div class="relative mb-3">
            <input
              type="email"
              v-model="inputEmail.value"
              pattern="^[a-zA-Z0-9._%+-]+@(gmail\.com|gmail\.com\.vn)$"
              placeholder="Địa chỉ email của bạn"
              :class="{
                ' outline-red-700 outline-2': inputEmail.error,
              }"
              class="form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full" />
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2 left-3">
              Địa chỉ email
            </p>
            <p class="invisible peer-invalid:visible text-pink-600 Caption pl-3">
              Vui lòng cung cấp một địa chỉ email hợp lệ
            </p>
          </div>

          <div class="relative mb-3">
            <input
              type="text"
              v-model="inputPhoneNumber.value"
              placeholder="Số điện thoại của bạn"
              maxlength="10"
              pattern="^0\d{9,10}$"
              :class="{
                ' outline-red-700 outline-2': inputPhoneNumber.error,
              }"
              class="form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full" />
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2 left-3">
              Số điện thoại
            </p>
            <p class="invisible peer-invalid:visible text-pink-600 Caption pl-3">
              Vui lòng cung cấp số điện thoại hợp lệ
            </p>
          </div>

          <div class="relative mb-3">
            <input
              type="text"
              v-model="inputPassword.value"
              pattern="^(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{8,}$"
              placeholder="Chứa ít nhất 8 ký tự bao gồm ( chứ viết hoa, chữ số và một ký tự đặc biệt )"
              :class="{
                ' outline-red-700 outline-2': inputPassword.error,
              }"
              class="form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full" />
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2 left-3">
              Mật khẩu
            </p>
            <p class="invisible peer-invalid:visible text-pink-600 Caption pl-3">
              Chứa ít nhất 8 ký tự bao gồm ( chứ viết hoa, chữ số và một ký tự đặc biệt )
            </p>
          </div>

          <div class="relative mb-3">
            <input
              type="text"
              v-model="enterPassword.value"
              placeholder="Nhập lại mật khẩu bạn đã nhập"
              :class="{
                ' outline-red-700 outline-2': enterPassword.error,
              }"
              class="form-input mb-2 placeholder:text-gray-400 placeholder:Caption outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full" />
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2 left-3">
              Nhập lại mật khẩu
            </p>
          </div>
          <p class="Body2 text-center text-pink-600 pl-3 mb-3" v-if="RoleRegister.error">Tài khoảng tồn tại</p>

          <button
            @click="Run()"
            type="submit"
            class="text-DomlantColor Button p-3 bg-AccentColor w-full rounded-xl mb-11">
            Đăng ký
          </button>
        </form>
        <p class="Body2 text-center">
          Đã có tài khoản ?
          <RouterLink to="/login"><strong class="underline text-SecondColor">Đăng nhập</strong> </RouterLink>
        </p>
      </div>

      <div class="background-img w-2/4 h-full"></div>
    </div>

    <!-- alert thông báo -->
    <div
      id="toastBox"
      class="pt-9 w-1/4 flex overflow-hidden flex-col-reverse gap-y-6 fixed top-0 right-0 show-alert transform transition-transform duration-1000 ease-in-out"></div>
    <!-- <div class=" pt-9 w-1/4 flex flex-col-reverse gap-y-6 fixed top-0 right-0 show-alert transform transition-transform duration-1000 ease-in-out">
      <div
        class="bg-teal-100 border-t-4 border-teal-500 rounded-b text-teal-900 px-4 py-3 shadow-md"
        role="alert" v-if="RoleRegister.success" :class="{ ' translate-x-full': RoleRegister.close }">
        <div class="flex">
          <div class="py-1"><svg class="fill-current h-6 w-6 text-teal-500 mr-4" xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20">
              <path
                d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 11V9h2v6H9v-4zm0-6h2v2H9V5z" />
            </svg></div>
          <div>
            <p class="font-bold">Đăng ký thành công</p>
            <p class="text-sm">Chào mừng bạn đã đăng ký thành công tài khoản trên Cỏ Studio</p>
          </div>
        </div>
      </div>
      <div
        class=" bg-red-100 border-t-4   border-red-500 rounded-b text-red-900 px-4 py-3 shadow-md"
        role="alert" v-if="RoleRegister.error" :class="{ ' translate-x-full': RoleRegister.close }">
  
        <div class="flex">
          <div class="py-1"><svg class="fill-current h-6 w-6 text-red-500 mr-4" xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20">
              <path
                d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 11V9h2v6H9v-4zm0-6h2v2H9V5z" />
            </svg></div>
          <div>
            <p class="font-bold">Đăng ký thất bại</p>
            <p class="text-sm">Tài khoản với địa chỉ email hoặc tên người dùng bạn đã cung cấp đã tồn tại trong hệ thống của chúng tôi.</p>
          </div>
        </div>
      </div>
    </div> -->
  </main>
</template>

<script>
  import { RouterLink, createWebHistory } from 'vue-router';
  import { useUserStore } from '../../../stores/Account.stores';
  // import { Router } from 'vue-router';
  export default {
    data() {
      return {
        inputFullName: {
          value: '',
        },
        inputUserName: {
          value: '',
          error: false,
        },
        inputEmail: {
          value: '',
          error: false,
        },
        inputPhoneNumber: {
          value: '',
          error: false,
        },
        inputPassword: {
          value: '',
          error: false,
        },
        enterPassword: {
          value: '',
          error: false,
        },
        role: 'user',
        provider: 'local',
        verified: true,
        enable: 1,
        RoleRegister: {
          error: false,
          success: false,
          close: true,
        },
        teamplateToast: {
          error: `<div
        class=" show-alert transform transition-transform error duration-1000 py-3 ease-in-out bg-red-100 rounded-b text-red-900 px-4 shadow-md"
        role="alert" v-if="RoleRegister.error" :class="{ 'error': RoleRegister.close }">
        <div class="flex">
          <div class="py-1"><svg class="fill-current h-6 w-6 text-red-500 mr-4" xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20">
              <path
                d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 11V9h2v6H9v-4zm0-6h2v2H9V5z" />
            </svg></div>
          <div>
            <p class="font-bold">Đăng ký thất bại</p>
            <p class="text-sm">Tài khoản với địa chỉ email hoặc tên người dùng bạn đã cung cấp đã tồn tại trong hệ thống của chúng tôi.</p>
          </div>
        </div>
      </div>`,
          success: `<div
        class="show-alert transform transition-transform duration-1000 py-3 ease-in-out bg-teal-100 success rounded-b text-teal-900 px-4 py-3 shadow-md"
        role="alert" v-if="RoleRegister.success" :class="{ ' translate-x-full': RoleRegister.close }">
        <div class="flex">
          <div class="py-1"><svg class="fill-current h-6 w-6 text-teal-500 mr-4" xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20">
              <path
                d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 11V9h2v6H9v-4zm0-6h2v2H9V5z" />
            </svg></div>
          <div>
            <p class="font-bold">Đăng ký thành công</p>
            <p class="text-sm">Chào mừng bạn đã đăng ký thành công tài khoản trên Cỏ Studio</p>
          </div>
        </div>
      </div>`,
        },
      };
    },
    methods: {
      CheckInputUserName() {
        const usernamePattern = /^[^*#&]+$/;
        if (!usernamePattern.test(this.inputUserName.value)) {
          this.inputUserName.error = true;
          this.inputUsername.value = '';
        } else {
          this.inputUserName.error = false;
        }
      },
      CheckInputEmail() {
        const emailPattern = /^[a-zA-Z0-9._%+-]+@(gmail\.com|gmail\.com\.vn)$/;
        if (!emailPattern.test(this.inputEmail.value)) {
          this.inputEmail.error = true;
          this.inputEmail.value = '';
        } else {
          this.inputEmail.error = false;
        }
      },
      CheckInputPhoneNumber() {
        const phoneNumberPattern = /^(0[1-9][0-9]{8,9})$/;
        if (!phoneNumberPattern.test(this.inputPhoneNumber.value)) {
          this.inputPhoneNumber.error = true;
          this.inputPhoneNumber.value = '';
        } else {
          this.inputPhoneNumber.error = false;
        }
      },
      CheckInputPassword() {
        const passwordPattern = /^(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{8,}$/;
        if (!passwordPattern.test(this.inputPassword.value)) {
          this.inputPassword.error = true;
          this.inputPassword.value = '';
        } else {
          this.inputPassword.error = false;
        }
      },
      CheckEnterPassword() {
        if (this.enterPassword.value == this.inputPassword.value) {
          this.enterPassword.error = false;
        } else {
          this.enterPassword.error = true;
        }
      },
      CheckForm() {
        this.CheckInputEmail();
        this.CheckInputPhoneNumber();
        this.CheckInputPassword();
        this.CheckEnterPassword();
        this.CheckInputUserName();
      },
      async Run() {
        this.CheckForm();
        if (
          this.inputEmail.error === false &&
          this.inputUserName.error === false &&
          this.inputPassword.error === false &&
          this.inputPhoneNumber.error === false &&
          this.enterPassword.error === false
        ) {
          await this.RegisterS();
          this.ShowToast();
          this.CheckStatus();
        }
      },
      CheckStatus() {
        if (this.RoleRegister.success == true) {
          setTimeout(() => {
            this.$router.push({ path: '/' });
          }, 2000);
        }
      },
      async RegisterS() {
        await useUserStore().Register(
          this.inputFullName.value,
          this.inputUserName.value,
          this.inputEmail.value,
          this.inputPassword.value,
          this.inputPhoneNumber.value,
          this.role,
          this.provider,
          this.verified,
          this.enable,
        );
        this.RoleRegister.error = useUserStore().RoleRegister.error;
        this.RoleRegister.success = useUserStore().RoleRegister.success;
      },
      ShowToast() {
        let toastBox = document.getElementById('toastBox');
        let toast = document.createElement('div');

        if (this.RoleRegister.error == true) {
          toast.innerHTML = this.teamplateToast.error;
        }

        if (this.RoleRegister.success == true) {
          toast.innerHTML = this.teamplateToast.success;
          this.RoleRegister.success == false;
        }
        toastBox.appendChild(toast);
        setTimeout(() => {
          toast.remove();
        }, 3000);
      },
    },

    mounted() {},
  };
</script>

<style scoped>
  .background-img {
    background: url('../../../../public/assets/images/SignUp/Image.png') no-repeat center/cover;
  }

  .alert {
    transform: translateX(22px);
  }

  .label-input {
    top: -2px;
    background-color: white;
    padding: 0 3px;
  }
</style>
