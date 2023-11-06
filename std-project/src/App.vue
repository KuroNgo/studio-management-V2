<template>
  <div v-if="show" class="banner"></div>
  <header-view v-if="showHeaderFooter" />
  <router-view></router-view>
  <!-- <router-view v-slot="{ Component }">
    <transition name="fade" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view> -->
  <footer-view v-if="showHeaderFooter" />
</template>

<script>
  import HeaderView from './components/HeaderView.vue';
  import FooterView from './components/FooterView.vue';
  export default {
    components: {
      HeaderView,
      FooterView,
    },
    computed: {
      showHeaderFooter() {
        // Truy cập giá trị showHeaderFooter từ meta data của tuyến hiện tại
        return this.$route.meta.showHeaderFooter;
      },
      show() {
        return this.$route.name == 'Trang chủ | Cỏ Studio';
      },
    },
  };
</script>
<style>
  .banner {
    width: 100%;
    height: 200px;
    border-bottom: 2px solid #ccc;
    background: url(../public/assets/images/banner.jpg) no-repeat center/cover;
  }
</style>
