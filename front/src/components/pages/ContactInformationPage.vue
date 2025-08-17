<template>
  <HomeTemplate>
    <div class="pages-section">
      <h2>Contact Information Page</h2>
      <div class="room-selection-page">
        <Breadcrumb :items="breadcrumbItems" :current-route="$route.path" />
      </div>
      <div class="reservation-component">
        <div class="contact-info">
          <form @submit.prevent="handleSubmit">
            <div class="contact-form">
              <div class="contact-title">CONTACT INFORMATION</div>
              <div class="form-group">
                <label for="title">Title</label>
                <select v-model="form.title" id="title" required>
                  <option disabled value="">Select...</option>
                  <option value="Mr">Mr.</option>
                  <option value="Ms">Ms.</option>
                  <option value="Mrs">Mrs.</option>
                </select>
              </div>

              <div class="form-group">
                <label for="name">Name</label>
                <input type="text" v-model="form.name" id="name" required />
              </div>

              <div class="form-group">
                <label for="email">Email Address</label>
                <input type="email" v-model="form.email" id="email" required />
              </div>
            </div>

            <button type="submit">PROCEED</button>
          </form>
        </div>
        <div class="room-summary">
          <div class="room-summary-header">
            {{ formattedDateCheckIn }} &#x2192; {{ formattedDateCheckOut }}
          </div>
          <div class="room-summary-night">{{ dateDifference }} NIGHT</div>
          <div class="room-summary-guest">ROOM: {{ guests }} GUEST</div>
          <img :src="room.imageUrl" alt="img" />
          <div class="room-summary-name">{{ room.name }}</div>
          <div class="price-info">
            <div>Room</div>
            <div>S${{ priceInfo.roomPrice }}</div>
          </div>
          <div class="price-info">
            <div>Tax & Service Charge</div>
            <div>S${{ priceInfo.taxServiceCharge }}</div>
          </div>
          <div class="price-info-total">
            <div>TOTAL</div>
            <div>S${{ priceInfo.total }}</div>
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
  name: "ContactInformationPage",
  components: {
    HomeTemplate,
    BookRoomForm,
    Breadcrumb,
  },
  data() {
    return {
      loading: true,
      room: {},
      checkIn: this.$route.query.checkIn || "",
      checkOut: this.$route.query.checkOut || "",
      guests: this.$route.query.guests || 1,
      form: {
        title: "",
        name: "",
        email: "",
      },
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
    handleSubmit() {
      let payload = {
        roomId: Number(this.$route.query.roomId),
        checkIn: this.checkIn,
        checkOut: this.checkOut,
        numberOfGuests: Number(this.guests),
        title: this.form.title,
        contactName: this.form.name,
        contactEmail: this.form.email,
      };

      apiService
        .post(API_ROUTE.BOOKING, payload)
        .then((response) => {
          this.$router.push(`${APP_ROUTE.CONFIRMATION}?id=${response.data.id}`);
        })
        .catch((error) => {
          alert("Failed to book the room. Please try again.");
        });
    },
  },
  watch: {
    // Watch for changes in route in case of direct navigation
    $route: "checkAuthentication",
  },
  async mounted() {
    try {
      const res = await apiService.get(
        `${API_ROUTE.ROOMS}/${this.$route.query.roomId}`
      );
      this.room = res.data;
    } catch (err) {
      alert("Error searching rooms");
    } finally {
      this.loading = false;
    }
  },
  computed: {
    formattedDateCheckIn() {
      if (!this.checkIn) return "Select date";
      const options = {
        weekday: "short",
        day: "2-digit",
        month: "short",
        year: "numeric",
      };
      return new Date(this.checkIn).toLocaleDateString("en-GB", options);
    },
    formattedDateCheckOut() {
      if (!this.checkOut) return "Select date";
      const options = {
        weekday: "short",
        day: "2-digit",
        month: "short",
        year: "numeric",
      };
      return new Date(this.checkOut).toLocaleDateString("en-GB", options);
    },
    dateDifference() {
      if (!this.checkIn || !this.checkOut) return 0;
      const checkInDate = new Date(this.checkIn);
      const checkOutDate = new Date(this.checkOut);
      return Math.ceil((checkOutDate - checkInDate) / (1000 * 60 * 60 * 24));
    },
    priceInfo() {
      if (!this.checkIn || !this.checkOut) return 0;
      const checkInDate = new Date(this.checkIn);
      const checkOutDate = new Date(this.checkOut);
      const night = Math.ceil(
        (checkOutDate - checkInDate) / (1000 * 60 * 60 * 24)
      );

      return {
        roomPrice: this.room.price * night,
        taxServiceCharge: this.room.price * night * 0.09,
        total: this.room.price * night + this.room.price * night * 0.09,
      };
    },
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

.reservation-component {
  display: flex;
  width: 100%;
  gap: 20px;
}

.contact-info {
  width: 60%;
}

.contact-form {
  padding: 20px;
  background-color: #f0f1eb;
  border-radius: 4px;
  margin-bottom: 20px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 12px;
}

.contact-title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 20px;
}

.form-group {
  display: flex;
  align-items: center;
  gap: 20px;
}

label {
  width: 120px;
  font-size: 14px;
  font-weight: 400;
}

input,
select {
  width: 100%;
  padding: 8px 0px 8px 0px;
  border: 1px solid #ccc;
  border-radius: 4px;
  background: none;
  border: none;
  border-radius: 0;
  border-bottom: 1px solid #000;
}

input:focus,
select:focus {
  outline: none;
  box-shadow: none;
}

button {
  padding: 10px 15px;
  background-color: #000;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.room-summary {
  width: 40%;
  padding: 20px;
  background-color: #f0f1ec;
  border-radius: 4px;
}

img {
  width: 100%;
  height: 210px;
  object-fit: cover;
  border-radius: 4px;
}

.price-info {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
  font-size: 16px;
  font-weight: 400;
}

.price-info-total {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
  font-size: 20px;
  font-weight: 500;
  margin-top: 24px;
}

.room-summary-header {
  font-size: 16px;
  font-weight: 500;
  margin-bottom: 4px;
  text-transform: uppercase;
}

.room-summary-night {
  font-size: 16px;
  font-weight: 500;
  margin-bottom: 10px;
}
.room-summary-guest {
  font-size: 24px;
  font-weight: 500;
  margin-bottom: 10px;
}

.room-summary-name {
  margin-top: 10px;
  font-size: 16px;
  font-weight: 500;
  text-transform: uppercase;
}
</style>
