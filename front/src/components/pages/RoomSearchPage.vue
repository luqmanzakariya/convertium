<template>
  <HomeTemplate>
    <div class="pages-section">
      <h2>BOOK A ROOM</h2>
      <BookRoomForm
        v-model:guests="guests"
        v-model:date="date"
        @searchRooms="searchRooms"
      />
    </div>
  </HomeTemplate>
</template>

<script>
import HomeTemplate from "../templates/HomeTemplate.vue";
import BookRoomForm from "../organisms/form/BookRoomForm.vue";

export default {
  name: "RoomSearchPages",
  components: { HomeTemplate, BookRoomForm },
  data() {
    return {
      guests: 2,
      date: new Date().toISOString().substr(0, 10),
    };
  },
  methods: {
    searchRooms() {
      console.log(`guest: ${this.guests}, date: ${this.date}`);
    },
  },
};
</script>

<style scoped>
.pages-section {
  margin: 20px 0;
}
</style>
