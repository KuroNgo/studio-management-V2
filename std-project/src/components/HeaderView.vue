<template>
    <main>
        <div class=" px-36 py-2 bg-slate-400 flex justify-between items-center">
            <div class="flex items-center gap-8">
                <form action="" :class="{ 'border-form': showBorder }"
                    class="flex items-center bg-white rounded-xl py-1 px-2 gap-2 w-72" onsubmit="return false">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="21" viewBox="0 0 20 21" fill="none">
                        <path
                            d="M19.7248 17.291L15.8298 13.4C15.7425 13.3131 15.6389 13.2442 15.525 13.1974C15.4111 13.1505 15.289 13.1266 15.1658 13.127H14.5298C15.8017 11.4972 16.404 9.4436 16.2139 7.38501C16.0239 5.32642 15.0558 3.41784 13.507 2.04842C11.9583 0.678998 9.94554 -0.0481282 7.87917 0.0153011C5.8128 0.0787304 3.84843 0.927938 2.3866 2.38977C0.924765 3.8516 0.0755566 5.81597 0.0121273 7.88234C-0.051302 9.94872 0.675824 11.9615 2.04524 13.5102C3.41466 15.059 5.32325 16.027 7.38184 16.2171C9.44043 16.4071 11.494 15.8048 13.1238 14.533V15.17C13.1234 15.2932 13.1474 15.4152 13.1942 15.5292C13.241 15.6431 13.3099 15.7467 13.3968 15.834L17.2908 19.728C17.3776 19.8152 17.4808 19.8844 17.5944 19.9316C17.708 19.9788 17.8298 20.0031 17.9528 20.0031C18.0758 20.0031 18.1977 19.9788 18.3113 19.9316C18.4249 19.8844 18.528 19.8152 18.6148 19.728L19.7198 18.623C19.8971 18.447 19.9972 18.2078 19.9981 17.958C19.9991 17.7082 19.9008 17.4683 19.7248 17.291ZM8.12481 13.123C7.13591 13.123 6.16921 12.8297 5.34696 12.2803C4.52472 11.7309 3.88385 10.95 3.50541 10.0364C3.12698 9.12277 3.02796 8.11744 3.22089 7.14753C3.41381 6.17763 3.89002 5.28672 4.58928 4.58745C5.28854 3.88819 6.17946 3.41199 7.14936 3.21906C8.11927 3.02613 9.1246 3.12515 10.0382 3.50359C10.9519 3.88203 11.7328 4.52289 12.2822 5.34514C12.8316 6.16738 13.1248 7.13408 13.1248 8.12299C13.1249 8.77976 12.9957 9.43013 12.7444 10.0369C12.4931 10.6437 12.1248 11.1951 11.6603 11.6595C11.1959 12.1239 10.6446 12.4923 10.0378 12.7436C9.43096 12.9949 8.78059 13.1241 8.12381 13.124L8.12481 13.123Z"
                            fill="#CCCCCC" />
                    </svg>
                    <input type="text" @focus="showBorder = true" @blur="showBorder = false" class="  outline-none w-full"
                        placeholder="Tìm kiếm">
                </form>
                <ul class="flex gap-8">
                    <RouterLink to="/"><li>Trang chủ</li></RouterLink>
                    <RouterLink to="/services"><li>Dịch vụ</li></RouterLink>
                    <RouterLink to="/product"><li>Thư viện</li></RouterLink>
                    <RouterLink to="/about"><li>Giới thiệu</li></RouterLink>
                    <RouterLink to="/contact"><li>Liên hệ</li></RouterLink>
                </ul>
            </div>
            <button class=" bg-white px-6 py-1 rounded-xl">Đăng nhập</button>
        </div>
    </main>
</template>

<script>

export default {
    data() {
        return {
            showBorder: false
        }
    },
    methods: {

    }
}

</script>

<style scoped>
.border-form {
    outline: 2px solid #000000;
}
</style>