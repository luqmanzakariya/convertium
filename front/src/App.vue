<template>
  <router-view></router-view>
</template>

<script>
export default {
  data() {
    return {
      isLoggedIn: !!localStorage.getItem("token"),
    };
  },
};
</script>

<style scoped></style>
