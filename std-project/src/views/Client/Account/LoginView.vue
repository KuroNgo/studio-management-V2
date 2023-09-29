<template>
  <div class="login-popup" v-if="showPopup" ref="loginPopup">
    <div class="overlay" @click="closePopup"></div>
    <div class="popup">
      <div class="left-section">
        <img src="/assets/images/login.jpg" alt="Login Image" />
      </div>
      <div class="right-section">
        <h2>Đăng nhập</h2>
        <h3>Chào mừng bạn !</h3>
        <button
          class="mx-auto flex items-center bg-white border border-gray-300 rounded-lg shadow-md px-6 py-2 text-sm font-medium text-gray-800 hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
        >
          <svg
            class="h-6 w-6 mr-2"
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            width="800px"
            height="800px"
            viewBox="-0.5 0 48 48"
            version="1.1"
          >
            <title>Google-color</title>
            <desc>Created with Sketch.</desc>
            <defs></defs>
            <g
              id="Icons"
              stroke="none"
              stroke-width="1"
              fill="none"
              fill-rule="evenodd"
            >
              <g id="Color-" transform="translate(-401.000000, -860.000000)">
                <g id="Google" transform="translate(401.000000, 860.000000)">
                  <path
                    d="M9.82727273,24 C9.82727273,22.4757333 10.0804318,21.0144 10.5322727,19.6437333 L2.62345455,13.6042667 C1.08206818,16.7338667 0.213636364,20.2602667 0.213636364,24 C0.213636364,27.7365333 1.081,31.2608 2.62025,34.3882667 L10.5247955,28.3370667 C10.0772273,26.9728 9.82727273,25.5168 9.82727273,24"
                    id="Fill-1"
                    fill="#FBBC05"
                  ></path>
                  <path
                    d="M23.7136364,10.1333333 C27.025,10.1333333 30.0159091,11.3066667 32.3659091,13.2266667 L39.2022727,6.4 C35.0363636,2.77333333 29.6954545,0.533333333 23.7136364,0.533333333 C14.4268636,0.533333333 6.44540909,5.84426667 2.62345455,13.6042667 L10.5322727,19.6437333 C12.3545909,14.112 17.5491591,10.1333333 23.7136364,10.1333333"
                    id="Fill-2"
                    fill="#EB4335"
                  ></path>
                  <path
                    d="M23.7136364,37.8666667 C17.5491591,37.8666667 12.3545909,33.888 10.5322727,28.3562667 L2.62345455,34.3946667 C6.44540909,42.1557333 14.4268636,47.4666667 23.7136364,47.4666667 C29.4455,47.4666667 34.9177955,45.4314667 39.0249545,41.6181333 L31.5177727,35.8144 C29.3995682,37.1488 26.7323182,37.8666667 23.7136364,37.8666667"
                    id="Fill-3"
                    fill="#34A853"
                  ></path>
                  <path
                    d="M46.1454545,24 C46.1454545,22.6133333 45.9318182,21.12 45.6113636,19.7333333 L23.7136364,19.7333333 L23.7136364,28.8 L36.3181818,28.8 C35.6879545,31.8912 33.9724545,34.2677333 31.5177727,35.8144 L39.0249545,41.6181333 C43.3393409,37.6138667 46.1454545,31.6490667 46.1454545,24"
                    id="Fill-4"
                    fill="#4285F4"
                  ></path>
                </g>
              </g>
            </g>
          </svg>
          <span>Continue with Google</span>
        </button>
        <h4>Hoặc đăng nhập với</h4>

        <form @submit.prevent="login">
          <div class="input-group">
            <div class="account-box">
              <label for="account" class="block text-gray-600">Tài khoản</label>
              <input
                type="text"
                id="username"
                v-model="username"
                placeholder="Enter your username"
              />
            </div>

            <div class="account-box">
              <label for="account" class="block text-gray-600">Mật khẩu</label>
              <input
                type="password"
                id="password"
                v-model="password"
                placeholder="Enter your password"
              />
            </div>

            <label class="radio-label">
      <input type="radio" class="custom-radio" name="radio-group">
      <span class="label-text">Nhớ mật khẩu</span>
      <a href="#" class="my-link">Quên mật khẩu ?</a>
    </label>
          </div>
          <button type="submit" class="mx-auto">Đăng nhập</button>
          <p>Không có tài khoản ? <a href="#">Đăng ký</a></p>      
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showPopup: false,
      username: "",
      password: "",
    };
  },
  methods: {
    togglePopup() {
      this.showPopup = !this.showPopup;
    },
    login() {
      // Thực hiện xử lý đăng nhập ở đây (có thể gửi dữ liệu đăng nhập đến máy chủ)
      // Sau khi đăng nhập thành công, bạn có thể đóng popup.
      this.closePopup();
    },
    closePopup() {
      this.showPopup = false;
      this.username = ""; // Xóa giá trị tài khoản
      this.password = ""; // Xóa giá trị mật khẩu
    },
  },
};
</script>

<style scoped>
.my-link{
    color:#AADFD1;
    padding-top: 20px;
    margin-left: 100px;
}
.popup {
  width: 1184px; /* Tổng chiều rộng 1184 px */
  height: 650px;
  background: white;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  z-index: 2;
  display: flex;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
a{
    color:#AADFD1;
}
p{
    text-align: center;
    margin-top: 20px;

}

.left-section {
  flex: 1;
  padding: 0px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f0f0f0;
}

.left-section img {
  width: 592px;
  height: 650px;
  object-fit: cover;
}

.right-section {
  flex: 1; /* Phần bên phải sẽ mở rộng nhiều hơn để chứa các trường đăng nhập */
  padding: 20px;
}
.right-section h2 {
  font-size: 34px;
  text-align: center; /* Căn giữa ngang */
  margin-top: 0; /* Loại bỏ khoảng trống trên đầu h2 nếu có */
  margin-bottom: 20px; /* Tạo khoảng trống phía dưới h2 để tạo khoảng cách */
}

.right-section h3 {
  font-size: 16px;
  color: #5a8d80;
  text-align: center; /* Căn giữa ngang */
  margin-top: 0; /* Loại bỏ khoảng trống trên đầu h2 nếu có */
  margin-bottom: 20px; /* Tạo khoảng trống phía dưới h2 để tạo khoảng cách */
}
.right-section h4 {
  padding-top: 20px;
  font-size: 12px;
  color: #999999;
  text-align: center; /* Căn giữa ngang */
  margin-top: 0; /* Loại bỏ khoảng trống trên đầu h2 nếu có */
  margin-bottom: 20px; /* Tạo khoảng trống phía dưới h2 để tạo khoảng cách */
}
/* Stylize form elements */
input[type="text"] {
  width: 374px;
  height: 48px;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 3px;
  display: block; /* Biến input thành phần tử block để chiếm toàn bộ chiều ngang */
  margin: 0 auto;
}

input[type="password"] {
  width: 374px;
  height: 48px;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 3px;
  display: block; /* Biến input thành phần tử block để chiếm toàn bộ chiều ngang */
  margin: 0 auto;
}
 input[type="radio"]
 {
    display: block; /* Biến input thành phần tử block để chiếm toàn bộ chiều ngang */
  margin-top: 20px;
    padding-left: 10px;
 }
label {
  padding-left: 100px;
}
span.label-text {
    margin-top: 20px;
}

button[type="submit"] {
  margin-top: 20px;
  background: var(--accent-color) 0% 0% no-repeat padding-box;
  box-shadow: 3px 3px 1px var(--unnamed-color-000000);
  background: #35665b 0% 0% no-repeat padding-box;
  box-shadow: 3px 3px 1px #000000;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 375px;
  height: 42px;
  border-radius: 12px;
  opacity: 1;
}
button{
    margin-top: 20px;
  background: var(--accent-color) 0% 0% no-repeat padding-box;
  box-shadow: 3px 3px 1px var(--unnamed-color-000000);
  background: #ffffff;
  box-shadow: 3px 3px 1px #ffffff;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 375px;
  height: 42px;
  border-radius: 12px;
  margin-left: 100px;
  opacity: 1;
}


button[type="submit"]:hover {
  background-color: #0056b3;
}
/* Định dạng CSS cho phần input */
.custom-input {
  border: none; /* Loại bỏ đường viền mặc định */
  border-bottom: 1px solid #ccc; /* Đường viền dưới cùng */
  padding: 10px; /* Thêm padding để tạo khoảng cách */
  outline: none; /* Loại bỏ đường viền khi tập trung vào input */
  width: 100%; /* Để input rộng hết phần tử cha */
}

/* Định dạng CSS cho phần account-box */
.account-box {
  position: relative;
}
.custom-radio {
    appearance: none; /* Loại bỏ giao diện mặc định */
    width: 20px;
    height: 20px;
    border: 1px solid #707070; /* Màu xanh */
    border-radius: 50%; /* Để nút radio tròn */
    outline: none;
    margin-right: 10px;
    background: #FFFFFF 0% 0% no-repeat padding-box;
  }

  /* Tạo kiểu dáng nút radio khi được chọn */
  .custom-radio:checked {
    background:#ffffff;
    border: 1px solid#AADFD1;;
  }

  /* Tạo kiểu dáng cho nhãn của nút radio */
  .radio-label {
    display: flex;
    align-items: center;
  }

  /* Định dạng văn bản bên cạnh nút radio */
  .label-text[data-v-2b619d6d] {
    color: #AADFD1;
}

/* Định dạng CSS cho đường viền */
</style>
