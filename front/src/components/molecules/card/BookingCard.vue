<template>
  <div :class="['booking-card', { past: isPast }]">
    <div class="action">
      <img
        v-if="booking.hotelDetail.imageUrl"
        :src="booking.hotelDetail.imageUrl"
        alt="Hotel Image"
        class="hotel-image"
      />
      <div class="hotel-detail-name">{{ booking.hotelDetail.name }}</div>
      <div class="hotel-detail-night">{{ dateDifference }} Night</div>
    </div>
    <div class="detail">
      <div class="booking-info">
        <div class="info-item">
          <div class="info-title">Confirmation Code:</div>
          <div class="info-value">{{ booking.confirmationCode }}</div>
        </div>
        <div class="info-item">
          <div class="info-title">Dates:</div>
          <div class="info-value">
            {{ formatDate(booking.checkIn) }} &#x2192;
            {{ formatDate(booking.checkOut) }}
          </div>
        </div>
        <div class="info-item">
          <div class="info-title">Guests:</div>
          <div class="info-value">{{ booking.numberOfGuests }}</div>
        </div>
        <div class="info-item">
          <div class="info-title">Total Price:</div>
          <div class="info-value">S${{ booking.totalPrice }}</div>
        </div>
        <div class="info-item">
          <div class="info-title">Contact Name:</div>
          <div class="info-value">{{ booking.contactName }}</div>
        </div>
        <div class="info-item">
          <div class="info-title">Contact Email:</div>
          <div class="info-value">{{ booking.contactEmail }}</div>
        </div>
      </div>
      <button
        v-if="!isPast"
        @click="$emit('cancel-booking', booking.id)"
        class="cancel-button"
        :disabled="booking.cancelled"
      >
        {{ booking.cancelled ? "Cancelled" : "Cancel Booking" }}
      </button>

      <div v-else class="booking-status">
        {{ booking.cancelled ? "Cancelled" : "Completed" }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "BookingCard",
  props: {
    booking: {
      type: Object,
      required: true,
    },
    isPast: {
      type: Boolean,
      default: false,
    },
  },
  methods: {
    formatDate(dateString) {
      const options = {
        weekday: "short",
        day: "2-digit",
        month: "short",
        year: "numeric",
      };
      return new Date(dateString).toLocaleDateString("en-GB", options);
    },
  },
  computed: {
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
.booking-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 20px;
  display: flex;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  gap: 20px;
}

.action {
  padding-right: 40px;
  border-right: 1px solid rgba(0, 0, 0, 0.1);
}

.booking-card.past {
  opacity: 0.8;
  background-color: #f9f9f9;
}

.booking-info h4 {
  margin-top: 0;
  color: #333;
  width: 100%;
}

.booking-info p {
  margin: 8px 0;
  color: #555;
  width: 100%;
  flex-grow: 1;
}

.cancel-button {
  margin-top: 15px;
  padding: 8px 16px;
  background-color: #ff4444;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  align-self: flex-start;
  width: 100%;
}

.cancel-button:hover {
  background-color: #cc0000;
}

.cancel-button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}

.booking-status {
  margin-top: 15px;
  padding: 8px 16px;
  background-color: #f0f0f0;
  border-radius: 4px;
  align-self: flex-start;
}

.hotel-detail-name {
  font-size: 20px;
  font-weight: 600;
}

.hotel-detail-night {
  font-size: 16px;
  font-weight: 400;
}

img {
  height: 120px;
  width: auto;
}

.info-item {
  display: flex;
  gap: 10px;
  text-align: left;
  font-size: 16px;
}
.info-title {
  color: #555;
  font-weight: 700;
  width: 180px;
}
.info-value {
  color: #555;
  font-weight: 500;
}

.detail {
  width: 100%;
}
</style>
