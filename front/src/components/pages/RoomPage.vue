<template>
  <HomeTemplate>
    <div class="pages-section">
      <h2>Select Room Page</h2>
      <div v-if="loading">Loading rooms...</div>
    </div>
  </HomeTemplate>
</template>

<script>
import axios from "axios";
import HomeTemplate from "../templates/HomeTemplate.vue";
import BookRoomForm from "../organisms/form/BookRoomForm.vue";

export default {
  name: "RoomPage",
  components: { HomeTemplate, BookRoomForm },
  data() {
    return {
      loading: true,
      rooms: [],
    };
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
  margin: 20px 0;
}
</style>
