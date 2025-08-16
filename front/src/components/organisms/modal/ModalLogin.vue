<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <h3>Please Login First</h3>
      <p>You need to be logged in to book a room.</p>
      <div class="modal-actions">
        <button @click="handleLogin" class="primary-button">Login</button>
        <button @click="handleCancel" class="secondary-button">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "LoginModal",
  props: {
    show: {
      type: Boolean,
      required: true,
    },
  },
  emits: ["login", "cancel"],
  methods: {
    handleLogin() {
      this.$emit("login");
    },
    handleCancel() {
      this.$emit("cancel");
    },
  },
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 2rem;
  border-radius: 8px;
  max-width: 400px;
  width: 90%;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.modal-content h3 {
  margin-top: 0;
  color: #333;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}

.primary-button {
  background-color: #000;
  color: white;
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.secondary-button {
  background-color: #f0f0f0;
  color: #333;
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>
