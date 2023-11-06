<template>
  <main class="px-36 mb-40">
    <h1 class="my-4">
      Dịch vụ /
      <strong class="heading4">Quay phim TVC, phim giới thiệu doanh nghiệp</strong>
    </h1>

    <div class="flex justify-between">
      <ul class="w-4/6 grid grid-cols-4 gap-4">
        <CardImageView v-for="x in dataImage" :url-image="x" />
      </ul>
      <div class="w-1/3">
        <div
          class="sticky top-20 mx-auto px-4 pt-4 pb-14 border border-solid border-BlackColor rounded-2xl"
          style="width: 77%">
          <h1 class="Heading4 text-SecondColor mb-6 text-center">Đặt Lịch</h1>
          <p class="Body1 mb-2">
            Bạn có thể đặt lịch trước và có thể đặt cọc trước nếu muốn, nếu như bạn đặt cọc trước có thể được ưu tiên
            sắp xếp lịch trước.
          </p>
          <p class="text-gray-300 mb-7">Studio hoạt động từ 7h đến 22h.</p>
          <p class="mb-4">Lịch</p>
          <VueDatePicker
            class="w-full mb-6"
            :inline="true"
            v-model="date"
            :min-date="new Date()"
            :enable-time-picker="false"
            placeholder="Select Date"></VueDatePicker>
          <button class="Body1 text-DomlantColor mb-6 bg-AccentColor w-full rounded-lg p-2">Đặt lịch</button>
          <button class="Body1 text-black bg-DomlantColor border border-solid border-BlackColor w-full rounded-lg p-2">
            Đặt cọc
          </button>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
  import VueDatePicker from '@vuepic/vue-datepicker';
  import '@vuepic/vue-datepicker/dist/main.css';
  import CardImageView from '../../../components/ContainerImages/CardImageView.vue';

  export default {
    components: { VueDatePicker, CardImageView },
    data() {
      return {
        date: null,
        dataImage: [
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
          '../../../../public/assets/images/Homepage/quayvideo2.jpg',
          '../../../../public/assets/images/Homepage/quayvideo3.jpg',
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
          '../../../../public/assets/images/Homepage/quayvideo2.jpg',
          '../../../../public/assets/images/Homepage/quayvideo3.jpg',
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
          '../../../../public/assets/images/Homepage/quayvideo2.jpg',
          '../../../../public/assets/images/Homepage/quayvideo3.jpg',
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
          '../../../../public/assets/images/Homepage/quayvideo2.jpg',
          '../../../../public/assets/images/Homepage/quayvideo3.jpg',
          '../../../../public/assets/images/Homepage/quayvideo1.jpg',
        ],
      };
    },
    mounted() {},
  };
</script>

<style>
  .dp__action_buttons {
    display: none;
  }

  .dp__selection_preview {
    font-size: 24px;
  }

  .dp__action_row {
    justify-content: center;
    width: 100%;
  }
  .dp__outer_menu_wrap {
    width: 100%;
  }
</style>
