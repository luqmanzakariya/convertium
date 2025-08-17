<template>
  <HomeTemplate>
    <div class="pages-section">
      <h2>Confirmation Page</h2>
      <Breadcrumb :items="breadcrumbItems" :current-route="$route.path" />
      <div class="header-title">YOUR BOOKING HAS BEEN CONFIRMED</div>
      <div class="reservation-info">
        <div>
          We have sent your booking confirmation to the email address that you
          have provided.
        </div>
        <div>
          Check-in/Check-out:
          <span class="bold"
            >{{ formattedDateCheckIn }} &#x2192;
            {{ formattedDateCheckOut }}</span
          >
        </div>
        <div>
          Booking Confirmation Number:
          <span class="bold">{{ booking.confirmationCode }}</span>
        </div>
        <div>
          Total Price for {{ dateDifference }} Night:
          <span class="bold">{{ booking.totalPrice }}</span>
        </div>
      </div>
      <div class="room-info">
        <div class="room-info-left">
          <div class="room-info-header">
            <div>
              <img :src="room.imageUrl" alt="Room Image" />
            </div>
            <div>
              <div>ROOM: {{ room.name }}</div>
              <div class="room-guest">{{ booking.numberOfGuests }} Guests</div>
            </div>
          </div>
          <div class="room-info-details">
            <div class="room-info-details-package">PACKAGE</div>
            <div class="room-info-item">
              <div>Room</div>
              <div>S${{ room.price }}</div>
            </div>
            <div class="room-info-item">
              <div>Tax & Service Charge (9%)</div>
              <div>S${{ booking.tax }}</div>
            </div>
            <div class="room-info-item">
              <div>TOTAL</div>
              <div>S${{ booking.totalPrice }}</div>
            </div>
          </div>
        </div>
        <div class="room-info-right">
          <div class="room-info-header">GUEST DETAILS</div>
          <div>
            <div>
              <span class="label">Name</span>: {{ booking.title }}.
              {{ booking.contactName }}
            </div>
            <div>
              <span class="label">Email</span>: {{ booking.contactEmail }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </HomeTemplate>
</template>

<script>
import HomeTemplate from "../templates/HomeTemplate.vue";
import BookRoomForm from "../organisms/form/BookRoomForm.vue";
import Breadcrumb from "../molecules/breadcrumb/Breadcrumb.vue";
import apiService from "../../service/api";
import { APP_ROUTE, API_ROUTE } from "../../constant/url";

export default {
  name: "ConfirmationPage",
  components: {
    HomeTemplate,
    BookRoomForm,
    Breadcrumb,
  },
  data() {
    return {
      loading: true,
      id: this.$route.query.id || "",
      booking: {},
      room: {},
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
    handleSubmit() {},
  },
  watch: {
    // Watch for changes in route in case of direct navigation
    $route: "checkAuthentication",
  },
  async mounted() {
    try {
      const bookingInfo = await apiService.get(
        `${API_ROUTE.BOOKING}/${this.$route.query.id}`
      );

      const roomInfo = await apiService.get(
        `${API_ROUTE.ROOMS}/${bookingInfo.data.roomId}`
      );

      this.booking = bookingInfo.data;
      this.room = roomInfo.data;
    } catch (err) {
      alert("Error searching rooms");
    } finally {
      this.loading = false;
    }
  },
  computed: {
    formattedDateCheckIn() {
      if (!this.booking.checkIn) return "Select date";
      const options = {
        weekday: "short",
        day: "2-digit",
        month: "short",
        year: "numeric",
      };
      return new Date(this.booking.checkIn).toLocaleDateString(
        "en-GB",
        options
      );
    },
    formattedDateCheckOut() {
      if (!this.booking.checkOut) return "Select date";
      const options = {
        weekday: "short",
        day: "2-digit",
        month: "short",
        year: "numeric",
      };
      return new Date(this.booking.checkOut).toLocaleDateString(
        "en-GB",
        options
      );
    },
    dateDifference() {
      if (!this.booking.checkIn || !this.booking.checkOut) return 0;
      const checkInDate = new Date(this.booking.checkIn);
      const checkOutDate = new Date(this.booking.checkOut);
      return Math.ceil((checkOutDate - checkInDate) / (1000 * 60 * 60 * 24));
    },
  },
};
</script>

<style scoped>
.pages-section {
  padding: 20px;
  text-align: left;
}

.header-title {
  font-size: 24px;
  font-weight: 500;
  text-align: center;
  text-transform: uppercase;
  margin-top: 32px;
  margin-bottom: 20px;
}

.reservation-info {
  text-align: center;
  font-weight: 400;
  font-size: 16px;
  margin-bottom: 20px;
}

.bold {
  font-weight: 700;
}

.room-info {
  border-radius: 4px;
  padding: 20px;
  background-color: #f0f1eb;
  display: flex;
  gap: 16px;
}

.room-info-left {
  width: 65%;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.room-info-header {
  display: flex;
  gap: 20px;
  font-size: 24px;
  font-weight: 500;
}

.room-guest {
  font-size: 16px;
  font-weight: 400;
}

.room-info-details {
  display: flex;
  flex-direction: column;
  gap: 8px;
  font-size: 16px;
}

.room-info-details-package {
  font-weight: 600;
  text-transform: uppercase;
  font-size: 20px;
}

.room-info-item {
  display: flex;
  justify-content: space-between;
}

.room-info-right {
  width: 35%;
  display: flex;
  flex-direction: column;
  gap: 8px;
  background-color: #e1e2d6;
  padding: 20px;
  border-radius: 4px;
}

.label {
  width: 60px;
  display: inline-block;
}

img {
  width: auto;
  height: 100px;
  object-fit: cover;
  border-radius: 4px;
}
</style>
