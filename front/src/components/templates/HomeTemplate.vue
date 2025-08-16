<template>
  <div class="container">
    <Masthead />
    <div class="header">
      <img src="../../assets/banner/home-banner.webp" alt="banner" class="banner-img"></img>
    </div>
    <slot></slot>
  </div>
</template>

<script>
  import Masthead from "../organisms/mashead/Mashead.vue";

  export default {
    name: "HomeTemplate",
    components: { Masthead },
  }
</script>

<style scoped>
  .container {
    min-height: 100dvh;
    text-align: center;
    padding: 0px 0px 0px;
    height: 100%;
  }
  .header {
    width: 100%;
  }

  .banner-img {
    width: 100%;
    height: 400px;
  }
</style>
