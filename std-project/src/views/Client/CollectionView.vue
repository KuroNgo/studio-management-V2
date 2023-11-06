<template>
  <main class="mx-36 mt-20 mb-40">
    <div class="">
      <ul class="grid grid-cols-3 grid-flow-row gap-4">
        <div class="" v-for="(x, i) in dataImage">
          <div class="h-96 rounded-3xl mb-5 relative" :style="{ background: 'url(' + x + ') center/cover no-repeat' }">
            <svg
              class="absolute right-3 bottom-6"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24.004"
              viewBox="0 0 24 24.004">
              <path
                id="Icon_awesome-search"
                data-name="Icon awesome-search"
                d="M23.674,20.753,19,16.079a1.124,1.124,0,0,0-.8-.328h-.764a9.746,9.746,0,1,0-1.688,1.688V18.2a1.124,1.124,0,0,0,.328.8l4.674,4.674a1.12,1.12,0,0,0,1.589,0l1.327-1.327A1.13,1.13,0,0,0,23.674,20.753Zm-13.923-5a6,6,0,1,1,6-6A6,6,0,0,1,9.751,15.751Z" />
            </svg>
          </div>
          <div class="whitespace-normal text-center">
            <h1 class="text-gray-400 Heading5">{{ name[i] }}</h1>
            <p class="text-gray-300">22/12/2020</p>
          </div>
        </div>
      </ul>
    </div>
  </main>
</template>

<script>
  import CardImageView from '../../components/ContainerImages/CardImageView.vue';

  export default {
    components: {
      CardImageView,
    },
    data() {
      return {
        dataImage: [
          '../../../../public/assets/images/Homepage/anhcuoi1.png',
          '../../../../public/assets/images/Homepage/ImageSlider.png',
          '../../../../public/assets/images/Homepage/anhcuoi2.jpg',
          '../../../../public/assets/images/Homepage/anhcuoi3.jpg',
          '../../../../public/assets/images/Homepage/anhchupdangoai1.jpg',
          '../../../../public/assets/images/Homepage/anhchupdangoai2.jpg',
          '../../../../public/assets/images/Homepage/anhchupdangoai3.jpg',
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
          '../../../../public/assets/images/Homepage/quayvideo2.jpg',
        ],
        name: [
          'Hoàng & Dinh',
          'Phong & Linh',
          "Hong's Family",
          'Dinh & Vy',
          'Duc & Alexa',
          'Quynh & Dinh',
          'Son & Thai',
          'Trang & Quan',
          'Trang & Quan',
        ],
      };
    },
  };
</script>

<style scoped></style>
