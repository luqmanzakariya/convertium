<template>
  <home-template>
    <div class="pages-section">
      <h2>LOGIN</h2>
      <LoginForm @success="handleLoginSuccess" />
      <div class="register-link">
        Don't have an account?
        <router-link to="/register">Register here</router-link>
      </div>
    </div>
  </home-template>
</template>

<script>
import HomeTemplate from "../templates/HomeTemplate.vue";
import LoginForm from "../organisms/form/LoginForm.vue";
import { APP_ROUTE } from "../../constant/url";

export default {
  name: "LoginPage",
  components: { HomeTemplate, LoginForm },
  methods: {
    handleLoginSuccess() {
      this.$router.push({
        path: APP_ROUTE.HOME,
      });
    },
  },
};
</script>

<style scoped>
.pages-section {
  margin: 20px 0;
  padding: 0px 0px 40px;
}

.register-link {
  margin-top: 16px;
  text-align: center;
}
</style>
