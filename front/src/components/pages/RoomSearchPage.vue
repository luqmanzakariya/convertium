<template>
  <HomeTemplate>
    <div class="pages-section">
      <h2>BOOK A ROOM</h2>
      <BookRoomForm
        v-model:guests="guests"
        v-model:dateCheckIn="dateCheckIn"
        v-model:dateCheckOut="dateCheckOut"
        @searchRooms="searchRooms"
      />
    </div>
  </HomeTemplate>
</template>

<script>
import HomeTemplate from "../templates/HomeTemplate.vue";
import BookRoomForm from "../organisms/form/BookRoomForm.vue";
import { APP_ROUTE } from "../../constant/url";

export default {
  name: "RoomSearchPages",
  components: { HomeTemplate, BookRoomForm },
  data() {
    return {
      guests: 2,
      dateCheckIn: new Date().toISOString().substr(0, 10),
      dateCheckOut: new Date().toISOString().substr(0, 10),
    };
  },
  methods: {
    async searchRooms() {
      try {
        this.$router.push({
          path: APP_ROUTE.ROOMS,
          query: {
            checkIn: this.dateCheckIn,
            checkOut: this.dateCheckOut,
            guests: this.guests,
          },
        });
      } catch (err) {
        alert("Error searching rooms");
      }
    },
  },
};
</script>

<style scoped>
.pages-section {
  margin: 20px 0;
}
</style>
