<template>
    <main>
        <div class="mb-28">
            <SlideHomeView />
        </div>
        <div class=" px-36">
            <h4 class="Heading4 mb-10">Đăng ký dịch vụ tại Cỏ Studio như thế nào ?</h4>
            <div class="flex flex-wrap justify-center gap-4 lg:flex-nowrap">
                <div class=" w-96 ">
                    <div class="" style="height: 500px;">
                        <img class="w-full h-full object-cover" src="../../../public/assets/images/BusLogin.png" alt="">
                    </div>
                    <h5 class="text-center Heading5 mb-3">Bước 1</h5>
                    <p>Đăng ký dịch vụ của studio thông qua form đăng ký.</p>
                </div>
                <div class=" w-96 ">
                    <div class="" style="height: 500px;">
                        <img class="w-full h-full object-cover" src="../../../public/assets/images/viewinbus.png" alt="">
                    </div>
                    <h5 class="text-center Heading5 mb-3">Bước 2</h5>
                    <p>Đến trực tiếp studio để được tư vấn và xem trực tiếp các sản phẩm của studio thực hiện, ký hợp đồng
                        và chuẩn bị các bước để thực hiện chụp, quay.</p>
                </div>
                <div class=" w-96 ">
                    <div class="" style="height: 500px;">
                        <img class="w-full h-full object-cover" src="../../../public/assets/images/BusLogin.png" alt="">
                    </div>
                    <h5 class="text-center Heading5 mb-3">Bước 3</h5>
                    <p>Quay, chụp ảnh chọn ảnh sau khi chụp.</p>
                </div>
                <div class=" w-96 ">
                    <div class="" style="height: 500px;">
                        <img class="w-full h-full object-cover" src="../../../public/assets/images/BusLogin.png" alt="">
                    </div>
                    <h5 class="text-center Heading5 mb-3">Bước 4</h5>
                    <p>Nhận file đã chỉnh sửa & duyệt in album để nhận sản phẩm cuối cùng.</p>
                </div>
            </div>
        </div>
        <div class=" px-36 mb-10">
            <div class="flex justify-between mb-10">
                <h4>Chụp Ảnh Cưới</h4>
                <a class="Button underline" href="">Xem thêm</a>
            </div>
        </div>

        <div class=" px-36 mb-10">
            <div class="flex justify-between mb-10">
                <h4>Quay Video Cưới</h4>
                <a class="Button underline" href="">Xem thêm</a>
            </div>
        </div>

        <div class=" px-36 mb-10">
            <div class="flex justify-between mb-10">
                <h4>Quay Video</h4>
                <a class="Button underline" href="">Xem thêm</a>
            </div>
        </div>
    </main>
</template>

<script>
import SlideHomeView from '../../components/Slides/SlidesHomePage/SlideHomeView.vue';
export default{
    components:{
        SlideHomeView
    }
}
</script>

<style scoped></style>