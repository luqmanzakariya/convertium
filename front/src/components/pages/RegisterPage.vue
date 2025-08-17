<template>
  <home-template>
    <div class="pages-section">
      <h2>REGISTER</h2>
      <RegisterForm @success="handleRegistrationSuccess" />
      <ToastNotification
        v-if="showToast"
        :show="showToast"
        message="Registration successful! Please login."
        type="success"
        @close="showToast = false"
      />
    </div>
  </home-template>
</template>

<script>
import HomeTemplate from "../templates/HomeTemplate.vue";
import RegisterForm from "../organisms/form/RegisterForm.vue";
import ToastNotification from "../atoms/toast/ToastNotification.vue";
import { APP_ROUTE } from "../../constant/url";

export default {
  name: "RegisterPage",
  components: { HomeTemplate, RegisterForm, ToastNotification },
  data() {
    return {
      showToast: false,
    };
  },
  methods: {
    handleRegistrationSuccess() {
      this.showToast = true;
      setTimeout(() => {
        this.$router.push({
          path: APP_ROUTE.LOGIN,
        });
      }, 3000);
    },
  },
};
</script>

<style scoped>
.pages-section {
  margin: 20px 0;
  padding: 0px 0px 40px;
}
</style>
