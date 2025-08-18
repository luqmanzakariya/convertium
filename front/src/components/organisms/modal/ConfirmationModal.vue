<template>
  <transition name="modal">
    <div v-if="show" class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">
              <h3>{{ title }}</h3>
            </slot>
          </div>

          <div class="modal-body">
            <slot name="body">
              {{ message }}
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <button class="modal-button cancel" @click="onCancel">
                {{ cancelText }}
              </button>
              <button class="modal-button confirm" @click="onConfirm">
                {{ confirmText }}
              </button>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "ConfirmationModal",
  props: {
    show: {
      type: Boolean,
      required: true,
    },
    title: {
      type: String,
      default: "Confirm Action",
    },
    message: {
      type: String,
      default: "Are you sure you want to perform this action?",
    },
    confirmText: {
      type: String,
      default: "Confirm",
    },
    cancelText: {
      type: String,
      default: "Cancel",
    },
  },
  methods: {
    onConfirm() {
      this.$emit("confirm");
    },
    onCancel() {
      this.$emit("cancel");
    },
  },
};
</script>

<style scoped>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100%;
}

.modal-container {
  width: 400px;
  max-width: 90%;
  margin: 0 auto;
  padding: 20px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
}

.modal-header h3 {
  margin-top: 0;
  color: #333;
}

.modal-body {
  margin: 20px 0;
  color: #555;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.modal-button {
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 500;
  transition: background-color 0.2s;
}

.modal-button.confirm {
  background-color: #ff4444;
  color: white;
  border: none;
}

.modal-button.confirm:hover {
  background-color: #cc0000;
}

.modal-button.cancel {
  background-color: #f0f0f0;
  color: #333;
  border: 1px solid #ddd;
}

.modal-button.cancel:hover {
  background-color: #e0e0e0;
}

/* Transition effects */
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}

.modal-enter-active .modal-container,
.modal-leave-active .modal-container {
  transform: scale(1.1);
}
</style>
