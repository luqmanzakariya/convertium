<template>
  <transition name="fade">
    <div v-if="show" class="toast" :class="type">
      <div class="toast-content">
        <span class="toast-message">{{ message }}</span>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "ToastNotification",
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    message: {
      type: String,
      default: "",
    },
    type: {
      type: String,
      default: "success",
      validator: (value) => ["success", "error", "warning"].includes(value),
    },
    duration: {
      type: Number,
      default: 3000,
    },
  },
  watch: {
    show(newVal) {
      if (newVal) {
        setTimeout(() => {
          this.$emit("close");
        }, this.duration);
      }
    },
  },
};
</script>

<style scoped>
.toast {
  position: fixed;
  bottom: 20px;
  right: 20px;
  padding: 12px 24px;
  border-radius: 4px;
  color: white;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 1000;
  display: flex;
  align-items: center;
}

.toast.success {
  background-color: #4caf50;
}

.toast.error {
  background-color: #f44336;
}

.toast.warning {
  background-color: #ff9800;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
