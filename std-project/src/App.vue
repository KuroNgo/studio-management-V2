<template>
  <header-view />
  <router-view></router-view>
  <!-- <router-view v-slot="{ Component }">
    <transition name="fade" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view> -->
  <!-- <footer-view /> -->
</template>

<script>
import HeaderView from './components/HeaderView.vue';
import FooterView from './components/FooterView.vue';
export default {
  components: {
    HeaderView, FooterView
  }
}
</script>
<style>

</style>
