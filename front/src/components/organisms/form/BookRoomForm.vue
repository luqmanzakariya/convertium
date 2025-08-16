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

      <DatePicker
        :modelValue="dateCheckIn"
        @update:modelValue="$emit('update:dateCheckIn', $event)"
      />

      <DatePicker
        :modelValue="dateCheckOut"
        @update:modelValue="$emit('update:dateCheckOut', $event)"
      />
    </div>

    <SubmitButton @click="$emit('searchRooms')" title="SEARCH FOR ROOMS" />
  </div>
</template>

<script>
import SubmitButton from "../../atoms/button/SubmitButton.vue";
import DatePicker from "../../atoms/date/DatePicker.vue";

export default {
  name: "BookRoomForm",
  components: { SubmitButton, DatePicker },
  props: ["guests", "dateCheckIn", "dateCheckOut"],
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
