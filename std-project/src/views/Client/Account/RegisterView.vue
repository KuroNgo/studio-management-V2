<template>
  <main class=" py-20 h-screen">
    <div class="flex h-full w-3/4 mx-auto border border-solid border-gray-500">
      <div class="w-2/4 mt-auto mb-auto px-32">
        <h4 class="Heading4 uppercase text-center mb-11">ĐĂNG KÝ</h4>
        <!-- <form>
          <label class="block">
            <span class="block text-sm font-medium text-slate-700">Email</span>
            <input type="email" class="peer" />
            <p class="mt-2 invisible peer-invalid:visible text-pink-600 text-sm">
              Please provide a valid email address.
            </p>
          </label>
        </form> -->
        
          <div class=" relative mb-3">
            <input type="text" v-model="inputUserName.value" placeholder="Không chứa các ký tự *,#,&" pattern="^[^*#&]+$"
            :class="{' outline-red-700' : inputUserName.error}"
              class="form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption  outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full">
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2  left-3 ">Tên tài
              khoản</p>
            <p class="invisible peer-invalid:visible text-pink-600 Caption pl-3">
              Tên tài khoảng không chứa các ký tự *,#,&
            </p>
          </div>

          <div class=" relative mb-3">
            <input type="text" v-model="inputEmail.value" pattern="^[a-zA-Z0-9._%+-]+@(gmail\.com|gmail\.com\.vn)$"
              placeholder="Địa chỉ email của bạn"
              :class="{' outline-red-700' : inputEmail.error}"
              class="form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption  outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full">
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2 left-3 ">Địa chỉ
              email</p>
            <p class="invisible peer-invalid:visible text-pink-600 Caption pl-3">
              Vui lòng cung cấp một địa chỉ email hợp lệ
            </p>
          </div>

          <div class=" relative mb-3">
            <input type="text" v-model="inputPhoneNumber.value" placeholder="Số điện thoại của bạn" pattern="^0\d{9,10}$"
            :class="{' outline-red-700' : inputPhoneNumber.error}"
              class="form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption  outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full">
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2  left-3 ">Số điện
              thoại</p>
            <p class="invisible peer-invalid:visible text-pink-600 Caption pl-3">
              Vui lòng cung cấp số điện thoại hợp lệ
            </p>
          </div>

          <div class=" relative mb-3">
            <input type="text" v-model="inputPassword.value" pattern="^(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{8,}$"
              placeholder="Chứa ít nhất 8 ký tự bao gồm ( chứ viết hoa, chữ số và một ký tự đặc biệt )"
              :class="{' outline-red-700' : inputPassword.error}"
              class=" form-input peer mb-2 placeholder:text-gray-400 placeholder:Caption  outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full">
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2  left-3 ">Mật khẩu
            </p>
            <p class="invisible peer-invalid:visible text-pink-600 Caption pl-3">
              Chứa ít nhất 8 ký tự bao gồm ( chứ viết hoa, chữ số và một ký tự đặc biệt )
            </p>
          </div>

          <div class=" relative mb-7">
            <input type="text" v-model="enterPassword.value" placeholder="Nhập lại mật khẩu bạn đã nhập"
            :class="{' outline-red-700' : enterPassword.error}"
              class="form-input mb-2 placeholder:text-gray-400 placeholder:Caption  outline outline-1 outline-SupportColor2 rounded-xl p-3 w-full">
            <p class="label-input Overline text-gray-400 absolute top-1/2 transform -translate-y-1/2  left-3 ">Nhập lại
              mật khẩu</p>
          </div>

          <button @click="CheckForm()"
            class=" text-DomlantColor Button p-3 bg-AccentColor w-full rounded-xl mb-11">Đăng
            nhập</button>
          <p class="Body2 text-center ">Đã có tài khoản ? <RouterLink to="/login"><strong
                class=" underline text-SecondColor">Đăng nhập</strong> </RouterLink>
          </p>
        
      </div>
      <div class="background-img w-2/4 h-full"></div>
    </div>
  </main>
</template>

<script>
import { RouterLink } from 'vue-router';
export default {
  data() {
    return {
      inputUserName: {
        value:'',
        error:false,
      },
      inputEmail: {
        value:'',
        error:false,
      },
      inputPhoneNumber: {
        value:'',
        error:false,
      },
      inputPassword: {
        value:'',
        error:false,
      },
      enterPassword:{
        value:'',
        error:false,
      },
      alertError:false,
    }
  },
  methods: {
    CheckInputUserName() {
      const usernamePattern = /^[^*#&]+$/;
      if (!usernamePattern.test(this.inputUserName.value)) {
        this.inputUserName.error=true;
        this.inputUsername.value = "";
      }else{
        this.inputUserName.error=false;
      }
    },
    CheckInputEmail() {
      const emailPattern = /^[a-zA-Z0-9._%+-]+@(gmail\.com|gmail\.com\.vn)$/;
      if (!emailPattern.test(this.inputEmail.value)) {
        this.inputEmail.error=true;
        this.inputEmail.value = "";
      }else{
        this.inputEmail.error=false;
      }
    },
    CheckInputPhoneNumber() {
      const phoneNumberPattern = /^(0[1-9][0-9]{8,9})$/;
      if (!phoneNumberPattern.test(this.inputPhoneNumber.value)) {
        this.inputPhoneNumber.error=true;
        this.inputPhoneNumber.value = "";
      }else{
        this.inputPhoneNumber.error=false;
      }
    },
    CheckInputPassword() {
      const passwordPattern = /^(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{8,}$/;
      if (!passwordPattern.test(this.inputPassword.value)) {
        this.inputPassword.error=true;
        this.inputPassword.value = "";
      }else{
        this.inputPassword.error=false;
      }
    },
    CheckEnterPassword(){
      if (this.enterPassword.value == this.inputPassword.value) {
        this.enterPassword.error=false;
      }else{
        this.enterPassword.error=true;
      }
    },
    CheckForm(){
      this.CheckInputEmail();
      this.CheckInputPhoneNumber();
      this.CheckInputPassword();
      this.CheckEnterPassword();
      this.CheckInputUserName();
      if (this.inputEmail.error == this.inputUserName.error == this.inputPassword.error == this.inputPhoneNumber.error == this.enterPassword.error == false) {
        alert("thành công")
      }
    }

  },
  mounted() {

  }
}
</script>

<style scoped>
.background-img {
  background: url("../../../../public/assets/images/Login/Image-1.png") no-repeat center/cover;

}
.border-error{

}
.label-input {
  top: -2px;
  background-color: white;
  padding: 0 3px;
}
</style>