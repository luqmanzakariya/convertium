<template>
  <div class="booking-form">
    <div class="form-fields">
      <div class="guests">
        <div class="input-with-icon">
          <span class="icon">👤</span>
          <select
            :value="guests"
            @change="$emit('update:guests', $event.target.value)"
          >
            <option v-for="n in 10" :key="n" :value="n">{{ n }}</option>
          </select>
        </div>
      </div>

      <div class="date-picker">
        <input
          type="text"
          :value="formattedDate"
          readonly
          @click="openDatePicker"
          class="visible-date"
        />
        <input
          type="date"
          :value="date"
          @change="updateDate"
          ref="dateInput"
          class="native-date"
        />
      </div>
    </div>

    <SubmitButton @click="$emit('searchRooms')" title="SEARCH FOR ROOMS" />
  </div>
</template>

<script>
import SubmitButton from "../../atoms/button/SubmitButton.vue";

export default {
  name: "BookRoomForm",
  components: { SubmitButton },
  props: ["guests", "date"],
  computed: {
    formattedDate() {
      if (!this.date) return "Select date";
      const options = {
        weekday: "short",
        day: "2-digit",
        month: "short",
        year: "numeric",
      };
      return new Date(this.date).toLocaleDateString("en-GB", options);
    },
  },
  methods: {
    updateDate(event) {
      this.$emit("update:date", event.target.value);
    },
    openDatePicker() {
      this.$refs.dateInput.showPicker();
    },
  },
};
</script>

<style scoped>
.booking-form {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 10px;
  padding: 16px;
}

.form-fields {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
}

.guests {
  width: 50%;
}

.date-picker {
  width: 50%;
  position: relative;
}

.date-picker .visible-date {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 16px;
  background: #f9f9f9;
  height: 39px;
  box-sizing: border-box;
  cursor: pointer;
}

.date-picker .native-date {
  position: absolute;
  top: 0;
  left: 0;
  width: 1px;
  height: 1px;
  opacity: 0;
  pointer-events: none;
}

.input-with-icon {
  display: flex;
  align-items: center;
  border: 1px solid #ccc;
  border-radius: 4px;
  background: #f9f9f9;
  height: 39px;
  padding: 0 8px;
}

.input-with-icon .icon {
  margin-right: 8px;
  font-size: 16px;
}

.input-with-icon select {
  border: none;
  outline: none;
  background: transparent;
  font-size: 16px;
  width: 100%;
  cursor: pointer;
}

button {
  padding: 10px 20px;
  background-color: black;
  color: white;
  border: none;
  cursor: pointer;
  width: 100%;
  border-radius: 4px;
}

button:hover {
  background-color: #555;
}
</style>
