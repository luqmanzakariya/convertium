<template>
  <div class="breadcrumb">
    <div
      v-for="(item, index) in items"
      :key="index"
      class="breadcrumb-item"
      :class="{
        active: isActive(index),
      }"
    >
      <div class="circle">{{ index + 1 }}</div>
      <div>{{ item.label }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Breadcrumb",
  props: {
    items: {
      type: Array,
      required: true,
      default: () => [],
    },
    currentRoute: {
      type: String,
      required: true,
    },
  },
  methods: {
    isActive(index) {
      return this.items[index].routes.includes(this.currentRoute);
    },
  },
};
</script>

<style scoped>
.breadcrumb {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  border-radius: 4px;
  background-color: #f0f1eb;
  padding: 16px;
  gap: 16px;
  overflow: auto;
  white-space: nowrap;
}

.breadcrumb-item {
  display: flex;
  align-items: center;
  gap: 12px;
  color: #666;
  font-weight: 600;
  text-transform: uppercase;
}

.breadcrumb-item .circle {
  border-radius: 100%;
  width: 32px;
  height: 32px;
  background-color: #ccc;
  color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 12px;
}

.breadcrumb-item.active {
  color: black;
}

.breadcrumb-item.active .circle {
  background-color: black;
}
</style>
