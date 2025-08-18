<template>
  <HomeTemplate>
    <div class="pages-section">
      <ConfirmationModal
        :show="showCancelModal"
        title="Cancel Booking"
        message="Are you sure you want to cancel this booking?"
        confirm-text="Yes, Cancel Booking"
        cancel-text="No, Keep Booking"
        @confirm="confirmCancel"
        @cancel="showCancelModal = false"
      />

      <ToastNotification
        v-if="showToast"
        :show="showToast"
        :message="toastMessage"
        :type="toastType"
        @close="showToast = false"
      />

      <div v-if="upcoming.length > 0" class="bookings-section">
        <h3>Upcoming Bookings</h3>
        <div class="booking-cards">
          <BookingCard
            v-for="booking in upcoming"
            :key="booking.id"
            :booking="booking"
            @cancel-booking="showCancelConfirmation"
          />
        </div>
      </div>
      <div v-else class="no-bookings">
        <p>No upcoming bookings found.</p>
      </div>

      <div v-if="past.length > 0" class="bookings-section">
        <h3>Past Bookings</h3>
        <div class="booking-cards">
          <BookingCard
            v-for="booking in past"
            :key="booking.id"
            :booking="booking"
            :is-past="true"
          />
        </div>
      </div>
      <div v-else class="no-bookings">
        <p>No past bookings found.</p>
      </div>
    </div>
  </HomeTemplate>
</template>

<script>
import apiService from "../../service/api";
import { API_ROUTE } from "../../constant/url";
import HomeTemplate from "../templates/HomeTemplate.vue";
import BookingCard from "../molecules/card/BookingCard.vue";
import ConfirmationModal from "../organisms/modal/ConfirmationModal.vue";
import ToastNotification from "../atoms/toast/ToastNotification.vue";

export default {
  name: "DashboardPage",
  components: {
    HomeTemplate,
    BookingCard,
    ConfirmationModal,
    ToastNotification,
  },
  data() {
    return {
      loading: true,
      id: this.$route.query.id || "",
      past: [],
      upcoming: [],
      room: {},
      showCancelModal: false,
      selectedBookingId: null,
      showToast: false,
      toastType: "",
      toastMessage: "",
    };
  },
  methods: {
    showCancelConfirmation(bookingId) {
      this.selectedBookingId = bookingId;
      this.showCancelModal = true;
    },
    async confirmCancel() {
      try {
        await apiService.delete(
          `${API_ROUTE.BOOKING}/${this.selectedBookingId}`
        );
        this.upcoming = this.upcoming.map((booking) =>
          booking.id === this.selectedBookingId
            ? { ...booking, cancelled: true }
            : booking
        );
        this.showCancelModal = false;
        this.toastType = "success";
        this.toastMessage = "Booking cancelled successfully";
        this.showToast = true;
      } catch (err) {
        this.showCancelModal = false;
        this.toastType = "error";
        this.toastMessage =
          "cannot cancel booking, you can only cancel 1 day before the current date";
        this.showToast = true;
        setTimeout(() => {
          this.showToast = false;
        }, 3000);
      }
    },
  },
  async mounted() {
    try {
      const bookingInfo = await apiService.get(`${API_ROUTE.BOOKING}`);

      const pastDataWithHotelDetails = await Promise.all(
        bookingInfo.data.past.map(async (past) => {
          let hotelDetail = await apiService.get(
            `${API_ROUTE.ROOMS}/${past.roomId}`
          );
          return {
            ...past,
            hotelDetail: hotelDetail.data,
          };
        })
      );

      const upcomingDataWithHotelDetails = await Promise.all(
        bookingInfo.data.upcoming.map(async (upcoming) => {
          let hotelDetail = await apiService.get(
            `${API_ROUTE.ROOMS}/${upcoming.roomId}`
          );
          return {
            ...upcoming,
            hotelDetail: hotelDetail.data,
          };
        })
      );

      this.past = pastDataWithHotelDetails || [];
      this.upcoming = upcomingDataWithHotelDetails || [];
    } catch (err) {
      this.$toast.error("Error loading bookings");
    } finally {
      this.loading = false;
    }
  },
};
</script>

<style scoped>
.pages-section {
  margin: 20px 0;
  padding: 0px 0px 40px;
}

.bookings-section {
  margin-top: 30px;
}

.bookings-section h3 {
  border-bottom: 1px solid #eee;
  padding-bottom: 10px;
  margin-bottom: 20px;
}

.booking-cards {
  display: flex;
  flex-direction: column;
  gap: 20px;
  padding: 20px;
}

.no-bookings {
  margin: 20px 0;
  color: #666;
  font-style: italic;
}
</style>
