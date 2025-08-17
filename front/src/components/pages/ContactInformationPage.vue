<template>
  <HomeTemplate>
    <div class="pages-section">
      <h2>Contact Information Page</h2>
      <div class="room-selection-page">
        <Breadcrumb :items="breadcrumbItems" :current-route="$route.path" />
      </div>
    </div>
  </HomeTemplate>
</template>

<script>
import HomeTemplate from "../templates/HomeTemplate.vue";
import BookRoomForm from "../organisms/form/BookRoomForm.vue";
import Breadcrumb from "../molecules/breadcrumb/Breadcrumb.vue";
import { APP_ROUTE } from "../../constant/url";

export default {
  name: "ContactInformationPage",
  components: {
    HomeTemplate,
    BookRoomForm,
    Breadcrumb,
  },
  data() {
    return {
      loading: true,
      rooms: [],
      checkIn: this.$route.query.checkIn || "",
      checkOut: this.$route.query.checkOut || "",
      guests: this.$route.query.guests || 1,
      breadcrumbItems: [
        {
          label: "SEARCH",
          routes: [
            APP_ROUTE.HOME,
            APP_ROUTE.ROOMS,
            APP_ROUTE.CONTACT_INFORMATION,
            APP_ROUTE.CONFIRMATION,
          ],
        },
        {
          label: "SELECT ROOM",
          routes: [
            APP_ROUTE.HOME,
            APP_ROUTE.ROOMS,
            APP_ROUTE.CONTACT_INFORMATION,
            APP_ROUTE.CONFIRMATION,
          ],
        },
        {
          label: "CONTACT INFORMATION",
          routes: [APP_ROUTE.CONTACT_INFORMATION, APP_ROUTE.CONFIRMATION],
        },
        {
          label: "CONFIRMATION",
          routes: [APP_ROUTE.CONFIRMATION],
        },
      ],
    };
  },
  created() {
    this.checkAuthentication();
  },
  methods: {
    checkAuthentication() {
      const token = localStorage.getItem("token");
      if (!token) {
        // Redirect to login with current path as return URL
        this.$router.push({
          path: APP_ROUTE.LOGIN,
          query: { redirect: this.$route.fullPath },
        });
      }
    },
  },
  watch: {
    // Watch for changes in route in case of direct navigation
    $route: "checkAuthentication",
  },
};
</script>

<style scoped>
.pages-section {
  padding: 20px;
  text-align: left;
}

.room-list {
  margin-top: 20px;
}
</style>
