<template>
  <HomeTemplate>
    <div class="pages-section">
      <h2>Select Room Page</h2>
      <div v-if="loading">Loading rooms...</div>
      <div class="room-selection-page">
        <Breadcrumb :items="breadcrumbItems" :current-route="$route.path" />

        <RoomBookInfoFilter
          :check-in="checkIn"
          :check-out="checkOut"
          :guests="guests"
          :sort-option="sortOption"
          :sort-options="sortOptions"
          @update:sortOption="sortOption = $event"
          @sort-change="handleSortChange"
        />

        <div class="room-list">
          <RoomCard
            v-for="room in sortedRooms"
            :key="room.id"
            :room="room"
            @book-room="bookRoom"
          />
        </div>

        <ModalLogin
          :show="showLoginModal"
          @login="redirectToLogin"
          @cancel="showLoginModal = false"
        />
      </div>
    </div>
  </HomeTemplate>
</template>

<script>
import axios from "axios";
import HomeTemplate from "../templates/HomeTemplate.vue";
import BookRoomForm from "../organisms/form/BookRoomForm.vue";
import Breadcrumb from "../molecules/breadcrumb/Breadcrumb.vue";
import RoomBookInfoFilter from "../molecules/filter/RoomBookInfoFilter.vue";
import RoomCard from "../molecules/card/RoomCard.vue";
import ModalLogin from "../organisms/modal/ModalLogin.vue";
import { APP_ROUTE } from "../../constant/url";

export default {
  name: "RoomPage",
  components: {
    HomeTemplate,
    BookRoomForm,
    Breadcrumb,
    RoomBookInfoFilter,
    RoomCard,
    ModalLogin,
  },
  data() {
    return {
      loading: true,
      rooms: [],
      checkIn: this.$route.query.checkIn || "",
      checkOut: this.$route.query.checkOut || "",
      guests: this.$route.query.guests || 1,
      showLoginModal: false,
      sortOption: "lowest",
      sortOptions: [
        { label: "Lowest Price", value: "lowest" },
        { label: "Highest Price", value: "highest" },
      ],
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
  methods: {
    bookRoom(roomId) {
      const token = localStorage.getItem("token");

      if (!token) {
        this.showLoginModal = true;
        return;
      }

      this.$router.push(`${APP_ROUTE.CONTACT_INFORMATION}?roomId=${roomId}`);
    },
    redirectToLogin() {
      this.showLoginModal = false;
      this.$router.push(`${APP_ROUTE.LOGIN}`);
    },
    handleSortChange(value) {
      this.sortOption = value;
    },
  },
  computed: {
    sortedRooms() {
      const rooms = [...this.rooms];

      if (this.sortOption === "lowest") {
        return rooms.sort((a, b) => a.price - b.price);
      } else if (this.sortOption === "highest") {
        return rooms.sort((a, b) => b.price - a.price);
      }

      return rooms;
    },
  },
  async mounted() {
    try {
      const res = await axios.get(`${import.meta.env.VITE_API_URL}/rooms`, {
        params: {
          checkIn: this.$route.query.checkIn,
          checkOut: this.$route.query.checkOut,
        },
      });
      this.rooms = res.data;
    } catch (err) {
      alert("Error searching rooms");
    } finally {
      this.loading = false;
    }
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
