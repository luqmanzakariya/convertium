<template>
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
      :value="modelValue"
      @change="updateDate"
      ref="nativeDateInput"
      class="native-date"
    />
  </div>
</template>

<script>
export default {
  name: "DatePicker",
  props: ["modelValue"],
  emits: ["update:modelValue"],
  computed: {
    formattedDate() {
      if (!this.modelValue) return "Select date";
      const options = {
        weekday: "short",
        day: "2-digit",
        month: "short",
        year: "numeric",
      };

      return new Date(this.modelValue).toLocaleDateString("en-GB", options);
    },
  },
  methods: {
    updateDate(event) {
      this.$emit("update:modelValue", event.target.value);
    },
    openDatePicker() {
      this.$refs.nativeDateInput.showPicker();
    },
  },
};
</script>

<style scoped>
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
</style>
