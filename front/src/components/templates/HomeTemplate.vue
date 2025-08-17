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
    min-height: calc(100dvh - 40px);
    text-align: center;
    padding: 0px 0px 0px;
    margin-bottom: 40px;
    height: 100%;
  }
  .header {
    width: 100%;
  }

  .banner-img {
    width: 100%;
    height: 400px;
  }

  @media (max-width: 768px) {
    .container {
      min-height: 100dvh;
      margin-bottom: 0;
    }
    .banner-img {
      height: 200px;
    }
  }
</style>
