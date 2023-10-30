<template>
    <div>
   <div class="centered-image">
    
        <img src="/assets/images/success.png" alt="Login Image" />
      </div>
      <div class="centered-content">
    <h1>All Done</h1>
    <p>Your password has been changed!</p>
    <button onclick="backToLogin()" >
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20">
        <path id="arrow_back_FILL0_wght400_GRAD0_opsz24" d="M164.781-788.75l7,7L170-780l-10-10,10-10,1.781,1.75-7,7H180v2.5Z" transform="translate(-160 800)" fill="#999"/>
      </svg>
      Back to Login
    </button>
  </div>

    </div>
</template>

<script>
export default {

}
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
p{
    margin-right: 156px;
    color: #999999;
}
h1{
   font-weight: bold;
   margin-right: 350px;
}
    button {
    
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: #5A8D80;
      color: #F8F8F8;
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
</style>