<template>
  <header-view v-if="showHeaderFooter" />
  <router-view></router-view>
  <!-- <router-view v-slot="{ Component }">
    <transition name="fade" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view> -->
  <footer-view />
</template>

<script>
import HeaderView from './components/HeaderView.vue';
import FooterView from './components/FooterView.vue';
export default {
  components: {
    HeaderView, FooterView
  },
  computed: {
    showHeaderFooter() {
      // Truy cập giá trị showHeaderFooter từ meta data của tuyến hiện tại
      return this.$route.meta.showHeaderFooter;
    }
  }

}
</script>
<style></style>
