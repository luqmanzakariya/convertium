<template>
  <div class="room-book-info">
    <div>{{ formattedDateCheckIn }} &#x2192; {{ formattedDateCheckOut }}</div>
    <div>
      <div>{{ dateDifference }} NIGHT | {{ guests }} GUEST</div>
      <div>
        SORT BY:
        <select :value="sortOption" @change="onSortChange($event.target.value)">
          <option
            v-for="option in sortOptions"
            :key="option.value"
            :value="option.value"
          >
            {{ option.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "RoomBookInfo",
  props: {
    checkIn: {
      type: String,
      default: "",
    },
    checkOut: {
      type: String,
      default: "",
    },
    guests: {
      type: [Number, String],
      default: 1,
    },
    sortOption: {
      type: String,
      default: "lowest",
    },
    sortOptions: {
      type: Array,
      default: () => [
        { label: "Lowest Price", value: "lowest" },
        { label: "Highest Price", value: "highest" },
      ],
    },
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
  },
  methods: {
    onSortChange(value) {
      this.$emit("update:sortOption", value);
      this.$emit("sort-change", value);
    },
  },
};
</script>

<style scoped>
.room-book-info {
  display: flex;
  justify-content: space-between;
  align-items: start;
  margin-bottom: 20px;
  border-radius: 4px;
}

select {
  border: none;
  text-transform: uppercase;
  font-size: 16px;
}

select:focus {
  outline: none;
  box-shadow: none;
}
</style>
