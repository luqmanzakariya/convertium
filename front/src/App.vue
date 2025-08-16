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
  methods: {
    logout() {
      localStorage.removeItem("token");
      this.isLoggedIn = false;
      this.$router.push("/");
    },
  },
};
</script>

<style scoped></style>
