<template>
  <div class="login-popup">
    <div class="overlay" @click="closePopup"></div>
    <div class="popup">
      <div class="left-section">
        <h2>Đăng ký</h2>
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
              <label for="email" class="block text-gray-600"
                >Địa chỉ email</label>
              <input
                type="text"
                id="email"
                v-model="email"
                placeholder="Địa chỉ email của bạn"
              />
            </div>
            <div class="account-box">
              <label for="phone" class="block text-gray-600"
                >Số điện thoại</label
              >
              <input
                type="text"
                id="phone"
                v-model="phone"
                placeholder="Không bắt buộc"
              />
            </div>
            <div class="account-box">
              <label for="password" class="block text-gray-600">Mật khẩu</label>
              <input
                type="password"
                id="password"
                v-model="password"
                placeholder="Chứa ít nhất 8 ký tự và"
              />
            </div>
            <div class="account-box">
              <label for="passwordConfirm" class="block text-gray-600"
                >Nhập lại mật khẩu</label
              >
              <input
                type="password"
                id="passwordConfirm"
                v-model="passwordConfirm"
                placeholder="Nhập lại mật khẩu bạn đã nhập"
              />
            </div>
          </div>
          <button type="submit" class="mx-auto">Đăng nhập</button>
          <p>Không có tài khoản ? <a href="#">Đăng ký</a></p>      
        </form>
       
      </div>
      <div class="right-section">
        <img src="/assets/images/login.jpg" alt="Login Image" />
      </div>
    </div>
  </div>
</template>


<script></script>

<style scoped>
.my-link {
  color: #aadfd1;
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
a {
  color: #aadfd1;
}
p {
  text-align: center;
  margin-top: 20px;
}

.right-section {
  flex: 1;
  padding: 0px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f0f0f0;
}

.right-section img {
  width: 592px;
  height: 650px;
  object-fit: cover;
}

.left-section {
  flex: 1; /* Phần bên phải sẽ mở rộng nhiều hơn để chứa các trường đăng nhập */
  padding: 20px;
}
.left-section h2 {
  font-size: 34px;
  text-align: center; /* Căn giữa ngang */
  margin-top: 0; /* Loại bỏ khoảng trống trên đầu h2 nếu có */
  margin-bottom: 20px; /* Tạo khoảng trống phía dưới h2 để tạo khoảng cách */
}

.left-section h3 {
  font-size: 16px;
  color: #5a8d80;
  text-align: center; /* Căn giữa ngang */
  margin-top: 0; /* Loại bỏ khoảng trống trên đầu h2 nếu có */
  margin-bottom: 20px; /* Tạo khoảng trống phía dưới h2 để tạo khoảng cách */
}
.left-section h4 {
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
input[type="radio"] {
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
button {
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
  padding: 8px;
}
.custom-radio {
  appearance: none; /* Loại bỏ giao diện mặc định */
  width: 20px;
  height: 20px;
  border: 1px solid #707070; /* Màu xanh */
  border-radius: 50%; /* Để nút radio tròn */
  outline: none;
  margin-right: 10px;
  background: #ffffff 0% 0% no-repeat padding-box;
}

/* Tạo kiểu dáng nút radio khi được chọn */
.custom-radio:checked {
  background: #ffffff;
  border: 1px solid#AADFD1;
}

/* Tạo kiểu dáng cho nhãn của nút radio */
.radio-label {
  display: flex;
  align-items: center;
}

/* Định dạng văn bản bên cạnh nút radio */
.label-text[data-v-2b619d6d] {
  color: #aadfd1;
}

/* Định dạng CSS cho đường viền */
</style>
