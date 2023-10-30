<template>
  <div>
    <div class="centered-image">
      <img src="/assets/images/sendcode.png" alt="Login Image" />
    </div>
    <div class="centered-content">
      <h1>Create New Password</h1>
      <p>Let's create a new password</p>
      <div class="input-group">
        <div class="account-box">
          <label class="lable">New Password</label>
          <input
            type="text"
            placeholder="At least 8 chacracters and one upper case "
          />
        </div>
        <div class="account-box">
          <label class="lable">Reset Password</label>
          <input
            type="text"
            placeholder="New password much match your new password "
          />
        </div>
        <RouterLink to="/"><button >Confirm</button></RouterLink> 
        <button onclick="backToLogin()" class="white-button">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <path
              id="arrow_back_FILL0_wght400_GRAD0_opsz24"
              d="M165.738-786.5l8.4,8.4L172-776l-12-12,12-12,2.137,2.1-8.4,8.4H184v3Z"
              transform="translate(-160 800)"
            />
          </svg>

          <h4>Back to login</h4>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style>
.centered-image {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%; /* Đảm bảo hình ảnh chiếm toàn bộ chiều cao của phần tử cha */
}
.centered-content {
  display: flex;
  flex-direction: column;
  align-items: center; /* Căn giữa theo chiều ngang */
  height: 100vh; /* Đảm bảo chiều cao của phần tử là 100% chiều cao của cửa sổ trình duyệt */
  text-align: center; /* Căn giữa văn bản theo chiều ngang */
}

.centered-content button {
  margin-top: 20px; /* Khoảng cách từ nút đến các phần tử trước đó */
}
.centered-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}
p {
  margin-right: 195px;
  color: #999999;
  padding-bottom: 15px;
}
h1 {
  font-weight: bold;
  margin-right: 240px;
}
h4 {
  color: #0f0402;
  font-weight: bold;
}
.white-button {
  background-color: #f8f8f8;
  border: 2px solid #aadfd1; /* Thêm viền xanh #81B5A8 */
  border-radius: 10px;
}
button {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #5a8d80;
  color: #f8f8f8;
  width: 414px; /* Độ rộng 441px */
  height: 56px;
  border: none;
  cursor: pointer;
  padding: 0;
  border-radius: 10px;
  margin-top: 20px; /* Điều chỉnh khoảng cách giữa nút và nội dung trên trang */
}

svg {
  width: 20px;
  height: 20px;
  margin-right: 20px; /* Khoảng cách giữa biểu tượng và văn bản */
}
.account-box {
  position: relative;
  padding-bottom: 20px;
}
input[type="text"] {
  width: 414px; /* Độ rộng 441px */
  height: 56px;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 10px;
  display: block; /* Biến input thành phần tử block để chiếm toàn bộ chiều ngang */
  margin: 0 auto;
} 
label{
  margin-right:290px ;
}
</style>
