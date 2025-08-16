<template>
  <home-template>
    <div class="pages-section">
      <h2>LOGIN</h2>
    </div>
  </home-template>
</template>

<script>
import HomeTemplate from "../templates/HomeTemplate.vue";

export default {
  name: "LoginPage",
  components: { HomeTemplate },
};
</script>

<style scoped>
.pages-section {
  margin: 20px 0;
}
</style>
